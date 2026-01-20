package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/utils"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

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

// GetMotoBrandsByCategoryID godoc
// @Summary      Get moto brands by category ID
// @Description  Returns a list of all moto brands by category ID
// @Tags         admin-moto-categories
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto category ID"
// @Success      200  {array}  model.AdminMotoBrandResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{id}/brands [get]
func (h *AdminHandler) GetMotoBrandsByCategoryID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("moto category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetMotoBrandsByCategoryID(ctx, id)
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

// // Moto Parameters handlers

// // GetMotoParameters godoc
// // @Summary      Get all moto parameters
// // @Description  Returns a list of all moto parameters
// // @Tags         admin-moto-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Success      200  {array}  model.AdminMotoParameterResponse
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters [get]
// func (h *AdminHandler) GetMotoParameters(c *fiber.Ctx) error {
// 	ctx := c.Context()
// 	data := h.service.GetMotoParameters(ctx)
// 	return utils.FiberResponse(c, data)
// }

// // CreateMotoParameter godoc
// // @Summary      Create a moto parameter
// // @Description  Creates a new moto parameter
// // @Tags         admin-moto-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        motoParameter  body      model.CreateMotoParameterRequest  true  "Moto parameter data"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters [post]
// func (h *AdminHandler) CreateMotoParameter(c *fiber.Ctx) error {
// 	var req model.CreateMotoParameterRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("invalid request body"),
// 		})
// 	}

// 	if err := h.validator.Validate(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  err,
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.CreateMotoParameter(ctx, &req)
// 	return utils.FiberResponse(c, data)
// }

// // UpdateMotoParameter godoc
// // @Summary      Update a moto parameter
// // @Description  Updates a moto parameter by ID
// // @Tags         admin-moto-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        id             path      int                            true  "Moto parameter ID"
// // @Param        motoParameter  body      model.UpdateMotoParameterRequest  true  "Moto parameter data"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters/{id} [put]
// func (h *AdminHandler) UpdateMotoParameter(c *fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	var req model.UpdateMotoParameterRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("invalid request body"),
// 		})
// 	}

// 	if err := h.validator.Validate(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  err,
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.UpdateMotoParameter(ctx, id, &req)
// 	return utils.FiberResponse(c, data)
// }

// // DeleteMotoParameter godoc
// // @Summary      Delete a moto parameter
// // @Description  Deletes a moto parameter by ID
// // @Tags         admin-moto-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Param        id   path      int  true  "Moto parameter ID"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters/{id} [delete]
// func (h *AdminHandler) DeleteMotoParameter(c *fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.DeleteMotoParameter(ctx, id)
// 	return utils.FiberResponse(c, data)
// }

// // Moto Parameter Values handlers

// // GetMotoParameterValues godoc
// // @Summary      Get moto parameter values
// // @Description  Returns a list of moto parameter values for a specific parameter
// // @Tags         admin-moto-parameter-values
// // @Produce      json
// // @Security     BearerAuth
// // @Param        moto_param_id   path      int  true  "Moto parameter ID"
// // @Success      200  {array}  model.AdminMotoParameterValueResponse
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values [get]
// func (h *AdminHandler) GetMotoParameterValues(c *fiber.Ctx) error {
// 	motoParamIdStr := c.Params("moto_param_id")
// 	motoParamId, err := strconv.Atoi(motoParamIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.GetMotoParameterValues(ctx, motoParamId)
// 	return utils.FiberResponse(c, data)
// }

// // CreateMotoParameterValue godoc
// // @Summary      Create a moto parameter value
// // @Description  Creates a new moto parameter value
// // @Tags         admin-moto-parameter-values
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        moto_param_id   path      int                                   true  "Moto parameter ID"
// // @Param        parameterValue  body      model.CreateMotoParameterValueRequest  true  "Moto parameter value data"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values [post]
// func (h *AdminHandler) CreateMotoParameterValue(c *fiber.Ctx) error {
// 	motoParamIdStr := c.Params("moto_param_id")
// 	motoParamId, err := strconv.Atoi(motoParamIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	var req model.CreateMotoParameterValueRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("invalid request body"),
// 		})
// 	}

// 	if err := h.validator.Validate(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  err,
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.CreateMotoParameterValue(ctx, motoParamId, &req)
// 	return utils.FiberResponse(c, data)
// }

// // UpdateMotoParameterValue godoc
// // @Summary      Update a moto parameter value
// // @Description  Updates a moto parameter value by ID
// // @Tags         admin-moto-parameter-values
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        moto_param_id   path      int                                   true  "Moto parameter ID"
// // @Param        id              path      int                                   true  "Moto parameter value ID"
// // @Param        parameterValue  body      model.UpdateMotoParameterValueRequest  true  "Moto parameter value data"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values/{id} [put]
// func (h *AdminHandler) UpdateMotoParameterValue(c *fiber.Ctx) error {
// 	motoParamIdStr := c.Params("moto_param_id")
// 	motoParamId, err := strconv.Atoi(motoParamIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter value id must be integer"),
// 		})
// 	}

// 	var req model.UpdateMotoParameterValueRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("invalid request body"),
// 		})
// 	}

// 	if err := h.validator.Validate(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  err,
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.UpdateMotoParameterValue(ctx, motoParamId, id, &req)
// 	return utils.FiberResponse(c, data)
// }

// // DeleteMotoParameterValue godoc
// // @Summary      Delete a moto parameter value
// // @Description  Deletes a moto parameter value by ID
// // @Tags         admin-moto-parameter-values
// // @Produce      json
// // @Security     BearerAuth
// // @Param        moto_param_id   path      int  true  "Moto parameter ID"
// // @Param        id              path      int  true  "Moto parameter value ID"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values/{id} [delete]
// func (h *AdminHandler) DeleteMotoParameterValue(c *fiber.Ctx) error {
// 	motoParamIdStr := c.Params("moto_param_id")
// 	motoParamId, err := strconv.Atoi(motoParamIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter value id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.DeleteMotoParameterValue(ctx, motoParamId, id)
// 	return utils.FiberResponse(c, data)
// }

// // Moto Category Parameters handlers

// // GetMotoCategoryParameters godoc
// // @Summary      Get moto category parameters
// // @Description  Returns a list of moto category parameters for a specific category
// // @Tags         admin-moto-category-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int  true  "Moto category ID"
// // @Success      200  {array}  model.AdminMotoCategoryParameterResponse
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-categories/{category_id}/parameters [get]
// func (h *AdminHandler) GetMotoCategoryParameters(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.GetMotoCategoryParameters(ctx, categoryId)
// 	return utils.FiberResponse(c, data)
// }

// // CreateMotoCategoryParameter godoc
// // @Summary      Create a moto category parameter
// // @Description  Creates a new moto category parameter
// // @Tags         admin-moto-category-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int                                     true  "Moto category ID"
// // @Param        parameter     body      model.CreateMotoCategoryParameterRequest  true  "Moto category parameter data"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-categories/{category_id}/parameters [post]
// func (h *AdminHandler) CreateMotoCategoryParameter(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	var req model.CreateMotoCategoryParameterRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("invalid request body"),
// 		})
// 	}

// 	if err := h.validator.Validate(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  err,
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.CreateMotoCategoryParameter(ctx, categoryId, &req)
// 	return utils.FiberResponse(c, data)
// }

// // UpdateMotoCategoryParameter godoc
// // @Summary      Update a moto category parameter
// // @Description  Updates a moto category parameter by ID
// // @Tags         admin-moto-category-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int                                     true  "Moto category ID"
// // @Param        parameter_id  path      int                                     true  "Moto parameter ID"
// // @Param        parameter     body      model.UpdateMotoCategoryParameterRequest  true  "Moto category parameter data"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-categories/{category_id}/parameters/{parameter_id} [put]
// func (h *AdminHandler) UpdateMotoCategoryParameter(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	parameterIdStr := c.Params("parameter_id")
// 	parameterId, err := strconv.Atoi(parameterIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	var req model.UpdateMotoCategoryParameterRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("invalid request body"),
// 		})
// 	}

// 	if err := h.validator.Validate(&req); err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  err,
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.UpdateMotoCategoryParameter(ctx, categoryId, parameterId, &req)
// 	return utils.FiberResponse(c, data)
// }

// // DeleteMotoCategoryParameter godoc
// // @Summary      Delete a moto category parameter
// // @Description  Deletes a moto category parameter by ID
// // @Tags         admin-moto-category-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int  true  "Moto category ID"
// // @Param        parameter_id  path      int  true  "Moto parameter ID"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/moto-categories/{category_id}/parameters/{parameter_id} [delete]
// func (h *AdminHandler) DeleteMotoCategoryParameter(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	parameterIdStr := c.Params("parameter_id")
// 	parameterId, err := strconv.Atoi(parameterIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("moto parameter id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.DeleteMotoCategoryParameter(ctx, categoryId, parameterId)
// 	return utils.FiberResponse(c, data)
// }
