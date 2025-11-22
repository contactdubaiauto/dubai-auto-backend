package service

import (
	"context"
	"dubai-auto/internal/model"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/utils"
	"strconv"
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

func (s *SocketService) GetUserAvatarAndUsername(userID int) (string, string, error) {
	return s.repo.GetUserAvatarAndUsername(userID)
}

func (s *SocketService) MessageWriteToDatabase(senderUserID int, status bool, msg model.MessageReceived) error {
	err := s.repo.MessageWriteToDatabase(senderUserID, status, msg)
	return err
}

func (s *SocketService) CheckUserExists(userID int) error {
	err := s.repo.CheckUserExists(userID)
	return err
}

func (s *SocketService) SendPushForMessage(senderUserID int, msg model.MessageReceived) error {
	return s.repo.SendPushForMessage(senderUserID, msg)
}

func (s *SocketService) GetActiveAdminsWithChatPermission() ([]int, error) {
	return s.repo.GetActiveAdminsWithChatPermission()
}

func (s *SocketService) GetConversations(userID int) model.Response {
	data, err := s.repo.GetConversations(userID)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: 500,
		}
	}

	return model.Response{
		Data:   data,
		Status: 200,
	}
}

func (s *SocketService) UpsertConversation(userID1 int, userID2 int) error {
	return s.repo.UpsertConversation(userID1, userID2)
}

func (s *SocketService) GetMessages(ctx context.Context, userID int, targetUserID, lastMessageID, limitStr string) model.Response {
	targetUserIDInt, err := strconv.Atoi(targetUserID)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: 400,
		}
	}

	lastID, limit := utils.CheckLastIDLimit(lastMessageID, limitStr)
	data, err := s.repo.GetMessages(ctx, userID, targetUserIDInt, lastID, limit)

	if err != nil {
		return model.Response{
			Error:  err,
			Status: 500,
		}
	}

	return model.Response{
		Data: data,
	}
}
