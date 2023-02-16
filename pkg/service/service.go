package service

import (
	"Brachium/pkg/entity"
	"Brachium/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GeneratedToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	CheckList
	Task
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		CheckList:     NewCheckListService(repo.CheckList),
	}
}
