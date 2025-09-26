package service

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"errors"

	"github.com/google/uuid"
)

type GroupService struct {
	groupRepo repo.GroupRepo
	userRepo  repo.UserRepo
}

func NewGroupService(gr repo.GroupRepo, ur repo.UserRepo) *GroupService {
	return &GroupService{
		groupRepo: gr,
		userRepo:  ur,
	}
}

func (s *GroupService) Create(g *dto.Group) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.New()
	}
	return s.groupRepo.Create(g)
}

func (s *GroupService) GetByID(id uuid.UUID) (*dto.Group, error) {
	if id == uuid.Nil {
		return nil, errors.New("не указан ID группы")
	}
	return s.groupRepo.GetByID(id)
}

func (s *GroupService) List() ([]*dto.Group, error) {
	return s.groupRepo.List()
}

func (s *GroupService) Update(g *dto.Group) error {
	if g.ID == uuid.Nil {
		return errors.New("не указан ID группы")
	}
	return s.groupRepo.Update(g)
}

func (s *GroupService) Delete(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("не указан ID группы")
	}
	return s.groupRepo.Delete(id)
}

func (s *GroupService) AddUserToGroup(userID, groupID uuid.UUID) error {
	if groupID == uuid.Nil || userID == uuid.Nil {
		return errors.New("не указан userID или groupID")
	}

	if _, err := s.groupRepo.GetByID(groupID); err != nil {
		return err
	}

	if _, err := s.userRepo.GetByID(userID); err != nil {
		return err
	}

	exists, err := s.groupRepo.ExistsUserInGroup(userID, groupID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("пользователь уже состоит в группе")
	}

	return s.groupRepo.AddUserToGroup(userID, groupID)
}

func (s *GroupService) RemoveUserFromGroup(userID, groupID uuid.UUID) error {
	if groupID == uuid.Nil || userID == uuid.Nil {
		return errors.New("не указан userID или groupID")
	}
	return s.groupRepo.RemoveUserFromGroup(userID, groupID)
}

func (s *GroupService) ListGroupsByUser(userID uuid.UUID) ([]*dto.Group, error) {
	if userID == uuid.Nil {
		return nil, errors.New("не указан userID")
	}
	return s.groupRepo.ListByUser(userID)
}

func (s *GroupService) ListUsersByGroup(groupID uuid.UUID) ([]*dto.User, error) {
	if groupID == uuid.Nil {
		return nil, errors.New("не указан groupID")
	}
	return s.groupRepo.ListByGroup(groupID)
}
