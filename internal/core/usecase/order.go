package usecase

import (
	"github.com/postech-fiap/production-api/internal/core/domain"
	"github.com/postech-fiap/production-api/internal/core/exception"
	"github.com/postech-fiap/production-api/internal/core/port"
)

type orderUseCase struct {
	orderRepository port.OrderRepositoryInterface
}

func NewOrderUserCase(orderRepository port.OrderRepositoryInterface) port.OrderUseCaseInterface {
	return &orderUseCase{
		orderRepository: orderRepository,
	}
}

func (o *orderUseCase) List() ([]domain.Order, error) {
	orders, err := o.orderRepository.List()
	if err != nil {
		return nil, exception.NewFailedDependencyException(err)
	}
	return orders, nil
}

func (o *orderUseCase) Insert(order *domain.Order) error {
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

func (o *orderUseCase) UpdateStatus(id int64, newStatus domain.Status) error {
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

func (o *orderUseCase) get(id int64) (*domain.Order, error) {
	order, err := o.orderRepository.Get(id)
	if err != nil {
		return nil, exception.NewFailedDependencyException(err)
	}
	return order, nil
}
