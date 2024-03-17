package repository

import (
	"github.com/google/uuid"
	"github.com/nanmenkaimak/film_library/internal/entity"
)

func (r *Repo) CreateUser(newUser entity.User) (uuid.UUID, error) {
	var userID uuid.UUID
	query := `insert into users (first_name, last_name, email, password, role_id)
				values ($1, $2, $3, $4, $5) returning id`
	err := r.main.Db.QueryRowx(query, newUser.FirstName, newUser.LastName,
		newUser.Email, newUser.Password, newUser.RoleID).Scan(&userID)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

func (r *Repo) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `select * from users where email = $1`
	err := r.replica.Db.QueryRowx(query, email).StructScan(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil
}