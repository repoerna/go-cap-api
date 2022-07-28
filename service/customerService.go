package service

import (
	"capi/domain"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(string) (*domain.Customer, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	// * add process here
	
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(CustomerID string) (*domain.Customer, *errs.AppErr) {
	return s.repository.FindByID(CustomerID)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}