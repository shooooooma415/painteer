package postgresql

import (
	"database/sql"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}