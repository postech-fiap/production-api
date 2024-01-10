package service

import (
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"github.com/postech-fiap/production-api/internal/core/port"
)

type orderService struct {
	orderRepository port.OrderRepositoryInterface
}

func NewOrderService(orderRepository port.OrderRepositoryInterface) port.OrderServiceInterface {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (o *orderService) List() ([]domain.Order, error) {
	orders, err := o.orderRepository.List()
	if err != nil {
		return nil, exception.NewFailedDependencyException(err)
	}
	return orders, nil
}

func (o *orderService) Insert(order *domain.Order) error {
	existentOrder, err := o.get(order.ID)
	if err != nil {
		return err
	}
	if existentOrder != nil {
		return exception.NewInvalidDataException("order already exists", nil)
	}

	order.Status = domain.PENDING
	err = o.orderRepository.Insert(order)
	if err != nil {
		return exception.NewFailedDependencyException(err)
	}
	return nil
}

func (o *orderService) UpdateStatus(id int64, newStatus domain.Status) error {
	order, err := o.get(id)
	if err != nil {
		return err
	}
	if order == nil {
		return exception.NewNotFoundException("order not found", nil)
	}

	isValidStatus := order.IsValidStatus(newStatus)
	if !isValidStatus {
		return exception.NewInvalidDataException("invalid status transition", nil)
	}

	order.Status = newStatus

	err = o.orderRepository.UpdateStatus(order)
	if err != nil {
		return exception.NewFailedDependencyException(err)
	}
	return nil
}

func (o *orderService) get(id int64) (*domain.Order, error) {
	order, err := o.orderRepository.Get(id)
	if err != nil {
		return nil, exception.NewFailedDependencyException(err)
	}
	return order, nil
}
