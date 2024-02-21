package dto

type OrderNewStatusMessage struct {
	ID     int64  `json:"id"`
	Number string `json:"number"`
	Status string `json:"status"`
}
