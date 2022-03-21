package eldenring

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiUrl = "https://eldenring.fanapis.com/api/"

type Client struct {
	http *http.Client
	url  string
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{},
		url:  apiUrl,
	}
}

func (c *Client) doRequest(ctx context.Context, endpoint, method string, expRes interface{}, queries ...map[string]string) error {
	callURL, err := url.Parse(c.url + endpoint)
	if err != nil {
		return err
	}
	if len(queries) > 0 {
		q := callURL.Query()
		for k, v := range queries[0] {
			q.Set(k, v)
		}
		callURL.RawQuery = q.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, method, callURL.String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	code := resp.StatusCode
	switch {
	case 200 <= code && code < 300:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("%s", body)
	}
}

func (c *Client) GetAmmos(ctx context.Context) ([]Ammo, error) {
	e := "/ammos"
	var ammos AmmosResult
	if err := c.doRequest(ctx, e, http.MethodGet, &ammos); err != nil {
		return nil, err
	}
	return ammos.Data, nil
}

func (c *Client) GetArmors(ctx context.Context) ([]Armor, error) {
	e := "/armors"
	var armors ArmorResult
	if err := c.doRequest(ctx, e, http.MethodGet, &armors); err != nil {
		return nil, err
	}
	return armors.Data, nil
}
