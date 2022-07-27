package service

import "capi/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerByID(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*domain.Customer, error) {
	return s.repo.FindByID(customerID)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
