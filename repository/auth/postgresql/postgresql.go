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

func (q *AuthRepositoryImpl) CreateUser(createUser model.CreateUser) (*model.User, error) {
	query := `
		INSERT INTO users (name, icon, auth_id) 
		VALUES ($1, $2, $3)
		RETURNING *
	`

	var resultUser model.User

	err := q.DB.QueryRow(query, createUser.UserName, createUser.Icon, createUser.AuthId).Scan(&resultUser.UserId, &resultUser.UserName,&resultUser.AuthId,&resultUser.Icon)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &resultUser, nil
}

func (q *AuthRepositoryImpl) FindUserByID(authId model.AuthId) (*model.UserId, error) {
	query := `
		SELECT user_id, auth_id
		FROM user
		WHERE auth_id = $1
	`

	var resultUserId model.UserId
	var resultAuthId model.AuthId

	err := q.DB.QueryRow(query, authId).Scan(&resultUserId, &resultAuthId)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	if resultAuthId != authId {
		return nil, fmt.Errorf("auth_id mismatch: expected %v, got %v", authId, resultAuthId)
	}
	return &resultUserId, nil
}

