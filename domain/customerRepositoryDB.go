package domain

import (
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

func (s CustomerRepositoryDB) FindAll() ([]Customer, error) {
	// connStr := "user=postgres dbname=banking sslmode=disable"
	// connStr := "postgres://postgres:postgres@localhost/banking?sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	query := "select * from customers"

	rows, err := s.db.Query(query)

	if err != nil {
		log.Println("error fetch data to customer table ", err.Error())
		return nil, err
	}

	var customers []Customer
	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("error scanning customer data ", err.Error())
		}

		customers = append(customers, c)
	}

	return customers, nil

}
