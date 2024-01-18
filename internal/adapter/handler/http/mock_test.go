package http

import (
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type orderUseCaseMock struct {
	mock.Mock
}

func (o *orderUseCaseMock) List() ([]domain.Order, error) {
	args := o.Called()
	res := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}
	if res != nil {
		return res.([]domain.Order), nil
	}
	return nil, nil
}

func (o *orderUseCaseMock) Insert(order *domain.Order) error {
	args := o.Called(order)
	return args.Error(0)
}

func (o *orderUseCaseMock) UpdateStatus(id int64, newStatus domain.Status) error {
	args := o.Called(id, newStatus)
	return args.Error(0)
}
