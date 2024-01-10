package dto

import "time"

type OrderInsertRequest struct {
	ID          int64                    `json:"id" binding:"required,gt=0"`
	Number      string                   `json:"number" binding:"required,gt=0"`
	CreatedDate time.Time                `json:"created_date" binding:"required"`
	Items       []OrderItemInsertRequest `json:"items" binding:"required,gt=0,dive"`
}

type OrderItemInsertRequest struct {
	Name     string `json:"name" binding:"required,gt=0"`
	Quantity int    `json:"quantity" binding:"required,gt=0"`
	Comment  string `json:"comment"`
}
