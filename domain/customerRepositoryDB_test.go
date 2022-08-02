package domain

import (
	"capi/errs"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// func TestNewCustomerRepositoryDB(t *testing.T) {
// 	type args struct {
// 		client *sqlx.DB
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want CustomerRepositoryDB
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewCustomerRepositoryDB(tt.args.client); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewCustomerRepositoryDB() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func NewMock() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	// defer db.Close()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	return sqlxDB, mock
}

func TestCustomerRepositoryDB_FindAll(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		args    args
		want    []Customer
		wantErr *errs.AppErr
	}{
		// TODO: Add test cases.
		{
			"succcess get data all customer",
			args{""},
			[]Customer{
				{"1", "User1", "Jakarta", "12345", "2022-01-01", "1"},
				{"2", "User2", "Surabaya", "67890", "2022-01-01", "1"},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := NewMock()
			repo := NewCustomerRepositoryDB(db)

			rows := mock.NewRows([]string{"customer_id", "name", "city", "zipcode", "date_of_birth", "status"}).AddRow(
				"1", "User1", "Jakarta", "12345", "2022-01-01", "1").AddRow("2", "User2", "Surabaya", "67890", "2022-01-01", "1")

			mock.ExpectQuery(`select \* from customers`).WillReturnRows(rows)
			got, got1 := repo.FindAll(tt.args.status)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerRepositoryDB.FindAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("CustomerRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestCustomerRepositoryDB_FindAll_should_return_error(t *testing.T) {
	type args struct {
		status string
	}

	tests := []struct {
		name    string
		args    args
		want    []Customer
		wantErr *errs.AppErr
	}{
		// TODO: Add test cases.
		{
			"succcess get data all customer",
			args{""},
			nil,
			errs.NewUnexpectedError("unexpected database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := NewMock()
			repo := NewCustomerRepositoryDB(db)

			mock.ExpectQuery(`select \* from customers`).WillReturnError(errors.New(""))
			got, got1 := repo.FindAll(tt.args.status)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerRepositoryDB.FindAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("CustomerRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}

}

func TestCustomerRepositoryDB_FindByID(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Customer
		want1  *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CustomerRepositoryDB{
				db: tt.fields.db,
			}
			got, got1 := s.FindByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerRepositoryDB.FindByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CustomerRepositoryDB.FindByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
