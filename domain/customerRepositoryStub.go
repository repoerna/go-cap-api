package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	Customers := []Customer{
		{"1", "Bruyne", "Banten", "15540", "2021-03-09", "1"},
		{"3", "Brown", "Surabaya", "15560", "2021-03-09", "1"},
		{"15", "Brown", "Surabaya", "15560", "2021-03-09", "1"},
		{"5", "Brown", "Surabaya", "15560", "2021-03-09", "1"},
	}

	return CustomerRepositoryStub{Customer: Customers}
}