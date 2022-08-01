package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"
const ACCESS_TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

type AccessTokenClaims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	Role       string   `json:"role"`
	jwt.StandardClaims
}

type AuthToken struct {
	token *jwt.Token
}

func (c AccessTokenClaims) IsUserRole() bool {
	return c.Role == "user"
}

func (c AccessTokenClaims) IsValidCustomerId(customerId string) bool {
	return c.CustomerId == customerId
}

func (c AccessTokenClaims) IsValidAccountId(accountId string) bool {
	if accountId != "" {
		accountFound := false
		for _, a := range c.Accounts {
			if a == accountId {
				accountFound = true
				break
			}
		}
		return accountFound
	}
	return true
}

func (c AccessTokenClaims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	if c.IsValidCustomerId(urlParams["customer_id"]) {
		return false
	}

	if !c.IsValidAccountId(urlParams["account_id"]) {
		return false
	}
	return true
}

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppErr)
}

func (l Login) ClaimsForAccessToken() AccessTokenClaims {
	if l.Accounts.Valid && l.CustomerId.Valid {
		return l.claimsForUser()
	} else {
		return l.claimsForAdmin()
	}
}

func (l Login) claimsForUser() AccessTokenClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	return AccessTokenClaims{
		CustomerId: l.CustomerId.String,
		Accounts:   accounts,
		Username:   l.Username,
		Role:       l.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

func (l Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		Username: l.Username,
		Role:     l.Role,
	}
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func (t AuthToken) NewAccessToken() (string, *errs.AppErr) {
	signedString, err := t.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Error("Failed while signing access token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}
