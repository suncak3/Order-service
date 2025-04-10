package usecase

import (
	"order-service/domain"
	"order-service/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService() *Service {
	return &Service{repository: repository.NewRepository()}
}

func (s *Service) GetAllOrders() ([]domain.Order, error) {
	return s.repository.GetAllOrders()
}

func (s *Service) GetOrderByID(id uint) (*domain.Order, error) {
	return s.repository.GetOrderByID(id)
}

func (s *Service) CreateOrder(Order domain.Order) (*domain.Order, error) {
	return s.repository.CreateOrder(Order)
}

func (s *Service) UpdateOrder(Order domain.Order) (*domain.Order, error) {
	return s.repository.UpdateOrder(Order)
}

func (s *Service) DeleteOrder(id uint) error {
	return s.repository.DeleteOrder(id)
}
