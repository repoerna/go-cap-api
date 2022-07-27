package domain

import (
	"capi/errs"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	db *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:postgres@localhost/banking?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (s CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	query := "select * from customers"

	rows, err := s.db.Query(query)

	if err != nil {
		log.Println("error fetch data to customer table ", err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	var customers []Customer
	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("error scanning customer data ", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

		customers = append(customers, c)
	}

	return customers, nil

}

func (s CustomerRepositoryDB) FindByID(id string) (*Customer, *errs.AppError) {

	query := "select * from customers where customer_id = $1"

	row := s.db.QueryRow(query, id)
	var c Customer

	err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("error scanning customer data ", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

	}

	return &c, nil

}
