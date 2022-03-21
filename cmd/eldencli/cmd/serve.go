package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
	"github.com/jon4hz/eldencli/internal/tui"
	"github.com/spf13/cobra"
)

const defaultServerPort = 2222

var serverFlags struct {
	host    string
	hostKey string
	port    int
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	Long:  `Run the server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return server(serverFlags.host, serverFlags.port, serverFlags.hostKey)
	},
}

func init() {
	serveCmd.Flags().IntVarP(&serverFlags.port, "port", "p", defaultServerPort, "Port to listen on")
	serveCmd.Flags().StringVarP(&serverFlags.host, "server", "s", "", "Host to listen on")
	serveCmd.Flags().StringVarP(&serverFlags.hostKey, "hostkey", "", ".ssh/alvarebot_hostkey_ed25519", "Path to hostkey file")
}

func server(host string, port int, hostKey string) error {
	serverTeaHandler := func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
		_, _, active := s.Pty()
		if !active {
			fmt.Println("no active terminal, skipping")
			return nil, nil
		}

		return tui.InitalModel(), []tea.ProgramOption{tea.WithAltScreen()}
	}

	// create a new ssh server
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(hostKey),
		wish.WithMiddleware(
			bm.Middleware(serverTeaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		return err
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	if err := s.Close(); err != nil {
		return err
	}
	return nil
}
