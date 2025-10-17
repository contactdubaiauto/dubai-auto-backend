package service

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
)

type SocketService struct {
	repo *repository.SocketRepository
}

func NewSocketService(repo *repository.SocketRepository) *SocketService {
	return &SocketService{repo}
}

func (s *SocketService) UpdateUserStatus(userID int, status bool) error {
	err := s.repo.UpdateUserStatus(userID, status)
	return err
}

func (s *SocketService) GetNewMessages(userID int) ([]model.UserMessage, error) {
	messages, err := s.repo.GetNewMessages(userID)
	return messages, err
}

func (s *SocketService) GetUserAvatar(userID int) string {
	avatar, err := s.repo.GetUserAvatar(userID)

	if err != nil {
		return ""
	}

	return avatar
}
