package mapper

import (
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/production-api/internal/core/domain"
)

func MapInsertDTOToDomain(orderDTO *dto.OrderInsertRequest) *domain.Order {
	items := make([]domain.Item, 0)
	for _, itemDTO := range orderDTO.Items {
		item := domain.Item{
			Name:     itemDTO.Name,
			Quantity: itemDTO.Quantity,
			Comment:  itemDTO.Comment,
		}
		items = append(items, item)
	}

	return &domain.Order{
		ID:          orderDTO.ID,
		Number:      orderDTO.Number,
		CreatedDate: orderDTO.CreatedDate,
		Items:       items,
	}
}
