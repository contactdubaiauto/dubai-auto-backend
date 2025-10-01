package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/files"
)

type AdminHandler struct {
	AdminService *service.AdminService
	validator    *auth.Validator
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{service, auth.NewValidator()}
}

// Application handlers

// GetApplications godoc
// @Summary      Get all applications
// @Description  Returns a list of all applications
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminApplicationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications [get]
func (h *AdminHandler) GetApplications(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetApplications(ctx)
	return utils.FiberResponse(c, data)
}

// GetApplication godoc
// @Summary      Get an application
// @Description  Returns an application by ID
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Application ID"
// @Success      200  {object}  model.AdminApplicationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications/{id} [get]
func (h *AdminHandler) GetApplication(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("application id must be integer"),
		})
	}
	ctx := c.Context()
	data := h.AdminService.GetApplication(ctx, id)
	return utils.FiberResponse(c, data)
}

// AcceptApplication godoc
// @Summary      Accept an application
// @Description  Accepts an application by ID
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Application ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications/{id}/accept [put]
func (h *AdminHandler) AcceptApplication(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("application id must be integer"),
		})
	}
	ctx := c.Context()
	data := h.AdminService.AcceptApplication(ctx, id)
	return utils.FiberResponse(c, data)
}

// RejectApplication godoc
// @Summary      Reject an application
// @Description  Rejects an application by ID
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Application ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications/{id}/reject [delete]
func (h *AdminHandler) RejectApplication(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("application id must be integer"),
		})
	}
	ctx := c.Context()
	data := h.AdminService.RejectApplication(ctx, id)
	return utils.FiberResponse(c, data)
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
// @Param        city  body      model.CreateNameRequest  true  "City data"
// @Success      200   {object}  model.SuccessWithId
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/cities [post]
func (h *AdminHandler) CreateCity(c *fiber.Ctx) error {
	var req model.CreateNameRequest
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
// @Param        city  body      model.CreateNameRequest  true  "City data"
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

	var req model.CreateNameRequest
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

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
// @Router       /api/v1/admin/brands/:brand_id/models [get]
func (h *AdminHandler) GetModels(c *fiber.Ctx) error {
	ctx := c.Context()
	brandIdStr := c.Params("brand_id")
	brandId, err := strconv.Atoi(brandIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}
	data := h.AdminService.GetModels(ctx, brandId)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("brand id must be integer"),
		})
	}

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

	data := h.AdminService.CreateModel(ctx, brandId, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("model id must be integer"),
		})
	}

	var req model.UpdateModelRequest

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
// @Router       /api/v1/admin/brands/:brand_id/models/{id} [delete]
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

// CreateBodyTypeImage godoc
// @Summary      Create a new body type image
// @Description  Creates a new body type image
// @Tags         admin-body-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int                          true  "Body type ID"
// @Param        image     formData  file    true   "body type image (max 1)"
// @Success      200       {object}  model.SuccessWithId
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}

	image := form.File["image"]

	if len(image) > 10 {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 10 files"),
		})
	}

	paths, status, err := files.SaveFiles(image, config.ENV.STATIC_PATH+"cars/body/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: status,
			Error:  err,
		})
	}

	data := h.AdminService.CreateBodyTypeImage(ctx, id, paths)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
	}

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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("body type image id must be integer"),
		})
	}
	data := h.AdminService.DeleteBodyTypeImage(ctx, id)
	return utils.FiberResponse(c, data)
}

// Regions handlers

// GetRegions godoc
// @Summary      Get all regions
// @Description  Returns a list of all regions
// @Tags         admin-regions
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Success      200  {array}  model.AdminCityResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions [get]
func (h *AdminHandler) GetRegions(c *fiber.Ctx) error {
	cityIdStr := c.Params("city_id")
	cityId, err := strconv.Atoi(cityIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}
	ctx := c.Context()
	data := h.AdminService.GetRegions(ctx, cityId)
	return utils.FiberResponse(c, data)
}

// CreateRegion godoc
// @Summary      Create a new region
// @Description  Creates a new region
// @Tags         admin-regions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Param        region  body      model.CreateNameRequest  true  "Region data"
// @Success      200    {object}  model.SuccessWithId
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions [post]
func (h *AdminHandler) CreateRegion(c *fiber.Ctx) error {
	cityIdStr := c.Params("city_id")
	cityId, err := strconv.Atoi(cityIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}

	var req model.CreateNameRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateRegion(ctx, cityId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateRegion godoc
// @Summary      Update a region
// @Description  Updates an existing region
// @Tags         admin-regions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Param        id        path      int  true  "Region ID"
// @Param        region    body      model.CreateNameRequest  true  "Region data"
// @Success      200       {object}  model.Success
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions/{id} [put]
func (h *AdminHandler) UpdateRegion(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("region id must be integer"),
		})
	}

	var req model.CreateNameRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateRegion(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteRegion godoc
// @Summary      Delete a region
// @Description  Deletes an existing region
// @Tags         admin-regions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Param        id        path      int  true  "Region ID"
// @Success      200       {object}  model.Success
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions/{id} [delete]
func (h *AdminHandler) DeleteRegion(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("region id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteRegion(ctx, id)
	return utils.FiberResponse(c, data)
}

// Transmission handlers

// GetTransmissions godoc
// @Summary      Get all transmissions
// @Description  Returns a list of all transmissions
// @Tags         admin-transmissions
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminTransmissionResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/transmissions [get]
func (h *AdminHandler) GetTransmissions(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetTransmissions(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateTransmission(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("transmission id must be integer"),
		})
	}

	var req model.CreateTransmissionRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateTransmission(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("transmission id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteTransmission(ctx, id)
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
	data := h.AdminService.GetEngines(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateEngine(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("engine id must be integer"),
		})
	}

	var req model.CreateEngineRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateEngine(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("engine id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteEngine(ctx, id)
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
	data := h.AdminService.GetDrivetrains(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateDrivetrain(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("drivetrain id must be integer"),
		})
	}

	var req model.CreateDrivetrainRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateDrivetrain(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("drivetrain id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteDrivetrain(ctx, id)
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
	data := h.AdminService.GetFuelTypes(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateFuelType(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("fuel type id must be integer"),
		})
	}

	var req model.CreateFuelTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateFuelType(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("fuel type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteFuelType(ctx, id)
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
	data := h.AdminService.GetMotoCategories(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateMotoCategory(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto category id must be integer"),
		})
	}

	var req model.UpdateMotoCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateMotoCategory(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteMotoCategory(ctx, id)
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
	data := h.AdminService.GetMotoBrands(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateMotoBrand(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto brand id must be integer"),
		})
	}

	var req model.UpdateMotoBrandRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateMotoBrand(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteMotoBrand(ctx, id)
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
	data := h.AdminService.GetMotoModels(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateMotoModel(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto model id must be integer"),
		})
	}

	var req model.UpdateMotoModelRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateMotoModel(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteMotoModel(ctx, id)
	return utils.FiberResponse(c, data)
}

// Moto Parameters handlers

// GetMotoParameters godoc
// @Summary      Get all moto parameters
// @Description  Returns a list of all moto parameters
// @Tags         admin-moto-parameters
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminMotoParameterResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters [get]
func (h *AdminHandler) GetMotoParameters(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetMotoParameters(ctx)
	return utils.FiberResponse(c, data)
}

// CreateMotoParameter godoc
// @Summary      Create a moto parameter
// @Description  Creates a new moto parameter
// @Tags         admin-moto-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        motoParameter  body      model.CreateMotoParameterRequest  true  "Moto parameter data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters [post]
func (h *AdminHandler) CreateMotoParameter(c *fiber.Ctx) error {
	var req model.CreateMotoParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateMotoParameter(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateMotoParameter godoc
// @Summary      Update a moto parameter
// @Description  Updates a moto parameter by ID
// @Tags         admin-moto-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      int                            true  "Moto parameter ID"
// @Param        motoParameter  body      model.UpdateMotoParameterRequest  true  "Moto parameter data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters/{id} [put]
func (h *AdminHandler) UpdateMotoParameter(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter id must be integer"),
		})
	}

	var req model.UpdateMotoParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateMotoParameter(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteMotoParameter godoc
// @Summary      Delete a moto parameter
// @Description  Deletes a moto parameter by ID
// @Tags         admin-moto-parameters
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Moto parameter ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters/{id} [delete]
func (h *AdminHandler) DeleteMotoParameter(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteMotoParameter(ctx, id)
	return utils.FiberResponse(c, data)
}

// Moto Parameter Values handlers

// GetMotoParameterValues godoc
// @Summary      Get moto parameter values
// @Description  Returns a list of moto parameter values for a specific parameter
// @Tags         admin-moto-parameter-values
// @Produce      json
// @Security     BearerAuth
// @Param        moto_param_id   path      int  true  "Moto parameter ID"
// @Success      200  {array}  model.AdminMotoParameterValueResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values [get]
func (h *AdminHandler) GetMotoParameterValues(c *fiber.Ctx) error {
	motoParamIdStr := c.Params("moto_param_id")
	motoParamId, err := strconv.Atoi(motoParamIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.GetMotoParameterValues(ctx, motoParamId)
	return utils.FiberResponse(c, data)
}

// CreateMotoParameterValue godoc
// @Summary      Create a moto parameter value
// @Description  Creates a new moto parameter value
// @Tags         admin-moto-parameter-values
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        moto_param_id   path      int                                   true  "Moto parameter ID"
// @Param        parameterValue  body      model.CreateMotoParameterValueRequest  true  "Moto parameter value data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values [post]
func (h *AdminHandler) CreateMotoParameterValue(c *fiber.Ctx) error {
	motoParamIdStr := c.Params("moto_param_id")
	motoParamId, err := strconv.Atoi(motoParamIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter id must be integer"),
		})
	}

	var req model.CreateMotoParameterValueRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateMotoParameterValue(ctx, motoParamId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateMotoParameterValue godoc
// @Summary      Update a moto parameter value
// @Description  Updates a moto parameter value by ID
// @Tags         admin-moto-parameter-values
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        moto_param_id   path      int                                   true  "Moto parameter ID"
// @Param        id              path      int                                   true  "Moto parameter value ID"
// @Param        parameterValue  body      model.UpdateMotoParameterValueRequest  true  "Moto parameter value data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values/{id} [put]
func (h *AdminHandler) UpdateMotoParameterValue(c *fiber.Ctx) error {
	motoParamIdStr := c.Params("moto_param_id")
	motoParamId, err := strconv.Atoi(motoParamIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter value id must be integer"),
		})
	}

	var req model.UpdateMotoParameterValueRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateMotoParameterValue(ctx, motoParamId, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteMotoParameterValue godoc
// @Summary      Delete a moto parameter value
// @Description  Deletes a moto parameter value by ID
// @Tags         admin-moto-parameter-values
// @Produce      json
// @Security     BearerAuth
// @Param        moto_param_id   path      int  true  "Moto parameter ID"
// @Param        id              path      int  true  "Moto parameter value ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-parameters/{moto_param_id}/values/{id} [delete]
func (h *AdminHandler) DeleteMotoParameterValue(c *fiber.Ctx) error {
	motoParamIdStr := c.Params("moto_param_id")
	motoParamId, err := strconv.Atoi(motoParamIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto parameter value id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteMotoParameterValue(ctx, motoParamId, id)
	return utils.FiberResponse(c, data)
}

// Moto Category Parameters handlers

// GetMotoCategoryParameters godoc
// @Summary      Get moto category parameters
// @Description  Returns a list of moto category parameters for a specific category
// @Tags         admin-moto-category-parameters
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int  true  "Moto category ID"
// @Success      200  {array}  model.AdminMotoCategoryParameterResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{category_id}/parameters [get]
func (h *AdminHandler) GetMotoCategoryParameters(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.GetMotoCategoryParameters(ctx, categoryId)
	return utils.FiberResponse(c, data)
}

// CreateMotoCategoryParameter godoc
// @Summary      Create a moto category parameter
// @Description  Creates a new moto category parameter
// @Tags         admin-moto-category-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int                                     true  "Moto category ID"
// @Param        parameter     body      model.CreateMotoCategoryParameterRequest  true  "Moto category parameter data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{category_id}/parameters [post]
func (h *AdminHandler) CreateMotoCategoryParameter(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	var req model.CreateMotoCategoryParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateMotoCategoryParameter(ctx, categoryId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateMotoCategoryParameter godoc
// @Summary      Update a moto category parameter
// @Description  Updates a moto category parameter by ID
// @Tags         admin-moto-category-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int                                     true  "Moto category ID"
// @Param        id            path      int                                     true  "Moto category parameter ID"
// @Param        parameter     body      model.UpdateMotoCategoryParameterRequest  true  "Moto category parameter data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{category_id}/parameters/{id} [put]
func (h *AdminHandler) UpdateMotoCategoryParameter(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto category parameter id must be integer"),
		})
	}

	var req model.UpdateMotoCategoryParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateMotoCategoryParameter(ctx, categoryId, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteMotoCategoryParameter godoc
// @Summary      Delete a moto category parameter
// @Description  Deletes a moto category parameter by ID
// @Tags         admin-moto-category-parameters
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int  true  "Moto category ID"
// @Param        id            path      int  true  "Moto category parameter ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/moto-categories/{category_id}/parameters/{id} [delete]
func (h *AdminHandler) DeleteMotoCategoryParameter(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("moto category parameter id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteMotoCategoryParameter(ctx, categoryId, id)
	return utils.FiberResponse(c, data)
}

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
	data := h.AdminService.GetComtransCategories(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateComtransCategory(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans category id must be integer"),
		})
	}

	var req model.UpdateComtransCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateComtransCategory(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteComtransCategory(ctx, id)
	return utils.FiberResponse(c, data)
}

// Comtrans Category Parameters handlers

// GetComtransCategoryParameters godoc
// @Summary      Get comtrans category parameters
// @Description  Returns a list of comtrans category parameters for a specific category
// @Tags         admin-comtrans-category-parameters
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int  true  "Comtrans category ID"
// @Success      200  {array}  model.AdminComtransCategoryParameterResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters [get]
func (h *AdminHandler) GetComtransCategoryParameters(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.GetComtransCategoryParameters(ctx, categoryId)
	return utils.FiberResponse(c, data)
}

// CreateComtransCategoryParameter godoc
// @Summary      Create a comtrans category parameter
// @Description  Creates a new comtrans category parameter
// @Tags         admin-comtrans-category-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int                                         true  "Comtrans category ID"
// @Param        parameter     body      model.CreateComtransCategoryParameterRequest  true  "Comtrans category parameter data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters [post]
func (h *AdminHandler) CreateComtransCategoryParameter(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	var req model.CreateComtransCategoryParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateComtransCategoryParameter(ctx, categoryId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateComtransCategoryParameter godoc
// @Summary      Update a comtrans category parameter
// @Description  Updates a comtrans category parameter by ID
// @Tags         admin-comtrans-category-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int                                         true  "Comtrans category ID"
// @Param        id            path      int                                         true  "Comtrans category parameter ID"
// @Param        parameter     body      model.UpdateComtransCategoryParameterRequest  true  "Comtrans category parameter data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters/{id} [put]
func (h *AdminHandler) UpdateComtransCategoryParameter(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans category parameter id must be integer"),
		})
	}

	var req model.UpdateComtransCategoryParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateComtransCategoryParameter(ctx, categoryId, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteComtransCategoryParameter godoc
// @Summary      Delete a comtrans category parameter
// @Description  Deletes a comtrans category parameter by ID
// @Tags         admin-comtrans-category-parameters
// @Produce      json
// @Security     BearerAuth
// @Param        category_id   path      int  true  "Comtrans category ID"
// @Param        id            path      int  true  "Comtrans category parameter ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-categories/{category_id}/parameters/{id} [delete]
func (h *AdminHandler) DeleteComtransCategoryParameter(c *fiber.Ctx) error {
	categoryIdStr := c.Params("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("category id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans category parameter id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteComtransCategoryParameter(ctx, categoryId, id)
	return utils.FiberResponse(c, data)
}

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
	data := h.AdminService.GetComtransBrands(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateComtransBrand(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans brand id must be integer"),
		})
	}

	var req model.UpdateComtransBrandRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateComtransBrand(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans brand id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteComtransBrand(ctx, id)
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
	data := h.AdminService.GetComtransModels(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateComtransModel(ctx, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans model id must be integer"),
		})
	}

	var req model.UpdateComtransModelRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateComtransModel(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans model id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteComtransModel(ctx, id)
	return utils.FiberResponse(c, data)
}

// Comtrans Parameters handlers

// GetComtransParameters godoc
// @Summary      Get all comtrans parameters
// @Description  Returns a list of all comtrans parameters
// @Tags         admin-comtrans-parameters
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminComtransParameterResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters [get]
func (h *AdminHandler) GetComtransParameters(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetComtransParameters(ctx)
	return utils.FiberResponse(c, data)
}

// CreateComtransParameter godoc
// @Summary      Create a comtrans parameter
// @Description  Creates a new comtrans parameter
// @Tags         admin-comtrans-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        comtransParameter  body      model.CreateComtransParameterRequest  true  "Comtrans parameter data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters [post]
func (h *AdminHandler) CreateComtransParameter(c *fiber.Ctx) error {
	var req model.CreateComtransParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateComtransParameter(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateComtransParameter godoc
// @Summary      Update a comtrans parameter
// @Description  Updates a comtrans parameter by ID
// @Tags         admin-comtrans-parameters
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id                 path      int                                true  "Comtrans parameter ID"
// @Param        comtransParameter  body      model.UpdateComtransParameterRequest  true  "Comtrans parameter data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters/{id} [put]
func (h *AdminHandler) UpdateComtransParameter(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans parameter id must be integer"),
		})
	}

	var req model.UpdateComtransParameterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateComtransParameter(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteComtransParameter godoc
// @Summary      Delete a comtrans parameter
// @Description  Deletes a comtrans parameter by ID
// @Tags         admin-comtrans-parameters
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Comtrans parameter ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters/{id} [delete]
func (h *AdminHandler) DeleteComtransParameter(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans parameter id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteComtransParameter(ctx, id)
	return utils.FiberResponse(c, data)
}

// Comtrans Parameter Values handlers

// GetComtransParameterValues godoc
// @Summary      Get comtrans parameter values
// @Description  Returns a list of comtrans parameter values for a specific parameter
// @Tags         admin-comtrans-parameter-values
// @Produce      json
// @Security     BearerAuth
// @Param        parameter_id   path      int  true  "Comtrans parameter ID"
// @Success      200  {array}  model.AdminComtransParameterValueResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values [get]
func (h *AdminHandler) GetComtransParameterValues(c *fiber.Ctx) error {
	parameterIdStr := c.Params("parameter_id")
	parameterId, err := strconv.Atoi(parameterIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("parameter id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.GetComtransParameterValues(ctx, parameterId)
	return utils.FiberResponse(c, data)
}

// CreateComtransParameterValue godoc
// @Summary      Create a comtrans parameter value
// @Description  Creates a new comtrans parameter value
// @Tags         admin-comtrans-parameter-values
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        parameter_id   path      int                                       true  "Comtrans parameter ID"
// @Param        parameterValue body      model.CreateComtransParameterValueRequest  true  "Comtrans parameter value data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values [post]
func (h *AdminHandler) CreateComtransParameterValue(c *fiber.Ctx) error {
	parameterIdStr := c.Params("parameter_id")
	parameterId, err := strconv.Atoi(parameterIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("parameter id must be integer"),
		})
	}

	var req model.CreateComtransParameterValueRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateComtransParameterValue(ctx, parameterId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateComtransParameterValue godoc
// @Summary      Update a comtrans parameter value
// @Description  Updates a comtrans parameter value by ID
// @Tags         admin-comtrans-parameter-values
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        parameter_id   path      int                                       true  "Comtrans parameter ID"
// @Param        id             path      int                                       true  "Comtrans parameter value ID"
// @Param        parameterValue body      model.UpdateComtransParameterValueRequest  true  "Comtrans parameter value data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values/{id} [put]
func (h *AdminHandler) UpdateComtransParameterValue(c *fiber.Ctx) error {
	parameterIdStr := c.Params("parameter_id")
	parameterId, err := strconv.Atoi(parameterIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("parameter id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans parameter value id must be integer"),
		})
	}

	var req model.UpdateComtransParameterValueRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateComtransParameterValue(ctx, parameterId, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteComtransParameterValue godoc
// @Summary      Delete a comtrans parameter value
// @Description  Deletes a comtrans parameter value by ID
// @Tags         admin-comtrans-parameter-values
// @Produce      json
// @Security     BearerAuth
// @Param        parameter_id   path      int  true  "Comtrans parameter ID"
// @Param        id             path      int  true  "Comtrans parameter value ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/comtrans-parameters/{parameter_id}/values/{id} [delete]
func (h *AdminHandler) DeleteComtransParameterValue(c *fiber.Ctx) error {
	parameterIdStr := c.Params("parameter_id")
	parameterId, err := strconv.Atoi(parameterIdStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("parameter id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("comtrans parameter value id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteComtransParameterValue(ctx, parameterId, id)
	return utils.FiberResponse(c, data)
}

// Service Type handlers

// GetServiceTypes godoc
// @Summary      Get all service types
// @Description  Returns a list of all service types
// @Tags         admin-service-types
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminServiceTypeResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/service-types [get]
func (h *AdminHandler) GetServiceTypes(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetServiceTypes(ctx)
	return utils.FiberResponse(c, data)
}

// CreateServiceType godoc
// @Summary      Create a service type
// @Description  Creates a new service type
// @Tags         admin-service-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        serviceType  body      model.CreateServiceTypeRequest  true  "Service type data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/service-types [post]
func (h *AdminHandler) CreateServiceType(c *fiber.Ctx) error {
	var req model.CreateServiceTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateServiceType(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateServiceType godoc
// @Summary      Update a service type
// @Description  Updates a service type by ID
// @Tags         admin-service-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id           path      int                          true  "Service type ID"
// @Param        serviceType  body      model.CreateServiceTypeRequest  true  "Service type data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/service-types/{id} [put]
func (h *AdminHandler) UpdateServiceType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("service type id must be integer"),
		})
	}

	var req model.CreateServiceTypeRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateServiceType(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteServiceType godoc
// @Summary      Delete a service type
// @Description  Deletes a service type by ID
// @Tags         admin-service-types
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Service type ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/service-types/{id} [delete]
func (h *AdminHandler) DeleteServiceType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("service type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteServiceType(ctx, id)
	return utils.FiberResponse(c, data)
}

// Service handlers

// GetServices godoc
// @Summary      Get all services
// @Description  Returns a list of all services
// @Tags         admin-services
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminServiceResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/services [get]
func (h *AdminHandler) GetServices(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetServices(ctx)
	return utils.FiberResponse(c, data)
}

// CreateService godoc
// @Summary      Create a service
// @Description  Creates a new service
// @Tags         admin-services
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        service  body      model.CreateServiceRequest  true  "Service data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/services [post]
func (h *AdminHandler) CreateService(c *fiber.Ctx) error {
	var req model.CreateServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateService(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateService godoc
// @Summary      Update a service
// @Description  Updates a service by ID
// @Tags         admin-services
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int                     true  "Service ID"
// @Param        service  body      model.CreateServiceRequest  true  "Service data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/services/{id} [put]
func (h *AdminHandler) UpdateService(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("service id must be integer"),
		})
	}

	var req model.CreateServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateService(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteService godoc
// @Summary      Delete a service
// @Description  Deletes a service by ID
// @Tags         admin-services
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Service ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/services/{id} [delete]
func (h *AdminHandler) DeleteService(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("service id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteService(ctx, id)
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
	data := h.AdminService.GetGenerations(ctx)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateGeneration(ctx, &req)
	return utils.FiberResponse(c, data)
}

// // CreateGenerationImage godoc
// // @Summary      Create a generation image
// // @Description  Creates a new generation image
// // @Tags         admin-generations
// // @Accept       json
// // @Produce      json
// // @Security     BearerAuth
// // @Param        id          path      int                        true  "Generation ID"
// // @Param        image       formData  file    true   "Generation image (max 1)"
// // @Success      200  {object}  model.SuccessWithId
// // @Failure      400  {object}  model.ResultMessage
// // @Failure      401  {object}  auth.ErrorResponse
// // @Failure      403  {object}  auth.ErrorResponse
// // @Failure      500  {object}  model.ResultMessage
// // @Router       /api/v1/admin/generations/{id}/images [post]
// func (h *AdminHandler) CreateGenerationImage(c *fiber.Ctx) error {
// 	ctx := c.Context()
// 	idStr := c.Params("id")
// 	id, err := strconv.Atoi(idStr)

// 	if err != nil {
// 		return utils.FiberResponse(c, &model.Response{
// 			Status: 400,
// 			Error:  errors.New("generation id must be integer"),
// 		})
// 	}

// 	form, _ := c.MultipartForm()

// 	if form == nil {
// 		return utils.FiberResponse(c, &model.Response{
// 			Status: 400,
// 			Error:  errors.New("didn't upload the files"),
// 		})
// 	}

// 	image := form.File["image"]

// 	if len(image) > 1 {
// 		return utils.FiberResponse(c, &model.Response{
// 			Status: 400,
// 			Error:  errors.New("must load maximum 1 file"),
// 		})
// 	}

// 	paths, status, err := files.SaveFiles(image, config.ENV.STATIC_PATH+"cars/generations/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)

// 	if err != nil {
// 		return utils.FiberResponse(c, &model.Response{
// 			Status: status,
// 			Error:  err,
// 		})
// 	}

// 	data := h.AdminService.CreateGenerationImage(ctx, id, paths)
// 	return utils.FiberResponse(c, data)
// }

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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	var req model.UpdateGenerationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateGeneration(ctx, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteGeneration(ctx, id)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.GetGenerationModifications(ctx, generationId)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	var req model.CreateGenerationModificationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateGenerationModification(ctx, generationId, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation modification id must be integer"),
		})
	}

	var req model.UpdateGenerationModificationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateGenerationModification(ctx, generationId, id, &req)
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
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation modification id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteGenerationModification(ctx, generationId, id)
	return utils.FiberResponse(c, data)
}

// Configuration handlers

// GetConfigurations godoc
// @Summary      Get all configurations
// @Description  Returns a list of all configurations
// @Tags         admin-configurations
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminConfigurationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/configurations [get]
func (h *AdminHandler) GetConfigurations(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetConfigurations(ctx)
	return utils.FiberResponse(c, data)
}

// CreateConfiguration godoc
// @Summary      Create a configuration
// @Description  Creates a new configuration
// @Tags         admin-configurations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        configuration  body      model.CreateConfigurationRequest  true  "Configuration data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/configurations [post]
func (h *AdminHandler) CreateConfiguration(c *fiber.Ctx) error {
	var req model.CreateConfigurationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateConfiguration(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateConfiguration godoc
// @Summary      Update a configuration
// @Description  Updates a configuration by ID
// @Tags         admin-configurations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      int                           true  "Configuration ID"
// @Param        configuration  body      model.UpdateConfigurationRequest  true  "Configuration data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/configurations/{id} [put]
func (h *AdminHandler) UpdateConfiguration(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("configuration id must be integer"),
		})
	}

	var req model.UpdateConfigurationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateConfiguration(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteConfiguration godoc
// @Summary      Delete a configuration
// @Description  Deletes a configuration by ID
// @Tags         admin-configurations
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Configuration ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/configurations/{id} [delete]
func (h *AdminHandler) DeleteConfiguration(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("configuration id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteConfiguration(ctx, id)
	return utils.FiberResponse(c, data)
}

// Color handlers

// GetColors godoc
// @Summary      Get all colors
// @Description  Returns a list of all colors
// @Tags         admin-colors
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminColorResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors [get]
func (h *AdminHandler) GetColors(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.AdminService.GetColors(ctx)
	return utils.FiberResponse(c, data)
}

// CreateColor godoc
// @Summary      Create a color
// @Description  Creates a new color
// @Tags         admin-colors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        color  body      model.CreateColorRequest  true  "Color data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors [post]
func (h *AdminHandler) CreateColor(c *fiber.Ctx) error {
	var req model.CreateColorRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.CreateColor(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateColor godoc
// @Summary      Update a color
// @Description  Updates a color by ID
// @Tags         admin-colors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                   true  "Color ID"
// @Param        color  body      model.UpdateColorRequest  true  "Color data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors/{id} [put]
func (h *AdminHandler) UpdateColor(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("color id must be integer"),
		})
	}

	var req model.UpdateColorRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.AdminService.UpdateColor(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteColor godoc
// @Summary      Delete a color
// @Description  Deletes a color by ID
// @Tags         admin-colors
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Color ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors/{id} [delete]
func (h *AdminHandler) DeleteColor(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("color id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.AdminService.DeleteColor(ctx, id)
	return utils.FiberResponse(c, data)
}
