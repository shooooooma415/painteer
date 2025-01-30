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

	err := q.DB.QueryRow(
		query,
		createUser.UserName,
		createUser.Icon,
		createUser.AuthId,
	).Scan(
		&resultUser.UserId,
		&resultUser.UserName,
		&resultUser.AuthId,
		&resultUser.Icon,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &resultUser, nil
}

func (q *AuthRepositoryImpl) FindUserByAuthID(authId model.AuthId) (*model.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE auth_id = $1
	`

	var resultUser model.User

	err := q.DB.QueryRow(
		query,
		authId,
	).Scan(
		&resultUser.UserId,
		&resultUser.UserName,
		&resultUser.AuthId,
		&resultUser.Icon,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &resultUser, nil
}

func (q *AuthRepositoryImpl) FindUserByUserID(userId model.UserId) (*model.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE user_id = $1
	`

	var resultUser model.User

	err := q.DB.QueryRow(
		query,
		userId,
	).Scan(
		&resultUser.UserId,
		&resultUser.UserName,
		&resultUser.AuthId,
		&resultUser.Icon,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &resultUser, nil
}
