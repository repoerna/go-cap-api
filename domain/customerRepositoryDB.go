package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (s CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var query string
	// var rows *sql.Rows
	var err error

	var customers []Customer

	if status == "" {
		query = "select * from customers"
		err = s.db.Select(&customers, query)
	} else {
		query = "select * from customers where status = $1"
		err = s.db.Select(&customers, query, status)
	}

	if err != nil {
		logger.Error("error fetch data to customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// 	err = sqlx.StructScan(rows, &customers)
	// 	if err != nil {
	// 		logger.Error("error scanning customer data " + err.Error())
	// 		return nil, errs.NewUnexpectedError("unexpected database error")
	// 	}

	return customers, nil

}

func (s CustomerRepositoryDB) FindByID(id string) (*Customer, *errs.AppError) {

	query := "select * from customers where customer_id = $1"

	// row := s.db.QueryRow(query, id)
	var c Customer

	err := s.db.Get(&c, query, id)

	// err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(err.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("error scanning customer data ", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}

	}

	return &c, nil

}
