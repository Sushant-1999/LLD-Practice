package service

//* This contains the business logic of the application : service layers

//* core business logic of the Splitwise

// splitwise/usecase/group_usecase.go
//* Service layers uses repo here it is abstracted in the service layers
//* Actual db implementation done at repository layer

import (
	"splitwise/models"
	"splitwise/repository"
)

type GroupService interface {
	CreateGroup(group models.Group) error
	GetGroups() ([]models.Group, error)
}

type groupService struct {
	groupRepo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) GroupService {
	return &groupService{
		groupRepo: repo,
	}
}

func (uc *groupService) CreateGroup(group models.Group) error {
	return uc.groupRepo.CreateGroup(group)
}

func (uc *groupService) GetGroups() ([]models.Group, error) {
	return uc.groupRepo.GetGroups()
}
