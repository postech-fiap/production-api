package domain

type Status string

const (
	PENDING    Status = "pending"
	RECEIVED   Status = "received"
	IN_PREPARE Status = "in_prepare"
	DONE       Status = "done"
	FINISHED   Status = "finished"
)
