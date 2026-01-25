package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/utils"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetMotorcycles godoc
// @Summary      Get all motorcycles
// @Description  Returns a list of all motorcycles
// @Tags         admin-motorcycles
// @Produce      json
// @Security     BearerAuth
// @Param        moderation_status  query  string  false  "Moderation Status"
// @Param        limit  query  string  false  "Limit"
// @Param        last_id  query  string  false  "Last ID"
// @Success      200  {array}  model.AdminMotoListItem
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/motorcycles [get]
func (h *AdminHandler) GetMotorcycles(c *fiber.Ctx) error {
	limit := c.Query("limit")
	lastID := c.Query("last_id")
	lastIDInt, limitInt := utils.CheckLastIDLimit(lastID, limit, "")
	moderationStatus := c.Query("moderation_status", "0")
	data := h.service.GetMotorcycles(c.Context(), limitInt, lastIDInt, moderationStatus)
	return utils.FiberResponse(c, data)
}

// GetMotorcycle godoc
// @Summary      Get a motorcycle by ID
// @Description  Returns a motorcycle by ID
// @Tags         admin-motorcycles
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "Motorcycle ID"
// @Success      200  {object}  model.GetMotorcyclesResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/motorcycles/{id} [get]
func (h *AdminHandler) GetMotorcycle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}
	data := h.service.GetMotorcycleByID(c.Context(), id)
	return utils.FiberResponse(c, data)
}

// DeleteMotorcycle godoc
// @Summary      Delete a motorcycle
// @Description  Deletes a motorcycle
// @Tags         admin-motorcycles
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "Motorcycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/motorcycles/{id} [delete]
func (h *AdminHandler) DeleteMotorcycle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}
	data := h.service.DeleteMotorcycle(c.Context(), id, "/images/motorcycles/"+idStr)
	return utils.FiberResponse(c, data)
}

// ModerateMotorcycleStatus godoc
// @Summary      Moderate a motorcycle
// @Description  Updates the moderation status of a motorcycle. If declined (status=3), sends push notification to the item's user.
// @Tags         admin-motorcycles
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.ModerateItemRequest  true  "Moderation request: id, status (1-pending, 2-accepted, 3-declined), title (optional), description (optional)"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/motorcycles/moderate [post]
func (h *AdminHandler) ModerateMotorcycleStatus(c *fiber.Ctx) error {
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
	data := h.service.ModerateMotorcycle(c.Context(), &req)
	return utils.FiberResponse(c, data)
}

// Moto Categories handlers

// GetMotoCategories godoc
// @Summary      Get all moto categories
// @Description  Returns a list of all moto categories
// @Tags         admin-moto-categories
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminMotoCategoryResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories [get]
func (h *AdminHandler) GetMotoCategories(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetMotoCategories(ctx)
	return utils.FiberResponse(c, data)
}

// CreateMotoCategory godoc
// @Summary      Create a moto category
// @Description  Creates a new moto category
// @Tags         admin-moto-categories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        motoCategory  body      model.CreateMotoCategoryRequest  true  "Moto category data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories [post]
func (h *AdminHandler) CreateMotoCategory(c *fiber.Ctx) error {
	var req model.CreateMotoCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.CreateMotoCategory(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateMotoCategory godoc
// @Summary      Update a moto category
// @Description  Updates a moto category by ID
// @Tags         admin-moto-categories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id            path      int                           true  "Moto category ID"
// @Param        motoCategory  body      model.UpdateMotoCategoryRequest  true  "Moto category data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{id} [put]
func (h *AdminHandler) UpdateMotoCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto category id must be integer"),
		})
	}

	var req model.UpdateMotoCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.UpdateMotoCategory(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteMotoCategory godoc
// @Summary      Delete a moto category
// @Description  Deletes a moto category by ID
// @Tags         admin-moto-categories
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto category ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{id} [delete]
func (h *AdminHandler) DeleteMotoCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteMotoCategory(ctx, id)
	return utils.FiberResponse(c, data)
}

// Moto Brands handlers

// GetMotoBrands godoc
// @Summary      Get all moto brands
// @Description  Returns a list of all moto brands
// @Tags         admin-moto-brands
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminMotoBrandResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-brands [get]
func (h *AdminHandler) GetMotoBrands(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetMotoBrands(ctx)
	return utils.FiberResponse(c, data)
}

// GetMotoModelsByBrandID godoc
// @Summary      Get moto models by brand ID
// @Description  Returns a list of all moto models by brand ID
// @Tags         admin-moto-brands
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto brand ID"
// @Success      200  {array}  model.AdminMotoModelResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-brands/{id}/models [get]
func (h *AdminHandler) GetMotoModelsByBrandID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetMotoModelsByBrandID(ctx, id)
	return utils.FiberResponse(c, data)
}

// CreateMotoBrand godoc
// @Summary      Create a moto brand
// @Description  Creates a new moto brand
// @Tags         admin-moto-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        motoBrand  body      model.CreateMotoBrandRequest  true  "Moto brand data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-brands [post]
func (h *AdminHandler) CreateMotoBrand(c *fiber.Ctx) error {
	var req model.CreateMotoBrandRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.CreateMotoBrand(ctx, &req)
	return utils.FiberResponse(c, data)
}

// CreateMotoBrandImage godoc
// @Summary      Create a new brand image
// @Description  Creates a new brand image
// @Tags         admin-moto-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto brand ID"
// @Param        image  formData  file  true  "Moto brand image"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-brands/{id}/images [post]
func (h *AdminHandler) CreateMotoBrandImage(c *fiber.Ctx) error {
	ctx := c.Context()
	form, _ := c.MultipartForm()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	data := h.service.CreateMotoBrandImage(ctx, form, id)
	return utils.FiberResponse(c, data)
}

// UpdateMotoBrand godoc
// @Summary      Update a moto brand
// @Description  Updates a moto brand by ID
// @Tags         admin-moto-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      int                        true  "Moto brand ID"
// @Param        motoBrand  body      model.UpdateMotoBrandRequest  true  "Moto brand data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-brands/{id} [put]
func (h *AdminHandler) UpdateMotoBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto brand id must be integer"),
		})
	}

	var req model.UpdateMotoBrandRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.UpdateMotoBrand(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteMotoBrand godoc
// @Summary      Delete a moto brand
// @Description  Deletes a moto brand by ID
// @Tags         admin-moto-brands
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto brand ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-brands/{id} [delete]
func (h *AdminHandler) DeleteMotoBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteMotoBrand(ctx, id)
	return utils.FiberResponse(c, data)
}

// Moto Models handlers

// GetMotoModels godoc
// @Summary      Get all moto models
// @Description  Returns a list of all moto models
// @Tags         admin-moto-models
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminMotoModelResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-models [get]
func (h *AdminHandler) GetMotoModels(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetMotoModels(ctx)
	return utils.FiberResponse(c, data)
}

// CreateMotoModel godoc
// @Summary      Create a moto model
// @Description  Creates a new moto model
// @Tags         admin-moto-models
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        motoModel  body      model.CreateMotoModelRequest  true  "Moto model data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-models [post]
func (h *AdminHandler) CreateMotoModel(c *fiber.Ctx) error {
	var req model.CreateMotoModelRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.CreateMotoModel(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateMotoModel godoc
// @Summary      Update a moto model
// @Description  Updates a moto model by ID
// @Tags         admin-moto-models
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id         path      int                        true  "Moto model ID"
// @Param        motoModel  body      model.UpdateMotoModelRequest  true  "Moto model data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-models/{id} [put]
func (h *AdminHandler) UpdateMotoModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto model id must be integer"),
		})
	}

	var req model.UpdateMotoModelRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.UpdateMotoModel(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteMotoModel godoc
// @Summary      Delete a moto model
// @Description  Deletes a moto model by ID
// @Tags         admin-moto-models
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto model ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-models/{id} [delete]
func (h *AdminHandler) DeleteMotoModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteMotoModel(ctx, id)
	return utils.FiberResponse(c, data)
}
