package usecase

import (
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type orderRepositoryMock struct {
	mock.Mock
}

func (o *orderRepositoryMock) List() ([]domain.Order, error) {
	args := o.Called()
	res := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}
	return res.([]domain.Order), nil
}

func (o *orderRepositoryMock) Get(id int64) (*domain.Order, error) {
	args := o.Called(id)
	res := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}
	if res != nil {
		return res.(*domain.Order), nil
	}
	return nil, nil
}

func (o *orderRepositoryMock) Insert(order *domain.Order) error {
	args := o.Called(order)
	return args.Error(0)
}

func (o *orderRepositoryMock) UpdateStatus(order *domain.Order) error {
	args := o.Called(order)
	return args.Error(0)
}

type orderQueuePublisherMock struct {
	mock.Mock
}

func (o *orderQueuePublisherMock) PublishNewStatus(order *domain.Order) error {
	args := o.Called(order)
	return args.Error(0)
}
