package dto

import (
	"time"
)

type OrderWrapperResponse struct {
	Orders []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	ID          int64               `json:"id"`
	Status      string              `json:"status"`
	Number      string              `json:"number"`
	CreatedDate time.Time           `json:"created_date"`
	Items       []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Comment  string `json:"comment"`
}
