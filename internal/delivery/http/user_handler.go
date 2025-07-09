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
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/profile/my-cars [get]
func (h *UserHandler) GetMyCars(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("id").(int)
	data := h.UserService.GetMyCars(&ctx, &userID)

	utils.GinResponse(c, data)
}

// GetProfileCars godoc
// @Summary      Get user's profile cars
// @Description  Returns the cars associated with the authenticated user's profile
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/profile/on-sale [get]
func (h *UserHandler) OnSale(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("id").(int)
	data := h.UserService.OnSale(&ctx, &userID)

	utils.GinResponse(c, data)
}

// Cancel cars godoc
// @Summary      Get user's cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/cancel [post]
func (h *UserHandler) Cancel(c *gin.Context) {
	// todo: delete images if exist
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}
	data := h.UserService.Cancel(&ctx, &id)
	utils.GinResponse(c, data)
}

// Cancel cars godoc
// @Summary      Get user's cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/delete [post]
func (h *UserHandler) Delete(c *gin.Context) {
	// todo: delete images if exists
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}
	data := h.UserService.Delete(&ctx, &id)
	utils.GinResponse(c, data)
}

// Dont sell godoc
// @Summary      Dont sell cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/dont-sell [post]
func (h *UserHandler) DontSell(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	userID := c.MustGet("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}
	data := h.UserService.DontSell(&ctx, &id, &userID)
	utils.GinResponse(c, data)
}

//	Sell godoc
//
// @Summary       Sell cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/sell [post]
func (h *UserHandler) Sell(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	userID := c.MustGet("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}
	data := h.UserService.Sell(&ctx, &id, &userID)
	utils.GinResponse(c, data)
}

// GetBrands godoc
// @Summary      Get car brands
// @Description  Returns a list of car brands, optionally filtered by text
// @Tags         users
// @Produce      json
// @Param        text  query     string  false  "Filter brands by text"
// @Success      200   {object}  model.GetBrandsResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands [get]
func (h *UserHandler) GetBrands(c *gin.Context) {
	text := c.Query("text")
	ctx := c.Request.Context()
	data := h.UserService.GetBrands(&ctx, text)
	utils.GinResponse(c, data)
}

// GetModifications godoc
// @Summary      Get generation modifications
// @Description  Returns a list of generation modifications
// @Tags         users
// @Produce      json
// @Param        generation_id  query     string  true  "Get modification's generation_id'"
// @Param        body_type_id  query     string  true  "Get modification's body_type_id'"
// @Param        fuel_type_id  query     string  true  "Get modification's fuel_type_id'"
// @Param        drivetrain_id  query     string  true  "Get modification's drivetrain_id'"
// @Param        transmission_id  query     string  true  "Get modification's transmission_id'"
// @Success      200   {array}  model.GetModificationsResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/modifications [get]
func (h *UserHandler) GetModifications(c *gin.Context) {
	generation_id := c.Query("generation_id")
	generationID, err := strconv.Atoi(generation_id)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("generation id must be integer"),
		})
		return
	}
	body_type_id := c.Query("body_type_id")
	bodyTypeID, err := strconv.Atoi(body_type_id)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("body type id must be integer"),
		})
		return
	}
	fuel_type_id := c.Query("fuel_type_id")
	fuelTypeID, err := strconv.Atoi(fuel_type_id)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("fuel type id must be integer"),
		})
		return
	}
	drivetrain_id := c.Query("drivetrain_id")
	drivetrainID, err := strconv.Atoi(drivetrain_id)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("drivetrain_id id must be integer"),
		})
		return
	}
	transmission_id := c.Query("transmission_id")
	transmissionID, err := strconv.Atoi(transmission_id)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("transmission_id id must be integer"),
		})
		return
	}
	ctx := c.Request.Context()
	data := h.UserService.GetModifications(
		&ctx, generationID, bodyTypeID, fuelTypeID, drivetrainID, transmissionID)

	utils.GinResponse(c, data)
}

// GetModelsByBrandID godoc
// @Summary      Get models by brand ID
// @Description  Returns a list of car models for a given brand ID, optionally filtered by text
// @Tags         users
// @Produce      json
// @Param        id    path      int     true   "Brand ID"
// @Param        text  query     string  false  "coroll"
// @Success      200   {object}  model.GetModelsResponse
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
// @Param        id  path  int  true  "brand id ID"
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
// @Param   brands            query   []string  false  "Filter by brand IDs"
// @Param   models            query   []string  false  "Filter by model IDs"
// @Param   regions           query   []string  false  "Filter by region IDs"
// @Param   cities            query   []string  false  "Filter by city IDs"
// @Param   generations       query   []string  false  "Filter by generation IDs"
// @Param   transmissions     query   []string  false  "Filter by transmission IDs"
// @Param   engines           query   []string  false  "Filter by engine IDs"
// @Param   drivetrains       query   []string  false  "Filter by drivetrain IDs"
// @Param   body_types        query   []string  false  "Filter by body type IDs"
// @Param   fuel_types        query   []string  false  "Filter by fuel type IDs"
// @Param   ownership_types   query   []string  false  "Filter by ownership type IDs"
// @Param   announcement_types query  []string  false  "Filter by announcement type IDs"
// @Param   year_from         query   string    false  "Filter by year from"
// @Param   year_to           query   string    false  "Filter by year to"
// @Param   exchange          query   string    false  "Filter by exchange"
// @Param   credit            query   string    false  "Filter by credit"
// @Param   right_hand_drive  query   string    false  "Filter by right hand drive"
// @Param   price_from        query   string    false  "Filter by price from"
// @Param   price_to          query   string    false  "Filter by price to"
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [get]
func (h *UserHandler) GetCars(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("id").(int)
	brands := pkg.QueryParamToArray(c.Query("brands"))
	models := pkg.QueryParamToArray(c.Query("models"))
	regions := pkg.QueryParamToArray(c.Query("regions"))
	cities := pkg.QueryParamToArray(c.Query("cities"))
	generations := pkg.QueryParamToArray(c.Query("generations"))
	transmissions := pkg.QueryParamToArray(c.Query("transmissions"))
	engines := pkg.QueryParamToArray(c.Query("engines"))
	drivetrains := pkg.QueryParamToArray(c.Query("drivetrains"))
	body_types := pkg.QueryParamToArray(c.Query("body_types"))
	fuel_types := pkg.QueryParamToArray(c.Query("fuel_types"))
	ownership_types := pkg.QueryParamToArray(c.Query("ownership_types"))
	announcement_types := pkg.QueryParamToArray(c.Query("announcement_types"))
	year_from := c.Query("year_from")
	year_to := c.Query("year_to")
	exchange := c.Query("exchange")
	credit := c.Query("credit")
	right_hand_drive := c.Query("right_hand_drive")
	price_from := c.Query("price_from")
	price_to := c.Query("price_to")

	data := h.UserService.GetCars(&ctx, userID, brands, models,
		regions, cities, generations, transmissions, engines, drivetrains,
		body_types, fuel_types, ownership_types, announcement_types,
		year_from, year_to, exchange, credit, right_hand_drive, price_from, price_to)

	utils.GinResponse(c, data)
}

// GetCarByID godoc
// @Summary      Get car by ID
// @Description  Returns a car by its ID
// @Tags         users
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id} [get]
func (h *UserHandler) GetCarByID(c *gin.Context) {
	idStr := c.Param("id")
	userID := c.MustGet("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}

	ctx := c.Request.Context()
	data := h.UserService.GetCarByID(&ctx, id, userID)

	utils.GinResponse(c, data)
}

// BuyCar godoc
// @Summary      Buy car
// @Description  Returns a status response message
// @Tags         users
// @Produce      json
// @Security 	 BearerAuth
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object} pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/buy [post]
func (h *UserHandler) BuyCar(c *gin.Context) {
	idStr := c.Param("id")
	userID := c.MustGet("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
		return
	}

	ctx := c.Request.Context()
	data := h.UserService.BuyCar(&ctx, id, userID)
	utils.GinResponse(c, data)
}

// CreateCar godoc
// @Summary      Create a car
// @Description  Creates a new car for the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        car  body      model.CreateCarRequest  true  "Car data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure		 403  {object}  pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [post]
func (h *UserHandler) CreateCar(c *gin.Context) {
	var car model.CreateCarRequest
	car.UserID = c.MustGet("id").(int)
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
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        car_id      path      int     true   "Car CAR_ID"
// @Param        images  formData  file    true   "Car images (max 10)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  pkg.ErrorResponse
// @Failure	 	 403  	 {object}  pkg.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/images [post]
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
