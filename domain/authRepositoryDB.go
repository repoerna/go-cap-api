package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryDb struct {
	db *sqlx.DB
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, *errs.AppErr) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, string_agg(a.account_id::varchar(255), ',') as account_numbers FROM users u
					LEFT JOIN accounts a ON a.customer_id = u.customer_id
					WHERE username = $1 and password = $2
					GROUP BY username, a.customer_id, role`
	err := d.db.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &login, nil
}
