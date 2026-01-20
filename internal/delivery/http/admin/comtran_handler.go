package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/utils"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Comtrans Categories handlers

// GetComtransCategories godoc
// @Summary      Get all comtrans categories
// @Description  Returns a list of all comtrans categories
// @Tags         admin-comtrans-categories
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminComtransCategoryResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories [get]
func (h *AdminHandler) GetComtransCategories(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetComtransCategories(ctx)
	return utils.FiberResponse(c, data)
}

// GetComtransBrandsByCategoryID godoc
// @Summary      Get comtrans brands by category ID
// @Description  Returns a list of comtrans brands for a specific category
// @Tags         admin-comtrans-brands
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int  true  "Comtrans category ID"
// @Success      200  {array}  model.AdminComtransBrandResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{category_id}/brands [get]
func (h *AdminHandler) GetComtransBrandsByCategoryID(c *fiber.Ctx) error {
	categoryIdStr := c.Params("id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetComtransBrandsByCategoryID(ctx, categoryId)
	return utils.FiberResponse(c, data)
}

// CreateComtransCategory godoc
// @Summary      Create a comtrans category
// @Description  Creates a new comtrans category
// @Tags         admin-comtrans-categories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        comtransCategory  body      model.CreateComtransCategoryRequest  true  "Comtrans category data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories [post]
func (h *AdminHandler) CreateComtransCategory(c *fiber.Ctx) error {
	var req model.CreateComtransCategoryRequest

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
	data := h.service.CreateComtransCategory(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateComtransCategory godoc
// @Summary      Update a comtrans category
// @Description  Updates a comtrans category by ID
// @Tags         admin-comtrans-categories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id                path      int                              true  "Comtrans category ID"
// @Param        comtransCategory  body      model.UpdateComtransCategoryRequest  true  "Comtrans category data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{id} [put]
func (h *AdminHandler) UpdateComtransCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans category id must be integer"),
		})
	}

	var req model.UpdateComtransCategoryRequest

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
	data := h.service.UpdateComtransCategory(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteComtransCategory godoc
// @Summary      Delete a comtrans category
// @Description  Deletes a comtrans category by ID
// @Tags         admin-comtrans-categories
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Comtrans category ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{id} [delete]
func (h *AdminHandler) DeleteComtransCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteComtransCategory(ctx, id)
	return utils.FiberResponse(c, data)
}

// // Comtrans Category Parameters handlers

// // GetComtransCategoryParameters godoc
// // @Summary      Get comtrans category parameters
// // @Description  Returns a list of comtrans category parameters for a specific category
// // @Tags         admin-comtrans-category-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int  true  "Comtrans category ID"
// // @Success      200  {array}  model.AdminComtransCategoryParameterResponse
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters [get]
// func (h *AdminHandler) GetComtransCategoryParameters(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.GetComtransCategoryParameters(ctx, categoryId)
// 	return utils.FiberResponse(c, data)
// }

// // CreateComtransCategoryParameter godoc
// // @Summary      Create a comtrans category parameter
// // @Description  Creates a new comtrans category parameter
// // @Tags         admin-comtrans-category-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int                                         true  "Comtrans category ID"
// // @Param        parameter     body      model.CreateComtransCategoryParameterRequest  true  "Comtrans category parameter data"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters [post]
// func (h *AdminHandler) CreateComtransCategoryParameter(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	var req model.CreateComtransCategoryParameterRequest

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
// 	data := h.service.CreateComtransCategoryParameter(ctx, categoryId, &req)
// 	return utils.FiberResponse(c, data)
// }

// // UpdateComtransCategoryParameter godoc
// // @Summary      Update a comtrans category parameter
// // @Description  Updates a comtrans category parameter by ID
// // @Tags         admin-comtrans-category-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int                                         true  "Comtrans category ID"
// // @Param        id            path      int                                         true  "Comtrans category parameter ID"
// // @Param        parameter     body      model.UpdateComtransCategoryParameterRequest  true  "Comtrans category parameter data"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters/{id} [put]
// func (h *AdminHandler) UpdateComtransCategoryParameter(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("comtrans category parameter id must be integer"),
// 		})
// 	}

// 	var req model.UpdateComtransCategoryParameterRequest

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
// 	data := h.service.UpdateComtransCategoryParameter(ctx, categoryId, id, &req)
// 	return utils.FiberResponse(c, data)
// }

// // DeleteComtransCategoryParameter godoc
// // @Summary      Delete a comtrans category parameter
// // @Description  Deletes a comtrans category parameter by ID
// // @Tags         admin-comtrans-category-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Param        category_id   path      int  true  "Comtrans category ID"
// // @Param        id            path      int  true  "Comtrans category parameter ID"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters/{id} [delete]
// func (h *AdminHandler) DeleteComtransCategoryParameter(c *fiber.Ctx) error {
// 	categoryIdStr := c.Params("category_id")
// 	categoryId, err := strconv.Atoi(categoryIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("category id must be integer"),
// 		})
// 	}

// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("comtrans category parameter id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.DeleteComtransCategoryParameter(ctx, categoryId, id)
// 	return utils.FiberResponse(c, data)
// }

// Comtrans Brands handlers

// GetComtransBrands godoc
// @Summary      Get all comtrans brands
// @Description  Returns a list of all comtrans brands
// @Tags         admin-comtrans-brands
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminComtransBrandResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-brands [get]
func (h *AdminHandler) GetComtransBrands(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetComtransBrands(ctx)
	return utils.FiberResponse(c, data)
}

// GetComtransModelsByBrandID godoc
// @Summary      Get comtrans models by brand ID
// @Description  Returns a list of comtrans models by brand ID
// @Tags         admin-comtrans-brands
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Comtrans brand ID"
// @Success      200  {array}  model.AdminComtransModelResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-brands/{id}/models [get]
func (h *AdminHandler) GetComtransModelsByBrandID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetComtransModelsByBrandID(ctx, id)
	return utils.FiberResponse(c, data)
}

// CreateComtransBrand godoc
// @Summary      Create a comtrans brand
// @Description  Creates a new comtrans brand
// @Tags         admin-comtrans-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        comtransBrand  body      model.CreateComtransBrandRequest  true  "Comtrans brand data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-brands [post]
func (h *AdminHandler) CreateComtransBrand(c *fiber.Ctx) error {
	var req model.CreateComtransBrandRequest

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
	data := h.service.CreateComtransBrand(ctx, &req)
	return utils.FiberResponse(c, data)
}

// CreateComtransBrandImage godoc
// @Summary      Create a new comtrans brand image
// @Description  Creates a new brand image
// @Tags         admin-comtrans-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Comtrans brand ID"
// @Param        image  formData  file  true  "Comtrans brand image"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-brands/{id}/images [post]
func (h *AdminHandler) CreateComtransBrandImage(c *fiber.Ctx) error {
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

	data := h.service.CreateComtransBrandImage(ctx, form, id)
	return utils.FiberResponse(c, data)
}

// UpdateComtransBrand godoc
// @Summary      Update a comtrans brand
// @Description  Updates a comtrans brand by ID
// @Tags         admin-comtrans-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      int                            true  "Comtrans brand ID"
// @Param        comtransBrand  body      model.UpdateComtransBrandRequest  true  "Comtrans brand data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-brands/{id} [put]
func (h *AdminHandler) UpdateComtransBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans brand id must be integer"),
		})
	}

	var req model.UpdateComtransBrandRequest

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
	data := h.service.UpdateComtransBrand(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteComtransBrand godoc
// @Summary      Delete a comtrans brand
// @Description  Deletes a comtrans brand by ID
// @Tags         admin-comtrans-brands
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Comtrans brand ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-brands/{id} [delete]
func (h *AdminHandler) DeleteComtransBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteComtransBrand(ctx, id)
	return utils.FiberResponse(c, data)
}

// Comtrans Models handlers

// GetComtransModels godoc
// @Summary      Get all comtrans models
// @Description  Returns a list of all comtrans models
// @Tags         admin-comtrans-models
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminComtransModelResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-models [get]
func (h *AdminHandler) GetComtransModels(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetComtransModels(ctx)
	return utils.FiberResponse(c, data)
}

// CreateComtransModel godoc
// @Summary      Create a comtrans model
// @Description  Creates a new comtrans model
// @Tags         admin-comtrans-models
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        comtransModel  body      model.CreateComtransModelRequest  true  "Comtrans model data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-models [post]
func (h *AdminHandler) CreateComtransModel(c *fiber.Ctx) error {
	var req model.CreateComtransModelRequest

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
	data := h.service.CreateComtransModel(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateComtransModel godoc
// @Summary      Update a comtrans model
// @Description  Updates a comtrans model by ID
// @Tags         admin-comtrans-models
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      int                            true  "Comtrans model ID"
// @Param        comtransModel  body      model.UpdateComtransModelRequest  true  "Comtrans model data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-models/{id} [put]
func (h *AdminHandler) UpdateComtransModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans model id must be integer"),
		})
	}

	var req model.UpdateComtransModelRequest

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
	data := h.service.UpdateComtransModel(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteComtransModel godoc
// @Summary      Delete a comtrans model
// @Description  Deletes a comtrans model by ID
// @Tags         admin-comtrans-models
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Comtrans model ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-models/{id} [delete]
func (h *AdminHandler) DeleteComtransModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteComtransModel(ctx, id)
	return utils.FiberResponse(c, data)
}

// // Comtrans Parameters handlers

// // GetComtransParameters godoc
// // @Summary      Get all comtrans parameters
// // @Description  Returns a list of all comtrans parameters
// // @Tags         admin-comtrans-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Success      200  {array}  model.AdminComtransParameterResponse
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters [get]
// func (h *AdminHandler) GetComtransParameters(c *fiber.Ctx) error {
// 	ctx := c.Context()
// 	data := h.service.GetComtransParameters(ctx)
// 	return utils.FiberResponse(c, data)
// }

// // CreateComtransParameter godoc
// // @Summary      Create a comtrans parameter
// // @Description  Creates a new comtrans parameter
// // @Tags         admin-comtrans-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        comtransParameter  body      model.CreateComtransParameterRequest  true  "Comtrans parameter data"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters [post]
// func (h *AdminHandler) CreateComtransParameter(c *fiber.Ctx) error {
// 	var req model.CreateComtransParameterRequest

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
// 	data := h.service.CreateComtransParameter(ctx, &req)
// 	return utils.FiberResponse(c, data)
// }

// // UpdateComtransParameter godoc
// // @Summary      Update a comtrans parameter
// // @Description  Updates a comtrans parameter by ID
// // @Tags         admin-comtrans-parameters
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        id                 path      int                                true  "Comtrans parameter ID"
// // @Param        comtransParameter  body      model.UpdateComtransParameterRequest  true  "Comtrans parameter data"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters/{id} [put]
// func (h *AdminHandler) UpdateComtransParameter(c *fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("comtrans parameter id must be integer"),
// 		})
// 	}

// 	var req model.UpdateComtransParameterRequest

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
// 	data := h.service.UpdateComtransParameter(ctx, id, &req)
// 	return utils.FiberResponse(c, data)
// }

// // DeleteComtransParameter godoc
// // @Summary      Delete a comtrans parameter
// // @Description  Deletes a comtrans parameter by ID
// // @Tags         admin-comtrans-parameters
// // @Produce      json
// // @Security     BearerAuth
// // @Param        id   path      int  true  "Comtrans parameter ID"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters/{id} [delete]
// func (h *AdminHandler) DeleteComtransParameter(c *fiber.Ctx) error {
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("comtrans parameter id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.DeleteComtransParameter(ctx, id)
// 	return utils.FiberResponse(c, data)
// }

// // Comtrans Parameter Values handlers

// // GetComtransParameterValues godoc
// // @Summary      Get comtrans parameter values
// // @Description  Returns a list of comtrans parameter values for a specific parameter
// // @Tags         admin-comtrans-parameter-values
// // @Produce      json
// // @Security     BearerAuth
// // @Param        parameter_id   path      int  true  "Comtrans parameter ID"
// // @Success      200  {array}  model.AdminComtransParameterValueResponse
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values [get]
// func (h *AdminHandler) GetComtransParameterValues(c *fiber.Ctx) error {
// 	parameterIdStr := c.Params("parameter_id")
// 	parameterId, err := strconv.Atoi(parameterIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("parameter id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.GetComtransParameterValues(ctx, parameterId)
// 	return utils.FiberResponse(c, data)
// }

// // CreateComtransParameterValue godoc
// // @Summary      Create a comtrans parameter value
// // @Description  Creates a new comtrans parameter value
// // @Tags         admin-comtrans-parameter-values
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        parameter_id   path      int                                       true  "Comtrans parameter ID"
// // @Param        parameterValue body      model.CreateComtransParameterValueRequest  true  "Comtrans parameter value data"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values [post]
// func (h *AdminHandler) CreateComtransParameterValue(c *fiber.Ctx) error {
// 	parameterIdStr := c.Params("parameter_id")
// 	parameterId, err := strconv.Atoi(parameterIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("parameter id must be integer"),
// 		})
// 	}

// 	var req model.CreateComtransParameterValueRequest

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
// 	data := h.service.CreateComtransParameterValue(ctx, parameterId, &req)
// 	return utils.FiberResponse(c, data)
// }

// // UpdateComtransParameterValue godoc
// // @Summary      Update a comtrans parameter value
// // @Description  Updates a comtrans parameter value by ID
// // @Tags         admin-comtrans-parameter-values
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        parameter_id   path      int                                       true  "Comtrans parameter ID"
// // @Param        id             path      int                                       true  "Comtrans parameter value ID"
// // @Param        parameterValue body      model.UpdateComtransParameterValueRequest  true  "Comtrans parameter value data"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values/{id} [put]
// func (h *AdminHandler) UpdateComtransParameterValue(c *fiber.Ctx) error {
// 	parameterIdStr := c.Params("parameter_id")
// 	parameterId, err := strconv.Atoi(parameterIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("parameter id must be integer"),
// 		})
// 	}

// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("comtrans parameter value id must be integer"),
// 		})
// 	}

// 	var req model.UpdateComtransParameterValueRequest

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
// 	data := h.service.UpdateComtransParameterValue(ctx, parameterId, id, &req)
// 	return utils.FiberResponse(c, data)
// }

// // DeleteComtransParameterValue godoc
// // @Summary      Delete a comtrans parameter value
// // @Description  Deletes a comtrans parameter value by ID
// // @Tags         admin-comtrans-parameter-values
// // @Produce      json
// // @Security     BearerAuth
// // @Param        parameter_id   path      int  true  "Comtrans parameter ID"
// // @Param        id             path      int  true  "Comtrans parameter value ID"
// // @Success      200  {object}  model.Success
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values/{id} [delete]
// func (h *AdminHandler) DeleteComtransParameterValue(c *fiber.Ctx) error {
// 	parameterIdStr := c.Params("parameter_id")
// 	parameterId, err := strconv.Atoi(parameterIdStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("parameter id must be integer"),
// 		})
// 	}

// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, model.Response{
// 			Status: 400,
// 			Error:  errors.New("comtrans parameter value id must be integer"),
// 		})
// 	}

// 	ctx := c.Context()
// 	data := h.service.DeleteComtransParameterValue(ctx, parameterId, id)
// 	return utils.FiberResponse(c, data)
// }
