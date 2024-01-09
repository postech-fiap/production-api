package dto

type UpdateStatusRequestURI struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateStatusRequestBody struct {
	Status string `json:"status" binding:"required"`
}
