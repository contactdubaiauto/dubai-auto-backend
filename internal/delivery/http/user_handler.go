package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

// GetProfileCars godoc
// @Summary      Get user's profile cars
// @Description  Returns the cars associated with the authenticated user's profile
// @Tags         users
// @Security     ApiKeyAuth
// @Produce      json
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/profile/cars [get]
func (h *UserHandler) GetProfileCars(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("id").(int)
	data := h.UserService.GetProfileCars(&ctx, &userID)
	// brands := c.Query("brands")
	// models := c.Query("models")
	// cities := c.Query("cities")
	// regions := c.Query("regions")

	utils.GinResponse(c, data)
}

// GetBrands godoc
// @Summary      Get car brands
// @Description  Returns a list of car brands, optionally filtered by text
// @Tags         users
// @Produce      json
// @Param        text  query     string  false  "Filter brands by text"
// @Success      200   {object}  model.ResultMessage
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands [get]
func (h *UserHandler) GetBrands(c *gin.Context) {
	text := c.Query("text")
	ctx := c.Request.Context()
	brands, err := h.UserService.GetBrands(&ctx, text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Status: 200,
		Data:   brands,
	})
}

// GetModelsByBrandID godoc
// @Summary      Get models by brand ID
// @Description  Returns a list of car models for a given brand ID, optionally filtered by text
// @Tags         users
// @Produce      json
// @Param        id    path      int     true   "Brand ID"
// @Param        text  query     string  false  "coroll"
// @Success      200   {array}  model.Model
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/models [get]
func (h *UserHandler) GetModelsByBrandID(c *gin.Context) {
	brandID := c.Param("id")
	text := c.Query("text")
	brandIDInt, err := strconv.ParseInt(brandID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx := c.Request.Context()
	data := h.UserService.GetModelsByBrandID(&ctx, brandIDInt, text)

	utils.GinResponse(c, data)
}

// GetGenerationsByModelID godoc
// @Summary      Get generations by model ID
// @Description  Returns a list of generations for a given model ID
// @Tags         users
// @Produce      json
// @Param        model_id  path  int  true  "Model ID"
// @Success      200   {array}  model.Generation
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/models/{model_id}/generations [get]
func (h *UserHandler) GetGenerationsByModelID(c *gin.Context) {
	modelID := c.Param("model_id")
	modelIDInt, err := strconv.ParseInt(modelID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	ctx := c.Request.Context()
	data := h.UserService.GetGenerationsByModelID(&ctx, modelIDInt)

	utils.GinResponse(c, data)
}

// GetBodyTypes godoc
// @Summary      Get body types
// @Description  Returns a list of car body types
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.BodyType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/body-types [get]
func (h *UserHandler) GetBodyTypes(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetBodyTypes(&ctx)
	utils.GinResponse(c, data)
}

// GetTransmissions godoc
// @Summary      Get transmissions
// @Description  Returns a list of car transmissions
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.Transmission
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/transmissions [get]
func (h *UserHandler) GetTransmissions(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetTransmissions(&ctx)
	utils.GinResponse(c, data)
}

// GetEngines godoc
// @Summary      Get engines
// @Description  Returns a list of car engines
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.Engine
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/engines [get]
func (h *UserHandler) GetEngines(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetEngines(&ctx)

	utils.GinResponse(c, data)
}

// GetDrivetrains godoc
// @Summary      Get drivetrains
// @Description  Returns a list of car drivetrains
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.Drivetrain
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/drivetrains [get]
func (h *UserHandler) GetDrivetrains(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetDrivetrains(&ctx)

	utils.GinResponse(c, data)
}

// GetFuelTypes godoc
// @Summary      Get fuel types
// @Description  Returns a list of car fuel types
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.FuelType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/fuel-types [get]
func (h *UserHandler) GetFuelTypes(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetFuelTypes(&ctx)

	utils.GinResponse(c, data)
}

// GetColors godoc
// @Summary      Get colors
// @Description  Returns a list of car colors
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.Color
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/colors [get]
func (h *UserHandler) GetColors(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetColors(&ctx)
	utils.GinResponse(c, data)
}

// GetCars godoc
// @Summary      Get cars
// @Description  Returns a list of cars
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [get]
func (h *UserHandler) GetCars(c *gin.Context) {
	ctx := c.Request.Context()
	data := h.UserService.GetCars(&ctx)
	// brands := c.Query("brands")
	// models := c.Query("models")
	// cities := c.Query("cities")
	// regions := c.Query("regions")

	utils.GinResponse(c, data)
}

// GetCarByID godoc
// @Summary      Get car by ID
// @Description  Returns a car by its ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{id} [get]
func (h *UserHandler) GetCarByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}

	ctx := c.Request.Context()
	data := h.UserService.GetCarByID(&ctx, id)

	utils.GinResponse(c, data)
}

// CreateCar godoc
// @Summary      Create a car
// @Description  Creates a new car for the authenticated user
// @Tags         users
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        car  body      model.CreateCarRequest  true  "Car data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [post]
func (h *UserHandler) CreateCar(c *gin.Context) {
	var car model.CreateCarRequest
	userID := c.MustGet("id").(int)
	car.UserID = int64(userID)
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&car); err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid request data" + err.Error()),
		})
		return
	}

	data := h.UserService.CreateCar(&ctx, &car)
	utils.GinResponse(c, data)
}

// CreateCarImages godoc
// @Summary      Upload car images
// @Description  Uploads images for a car (max 10 files)
// @Tags         users
// @Security     ApiKeyAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id      path      int     true   "Car ID"
// @Param        images  formData  file    true   "Car images (max 10)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{id}/images [post]
func (h *UserHandler) CreateCarImages(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid car ID"),
		})
		return
	}

	form, _ := c.MultipartForm()

	if form == nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
		return
	}
	images := form.File["images"]

	if len(images) > 10 {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 10 files"),
		})
		return
	}

	paths, status, err := pkg.SaveFiles(images, config.ENV.STATIC_PATH+"cars/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: status,
			Error:  err,
		})
		return
	}

	data := h.UserService.CreateCarImages(&ctx, id, paths)
	utils.GinResponse(c, data)
}
