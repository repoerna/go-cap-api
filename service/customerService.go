package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	// * add process here	
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(CustomerID string) (*dto.CustomerResponse, *errs.AppErr) {
	cust, err := s.repository.FindByID(CustomerID)
	if err != nil {
		return nil, err
	}

	response := cust.ToDTO()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}