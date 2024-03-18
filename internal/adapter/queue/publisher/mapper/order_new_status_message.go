package mapper

import (
	"github.com/postech-fiap/production-api/internal/adapter/queue/publisher/dto"
	"github.com/postech-fiap/production-api/internal/core/domain"
)

func DomainToOrderNewStatusMessage(orderDomain *domain.Order) *dto.OrderNewStatusMessage {
	return &dto.OrderNewStatusMessage{
		ID:     orderDomain.ID,
		Number: orderDomain.Number,
		Status: string(orderDomain.Status),
	}
}
