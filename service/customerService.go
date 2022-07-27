package service

import (
	"capi/domain"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomerByID(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindByID(customerID)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
