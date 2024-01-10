package mapper

import (
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/production-api/internal/core/domain"
)

func MapDomainToOrderWrapperDto(ordersDomain []domain.Order) *dto.OrderWrapperResponse {
	orders := make([]dto.OrderResponse, 0)

	for _, orderDomain := range ordersDomain {
		order := mapDomainToOrderDto(&orderDomain)
		orders = append(orders, *order)
	}

	return &dto.OrderWrapperResponse{
		Orders: orders,
	}
}

func mapDomainToOrderDto(orderDomain *domain.Order) *dto.OrderResponse {
	items := make([]dto.OrderItemResponse, 0)

	for _, itemDomain := range orderDomain.Items {
		item := dto.OrderItemResponse{
			Name:     itemDomain.Name,
			Quantity: itemDomain.Quantity,
			Comment:  itemDomain.Comment,
		}
		items = append(items, item)
	}

	return &dto.OrderResponse{
		ID:          orderDomain.ID,
		Status:      string(orderDomain.Status),
		Number:      orderDomain.Number,
		CreatedDate: orderDomain.CreatedDate,
		Items:       items,
	}
}
