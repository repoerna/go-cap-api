package domain

import (
	"capi/dto"
	"capi/errs"
)

type Customer struct {
	//`` berfungsi untuk penamaan dalam json di postman
	ID          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string `json:"date_of_birth" xml:"dateofbirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

//I untuk interface didepan nama
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindByID(string) (*Customer, *errs.AppErr)
}

func (c Customer) convertStatusName()string{
	statusName := "active"
	if c.Status == "0"{
		statusName = "inactive"
	}
	return statusName
}
func (c Customer) ToDTO() dto.CustomerResponse{
	
	return dto.CustomerResponse{
		ID: c.ID,
		Name: c.Name,
		DateOfBirth: c.DateOfBirth,
		City: c.City,
		ZipCode: c.ZipCode,
		Status: c.convertStatusName(),
	}
}
