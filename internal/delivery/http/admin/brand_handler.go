package http

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/files"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Brands handlers

// GetBrands godoc
// @Summary      Get all brands
// @Description  Returns a list of all brands
// @Tags         admin-brands
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminBrandResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/brands [get]
func (h *AdminHandler) GetBrands(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetBrands(ctx)
	return utils.FiberResponse(c, data)
}

// CreateBrand godoc
// @Summary      Create a new brand
// @Description  Creates a new brand
// @Tags         admin-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        brand  body      model.CreateBrandRequest  true  "Brand data"
// @Success      200    {object}  model.SuccessWithId
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/brands [post]
func (h *AdminHandler) CreateBrand(c *fiber.Ctx) error {
	var req model.CreateBrandRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateBrand(ctx, &req)
	return utils.FiberResponse(c, data)
}

// CreateBrandImage godoc
// @Summary      Create a new brand image
// @Description  Creates a new brand image
// @Tags         admin-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Brand ID"
// @Param        image  formData  file  true  "Brand image"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/{id}/images [post]
func (h *AdminHandler) CreateBrandImage(c *fiber.Ctx) error {
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

	data := h.service.CreateBrandImage(ctx, form, id)
	return utils.FiberResponse(c, data)
}

// UpdateBrand godoc
// @Summary      Update a brand
// @Description  Updates an existing brand
// @Tags         admin-brands
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                       true  "Brand ID"
// @Param        brand  body      model.CreateBrandRequest  true  "Brand data"
// @Success      200    {object}  model.Success
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/{id} [put]
func (h *AdminHandler) UpdateBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	var req model.CreateBrandRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.UpdateBrand(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteBrand godoc
// @Summary      Delete a brand
// @Description  Deletes a brand by ID
// @Tags         admin-brands
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Brand ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/{id} [delete]
func (h *AdminHandler) DeleteBrand(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteBrand(ctx, id)
	return utils.FiberResponse(c, data)
}

// Models handlers

// GetModels godoc
// @Summary      Get all models
// @Description  Returns a list of all models
// @Tags         admin-models
// @Produce      json
// @Security     BearerAuth
// @Param        brand_id   path      int                       true  "Brand ID"
// @Success      200  {array}  model.AdminModelResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/:brand_id/models [get]
func (h *AdminHandler) GetModels(c *fiber.Ctx) error {
	ctx := c.Context()
	brandIdStr := c.Params("brand_id")
	brandId, err := strconv.Atoi(brandIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}
	data := h.service.GetModels(ctx, brandId)
	return utils.FiberResponse(c, data)
}

// CreateModel godoc
// @Summary      Create a new model
// @Description  Creates a new model
// @Tags         admin-models
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        model  body      model.CreateModelRequest  true  "Model data"
// @Success      200    {object}  model.SuccessWithId
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/:brand_id/models [post]
func (h *AdminHandler) CreateModel(c *fiber.Ctx) error {
	var req model.CreateModelRequest
	ctx := c.Context()
	brandIdStr := c.Params("brand_id")
	brandId, err := strconv.Atoi(brandIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateModel(ctx, brandId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateModel godoc
// @Summary      Update a model
// @Description  Updates an existing model
// @Tags         admin-models
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                       true  "Model ID"
// @Param        model  body      model.UpdateModelRequest  true  "Model data"
// @Success      200    {object}  model.Success
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/:brand_id/models/{id} [put]
func (h *AdminHandler) UpdateModel(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("model id must be integer"),
		})
	}

	var req model.UpdateModelRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.UpdateModel(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteModel godoc
// @Summary      Delete a model
// @Description  Deletes a model by ID
// @Tags         admin-models
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Model ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/:brand_id/models/{id} [delete]
func (h *AdminHandler) DeleteModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteModel(ctx, id)
	return utils.FiberResponse(c, data)
}

// Body Types handlers

// GetBodyTypes godoc
// @Summary      Get all body types
// @Description  Returns a list of all body types
// @Tags         admin-body-types
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminBodyTypeResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/body-types [get]
func (h *AdminHandler) GetBodyTypes(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetBodyTypes(ctx)
	return utils.FiberResponse(c, data)
}

// CreateBodyType godoc
// @Summary      Create a new body type
// @Description  Creates a new body type
// @Tags         admin-body-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        bodyType  body      model.CreateBodyTypeRequest  true  "Body type data"
// @Success      200       {object}  model.SuccessWithId
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/body-types [post]
func (h *AdminHandler) CreateBodyType(c *fiber.Ctx) error {
	var req model.CreateBodyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateBodyType(ctx, &req)
	return utils.FiberResponse(c, data)
}

// CreateBodyTypeImage godoc
// @Summary      Create a new body type image
// @Description  Creates a new body type image
// @Tags         admin-body-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int                          true  "Body type ID"
// @Param        image     formData  file    true   "body type image (max 1)"
// @Success      200       {object}  model.Success
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/body-types/{id} [post]
func (h *AdminHandler) CreateBodyTypeImage(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}

	image := form.File["image"]

	if len(image) > 1 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 files"),
		})
	}

	path, err := files.SaveOriginal(image[0], config.ENV.STATIC_PATH+"cars/body/"+strconv.Itoa(id))

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 500,
			Error:  err,
		})
	}

	data := h.service.CreateBodyTypeImage(ctx, id, path)
	return utils.FiberResponse(c, data)
}

// UpdateBodyType godoc
// @Summary      Update a body type
// @Description  Updates an existing body type
// @Tags         admin-body-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int                          true  "Body type ID"
// @Param        bodyType  body      model.CreateBodyTypeRequest  true  "Body type data"
// @Success      200       {object}  model.Success
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/body-types/{id} [put]
func (h *AdminHandler) UpdateBodyType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

	var req model.CreateBodyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.UpdateBodyType(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteBodyType godoc
// @Summary      Delete a body type
// @Description  Deletes a body type by ID
// @Tags         admin-body-types
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Body type ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/body-types/{id} [delete]
func (h *AdminHandler) DeleteBodyType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteBodyType(ctx, id)
	return utils.FiberResponse(c, data)
}

// DeleteBodyTypeImage godoc
// @Summary      Delete a body type image
// @Description  Deletes a body type image by ID
// @Tags         admin-body-types
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Body type image ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/body-types/{id}/images [delete]
func (h *AdminHandler) DeleteBodyTypeImage(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("body type image id must be integer"),
		})
	}
	data := h.service.DeleteBodyTypeImage(ctx, id)
	return utils.FiberResponse(c, data)
}

// CreateTransmission godoc
// @Summary      Create a transmission
// @Description  Creates a new transmission
// @Tags         admin-transmissions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        transmission  body      model.CreateTransmissionRequest  true  "Transmission data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/transmissions [post]
func (h *AdminHandler) CreateTransmission(c *fiber.Ctx) error {
	var req model.CreateTransmissionRequest

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
	data := h.service.CreateTransmission(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateTransmission godoc
// @Summary      Update a transmission
// @Description  Updates a transmission by ID
// @Tags         admin-transmissions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id            path      int                           true  "Transmission ID"
// @Param        transmission  body      model.CreateTransmissionRequest  true  "Transmission data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/transmissions/{id} [put]
func (h *AdminHandler) UpdateTransmission(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("transmission id must be integer"),
		})
	}

	var req model.CreateTransmissionRequest

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
	data := h.service.UpdateTransmission(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteTransmission godoc
// @Summary      Delete a transmission
// @Description  Deletes a transmission by ID
// @Tags         admin-transmissions
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Transmission ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/transmissions/{id} [delete]
func (h *AdminHandler) DeleteTransmission(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("transmission id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteTransmission(ctx, id)
	return utils.FiberResponse(c, data)
}

// Engine handlers

// GetEngines godoc
// @Summary      Get all engines
// @Description  Returns a list of all engines
// @Tags         admin-engines
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminEngineResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/engines [get]
func (h *AdminHandler) GetEngines(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetEngines(ctx)
	return utils.FiberResponse(c, data)
}

// CreateEngine godoc
// @Summary      Create an engine
// @Description  Creates a new engine
// @Tags         admin-engines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        engine  body      model.CreateEngineRequest  true  "Engine data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/engines [post]
func (h *AdminHandler) CreateEngine(c *fiber.Ctx) error {
	var req model.CreateEngineRequest

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
	data := h.service.CreateEngine(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateEngine godoc
// @Summary      Update an engine
// @Description  Updates an engine by ID
// @Tags         admin-engines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path      int                     true  "Engine ID"
// @Param        engine  body      model.CreateEngineRequest  true  "Engine data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/engines/{id} [put]
func (h *AdminHandler) UpdateEngine(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("engine id must be integer"),
		})
	}

	var req model.CreateEngineRequest

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
	data := h.service.UpdateEngine(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteEngine godoc
// @Summary      Delete an engine
// @Description  Deletes an engine by ID
// @Tags         admin-engines
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Engine ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/engines/{id} [delete]
func (h *AdminHandler) DeleteEngine(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("engine id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteEngine(ctx, id)
	return utils.FiberResponse(c, data)
}

// Drivetrain handlers

// GetDrivetrains godoc
// @Summary      Get all drivetrains
// @Description  Returns a list of all drivetrains
// @Tags         admin-drivetrains
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminDrivetrainResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/drivetrains [get]
func (h *AdminHandler) GetDrivetrains(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetDrivetrains(ctx)
	return utils.FiberResponse(c, data)
}

// CreateDrivetrain godoc
// @Summary      Create a drivetrain
// @Description  Creates a new drivetrain
// @Tags         admin-drivetrains
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        drivetrain  body      model.CreateDrivetrainRequest  true  "Drivetrain data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/drivetrains [post]
func (h *AdminHandler) CreateDrivetrain(c *fiber.Ctx) error {
	var req model.CreateDrivetrainRequest

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
	data := h.service.CreateDrivetrain(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateDrivetrain godoc
// @Summary      Update a drivetrain
// @Description  Updates a drivetrain by ID
// @Tags         admin-drivetrains
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path      int                         true  "Drivetrain ID"
// @Param        drivetrain  body      model.CreateDrivetrainRequest  true  "Drivetrain data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/drivetrains/{id} [put]
func (h *AdminHandler) UpdateDrivetrain(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("drivetrain id must be integer"),
		})
	}

	var req model.CreateDrivetrainRequest

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
	data := h.service.UpdateDrivetrain(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteDrivetrain godoc
// @Summary      Delete a drivetrain
// @Description  Deletes a drivetrain by ID
// @Tags         admin-drivetrains
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Drivetrain ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/drivetrains/{id} [delete]
func (h *AdminHandler) DeleteDrivetrain(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("drivetrain id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteDrivetrain(ctx, id)
	return utils.FiberResponse(c, data)
}

// Fuel Type handlers

// GetFuelTypes godoc
// @Summary      Get all fuel types
// @Description  Returns a list of all fuel types
// @Tags         admin-fuel-types
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminFuelTypeResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/fuel-types [get]
func (h *AdminHandler) GetFuelTypes(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetFuelTypes(ctx)
	return utils.FiberResponse(c, data)
}

// CreateFuelType godoc
// @Summary      Create a fuel type
// @Description  Creates a new fuel type
// @Tags         admin-fuel-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        fuelType  body      model.CreateFuelTypeRequest  true  "Fuel type data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/fuel-types [post]
func (h *AdminHandler) CreateFuelType(c *fiber.Ctx) error {
	var req model.CreateFuelTypeRequest

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
	data := h.service.CreateFuelType(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateFuelType godoc
// @Summary      Update a fuel type
// @Description  Updates a fuel type by ID
// @Tags         admin-fuel-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int                       true  "Fuel type ID"
// @Param        fuelType  body      model.CreateFuelTypeRequest  true  "Fuel type data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/fuel-types/{id} [put]
func (h *AdminHandler) UpdateFuelType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("fuel type id must be integer"),
		})
	}

	var req model.CreateFuelTypeRequest

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
	data := h.service.UpdateFuelType(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteFuelType godoc
// @Summary      Delete a fuel type
// @Description  Deletes a fuel type by ID
// @Tags         admin-fuel-types
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Fuel type ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/fuel-types/{id} [delete]
func (h *AdminHandler) DeleteFuelType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("fuel type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteFuelType(ctx, id)
	return utils.FiberResponse(c, data)
}

// Generation handlers

// GetGenerations godoc
// @Summary      Get all generations
// @Description  Returns a list of all generations
// @Tags         admin-generations
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminGenerationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations [get]
func (h *AdminHandler) GetGenerations(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetGenerations(ctx)
	return utils.FiberResponse(c, data)
}

// GetGenerationsByModel godoc
// @Summary      Get generations by model ID
// @Description  Returns a list of generations for a given model ID within a specific brand
// @Tags         admin-generations
// @Produce      json
// @Security     BearerAuth
// @Param        brand_id  path  int  true  "Brand ID"
// @Param        model_id  path  int  true  "Model ID"
// @Success      200  {array}  model.AdminGenerationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/brands/{brand_id}/models/{model_id}/generations [get]
func (h *AdminHandler) GetGenerationsByModel(c *fiber.Ctx) error {
	brandIdStr := c.Params("brand_id")
	modelIdStr := c.Params("model_id")

	brandId, err := strconv.Atoi(brandIdStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	modelId, err := strconv.Atoi(modelIdStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetGenerationsByModel(ctx, brandId, modelId)
	return utils.FiberResponse(c, data)
}

// CreateGeneration godoc
// @Summary      Create a generation
// @Description  Creates a new generation
// @Tags         admin-generations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        generation  body      model.CreateGenerationRequest  true  "Generation data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations [post]
func (h *AdminHandler) CreateGeneration(c *fiber.Ctx) error {
	var req model.CreateGenerationRequest

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
	data := h.service.CreateGeneration(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateGeneration godoc
// @Summary      Update a generation
// @Description  Updates a generation by ID
// @Tags         admin-generations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path      int                        true  "Generation ID"
// @Param        generation  body      model.UpdateGenerationRequest  true  "Generation data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{id} [put]
func (h *AdminHandler) UpdateGeneration(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	var req model.UpdateGenerationRequest

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
	data := h.service.UpdateGeneration(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// CreateGenerationImage godoc
// @Summary      Create a new generation image
// @Description  Creates a new generation image
// @Tags         admin-generations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true  "Generation ID"
// @Param        image     formData  file    true   "generation image (max 1)"
// @Success      200       {object}  model.SuccessWithId
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{id}/images [post]
func (h *AdminHandler) CreateGenerationImage(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	form, _ := c.MultipartForm()
	data := h.service.CreateGenerationImage(ctx, form, id)
	return utils.FiberResponse(c, data)
}

// DeleteGeneration godoc
// @Summary      Delete a generation
// @Description  Deletes a generation by ID
// @Tags         admin-generations
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Generation ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{id} [delete]
func (h *AdminHandler) DeleteGeneration(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteGeneration(ctx, id)
	return utils.FiberResponse(c, data)
}

// Generation Modification handlers

// GetGenerationModifications godoc
// @Summary      Get generation modifications
// @Description  Returns a list of generation modifications for a specific generation
// @Tags         admin-generation-modifications
// @Produce      json
// @Security     BearerAuth
// @Param        generation_id   path      int  true  "Generation ID"
// @Success      200  {array}  model.AdminGenerationModificationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{generation_id}/ [get]
func (h *AdminHandler) GetGenerationModifications(c *fiber.Ctx) error {
	generationIdStr := c.Params("generation_id")
	generationId, err := strconv.Atoi(generationIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetGenerationModifications(ctx, generationId)
	return utils.FiberResponse(c, data)
}

// CreateGenerationModification godoc
// @Summary      Create a generation modification
// @Description  Creates a new generation modification
// @Tags         admin-generation-modifications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        generation_id   path      int                                    true  "Generation ID"
// @Param        modification    body      model.CreateGenerationModificationRequest  true  "Generation modification data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{generation_id}/ [post]
func (h *AdminHandler) CreateGenerationModification(c *fiber.Ctx) error {
	generationIdStr := c.Params("generation_id")
	generationId, err := strconv.Atoi(generationIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	var req model.CreateGenerationModificationRequest

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
	data := h.service.CreateGenerationModification(ctx, generationId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateGenerationModification godoc
// @Summary      Update a generation modification
// @Description  Updates a generation modification by ID
// @Tags         admin-generation-modifications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        generation_id   path      int                                    true  "Generation ID"
// @Param        id              path      int                                    true  "Generation modification ID"
// @Param        modification    body      model.UpdateGenerationModificationRequest  true  "Generation modification data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{generation_id}/{id} [put]
func (h *AdminHandler) UpdateGenerationModification(c *fiber.Ctx) error {
	generationIdStr := c.Params("generation_id")
	generationId, err := strconv.Atoi(generationIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation modification id must be integer"),
		})
	}

	var req model.UpdateGenerationModificationRequest

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
	data := h.service.UpdateGenerationModification(ctx, generationId, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteGenerationModification godoc
// @Summary      Delete a generation modification
// @Description  Deletes a generation modification by ID
// @Tags         admin-generation-modifications
// @Produce      json
// @Security     BearerAuth
// @Param        generation_id   path      int  true  "Generation ID"
// @Param        id              path      int  true  "Generation modification ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/generations/{generation_id}/{id} [delete]
func (h *AdminHandler) DeleteGenerationModification(c *fiber.Ctx) error {
	generationIdStr := c.Params("generation_id")
	generationId, err := strconv.Atoi(generationIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("generation modification id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteGenerationModification(ctx, generationId, id)
	return utils.FiberResponse(c, data)
}
