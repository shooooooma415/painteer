package postgresql

import "database/sql"

type GroupRepositoryImpl struct {
	DB *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{DB: db}
}

