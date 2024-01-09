package domain

import "time"

type Order struct {
	ID          int64
	Status      Status
	Number      string
	CreatedDate time.Time
	Items       []Item
}

func (o *Order) IsValidStatus(newStatus Status) bool {
	actualStatus := o.Status

	switch newStatus {
	case RECEIVED:
		return actualStatus == PENDING
	case IN_PREPARE:
		return actualStatus == RECEIVED
	case DONE:
		return actualStatus == IN_PREPARE
	case FINISHED:
		return actualStatus == DONE
	default:
		return false
	}
}
