package service

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/pkg/firebase"
	"strconv"
	"time"
)

type SocketService struct {
	repo            *repository.SocketRepository
	firebaseService *firebase.FirebaseService
}

func NewSocketService(repo *repository.SocketRepository, firebaseService *firebase.FirebaseService) *SocketService {
	return &SocketService{repo, firebaseService}
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

func (s *SocketService) MessageWriteToDatabase(senderUserID int, status bool, msg model.MessageReceived) error {
	err := s.repo.MessageWriteToDatabase(senderUserID, status, msg)

	if err != nil {
		return err
	}

	// todo: send push notification to the user
	token, err := s.repo.GetUserToken(msg.TargetUserID)

	if err != nil {
		return err
	}

	messageData := map[string]string{
		"message": msg.Message,
		"type":    strconv.Itoa(msg.Type),
		"time":    msg.Time.Format(time.RFC3339),
	}

	s.firebaseService.SendToToken(token, "New Message", "You have a new message", messageData)

	return nil
}

func (s *SocketService) CheckUserExists(userID int) error {
	err := s.repo.CheckUserExists(userID)
	return err
}
