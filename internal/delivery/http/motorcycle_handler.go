package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type MotorcycleHandler struct {
	service *service.MotorcycleService
}

func NewMotorcycleHandler(service *service.MotorcycleService) *MotorcycleHandler {
	return &MotorcycleHandler{service}
}

// GetMotorcycleCategories godoc
// @Summary Get motorcycle categories
// @Description Get motorcycle categories
// @Tags motorcycles
// @Accept json
// @Produce json
// @Success 200 {array} model.GetMotorcycleCategoriesResponse
// @Failure 500 {object} model.ResultMessage
// @Router /motorcycles/categories [get]

func (h *MotorcycleHandler) GetMotorcycleCategories(c *fiber.Ctx) error {

	ctx := c.Context()
	categories, err := h.service.GetMotorcycleCategories(ctx)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 500,
			Error:  err,
		})
	}

	return utils.FiberResponse(c, &model.Response{
		Status: 200,
		Data:   categories,
	})
}

// GetMotorcycleParameters godoc
// @Summary Get motorcycle parameters
// @Description Get motorcycle parameters
// @Tags motorcycles
// @Accept json
// @Produce json
// @Param category_id path string true "Category ID"
// @Success 200 {array} model.GetMotorcycleParametersResponse
// @Failure 500 {object} model.ResultMessage
// @Router /motorcycles/categories/{category_id}/parameters [get]
func (h *MotorcycleHandler) GetMotorcycleParameters(c *fiber.Ctx) error {
	ctx := c.Context()
	categoryID := c.Params("category_id")
	parameters, err := h.service.GetMotorcycleParameters(ctx, categoryID)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 500,
			Error:  err,
		})
	}

	return utils.FiberResponse(c, &model.Response{
		Status: 200,
		Data:   parameters,
	})
}

// GetMotorcycleBrands godoc
// @Summary Get motorcycle brands
// @Description Get motorcycle brands
// @Tags motorcycles
// @Accept json
// @Produce json
// @Param category_id path string true "Category ID"
// @Success 200 {array} model.GetMotorcycleBrandsResponse
// @Failure 500 {object} model.ResultMessage
// @Router /motorcycles/categories/{category_id}/brands [get]
func (h *MotorcycleHandler) GetMotorcycleBrands(c *fiber.Ctx) error {
	ctx := c.Context()
	categoryID := c.Params("category_id")
	brands, err := h.service.GetMotorcycleBrands(ctx, categoryID)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 500,
			Error:  err,
		})
	}

	return utils.FiberResponse(c, &model.Response{
		Status: 200,
		Data:   brands,
	})
}

// GetMotorcycleModelsByBrandID godoc
// @Summary Get motorcycle models by brand ID
// @Description Get motorcycle models by brand ID
// @Tags motorcycles
// @Accept json
// @Produce json
// @Param category_id path string true "Category ID"
// @Param brand_id path string true "Brand ID"
// @Success 200 {array} model.GetMotorcycleModelsResponse
// @Failure 500 {object} model.ResultMessage
// @Router /motorcycles/categories/{category_id}/brands/{brand_id}/models [get]
func (h *MotorcycleHandler) GetMotorcycleModelsByBrandID(c *fiber.Ctx) error {
	ctx := c.Context()
	categoryID := c.Params("category_id")
	brandID := c.Params("brand_id")
	models, err := h.service.GetMotorcycleModelsByBrandID(ctx, categoryID, brandID)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 500,
			Error:  err,
		})
	}

	return utils.FiberResponse(c, &model.Response{
		Status: 200,
		Data:   models,
	})
}

// CreateMotorcycle godoc
// @Summary Create motorcycle
// @Description Create motorcycle
// @Tags motorcycles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param motorcycle body model.CreateMotorcycleRequest true "Motorcycle"
// @Success 200 {object} model.SuccessWithId
// @Failure 500 {object} model.ResultMessage
// @Failure 400 {object} model.ResultMessage
// @Router /motorcycles [post]
func (h *MotorcycleHandler) CreateMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	var motorcycle model.CreateMotorcycleRequest

	if err := c.BodyParser(&motorcycle); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	userID := c.Locals("id").(int)

	createdMotorcycle, err := h.service.CreateMotorcycle(ctx, motorcycle, userID)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 500,
			Error:  err,
		})
	}

	return utils.FiberResponse(c, &model.Response{
		Status: 200,
		Data:   createdMotorcycle,
	})
}

// GetMotorcycles godoc
// @Summary Get motorcycles
// @Description Get motorcycles
// @Tags motorcycles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.GetMotorcyclesResponse
// @Failure 500 {object} model.ResultMessage
// @Router /motorcycles [get]
func (h *MotorcycleHandler) GetMotorcycles(c *fiber.Ctx) error {
	ctx := c.Context()
	motorcycles, err := h.service.GetMotorcycles(ctx)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 500,
			Error:  err,
		})
	}

	return utils.FiberResponse(c, &model.Response{
		Status: 200,
		Data:   motorcycles,
	})
}
