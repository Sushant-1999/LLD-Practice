package repository

import "splitwise/models"

// interface for abstracting database operations
type GroupRepository interface {
	CreateGroup(group models.Group) error
	GetGroups() ([]models.Group, error)
}

type groupRepository struct {
	groups []models.Group
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{
		groups: []models.Group{},
	}
}

func (r *groupRepository) CreateGroup(group models.Group) error {
	r.groups = append(r.groups, group)
	return nil
}

func (r *groupRepository) GetGroups() ([]models.Group, error) {
	return r.groups, nil
}
