package postgresql

import (
	"database/sql"
	"painteer/model"
)

type PostingRepositoryImpl struct {
	DB *sql.DB
}

func NewPostingRepository(db *sql.DB) *PostingRepositoryImpl {
	return &PostingRepositoryImpl{DB: db}
}
