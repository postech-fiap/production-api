package http

import (
	"encoding/json"
	"errors"
	"github.com/postech-fiap/production-api/internal/adapter/handler/http/dto"
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"github.com/postech-fiap/production-api/internal/core/port"
	"time"
)

const mockOrderId int64 = 1

const mockInsertOrderRequest = `
{
    "id": 1,
    "number": "0001",
    "created_date": "2024-01-17T00:00:00Z",
    "items": [
        {
            "name": "Hamburger",
            "quantity": 1,
            "comment": "No lettuce"
        }
    ]
}
`

const mockUpdateOrderStatusRequest = `{"status": "done"}`

const mockStatus = domain.DONE

func mockDomainOrderList() []domain.Order {
	return []domain.Order{
		{
			ID:          1,
			Status:      domain.DONE,
			Number:      "0001",
			CreatedDate: time.Date(2024, 01, 12, 0, 0, 0, 0, time.UTC),
			Items: []domain.Item{
				{
					Name:     "Hamburger",
					Quantity: 1,
					Comment:  "No lettuce",
				},
			},
		},
	}
}

func mockOrderListResponse() string {
	response := dto.OrderWrapperResponse{
		Orders: []dto.OrderResponse{
			{
				ID:          1,
				Status:      "done",
				Number:      "0001",
				CreatedDate: time.Date(2024, 01, 12, 0, 0, 0, 0, time.UTC),
				Items: []dto.OrderItemResponse{
					{
						Name:     "Hamburger",
						Quantity: 1,
						Comment:  "No lettuce",
					},
				},
			},
		},
	}
	marshal, _ := json.Marshal(response)
	return string(marshal)
}

func mockFailedDependencyException() port.CustomExceptionInterface {
	return exception.NewFailedDependencyException(errors.New("forced exception"))
}
