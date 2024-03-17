package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type GenerateTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RenewTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewTokenResponse struct {
	Token string `json:"token"`
}

type JwtUserToken struct {
	Token        string
	RefreshToken string
}

type JwtRenewToken struct {
	Token string
}

type MyCustomClaims struct {
	RoleID int       `json:"role_id"`
	UserId uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type RegisterUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type DeleteActorRequest struct {
	ID uuid.UUID
}

type DeleteFilmRequest struct {
	ID uuid.UUID
}

type CreateFilmRequest struct {
	ID uuid.UUID
}

type GetFilmsRequest struct {
	Sorting string
	Name    string
}

type GetActorRequest struct {
	Name    string
}
