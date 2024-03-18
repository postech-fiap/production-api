package dto

type OrderNewStatusMessage struct {
	ID     int64  `json:"id_pedido"`
	Number string `json:"numero_pedido"`
	Status string `json:"status_pedido"`
}
