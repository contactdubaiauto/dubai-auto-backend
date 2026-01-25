package service

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
	"dubai-auto/pkg/files"
)

// Vehicles (admin) service methods.

func (s *AdminService) GetVehicles(ctx *fasthttp.RequestCtx, limit, lastID int, moderationStatus string) model.Response {
	moderationStatusInt, err := strconv.Atoi(moderationStatus)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}

	vehicles, err := s.repo.GetVehicles(ctx, limit, lastID, moderationStatusInt)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: vehicles}
}

func (s *AdminService) GetVehicleByID(ctx *fasthttp.RequestCtx, vehicleID int) model.Response {
	vehicle, err := s.repo.GetVehicleByID(ctx, vehicleID)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusNotFound}
	}
	return model.Response{Data: vehicle}
}

func (s *AdminService) CreateVehicle(ctx *fasthttp.RequestCtx, req *model.AdminCreateVehicleRequest) model.Response {
	id, err := s.repo.CreateVehicle(ctx, req)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusBadRequest}
	}
	return model.Response{Data: model.SuccessWithId{Id: id, Message: "Vehicle created successfully"}}
}

func (s *AdminService) UpdateVehicleStatus(ctx *fasthttp.RequestCtx, vehicleID int, req *model.AdminUpdateVehicleStatusRequest) model.Response {

	if req == nil {
		return model.Response{Error: errors.New("invalid request data"), Status: http.StatusBadRequest}
	}

	err := s.repo.UpdateVehicleStatus(ctx, vehicleID, req.Status)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	return model.Response{Data: model.Success{Message: "Vehicle status updated successfully"}}
}

func (s *AdminService) DeleteVehicle(ctx *fasthttp.RequestCtx, vehicleID int, dir string) model.Response {
	err := s.repo.DeleteVehicle(ctx, vehicleID)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}
	_ = files.RemoveFolder(dir)
	return model.Response{Data: model.Success{Message: "Vehicle deleted successfully"}}
}

// ModerateVehicle updates the moderation status of a vehicle.
// If status is declined (3), sends push notification and inserts notification record.
func (s *AdminService) ModerateVehicle(ctx *fasthttp.RequestCtx, req *model.ModerateItemRequest) model.Response {
	if req == nil {
		return model.Response{Error: errors.New("invalid request data"), Status: http.StatusBadRequest}
	}

	userID, err := s.repo.ModerateVehicle(ctx, req.ID, req.Status)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	// If declined, send notification
	if req.Status == 3 {
		s.sendModerationNotification(context.Background(), userID, "car", req.ID, req.Title, req.Description)
	}

	return model.Response{Data: model.Success{Message: "Vehicle moderation status updated successfully"}}
}

// sendModerationNotification sends push notification and inserts notification record when item is declined.
func (s *AdminService) sendModerationNotification(ctx context.Context, userID int, itemType string, itemID int, title, description string) {
	// Get user's role ID
	userRoleID, err := s.repo.GetUserRoleID(ctx, userID)
	if err != nil {
		return
	}

	// Set default title and description if not provided
	if title == "" {
		title = "Listing Declined"
	}
	if description == "" {
		description = "Your listing has been declined by moderation."
	}

	// Insert notification record
	_ = s.repo.InsertNotification(ctx, "moderation_declined", userRoleID, userID, itemType, itemID, title, description)

	// Send push notification
	token, err := s.repo.GetUserDeviceToken(ctx, userID)
	if err != nil || token == "" {
		return
	}

	data := map[string]string{
		"type":      "moderation",
		"item_type": itemType,
		"item_id":   strconv.Itoa(itemID),
	}

	_, _ = s.firebase.SendSimpleNotification(token, title, description, data)
}
