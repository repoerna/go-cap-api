package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppErr)
}

type DefaultAuthService struct {
	repository domain.AuthRepository
	// rolePermissions domain.RolePermissions
}

func NewAuthService(repository domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repository}
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppErr) {
	var appErr *errs.AppErr
	var login *domain.Login

	if login, appErr = s.repository.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)

	var accessToken string
	if accessToken, appErr = authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	}

	return &dto.LoginResponse{AccessToken: accessToken}, nil
}
