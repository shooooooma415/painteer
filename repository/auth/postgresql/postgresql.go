package postgresql

import (
	"database/sql"
	"fmt"
	"painteer/model"
)

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}

func (q *AuthRepositoryImpl) CreateUserQuery(createUser model.CreateUser) (*model.UserId, error) {
	query := `
		INSERT INTO users (name,icon,auth_id) 
		VALUES ($1,$2,$3)
		RETURNING id, name
	`

	var userId model.UserId
	var userName model.UserName

	err := q.DB.QueryRow(query, createUser).Scan(&userId, &userName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &userId, nil
}
