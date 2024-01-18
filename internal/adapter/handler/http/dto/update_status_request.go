package dto

type UpdateStatusRequestURI struct {
	ID int64 `uri:"id" binding:"required,gt=0"`
}

type UpdateStatusRequestBody struct {
	Status string `json:"status" binding:"required,gt=0"`
}
