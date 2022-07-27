package domain

import "capi/errs"

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	//  status -> "1", "0", ""
	FindAll(string) ([]Customer, *errs.AppError)
	FindByID(string) (*Customer, *errs.AppError)
}
