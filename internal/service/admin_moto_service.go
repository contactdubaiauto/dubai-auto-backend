package service

import (
	"context"
	"errors"
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

// ModerateMotorcycle updates the moderation status of a motorcycle.
// If status is declined (3), sends push notification and inserts notification record.
func (s *AdminService) ModerateMotorcycle(ctx context.Context, req *model.ModerateItemRequest) model.Response {
	if req == nil {
		return model.Response{Error: errors.New("invalid request data"), Status: http.StatusBadRequest}
	}

	userID, err := s.repo.ModerateMotorcycle(ctx, req.ID, req.Status)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	// If declined, send notification
	if req.Status == 3 {
		s.sendModerationNotification(ctx, userID, "moto", req.ID, req.Title, req.Description)
	}

	return model.Response{Data: model.Success{Message: "Motorcycle moderation status updated successfully"}}
}
