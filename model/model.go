package model

type Tip struct {
	Id       int      `json:"id"`
	Match    []string `json:"match"`
	Analysis string   `json:"analysis"`
	Bookie   string   `json:"bookie"`
	Rate     float32  `json:"rate"`
	Stake    int      `json:"stake"`
	Pick     string   `json:"pick"`
}
