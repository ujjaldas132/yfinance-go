package models

type HistoryData struct {
	Datetime int64   `json:"datetime"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	AdjClose float64 `json:"adj_close"`
}
