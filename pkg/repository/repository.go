package repository

import (
	"Brachium/pkg/entity"
	"database/sql"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
}

type CheckList interface {
	Create(userId int, list entity.CheckList) (int, error)
	GetAll(userId int) ([]entity.CheckList, error)
	GetById(userId, listId int) (entity.CheckList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entity.UpdateCheckListInput) error
}

type Task interface {
}

type Repository struct {
	Authorization
	CheckList
	Task
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		CheckList:     NewCheckListPostgres(db),
	}
}
