package service

import (
	"Brachium/pkg/entity"
	"Brachium/pkg/repository"
)

type CheckListService struct {
	repo repository.CheckList
}

func NewCheckListService(repo repository.CheckList) *CheckListService {
	return &CheckListService{repo: repo}
}

func (s *CheckListService) Create(userId int, list entity.CheckList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *CheckListService) GetAll(userId int) ([]entity.CheckList, error) {
	return s.repo.GetAll(userId)
}

func (s *CheckListService) GetById(userId, listId int) (entity.CheckList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *CheckListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *CheckListService) Update(userId, listId int, input entity.UpdateCheckListInput) error {
	return s.repo.Update(userId, listId, input)
}
