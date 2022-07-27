package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "User1", "City1", "ZipCode1", "2022-01-01", "1"},
		{"2", "User2", "City2", "ZipCode2", "2022-01-01", "1"},
		{"3", "User3", "City3", "ZipCode3", "2022-01-01", "1"},
		{"4", "User4", "City4", "ZipCode4", "2022-01-01", "1"},
	}

	return CustomerRepositoryStub{customer: customers}
}
