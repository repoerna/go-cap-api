package dto

type CustomerResponse struct {
	ID          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zipCode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}