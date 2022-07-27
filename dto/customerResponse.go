package dto

type CustomerResponse struct {
	ID          string `json:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
