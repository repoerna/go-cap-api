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
		{"3", "Owen", "Surabaya", "18960", "2021-06-09", "1"},
		{"15", "Towel", "Serang", "15546", "2021-05-09", "1"},
		{"5", "Brown", "Jakarta", "67560", "2021-04-09", "1"},
	}

	return CustomerRepositoryStub{Customer: Customers}
}