package service

import (
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
