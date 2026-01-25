package service

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"dubai-auto/internal/model"
	"dubai-auto/pkg/files"
)

func (s *AdminService) GetComtrans(ctx context.Context, limit, lastID int, moderationStatus string) model.Response {
	moderationStatusInt, err := strconv.Atoi(moderationStatus)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}
	list, err := s.repo.GetComtrans(ctx, limit, lastID, moderationStatusInt)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return model.Response{Data: list}
}

func (s *AdminService) GetComtranByID(ctx context.Context, id int) model.Response {
	com, err := s.repo.GetComtranByID(ctx, id)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusNotFound}
	}
	return model.Response{Data: com}
}

func (s *AdminService) DeleteComtran(ctx context.Context, id int, dir string) model.Response {
	err := s.repo.DeleteComtran(ctx, id)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	_ = files.RemoveFolder(dir)
	return model.Response{Data: model.Success{Message: "Comtran deleted successfully"}}
}

// ModerateComtran updates the moderation status of a comtran.
// If status is declined (3), sends push notification and inserts notification record.
func (s *AdminService) ModerateComtran(ctx context.Context, req *model.ModerateItemRequest) model.Response {
	if req == nil {
		return model.Response{Error: errors.New("invalid request data"), Status: http.StatusBadRequest}
	}

	userID, err := s.repo.ModerateComtran(ctx, req.ID, req.Status)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	// If declined, send notification
	if req.Status == 3 {
		s.sendModerationNotification(ctx, userID, "comtran", req.ID, req.Title, req.Description)
	}

	return model.Response{Data: model.Success{Message: "Comtran moderation status updated successfully"}}
}
