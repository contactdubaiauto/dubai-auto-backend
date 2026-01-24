package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/model"
	"dubai-auto/internal/utils"
)

// GetCars godoc
// @Summary      Get all cars
// @Description  Returns a list of all cars
// @Tags         admin-cars
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminVehicleListItem
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cars [get]
// Admin cars handlers
func (h *AdminHandler) GetVehicles(c *fiber.Ctx) error {
	limit := c.Query("limit")
	lastID := c.Query("last_id")

	lastIDInt, limitInt := utils.CheckLastIDLimit(lastID, limit, "")
	data := h.service.GetVehicles(c.Context(), limitInt, lastIDInt)
	return utils.FiberResponse(c, data)
}

// GetVehicle godoc
// @Summary      Get a vehicle by ID
// @Description  Returns a vehicle by ID
// @Tags         admin-cars
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "Vehicle ID"
// @Success      200  {object}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cars/{id} [get]
func (h *AdminHandler) GetVehicle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("vehicle id must be integer"),
		})
	}

	data := h.service.GetVehicleByID(c.Context(), id)
	return utils.FiberResponse(c, data)
}

// CreateVehicle godoc
// @Summary      Create a vehicle
// @Description  Creates a new vehicle
// @Tags         admin-cars
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        vehicle  body      model.AdminCreateVehicleRequest  true  "Vehicle"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cars [post]
func (h *AdminHandler) CreateVehicle(c *fiber.Ctx) error {
	req := &model.AdminCreateVehicleRequest{}

	if err := c.BodyParser(req); err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}

	data := h.service.CreateVehicle(c.Context(), req)
	return utils.FiberResponse(c, data)
}

// UpdateVehicle godoc
// @Summary      Update a vehicle
// @Description  Updates a vehicle
// @Tags         admin-cars
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "Vehicle ID"
// @Param        vehicle  body      model.AdminUpdateVehicleStatusRequest  true  "Vehicle"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cars/{id} [put]
func (h *AdminHandler) UpdateVehicle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("vehicle id must be integer"),
		})
	}

	req := &model.AdminUpdateVehicleStatusRequest{}

	if err := c.BodyParser(req); err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}

	data := h.service.UpdateVehicleStatus(c.Context(), id, req)
	return utils.FiberResponse(c, data)
}

// DeleteVehicle godoc
// @Summary      Delete a vehicle
// @Description  Deletes a vehicle
// @Tags         admin-cars
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "Vehicle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cars/{id} [delete]
func (h *AdminHandler) DeleteVehicle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("vehicle id must be integer"),
		})
	}

	data := h.service.DeleteVehicle(c.Context(), id, "/images/cars/"+idStr)
	return utils.FiberResponse(c, data)
}
