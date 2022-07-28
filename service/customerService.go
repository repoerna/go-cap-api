package service

import "capi/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	// * add process here
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(CustomerID string) (*domain.Customer, error) {
	return s.repository.FindByID(CustomerID)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}