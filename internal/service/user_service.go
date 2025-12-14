package service

import (
	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/repo"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(r repo.UserRepo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(u *dto.User) error {
	if u.Name == nil || *u.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if u.OwnerID != nil {
		if _, err := s.repo.GetByID(*u.OwnerID); err != nil {
			return fmt.Errorf("owner not found: %w", err)
		}
	}

	if u.Email != nil && *u.Email != "" {
		exists, err := s.repo.ExistsByEmail(*u.Email)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("email already exists")
		}
	}

	return s.repo.Create(u)
}

func (s *UserService) GetByID(id uuid.UUID) (*dto.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(u *dto.User) error {
	if u.ID == uuid.Nil {
		return errors.New("не указан ID пользователя")
	}

	if u.OwnerID != nil {
		if _, err := s.repo.GetByID(*u.OwnerID); err != nil {
			return fmt.Errorf("owner not found: %w", err)
		}
		if err := s.repo.UpdateOwner(u.ID, *u.OwnerID); err != nil {
			return err
		}
	}

	if u.Email != nil && *u.Email != "" {
		exists, err := s.repo.ExistsByEmail(*u.Email)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("email already exists")
		}
		if err := s.repo.UpdateEmail(u.ID, *u.Email); err != nil {
			return err
		}
	}

	if u.Name != nil && *u.Name != "" {
		if err := s.repo.UpdateName(u.ID, *u.Name); err != nil {
			return err
		}
	}

	return nil
}

func (s *UserService) ListOwnedUsers(ownerID uuid.UUID) ([]*dto.User, error) {
	if ownerID == uuid.Nil {
		return nil, errors.New("не указан ownerID")
	}
	return s.repo.ListByOwner(ownerID)
}

func (s *UserService) MakePermanent(tempUserID uuid.UUID) error {
	if tempUserID == uuid.Nil {
		return errors.New("не указан tempUserID")
	}

	user, err := s.repo.GetByID(tempUserID)
	if err != nil {
		return err
	}

	if user.OwnerID != nil && *user.OwnerID != uuid.Nil {
		return errors.New("пользователь уже сохранён")
	}

	ownerID, err := s.repo.GetOwnerForTempUser(tempUserID)
	if err != nil {
		return err
	}

	return s.repo.UpdateOwner(tempUserID, ownerID)
}

func (s *UserService) Delete(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("не указан ID пользователя")
	}
	return s.repo.Delete(id)
}

func (s *UserService) DeleteUserWithOwned(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("не указан ID пользователя")
	}

	// удаляем всех, кого владелец создал
	if err := s.repo.DeleteByOwner(id); err != nil {
		return err
	}

	// удаляем самого пользователя
	return s.repo.Delete(id)
}
