package domain

import "capi/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindByID(string) (*Customer, *errs.AppError)
}
