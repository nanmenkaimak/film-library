package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	RoleID      int       `json:"role_id" db:"role_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

