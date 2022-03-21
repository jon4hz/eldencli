package eldenring

type EffectPair struct {
	Name   string `json:"name,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

type Ammos struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Image       string      `json:"image,omitempty"`
	Description string      `json:"description,omitempty"`
	Type        string      `json:"type,omitempty"`
	Passive     string      `json:"passive,omitempty"`
	AttackPower *EffectPair `json:"attack_power,omitempty"`
}

type APIResult struct {
	Success bool  `json:"success,omitempty"`
	Count   int64 `json:"count,omitempty"`
}

type AmmosResult struct {
	APIResult
	Data []Ammos `json:"data,omitempty"`
}

type Armor struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Image       string      `json:"image,omitempty"`
	Description string      `json:"description,omitempty"`
	Category    string      `json:"category,omitempty"`
	Weight      int64       `json:"weight,omitempty"`
	DMGNegation *EffectPair `json:"dmg_negation,omitempty"`
	Resistance  *EffectPair `json:"resistance,omitempty"`
}

type ArmorResult struct {
	APIResult
	Data []Armor `json:"data,omitempty"`
}
