package service

import (
	"context"
	"net/http"
	"strconv"

	"dubai-auto/internal/model"
	"dubai-auto/pkg/files"
)

func (s *AdminService) GetMotorcycles(ctx context.Context, limit, lastID int, moderationStatus string) model.Response {
	moderationStatusInt, err := strconv.Atoi(moderationStatus)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}
	list, err := s.repo.GetMotorcycles(ctx, limit, lastID, moderationStatusInt)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return model.Response{Data: list}
}

func (s *AdminService) GetMotorcycleByID(ctx context.Context, id int) model.Response {
	m, err := s.repo.GetMotorcycleByID(ctx, id)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusNotFound}
	}
	return model.Response{Data: m}
}

func (s *AdminService) DeleteMotorcycle(ctx context.Context, id int, dir string) model.Response {
	err := s.repo.DeleteMotorcycle(ctx, id)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	_ = files.RemoveFolder(dir)
	return model.Response{Data: model.Success{Message: "Motorcycle deleted successfully"}}
}
