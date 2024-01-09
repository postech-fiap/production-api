package port

import "github.com/postech-fiap/producao/internal/core/domain"

type OrderRepositoryInterface interface {
	List() ([]domain.Order, error)
	Get(id int64) (*domain.Order, error)
	Insert(order *domain.Order) error
	UpdateStatus(order *domain.Order) error
}

type OrderServiceInterface interface {
	List() ([]domain.Order, error)
	Insert(order *domain.Order) error
	UpdateStatus(id int64, newStatus domain.Status) error
}
