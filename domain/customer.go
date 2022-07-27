package domain

type Customer struct {
	//`` berfungsi untuk penamaan dalam json di postman
	ID          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string `json:"date_of_birth" xml:"dateofbirth"`
	Status      string `json:"status" xml:"status"`
}

//I untuk interface didepan nama
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindByID(string) (*Customer, error)
}
