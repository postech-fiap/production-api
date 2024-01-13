package usecase

import (
	"errors"
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"time"
)

const mockOrderID int64 = 1

func mockOrder(status domain.Status) *domain.Order {
	return &domain.Order{
		ID:          mockOrderID,
		Status:      status,
		Number:      "0001",
		CreatedDate: time.Date(2024, 01, 12, 0, 0, 0, 0, time.UTC),
		Items: []domain.Item{
			{
				Name:     "Hamburger",
				Quantity: 1,
				Comment:  "No lettuce",
			},
		},
	}
}

func mockOrders() []domain.Order {
	return []domain.Order{
		*mockOrder(domain.PENDING),
	}
}

func mockError() error {
	return errors.New("forced exception")
}

func mockFailedDependencyException() error {
	return exception.NewFailedDependencyException(errors.New("forced exception"))
}
