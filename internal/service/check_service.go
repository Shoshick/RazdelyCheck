package service

import (
	"errors"
	"time"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"

	"github.com/google/uuid"
)

type CheckService struct {
	checkRepo repo.CheckRepo
	groupRepo repo.GroupRepo
	itemRepo  repo.ItemRepo
}

func NewCheckService(cRepo repo.CheckRepo, gRepo repo.GroupRepo) *CheckService {
	return &CheckService{
		checkRepo: cRepo,
		groupRepo: gRepo,
	}
}

func (s *CheckService) Create(userID uuid.UUID, groupID *uuid.UUID, totalSum int64) (*dto.Check, error) {
	var gid uuid.UUID

	if groupID == nil {

		gid = uuid.New()
		group := &dto.Group{ID: gid}
		if err := s.groupRepo.Create(group); err != nil {
			return nil, err
		}

		if err := s.groupRepo.AddUserToGroup(userID, gid); err != nil {
			return nil, err
		}
	} else {

		gid = *groupID
		exists, err := s.groupRepo.ExistsUserInGroup(userID, gid)
		if err != nil {
			return nil, err
		}
		if !exists {
			if err := s.groupRepo.AddUserToGroup(userID, gid); err != nil {
				return nil, err
			}
		}
	}

	c := &dto.Check{
		ID:        uuid.New(),
		UserID:    userID,
		GroupID:   gid,
		TotalSum:  totalSum,
		CreatedAt: time.Now(),
	}

	if err := s.checkRepo.Create(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CheckService) Delete(id, currentUserID uuid.UUID, isAdmin bool) error {
	c, err := s.checkRepo.GetByID(id)
	if err != nil {
		return err
	}

	if c.UserID != currentUserID && !isAdmin {
		return errors.New("only owner or admin can delete check")
	}

	return s.checkRepo.Delete(id)
}

func (s *CheckService) UpdateTotalSum(userID, checkID uuid.UUID, totalSum int64) (*dto.Check, error) {

	if totalSum < 0 {
		return nil, errors.New("total sum cannot be negative")
	}

	check, err := s.checkRepo.GetByID(checkID)
	if err != nil {
		return nil, err
	}

	if check.UserID != userID {
		return nil, errors.New("forbidden")
	}

	err = s.checkRepo.UpdateTotalSum(checkID)
	if err != nil {
		return nil, err
	}

	check.TotalSum = totalSum
	return check, nil
}

func (s *CheckService) ListByUser(userID uuid.UUID) ([]*dto.Check, error) {
	return s.checkRepo.ListByUserID(userID)
}

func (s *CheckService) GetItemsFromCheck(userID, checkID uuid.UUID) ([]*dto.Item, error) {

	check, err := s.checkRepo.GetByID(checkID)
	if err != nil {
		return nil, err
	}
	if check.UserID != userID {
		return nil, errors.New("forbidden")
	}

	items, err := s.itemRepo.ListByCheckID(checkID)
	if err != nil {
		return nil, err
	}

	return items, nil
}
