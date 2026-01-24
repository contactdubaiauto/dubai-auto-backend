package http

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/files"
)

type UserHandler struct {
	UserService       *service.UserService
	MotorcycleService *service.MotorcycleService
	ComtransService   *service.ComtransService
	validator         *auth.Validator
}

func NewUserHandler(userService *service.UserService, motorcycleService *service.MotorcycleService, comtransService *service.ComtransService, validator *auth.Validator) *UserHandler {
	return &UserHandler{
		UserService:       userService,
		MotorcycleService: motorcycleService,
		ComtransService:   comtransService,
		validator:         validator,
	}
}

// GetCars godoc
// @Summary      Get cars
// @Description  Returns a list of cars
// @Tags         car
// @Produce      json
// @Security 	 BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param   brands            query   string  false  "Filter by brand IDs"
// @Param   models            query   string  false  "Filter by model IDs"
// @Param   regions           query   string  false  "Filter by region IDs"
// @Param   cities            query   string  false  "Filter by city IDs"
// @Param   generations       query   string  false  "Filter by generation IDs"
// @Param   colors       	  query   string  false  "Filter by color IDs"
// @Param   crash       	  query   string  false  "Filter by crash status, true or empty"
// @Param   transmissions     query   string  false  "Filter by transmission IDs"
// @Param   engines           query   string  false  "Filter by engine IDs"
// @Param   drivetrains       query   string  false  "Filter by drivetrain IDs"
// @Param   body_types        query   string  false  "Filter by body type IDs"
// @Param   fuel_types        query   string  false  "Filter by fuel type IDs"
// @Param   trade_in          query   string  false  "Filter by trade_in id, from 1 to 5"
// @Param   owners            query   string  false  "Filter by owners id, from 1 to 4"
// @Param   ownership_types   query   string  false  "Filter by ownership type IDs"
// @Param   year_from         query   string  false  "Filter by year from"
// @Param   year_to           query   string  false  "Filter by year to"
// @Param   credit            query   string  false  "Filter by credit"
// @Param  	new   		      query   string  false  "true or false new"
// @Param  	wheel   		  query   string  false  "true or false wheel"
// @Param  	odometer   	      query   string  false  "Filter by odometer"
// @Param  	dealers   	      query   string  false  "Filter by dealers"
// @Param   price_from        query   string  false  "Filter by price from"
// @Param   price_to          query   string  false  "Filter by price to"
// @Param   limit             query   string  false  "Limit"
// @Param   last_id           query   string  false  "Last item ID"
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [get]
func (h *UserHandler) GetCars(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	userID := c.Locals("id").(int)
	fmt.Println(userID)
	brands := auth.QueryParamToArray(c.Query("brands"))
	dealers := auth.QueryParamToArray(c.Query("dealers"))
	models := auth.QueryParamToArray(c.Query("models"))
	regions := auth.QueryParamToArray(c.Query("regions"))
	cities := auth.QueryParamToArray(c.Query("cities"))
	generations := auth.QueryParamToArray(c.Query("generations"))
	colors := auth.QueryParamToArray(c.Query("colors"))
	transmissions := auth.QueryParamToArray(c.Query("transmissions"))
	engines := auth.QueryParamToArray(c.Query("engines"))
	drivetrains := auth.QueryParamToArray(c.Query("drivetrains"))
	body_types := auth.QueryParamToArray(c.Query("body_types"))
	fuel_types := auth.QueryParamToArray(c.Query("fuel_types"))
	ownership_types := auth.QueryParamToArray(c.Query("ownership_types"))
	targetUserID := c.Query("user_id")
	year_from := c.Query("year_from")
	limit := c.Query("limit")
	lastID := c.Query("last_id")
	odometer := c.Query("odometer")
	year_to := c.Query("year_to")
	tradeIn := c.Query("trade_in")
	credit := c.Query("credit")
	crash := c.Query("crash")
	owners := c.Query("owners")
	price_from := c.Query("price_from")
	price_to := c.Query("price_to")
	wheelQ := c.Query("wheel")
	newQ := c.Query("new")

	lastIDInt, limitInt := utils.CheckLastIDLimit(lastID, limit, "")
	data := h.UserService.GetCars(c.Context(), userID, targetUserID, brands, models,
		regions, cities, generations, transmissions, engines, drivetrains,
		body_types, fuel_types, ownership_types, colors, dealers,
		year_from, year_to, credit, price_from, price_to,
		tradeIn, owners, crash, odometer, newQ, wheelQ, limitInt, lastIDInt, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetPriceRecommendation godoc
// @Summary      Get price recommendation
// @Description  Returns a price recommendation
// @Tags         car
// @Produce      json
// @Security 	 BearerAuth
// @Param        brand_id          query   string  true  "Brand ID"
// @Param        model_id          query   string  true  "Model ID"
// @Param        year              query   string  true  "Year"
// @Param        modification_id   query   string  false  "Modification ID"
// @Param        city_id           query   string  false  "City ID"
// @Param        odometer          query   string  false  "Odometer"
// @Router       /api/v1/users/cars/price-recommendation [get]
// @Success      200  {object}  model.GetPriceRecommendationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
func (h *UserHandler) GetPriceRecommendation(c *fiber.Ctx) error {

	brandID := c.Query("brand_id")
	modelID := c.Query("model_id")
	modificationID := c.Query("modification_id")
	cityID := c.Query("city_id")
	year := c.Query("year")
	odometer := c.Query("odometer")

	if brandID == "" || modelID == "" || year == "" {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data"),
		})
	}

	data := h.UserService.GetPriceRecommendation(c.Context(), model.GetPriceRecommendationRequest{
		BrandID:        brandID,
		ModelID:        modelID,
		Year:           year,
		ModificationID: modificationID,
		CityID:         cityID,
		Odometer:       odometer,
	})
	return utils.FiberResponse(c, data)
}

// GetCarByID godoc
// @Summary      Get car by ID
// @Description  Returns a car by its ID
// @Tags         car
// @Produce      json
// @Security 	 BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id} [get]
func (h *UserHandler) GetCarByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	nameColumn := c.Locals("lang").(string)
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.GetCarByID(c.Context(), id, userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetEditCarByID godoc
// @Summary      Get Edit car by ID
// @Description  Returns a car by its ID
// @Tags         car
// @Produce      json
// @Security 	 BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.GetEditCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/edit [get]
func (h *UserHandler) GetEditCarByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	nameColumn := c.Locals("lang").(string)
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.GetEditCarByID(c.Context(), id, userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// BuyCar godoc
// @Summary      Buy car
// @Description  Returns a status response message
// @Tags         car
// @Produce      json
// @Security 	 BearerAuth
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/buy [post]
func (h *UserHandler) BuyCar(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.BuyCar(c.Context(), id, userID)
	return utils.FiberResponse(c, data)
}

// CreateCar godoc
// @Summary      Create a car
// @Description  Creates a new car for the authenticated user
// @Tags         car
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        car  body      model.CreateCarRequest  true  "Car data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [post]
func (h *UserHandler) CreateCar(c *fiber.Ctx) error {
	var car model.CreateCarRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&car); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data" + err.Error()),
		})
	}

	if err := h.validator.Validate(car); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.CreateCar(c.Context(), &car, userID)
	return utils.FiberResponse(c, data)
}

// UpdateCar godoc
// @Summary      Update a car
// @Description  Updates an existing car for the authenticated user
// @Tags         car
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        car  body      model.UpdateCarRequest  true  "Car data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/cars [put]
func (h *UserHandler) UpdateCar(c *fiber.Ctx) error {
	var car model.UpdateCarRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&car); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data" + err.Error()),
		})
	}

	if err := h.validator.Validate(car); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.UpdateCar(c.Context(), &car, userID)
	return utils.FiberResponse(c, data)
}

// CreateCarImages godoc
// @Summary      Upload car images
// @Description  Uploads images for a car (max 10 files)
// @Tags         car
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        car_id      path      int     true   "Car CAR_ID"
// @Param        images  formData  file    true   "Car images (max 10)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/images [post]
func (h *UserHandler) CreateCarImages(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid car ID"),
		})
	}

	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}

	images := form.File["images"]

	if len(images) > 10 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 10 files"),
		})
	}

	paths, status, err := files.SaveFiles(images, config.ENV.STATIC_PATH+"cars/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: status,
			Error:  err,
		})
	}

	data := h.UserService.CreateCarImages(c.Context(), id, paths)
	return utils.FiberResponse(c, data)
}

// CreateCarVideos godoc
// @Summary      Upload car videos
// @Description  Uploads videos for a car (max 1 files)
// @Tags         car
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        car_id      path      int     true   "Car CAR_ID"
// @Param        videos  formData  file    true   "Car videos (max 10)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/videos [post]
func (h *UserHandler) CreateCarVideos(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid car ID"),
		})
	}

	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}

	videos := form.File["videos"]

	if len(videos) > 1 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 file(s)"),
		})
	}

	// path, err := pkg.SaveVideos(videos[0], config.ENV.STATIC_PATH+"cars/"+idStr+"/videos") // if have ffmpeg on server
	path, err := files.SaveOriginal(videos[0], config.ENV.STATIC_PATH+"cars/"+idStr+"/videos")

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.UserService.CreateCarVideos(c.Context(), id, path)
	return utils.FiberResponse(c, data)
}

// CreateMessageFile godoc
// @Summary      Upload message file
// @Description  Uploads file for a message (max 1 file)
// @Tags         messages
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file    true   "Message file (max 1)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/messages/files [post]
func (h *UserHandler) CreateMessageFile(c *fiber.Ctx) error {

	senderID := c.Locals("id").(int)
	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}

	videos := form.File["file"]

	if len(videos) > 1 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 file(s)"),
		})
	}

	path, err := files.SaveOriginal(videos[0], config.ENV.STATIC_PATH+"messages/"+fmt.Sprintf("%d", senderID)+"/files")

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.UserService.CreateMessageFile(c.Context(), senderID, path)
	return utils.FiberResponse(c, data)
}

// Cancel cars godoc
// @Summary      Get user's cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         car
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/cancel [post]
func (h *UserHandler) Cancel(c *fiber.Ctx) error {
	// todo: delete images if exist

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.Cancel(c.Context(), &id, "./images/cars/"+idStr)
	return utils.FiberResponse(c, data)
}

// Dont sell godoc
// @Summary      Dont sell cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         car
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/dont-sell [post]
func (h *UserHandler) DontSell(c *fiber.Ctx) error {

	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.DontSell(c.Context(), &id, &userID)
	return utils.FiberResponse(c, data)
}

//	Sell godoc
//
// @Summary       Sell cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         car
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id}/sell [post]
func (h *UserHandler) Sell(c *fiber.Ctx) error {

	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.Sell(c.Context(), &id, &userID)
	return utils.FiberResponse(c, data)
}

// DeleteCarImage godoc
// @Summary      Delete car image
// @Description  Deletes a car image by car ID and image path
// @Tags         car
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      int     true   "Car ID"
// @Param        image body      model.DeleteCarImageRequest true "Image path"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{id}/images [delete]
func (h *UserHandler) DeleteCarImage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	carID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	var req model.DeleteCarImageRequest

	if err := c.BodyParser(&req); err != nil || req.Image == "" {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid image path in request body"),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	// Remove from DB
	resp := h.UserService.DeleteCarImage(c.Context(), carID, req.Image)

	if resp.Error == nil {
		// Remove from disk (ignore error, as file may not exist)
		_ = files.RemoveFile(req.Image)
	}
	return utils.FiberResponse(c, resp)
}

// DeleteCarVideo godoc
// @Summary      Delete car video
// @Description  Deletes a car video by car ID and video path
// @Tags         car
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id    path      int     true   "Car ID"
// @Param        video body      model.DeleteCarVideoRequest true "Video path"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/cars/{id}/videos [delete]
func (h *UserHandler) DeleteCarVideo(c *fiber.Ctx) error {
	idStr := c.Params("id")
	carID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	var video model.DeleteCarVideoRequest

	if err := c.BodyParser(&video); err != nil || video.Video == "" {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid video path in request body"),
		})
	}

	if err := h.validator.Validate(video); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	// Remove from DB
	resp := h.UserService.DeleteCarVideo(c.Context(), carID, video.Video)

	if resp.Error == nil {
		// pkg.RemoveFile(req.Video[:5]) // use it if have car's multiple videos
		files.RemoveFolder(config.ENV.STATIC_PATH + "cars/" + idStr + "/videos")

	}
	return utils.FiberResponse(c, resp)
}

// Cancel cars godoc
// @Summary      Get user's cars
// @Description  Returns the cars associated with the authenticated user's
// @Tags         car
// @Security     BearerAuth
// @Produce      json
// @Param        car_id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/cars/{car_id} [delete]
func (h *UserHandler) DeleteCar(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	data := h.UserService.DeleteCar(c.Context(), &id, "/images/cars/"+idStr)
	return utils.FiberResponse(c, data)
}

// GetBrands godoc
// @Summary      Get car brands
// @Description  Returns a list of car brands, optionally filtered by text
// @Tags         brand
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        text  query     string  false  "Filter brands by text"
// @Success      200   {array}  model.GetBrandsResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands [get]
func (h *UserHandler) GetBrands(c *fiber.Ctx) error {
	text := c.Query("text")
	nameColumn := c.Locals("lang").(string)

	data := h.UserService.GetBrands(c.Context(), text, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetFilterBrands godoc
// @Summary      Get car brands
// @Description  Returns a list of car brands, optionally filtered by text
// @Tags         filter
// @Produce      json
// @Param        text  query     string  false  "Filter brands by text"
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200   {object}  model.GetFilterBrandsResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/filter-brands [get]
func (h *UserHandler) GetFilterBrands(c *fiber.Ctx) error {
	text := c.Query("text")
	nameColumn := c.Locals("lang").(string)

	data := h.UserService.GetFilterBrands(c.Context(), text, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetBrands godoc
// @Summary      Get car cities
// @Description  Returns a list of car cities, optionally filtered by text
// @Tags         filter
// @Produce      json
// @Param        text  query     string  false  "Filter cities by text"
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200   {array}  model.GetCitiesResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/cities [get]
func (h *UserHandler) GetCities(c *fiber.Ctx) error {
	text := c.Query("text")

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetCities(c.Context(), text, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetModelsByBrandID godoc
// @Summary      Get models by brand ID for create cars
// @Description  Returns a list of car models for a given brand ID, optionally filtered by text
// @Tags         brand
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        id    path      int     true   "Brand ID"
// @Param        text  query     string  false  "coroll"
// @Success      200   {array}  model.Model
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/models [get]
func (h *UserHandler) GetModelsByBrandID(c *fiber.Ctx) error {
	brandID := c.Params("id")
	text := c.Query("text")
	nameColumn := c.Locals("lang").(string)
	brandIDInt, err := strconv.ParseInt(brandID, 10, 64)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.GetModelsByBrandID(c.Context(), brandIDInt, text, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetFilterModelsByBrandID godoc
// @Summary      Get filter models by brand ID
// @Description  Returns a list of car models for a given brand ID, optionally filtered by text
// @Tags         brand
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        id    path      int     true   "Brand ID"
// @Param        text  query     string  false  "coroll"
// @Success      200   {object}  model.GetFilterModelsResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/filter-models [get]
func (h *UserHandler) GetFilterModelsByBrandID(c *fiber.Ctx) error {
	brandID := c.Params("id")
	text := c.Query("text")
	nameColumn := c.Locals("lang").(string)
	brandIDInt, err := strconv.ParseInt(brandID, 10, 64)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.GetFilterModelsByBrandID(c.Context(), brandIDInt, text, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetGenerationsByModelID godoc
// @Summary      Get generations by model ID
// @Description  Returns a list of generations for a given model ID
// @Tags         brand
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        id  path  int  true  "brand id ID"
// @Param        model_id  path  int  true  "Model ID"
// @Param   	 year   		query   string    	true  "Selected year"
// @Param   	 wheel   		query   string    	true  "true or false wheel"
// @Param   	 body_type_id   query   string    	true  "the selected body type's ID"
// @Success      200   {array}  model.Generation
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/models/{model_id}/generations [get]
func (h *UserHandler) GetGenerationsByModelID(c *fiber.Ctx) error {
	modelID := c.Params("model_id")
	year := c.Query("year")
	nameColumn := c.Locals("lang").(string)
	bodyTypeID := c.Query("body_type_id")
	wheel := true

	if c.Query("wheel", "true") == "false" {
		wheel = false
	}

	modelIDInt, err := strconv.Atoi(modelID)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.GetGenerationsByModelID(c.Context(), modelIDInt, wheel, year, bodyTypeID, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetGenerationsByModelID godoc
// @Summary      Get generations by model ID
// @Description  Returns a list of generations for a given model ID
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        models			query		string		true  "Model IDs"
// @Success      200   {array}  model.Generation
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/models/generations [get]
func (h *UserHandler) GetGenerationsByModels(c *fiber.Ctx) error {
	models, err := auth.QueryParamToIntArray(c.Query("models"))
	nameColumn := c.Locals("lang").(string)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.UserService.GetGenerationsByModels(c.Context(), models, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetYearsByModelID godoc
// @Summary      Get years by model ID
// @Description  Returns a list of years for a given model ID
// @Tags         brand
// @Produce      json
// @Param        id  path  int  true  "Brand ID"
// @Param        model_id  path  int  true  "Model ID"
// @Param        wheel  query  string  true  "the wheel true or false"
// @Success      200   {array}  int
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/models/{model_id}/years [get]
func (h *UserHandler) GetYearsByModelID(c *fiber.Ctx) error {
	modelID := c.Params("model_id")
	wheel := true
	if c.Query("wheel", "true") == "false" {
		wheel = false
	}

	modelIDInt, err := strconv.ParseInt(modelID, 10, 64)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.GetYearsByModelID(c.Context(), modelIDInt, wheel)
	return utils.FiberResponse(c, data)
}

// GetBodysByModelID godoc
// @Summary      Get bodys by model ID
// @Description  Returns a list of bodys for a given model ID
// @Tags         brand
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        id        	path    int  		true  "Brand ID"
// @Param        model_id  	path    int  		true  "Model ID"
// @Param   	 year   	query   string    	true  "Selected year"
// @Param   	 wheel   	query   string    	true  "true or false wheel"
// @Success      200   {array}  model.BodyType
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/brands/{id}/models/{model_id}/body-types [get]
func (h *UserHandler) GetBodyTypesByModelID(c *fiber.Ctx) error {
	modelID := c.Params("model_id")
	year := c.Query("year")
	wheel := true
	nameColumn := c.Locals("lang").(string)

	if c.Query("wheel", "true") == "false" {
		wheel = false
	}

	modelIDInt, err := strconv.Atoi(modelID)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.GetBodysByModelID(c.Context(), modelIDInt, wheel, year, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetBodyTypes godoc
// @Summary      Get body types
// @Description  Returns a list of car body types
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.BodyType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/body-types [get]
func (h *UserHandler) GetBodyTypes(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetBodyTypes(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetTransmissions godoc
// @Summary      Get transmissions
// @Description  Returns a list of car transmissions
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.Transmission
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/transmissions [get]
func (h *UserHandler) GetTransmissions(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetTransmissions(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetEngines godoc
// @Summary      Get engines
// @Description  Returns a list of car engines
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.Engine
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/engines [get]
func (h *UserHandler) GetEngines(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetEngines(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetDrivetrains godoc
// @Summary      Get drivetrains
// @Description  Returns a list of car drivetrains
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.Drivetrain
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/drivetrains [get]
func (h *UserHandler) GetDrivetrains(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetDrivetrains(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetFuelTypes godoc
// @Summary      Get fuel types
// @Description  Returns a list of car fuel types
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.FuelType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/fuel-types [get]
func (h *UserHandler) GetFuelTypes(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetFuelTypes(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetColors godoc
// @Summary      Get colors
// @Description  Returns a list of car colors
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.Color
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/colors [get]
func (h *UserHandler) GetColors(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetColors(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetCountries godoc
// @Summary      Get countries
// @Description  Returns a list of countries
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.Country
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/countries [get]
func (h *UserHandler) GetCountries(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetCountries(c.Context(), nameColumn)
	return utils.FiberResponse(c, data)
}

// GetHome godoc
// @Summary      Get home
// @Description  Returns a list of car home
// @Tags         filter
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {object}  model.Home
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/home [get]
func (h *UserHandler) GetHome(c *fiber.Ctx) error {

	userID := c.Locals("id").(int)
	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetHome(c.Context(), userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// Liked cars
// @Summary      My liked cars
// @Description  Liked cars
// @Tags         like
// @Security     BearerAuth
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/likes [get]
func (h *UserHandler) Likes(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	nameColumn := c.Locals("lang").(string)
	data := h.UserService.Likes(c.Context(), &userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// Like item godoc
// @Summary      Crate liked item
// @Description  User like a item
// @Tags         like
// @Security     BearerAuth
// @Produce      json
// @Param        item_id   path      int  true  "Item ID"
// @Param        item_type   query      string  true  "Item Type (car, motorcycle, comtran), default: car"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/likes/{item_id} [post]
func (h *UserHandler) ItemLike(c *fiber.Ctx) error {
	itemIDStr := c.Params("item_id")
	userID := c.Locals("id").(int)
	itemType := c.Query("item_type", "car")
	data := h.UserService.ItemLike(c.Context(), userID, itemIDStr, itemType)
	return utils.FiberResponse(c, data)
}

// remove Like car godoc
// @Summary      remove Crate liked item
// @Description  User like a item
// @Tags         like
// @Security     BearerAuth
// @Produce      json
// @Param        item_id   path      int  true  "Item ID"
// @Param        item_type   query      string  true  "Item Type (car, motorcycle, comtran), default: car"
// @Success      200  {object}  model.Success
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/likes/{item_id} [delete]
func (h *UserHandler) RemoveLike(c *fiber.Ctx) error {
	// todo: delete images if exist

	itemIDStr := c.Params("item_id")
	itemType := c.Query("item_type", "car")
	userID := c.Locals("id").(int)
	data := h.UserService.RemoveLike(c.Context(), userID, itemIDStr, itemType)
	return utils.FiberResponse(c, data)
}

// GetProfileCars godoc
// @Summary      Get user's profile cars
// @Description  Returns the cars associated with the authenticated user's profile
// @Tags         profile
// @Security     BearerAuth
// @Produce      json
// @Param        limit   query      string  false  "Limit"
// @Param        last_id   query      string  false  "Last item ID"
// @Param   Accept-Language  header  string  false  "Language"
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/profile/my-cars [get]
func (h *UserHandler) GetMyCars(c *fiber.Ctx) error {

	limit := c.Query("limit")
	lastID := c.Query("last_id")
	userID := c.Locals("id").(int)
	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetMyCars(c.Context(), userID, limit, lastID, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetProfileCars godoc
// @Summary      Get user's profile cars
// @Description  Returns the cars associated with the authenticated user's profile
// @Tags         profile
// @Security     BearerAuth
// @Produce      json
// @Param   Accept-Language  header  string  false  "Language"
// @Param        limit   query      string  false  "Limit"
// @Param        last_id   query      string  false  "Last item ID"
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object} model.ResultMessage
// @Failure      401  {object} auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object} model.ResultMessage
// @Failure      500  {object} model.ResultMessage
// @Router       /api/v1/users/profile/on-sale [get]
func (h *UserHandler) OnSale(c *fiber.Ctx) error {

	limit := c.Query("limit")
	lastID := c.Query("last_id")
	userID := c.Locals("id").(int)
	nameColumn := c.Locals("lang").(string)
	data := h.UserService.OnSale(c.Context(), userID, limit, lastID, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Returns a list of user profile
// @Tags         profile
// @Produce      json
// @Security     BearerAuth
// @Success      200   {object}  model.GetProfileResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/users/profile [get]
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {

	userID := c.Locals("id").(int)
	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetProfile(c.Context(), userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// UpdateProfile godoc
// @Summary      Update user profile
// @Description  Updates the authenticated user's profile information
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        profile  body      model.UpdateProfileRequest  true  "Profile data"
// @Success      200      {object}  model.Success
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure		 403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/users/profile [put]
func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	var profile model.UpdateProfileRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&profile); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(profile); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.UpdateProfile(c.Context(), userID, &profile)
	return utils.FiberResponse(c, data)
}

// UploadProfileAvatar godoc
// @Summary      Upload profile avatar
// @Description  Uploads the authenticated user's profile avatar image
// @Tags         profile
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        avatar  formData  file  true  "Avatar image (max 1)"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/profile/avatar [post]
func (h *UserHandler) UploadProfileAvatar(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	form, err := c.MultipartForm()
	if err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}
	data := h.UserService.UploadProfileAvatar(c.Context(), form, userID)
	return utils.FiberResponse(c, data)
}

// DeleteProfileAvatar godoc
// @Summary      Delete profile avatar
// @Description  Deletes the authenticated user's profile avatar
// @Tags         profile
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  model.Success
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/profile/avatar [delete]
func (h *UserHandler) DeleteProfileAvatar(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	data := h.UserService.DeleteProfileAvatar(c.Context(), userID)
	return utils.FiberResponse(c, data)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Returns a single user by ID
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.ThirdPartyGetProfileRes
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")

	nameColumn := c.Locals("lang").(string)
	data := h.UserService.GetUserByID(c.Context(), userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetDealers godoc
// @Summary      Get dealers
// @Description  Returns a list of dealers
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.ThirdPartyUserResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/dealers [get]
func (h *UserHandler) GetDealers(c *fiber.Ctx) error {
	data := h.UserService.GetDealers(c.Context())
	return utils.FiberResponse(c, data)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Returns a single user by ID
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Param        from_id   query      int  false  "From ID"
// @Param        to_id   query      int  false  "To ID"
// @Param        role_id   query      int  false  "Role ID"
// @Param        search   query      string  false  "Search"
// @Success      200  {array}  model.ThirdPartyUserResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/third-party [get]
func (h *UserHandler) GetThirdPartyUsers(c *fiber.Ctx) error {
	fromID := c.Query("from_id")
	toID := c.Query("to_id")
	roleID := c.Query("role_id")
	search := c.Query("search")

	data := h.UserService.GetThirdPartyUsers(c.Context(), roleID, fromID, toID, search)
	return utils.FiberResponse(c, data)
}

// CreateReport godoc
// @Summary      Create a report
// @Description  Creates a new report for the authenticated user
// @Tags         report
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        report  body      model.CreateReportRequest  true  "Report data"
// @Success      200     {object}  model.SuccessWithId
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure      403     {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/reports [post]
func (h *UserHandler) CreateReport(c *fiber.Ctx) error {
	var report model.CreateReportRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&report); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(report); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.CreateReport(c.Context(), &report, userID)
	return utils.FiberResponse(c, data)
}

// GetReports godoc
// @Summary      Get user reports
// @Description  Returns a list of reports created by the authenticated user
// @Tags         report
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.GetReportsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/reports [get]
func (h *UserHandler) GetReports(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	data := h.UserService.GetReports(c.Context(), userID)
	return utils.FiberResponse(c, data)
}

// CreateItemReports godoc
// @Summary      Create item report
// @Description  Creates a new report for a specific item (car, moto, or comtran)
// @Tags         report
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        report  body      model.CreateItemReportRequest  true  "Item Report data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/item-reports [post]
func (h *UserHandler) CreateItemReports(c *fiber.Ctx) error {
	var report model.CreateItemReportRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&report); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(report); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.UserService.CreateItemReport(c.Context(), &report, userID)
	return utils.FiberResponse(c, data)
}

// Motorcycle handlers

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
// @Router /api/v1/users/motorcycles [post]
func (h *UserHandler) CreateMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	var motorcycle model.CreateMotorcycleRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&motorcycle); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	if err := h.validator.Validate(motorcycle); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	return utils.FiberResponse(c, h.MotorcycleService.CreateMotorcycle(ctx, motorcycle, userID))
}

// GetMotorcycles godoc
// @Summary      Get motorcycles
// @Description  Returns a list of motorcycles
// @Tags         motorcycles
// @Produce      json
// @Security BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param   brands            query   string  false  "Filter by brand IDs"
// @Param   models            query   string  false  "Filter by model IDs"
// @Param   regions           query   string  false  "Filter by region IDs"
// @Param   cities            query   string  false  "Filter by city IDs"
// @Param   generations       query   string  false  "Filter by generation IDs"
// @Param   colors       	  query   string  false  "Filter by color IDs"
// @Param   crash       	  query   string  false  "Filter by crash status, true or empty"
// @Param   transmissions     query   string  false  "Filter by transmission IDs"
// @Param   engines           query   string  false  "Filter by engine IDs"
// @Param   drivetrains       query   string  false  "Filter by drivetrain IDs"
// @Param   body_types        query   string  false  "Filter by body type IDs"
// @Param   fuel_types        query   string  false  "Filter by fuel type IDs"
// @Param   trade_in          query   string  false  "Filter by trade_in id, from 1 to 5"
// @Param   owners            query   string  false  "Filter by owners id, from 1 to 4"
// @Param   ownership_types   query   string  false  "Filter by ownership type IDs"
// @Param   year_from         query   string  false  "Filter by year from"
// @Param   year_to           query   string  false  "Filter by year to"
// @Param   credit            query   string  false  "Filter by credit"
// @Param  	new   		      query   string  false  "true or false new"
// @Param  	wheel   		  query   string  false  "true or false wheel"
// @Param  	odometer   	      query   string  false  "Filter by odometer"
// @Param  	dealers   	      query   string  false  "Filter by dealers"
// @Param   price_from        query   string  false  "Filter by price from"
// @Param   price_to          query   string  false  "Filter by price to"
// @Param   limit             query   string  false  "Limit"
// @Param   last_id           query   string  false  "Last item ID"
// @Success      200  {array}  model.GetMotorcycleResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles [get]
func (h *UserHandler) GetMotorcycles(c *fiber.Ctx) error {

	nameColumn := c.Locals("lang").(string)
	userID := c.Locals("id").(int)
	brands := auth.QueryParamToArray(c.Query("brands"))
	dealers := auth.QueryParamToArray(c.Query("dealers"))
	models := auth.QueryParamToArray(c.Query("models"))
	regions := auth.QueryParamToArray(c.Query("regions"))
	cities := auth.QueryParamToArray(c.Query("cities"))
	generations := auth.QueryParamToArray(c.Query("generations"))
	colors := auth.QueryParamToArray(c.Query("colors"))
	transmissions := auth.QueryParamToArray(c.Query("transmissions"))
	engines := auth.QueryParamToArray(c.Query("engines"))
	drivetrains := auth.QueryParamToArray(c.Query("drivetrains"))
	body_types := auth.QueryParamToArray(c.Query("body_types"))
	fuel_types := auth.QueryParamToArray(c.Query("fuel_types"))
	ownership_types := auth.QueryParamToArray(c.Query("ownership_types"))
	year_from := c.Query("year_from")
	limit := c.Query("limit")
	lastID := c.Query("last_id")
	odometer := c.Query("odometer")
	year_to := c.Query("year_to")
	tradeIn := c.Query("trade_in")
	credit := c.Query("credit")
	crash := c.Query("crash")
	owners := c.Query("owners")
	price_from := c.Query("price_from")
	price_to := c.Query("price_to")
	wheelQ := c.Query("wheel")
	newQ := c.Query("new")

	lastIDInt, limitInt := utils.CheckLastIDLimit(lastID, limit, "")
	data := h.MotorcycleService.GetMotorcycles(c.Context(), userID, brands, models,
		regions, cities, generations, transmissions, engines, drivetrains,
		body_types, fuel_types, ownership_types, colors, dealers,
		year_from, year_to, credit, price_from, price_to,
		tradeIn, owners, crash, odometer, newQ, wheelQ, limitInt, lastIDInt, nameColumn)
	return utils.FiberResponse(c, data)
}

// BuyMotorcycle godoc
// @Summary      Buy motorcycle
// @Description  Buys a motorcycle (transfers ownership)
// @Tags         motorcycles
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/buy [post]
func (h *UserHandler) BuyMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.MotorcycleService.BuyMotorcycle(ctx, id, userID))
}

// CancelMotorcycle godoc
// @Summary      Cancel motorcycle
// @Description  Cancels/deletes a motorcycle listing
// @Tags         motorcycles
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/cancel [post]
func (h *UserHandler) CancelMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	dir := config.ENV.STATIC_PATH + "motorcycles/" + idStr
	return utils.FiberResponse(c, h.MotorcycleService.CancelMotorcycle(ctx, &id, dir))
}

// DontSellMotorcycle godoc
// @Summary      Set motorcycle as not for sale
// @Description  Updates motorcycle status to not for sale
// @Tags         motorcycles
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/dont-sell [post]
func (h *UserHandler) DontSellMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.MotorcycleService.DontSellMotorcycle(ctx, id, userID))
}

// SellMotorcycle godoc
// @Summary      Set motorcycle for sale
// @Description  Updates motorcycle status to for sale
// @Tags         motorcycles
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/sell [post]
func (h *UserHandler) SellMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.MotorcycleService.SellMotorcycle(ctx, id, userID))
}

// CreateMotorcycleImages godoc
// @Summary      Upload motorcycle images
// @Description  Uploads images for a motorcycle (max 10 files)
// @Tags         motorcycles
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id       path      int    true  "Motorcycle ID"
// @Param        images   formData  file   true  "Motorcycle images (max 10)"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/images [post]
func (h *UserHandler) CreateMotorcycleImages(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}
	form, _ := c.MultipartForm()
	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}
	images := form.File["images"]
	if len(images) == 0 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("images are required"),
		})
	}
	if len(images) > 10 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 10 files"),
		})
	}
	paths, status, err := files.SaveFiles(images, config.ENV.STATIC_PATH+"motorcycles/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)
	if err != nil {
		return utils.FiberResponse(c, model.Response{Status: status, Error: err})
	}
	data := h.MotorcycleService.CreateMotorcycleImages(ctx, id, paths)
	return utils.FiberResponse(c, data)
}

// CreateMotorcycleVideos godoc
// @Summary      Upload motorcycle videos
// @Description  Uploads videos for a motorcycle (max 1 file)
// @Tags         motorcycles
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id       path      int    true  "Motorcycle ID"
// @Param        videos   formData  file   true  "Motorcycle videos (max 1)"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/videos [post]
func (h *UserHandler) CreateMotorcycleVideos(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}
	form, _ := c.MultipartForm()
	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}
	videos := form.File["videos"]
	if len(videos) == 0 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("videos are required"),
		})
	}
	if len(videos) > 1 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 file(s)"),
		})
	}
	path, err := files.SaveOriginal(videos[0], config.ENV.STATIC_PATH+"motorcycles/"+idStr+"/videos")
	if err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}
	data := h.MotorcycleService.CreateMotorcycleVideos(ctx, id, path)
	return utils.FiberResponse(c, data)
}

// UpdateMotorcycle godoc
// @Summary      Update a motorcycle
// @Description  Updates an existing motorcycle for the authenticated user
// @Tags         motorcycles
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        motorcycle  body      model.UpdateMotorcycleRequest  true  "Motorcycle data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles [put]
func (h *UserHandler) UpdateMotorcycle(c *fiber.Ctx) error {
	var motorcycle model.UpdateMotorcycleRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&motorcycle); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data" + err.Error()),
		})
	}

	if err := h.validator.Validate(motorcycle); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.MotorcycleService.UpdateMotorcycle(c.Context(), &motorcycle, userID)
	return utils.FiberResponse(c, data)
}

// DeleteMotorcycleImage godoc
// @Summary      Delete motorcycle image
// @Description  Deletes an image from a motorcycle
// @Tags         motorcycles
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Param        image_id   query      int  true  "Image ID"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/images [delete]
func (h *UserHandler) DeleteMotorcycleImage(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	motorcycleID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	imageIDStr := c.Query("image_id")
	imageID, err := strconv.Atoi(imageIDStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("image id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.MotorcycleService.DeleteMotorcycleImage(ctx, motorcycleID, imageID))
}

// DeleteMotorcycleVideo godoc
// @Summary      Delete motorcycle video
// @Description  Deletes a video from a motorcycle
// @Tags         motorcycles
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Param        video_id   query      int  true  "Video ID"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id}/videos [delete]
func (h *UserHandler) DeleteMotorcycleVideo(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	motorcycleID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	videoIDStr := c.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("video id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.MotorcycleService.DeleteMotorcycleVideo(ctx, motorcycleID, videoID))
}

// DeleteMotorcycle godoc
// @Summary      Delete motorcycle
// @Description  Deletes a motorcycle and its associated files
// @Tags         motorcycles
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Motorcycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/motorcycles/{id} [delete]
func (h *UserHandler) DeleteMotorcycle(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("motorcycle id must be integer"),
		})
	}

	dir := config.ENV.STATIC_PATH + "motorcycles/" + idStr
	return utils.FiberResponse(c, h.MotorcycleService.DeleteMotorcycle(ctx, id, dir))
}

// Comtrans handlers

// CreateComtrans godoc
// @Summary Create commercial transport
// @Description Create commercial transport
// @Tags comtrans
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param comtrans body model.CreateComtransRequest true "Commercial Transport"
// @Success 200 {object} model.SuccessWithId
// @Failure 500 {object} model.ResultMessage
// @Failure 400 {object} model.ResultMessage
// @Router /api/v1/users/comtrans [post]
func (h *UserHandler) CreateComtrans(c *fiber.Ctx) error {
	ctx := c.Context()
	var comtrans model.CreateComtransRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&comtrans); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	if err := h.validator.Validate(comtrans); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	return utils.FiberResponse(c, h.ComtransService.CreateComtrans(ctx, comtrans, userID))
}

// GetComtrans godoc
// @Summary      Get comtrans
// @Description  Returns a list of comtrans
// @Tags         comtrans
// @Produce      json
// @Security BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param   brands            query   string  false  "Filter by brand IDs"
// @Param   models            query   string  false  "Filter by model IDs"
// @Param   regions           query   string  false  "Filter by region IDs"
// @Param   cities            query   string  false  "Filter by city IDs"
// @Param   generations       query   string  false  "Filter by generation IDs"
// @Param   colors       	  query   string  false  "Filter by color IDs"
// @Param   crash       	  query   string  false  "Filter by crash status, true or empty"
// @Param   transmissions     query   string  false  "Filter by transmission IDs"
// @Param   engines           query   string  false  "Filter by engine IDs"
// @Param   drivetrains       query   string  false  "Filter by drivetrain IDs"
// @Param   body_types        query   string  false  "Filter by body type IDs"
// @Param   fuel_types        query   string  false  "Filter by fuel type IDs"
// @Param   trade_in          query   string  false  "Filter by trade_in id, from 1 to 5"
// @Param   owners            query   string  false  "Filter by owners id, from 1 to 4"
// @Param   ownership_types   query   string  false  "Filter by ownership type IDs"
// @Param   year_from         query   string  false  "Filter by year from"
// @Param   year_to           query   string  false  "Filter by year to"
// @Param   credit            query   string  false  "Filter by credit"
// @Param  	new   		      query   string  false  "true or false new"
// @Param  	wheel   		  query   string  false  "true or false wheel"
// @Param  	odometer   	      query   string  false  "Filter by odometer"
// @Param  	dealers   	      query   string  false  "Filter by dealers"
// @Param   price_from        query   string  false  "Filter by price from"
// @Param   price_to          query   string  false  "Filter by price to"
// @Param   limit             query   string  false  "Limit"
// @Param   last_id           query   string  false  "Last item ID"
// @Success      200  {array}  model.GetCarsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans [get]
func (h *UserHandler) GetComtrans(c *fiber.Ctx) error {
	nameColumn := c.Locals("lang").(string)
	userID := c.Locals("id").(int)
	brands := auth.QueryParamToArray(c.Query("brands"))
	dealers := auth.QueryParamToArray(c.Query("dealers"))
	models := auth.QueryParamToArray(c.Query("models"))
	regions := auth.QueryParamToArray(c.Query("regions"))
	cities := auth.QueryParamToArray(c.Query("cities"))
	generations := auth.QueryParamToArray(c.Query("generations"))
	colors := auth.QueryParamToArray(c.Query("colors"))
	transmissions := auth.QueryParamToArray(c.Query("transmissions"))
	engines := auth.QueryParamToArray(c.Query("engines"))
	drivetrains := auth.QueryParamToArray(c.Query("drivetrains"))
	body_types := auth.QueryParamToArray(c.Query("body_types"))
	fuel_types := auth.QueryParamToArray(c.Query("fuel_types"))
	ownership_types := auth.QueryParamToArray(c.Query("ownership_types"))
	targetUserID := c.Query("user_id")
	year_from := c.Query("year_from")
	limit := c.Query("limit")
	lastID := c.Query("last_id")
	odometer := c.Query("odometer")
	year_to := c.Query("year_to")
	tradeIn := c.Query("trade_in")
	credit := c.Query("credit")
	crash := c.Query("crash")
	owners := c.Query("owners")
	price_from := c.Query("price_from")
	price_to := c.Query("price_to")
	wheelQ := c.Query("wheel")
	newQ := c.Query("new")
	var wheel *bool
	var new *bool

	if newQ != "" {
		if newQ == "false" {
			tmp := false
			new = &tmp
		} else {
			tmp := true
			new = &tmp
		}

	}

	if wheelQ != "" {
		if wheelQ == "false" {
			tmp := false
			wheel = &tmp
		} else {
			tmp := true
			wheel = &tmp
		}

	}

	lastIDInt, limitInt := utils.CheckLastIDLimit(lastID, limit, "")
	data := h.ComtransService.GetComtrans(c.Context(), userID, targetUserID, brands, models,
		regions, cities, generations, transmissions, engines, drivetrains,
		body_types, fuel_types, ownership_types, colors, dealers,
		year_from, year_to, credit, price_from, price_to,
		tradeIn, owners, crash, odometer, new, wheel, limitInt, lastIDInt, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetEditComtransByID godoc
// @Summary      Get Edit comtrans by ID
// @Description  Returns a comtrans by its ID for editing
// @Tags         comtrans
// @Produce      json
// @Security 	 BearerAuth
// @Param   Accept-Language  header  string  false  "Language"
// @Param        id   path      int  true  "Comtrans ID"
// @Success      200  {object}  model.GetEditComtransResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object} auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/edit [get]
func (h *UserHandler) GetEditComtransByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	nameColumn := c.Locals("lang").(string)
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("comtrans id must be integer"),
		})
	}

	data := h.ComtransService.GetEditComtransByID(c.Context(), id, userID, nameColumn)
	return utils.FiberResponse(c, data)
}

// GetComtransCategories godoc
// @Summary Get commercial transport categories
// @Description Get commercial transport categories
// @Tags comtrans
// @Accept json
// @Produce json
// @Param   Accept-Language  header  string  false  "Language"
// @Success 200 {array} model.GetComtransCategoriesResponse
// @Failure 500 {object} model.ResultMessage
// @Router /api/v1/users/comtrans/categories [get]
func (h *UserHandler) GetComtransCategories(c *fiber.Ctx) error {
	ctx := c.Context()
	lang := c.Locals("lang").(string)
	return utils.FiberResponse(c, h.ComtransService.GetComtransCategories(ctx, lang))
}

// GetComtransBrands godoc
// @Summary Get commercial transport brands
// @Description Get commercial transport brands
// @Tags comtrans
// @Accept json
// @Produce json
// @Param   Accept-Language  header  string  false  "Language"
// @Success 200 {array} model.GetComtransBrandsResponse
// @Failure 500 {object} model.ResultMessage
// @Router /api/v1/users/comtrans/brands [get]
func (h *UserHandler) GetComtransBrands(c *fiber.Ctx) error {
	ctx := c.Context()
	lang := c.Locals("lang").(string)
	return utils.FiberResponse(c, h.ComtransService.GetComtransBrands(ctx, lang))
}

// GetComtransEngines godoc
// @Summary Get commercial transport engines
// @Description Get commercial transport engines
// @Tags comtrans
// @Accept json
// @Produce json
// @Param   Accept-Language  header  string  false  "Language"
// @Success 200 {array} model.GetComtransModelsResponse
// @Failure 500 {object} model.ResultMessage
// @Router /api/v1/users/comtrans/engines [get]
func (h *UserHandler) GetComtransEngines(c *fiber.Ctx) error {
	ctx := c.Context()
	lang := c.Locals("lang").(string)
	return utils.FiberResponse(c, h.ComtransService.GetComtransEngines(ctx, lang))
}

// GetComtransModelsByBrandID godoc
// @Summary Get commercial transport models by brand ID
// @Description Get commercial transport models by brand ID
// @Tags comtrans
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Commercial Transport Brand ID"
// @Param   Accept-Language  header  string  false  "Language"
// @Success 200 {array} model.GetComtransModelsResponse
// @Failure 500 {object} model.ResultMessage
// @Router /api/v1/users/comtrans/brands/{id}/models [get]
func (h *UserHandler) GetComtransModelsByBrandID(c *fiber.Ctx) error {
	ctx := c.Context()
	lang := c.Locals("lang").(string)
	brandID := c.Params("id")
	return utils.FiberResponse(c, h.ComtransService.GetComtransModelsByBrandID(ctx, brandID, lang))
}

// BuyComtrans godoc
// @Summary      Buy commercial transport
// @Description  Buys a commercial transport (transfers ownership)
// @Tags         comtrans
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/buy [post]
func (h *UserHandler) BuyComtrans(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.ComtransService.BuyComtrans(ctx, id, userID))
}

// CancelComtrans godoc
// @Summary      Cancel commercial transport
// @Description  Cancels/deletes a commercial transport listing
// @Tags         comtrans
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/cancel [post]
func (h *UserHandler) CancelComtrans(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	dir := config.ENV.STATIC_PATH + "comtrans/" + idStr
	return utils.FiberResponse(c, h.ComtransService.CancelComtrans(ctx, &id, dir))
}

// DontSellComtrans godoc
// @Summary      Set commercial transport as not for sale
// @Description  Updates commercial transport status to not for sale
// @Tags         comtrans
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/dont-sell [post]
func (h *UserHandler) DontSellComtrans(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.ComtransService.DontSellComtrans(ctx, id, userID))
}

// SellComtrans godoc
// @Summary      Set commercial transport for sale
// @Description  Updates commercial transport status to for sale
// @Tags         comtrans
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/sell [post]
func (h *UserHandler) SellComtrans(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID := c.Locals("id").(int)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.ComtransService.SellComtrans(ctx, id, userID))
}

// CreateComtransImages godoc
// @Summary      Upload commercial transport images
// @Description  Uploads images for a commercial transport (max 10 files)
// @Tags         comtrans
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id       path      int    true  "Commercial Transport ID"
// @Param        images   formData  file   true  "Commercial transport images (max 10)"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/images [post]
func (h *UserHandler) CreateComtransImages(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}
	form, _ := c.MultipartForm()
	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}
	images := form.File["images"]
	if len(images) == 0 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("images are required"),
		})
	}
	if len(images) > 10 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 10 files"),
		})
	}
	paths, status, err := files.SaveFiles(images, config.ENV.STATIC_PATH+"comtrans/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)
	if err != nil {
		return utils.FiberResponse(c, model.Response{Status: status, Error: err})
	}
	data := h.ComtransService.CreateComtransImages(ctx, id, paths)
	return utils.FiberResponse(c, data)
}

// CreateComtransVideos godoc
// @Summary      Upload commercial transport videos
// @Description  Uploads videos for a commercial transport (max 1 file)
// @Tags         comtrans
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id       path      int    true  "Commercial Transport ID"
// @Param        videos   formData  file   true  "Commercial transport videos (max 1)"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/videos [post]
func (h *UserHandler) CreateComtransVideos(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}
	form, _ := c.MultipartForm()
	if form == nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the files"),
		})
	}
	videos := form.File["videos"]
	if len(videos) == 0 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("videos are required"),
		})
	}
	if len(videos) > 1 {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 file(s)"),
		})
	}
	path, err := files.SaveOriginal(videos[0], config.ENV.STATIC_PATH+"comtrans/"+idStr+"/videos")
	if err != nil {
		return utils.FiberResponse(c, model.Response{Status: 400, Error: err})
	}
	data := h.ComtransService.CreateComtransVideos(ctx, id, path)
	return utils.FiberResponse(c, data)
}

// UpdateComtrans godoc
// @Summary      Update a commercial transport
// @Description  Updates an existing commercial transport for the authenticated user
// @Tags         comtrans
// @Accept       json
// @Produce      json
// @Security 	 BearerAuth
// @Param        comtrans  body      model.UpdateComtransRequest  true  "Commercial Transport data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure		 403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans [put]
func (h *UserHandler) UpdateComtrans(c *fiber.Ctx) error {
	var comtrans model.UpdateComtransRequest
	userID := c.Locals("id").(int)

	if err := c.BodyParser(&comtrans); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data" + err.Error()),
		})
	}

	if err := h.validator.Validate(comtrans); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.ComtransService.UpdateComtrans(c.Context(), &comtrans, userID)
	return utils.FiberResponse(c, data)
}

// DeleteComtransImage godoc
// @Summary      Delete commercial transport image
// @Description  Deletes an image from a commercial transport
// @Tags         comtrans
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Param        image_id   query      int  true  "Image ID"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/images [delete]
func (h *UserHandler) DeleteComtransImage(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	comtransID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	imageIDStr := c.Query("image_id")
	imageID, err := strconv.Atoi(imageIDStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("image id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.ComtransService.DeleteComtransImage(ctx, comtransID, imageID))
}

// DeleteComtransVideo godoc
// @Summary      Delete commercial transport video
// @Description  Deletes a video from a commercial transport
// @Tags         comtrans
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Param        video_id   query      int  true  "Video ID"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id}/videos [delete]
func (h *UserHandler) DeleteComtransVideo(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	comtransID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	videoIDStr := c.Query("video_id")
	videoID, err := strconv.Atoi(videoIDStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("video id must be integer"),
		})
	}

	return utils.FiberResponse(c, h.ComtransService.DeleteComtransVideo(ctx, comtransID, videoID))
}

// DeleteComtrans godoc
// @Summary      Delete commercial transport
// @Description  Deletes a commercial transport and its associated files
// @Tags         comtrans
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Commercial Transport ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/users/comtrans/{id} [delete]
func (h *UserHandler) DeleteComtrans(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("commercial transport id must be integer"),
		})
	}

	dir := config.ENV.STATIC_PATH + "comtrans/" + idStr
	return utils.FiberResponse(c, h.ComtransService.DeleteComtrans(ctx, id, dir))
}
