package eldenring

type EffectPair struct {
	Name   string `json:"name,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type Ammo struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Image       string       `json:"image,omitempty"`
	Description string       `json:"description,omitempty"`
	Type        string       `json:"type,omitempty"`
	Passive     string       `json:"passive,omitempty"`
	AttackPower []EffectPair `json:"attackPower,omitempty"`
}

type APIResult struct {
	Success bool `json:"success,omitempty"`
	Count   int  `json:"count,omitempty"`
}

type AmmosResult struct {
	APIResult
	Data []Ammo `json:"data,omitempty"`
}

type Armor struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Image       string      `json:"image,omitempty"`
	Description string      `json:"description,omitempty"`
	Category    string      `json:"category,omitempty"`
	Weight      int         `json:"weight,omitempty"`
	DMGNegation *EffectPair `json:"dmg_negation,omitempty"`
	Resistance  *EffectPair `json:"resistance,omitempty"`
}

type ArmorResult struct {
	APIResult
	Data []Armor `json:"data,omitempty"`
}
