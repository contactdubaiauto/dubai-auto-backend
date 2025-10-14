package service

import "dubai-auto/internal/repository"

type SocketService struct {
	repo *repository.SocketRepository
}

func NewSocketService(repo *repository.SocketRepository) *SocketService {
	return &SocketService{repo}
}
