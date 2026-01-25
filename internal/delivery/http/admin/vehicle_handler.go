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
// @Param        moderation_status  query  string  false  "Moderation Status"
// @Param        limit  query  string  false  "Limit"
// @Param        last_id  query  string  false  "Last ID"
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
	moderationStatus := c.Query("moderation_status", "0")

	lastIDInt, limitInt := utils.CheckLastIDLimit(lastID, limit, "")
	data := h.service.GetVehicles(c.Context(), limitInt, lastIDInt, moderationStatus)
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

// ModerateVehicleStatus godoc
// @Summary      Moderate a vehicle
// @Description  Updates the moderation status of a vehicle. If declined (status=3), sends push notification to the item's user.
// @Tags         admin-cars
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.ModerateItemRequest  true  "Moderation request: id, status (1-pending, 2-accepted, 3-declined), title (optional), description (optional)"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/cars/moderate [post]
func (h *AdminHandler) ModerateVehicleStatus(c *fiber.Ctx) error {
	var req model.ModerateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}
	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}
	data := h.service.ModerateVehicle(c.Context(), &req)
	return utils.FiberResponse(c, data)
}
