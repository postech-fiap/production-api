package dto

import "time"

type NewOrderMessage struct {
	ID          int64                 `json:"id" validate:"required,gt=0"`
	Number      string                `json:"number" validate:"required,gt=0"`
	CreatedDate time.Time             `json:"created_date" validate:"required"`
	Items       []NewOrderItemMessage `json:"items" validate:"required,gt=0,dive"`
}

type NewOrderItemMessage struct {
	Name     string `json:"name" validate:"required,gt=0"`
	Quantity int    `json:"quantity" validate:"required,gt=0"`
	Comment  string `json:"comment"`
}
