package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
)

type AdminHandler struct {
	AdminService *service.AdminService
	validator    *auth.Validator
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{service, auth.New()}
}

// Cities handlers

// GetCities godoc
// @Summary      Get all cities
// @Description  Returns a list of all cities
// @Tags         admin-cities
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminCityResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cities [get]
func (h *AdminHandler) GetCities(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetCities(ctx)
	return utils.FiberResponse(c, data)
}

// CreateCity godoc
// @Summary      Create a new city
// @Description  Creates a new city
// @Tags         admin-cities
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city  body      model.CreateCityRequest  true  "City data"
// @Success      200   {object}  model.SuccessWithId
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/cities [post]
func (h *AdminHandler) CreateCity(c *fiber.Ctx) error {
	var req model.CreateCityRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.CreateCity(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateCity godoc
// @Summary      Update a city
// @Description  Updates an existing city
// @Tags         admin-cities
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                      true  "City ID"
// @Param        city  body      model.UpdateCityRequest  true  "City data"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{id} [put]
func (h *AdminHandler) UpdateCity(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}

	var req model.UpdateCityRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.UpdateCity(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteCity godoc
// @Summary      Delete a city
// @Description  Deletes a city by ID
// @Tags         admin-cities
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "City ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{id} [delete]
func (h *AdminHandler) DeleteCity(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteCity(ctx, id)
	return utils.FiberResponse(c, data)
}

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
	data := h.AdminService.GetBrands(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.CreateBrand(ctx, &req)
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
// @Param        brand  body      model.UpdateBrandRequest  true  "Brand data"
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	var req model.UpdateBrandRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.UpdateBrand(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteBrand(ctx, id)
	return utils.FiberResponse(c, data)
}

// Models handlers

// GetModels godoc
// @Summary      Get all models
// @Description  Returns a list of all models
// @Tags         admin-models
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminModelResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/models [get]
func (h *AdminHandler) GetModels(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetModels(ctx)
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
// @Router       /api/v1/admin/models [post]
func (h *AdminHandler) CreateModel(c *fiber.Ctx) error {
	var req model.CreateModelRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.CreateModel(ctx, &req)
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
// @Router       /api/v1/admin/models/{id} [put]
func (h *AdminHandler) UpdateModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("model id must be integer"),
		})
	}

	var req model.UpdateModelRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.UpdateModel(ctx, id, &req)
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
// @Router       /api/v1/admin/models/{id} [delete]
func (h *AdminHandler) DeleteModel(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteModel(ctx, id)
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
	data := h.AdminService.GetBodyTypes(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.CreateBodyType(ctx, &req)
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
// @Param        bodyType  body      model.UpdateBodyTypeRequest  true  "Body type data"
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

	var req model.UpdateBodyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.AdminService.UpdateBodyType(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteBodyType(ctx, id)
	return utils.FiberResponse(c, data)
}
