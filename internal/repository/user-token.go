package repository

import (
	"github.com/nanmenkaimak/film_library/internal/entity"
)

func (r *Repo) CreateUserToken(userToken entity.UserToken) error {
	query := `insert into user_tokens (user_id, token, refresh_token, expires_at)
				values ($1, $2, $3, $4)`
	err := r.main.Db.QueryRowx(query, userToken.UserID, userToken.Token, userToken.RefreshToken, userToken.ExpiresAt).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateUserToken(userToken entity.UserToken) error {
	query := `update user_tokens set token = $1 where user_id = $2`

	err := r.main.Db.QueryRowx(query, userToken.Token, userToken.UserID).Err()
	if err != nil {
		return err
	}
	return nil
}
