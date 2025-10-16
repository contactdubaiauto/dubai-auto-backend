package http

import (
	"errors"
	"strconv"
	"strings"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/files"

	"github.com/gofiber/fiber/v2"
)

type ThirdPartyHandler struct {
	service *service.ThirdPartyService
}

func NewThirdPartyHandler(service *service.ThirdPartyService) *ThirdPartyHandler {
	return &ThirdPartyHandler{service}
}

// Profile godoc
// @Summary      Profile
// @Description  Returns a profile
// @Tags         third-party
// @Produce      json
// @Security     BearerAuth
// @Param        profile  body      model.ThirdPartyProfileReq  true  "Profile"
// @Success      200      {object}  model.Success
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/third-party/profile [post]
func (h *ThirdPartyHandler) Profile(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	profile := &model.ThirdPartyProfileReq{}

	if err := c.BodyParser(profile); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.Profile(ctx, id, *profile)
	return utils.FiberResponse(c, &data)
}

// Profile godoc
// @Summary      First login
// @Description  Returns a first login
// @Tags         third-party
// @Produce      json
// @Security     BearerAuth
// @Param        profile  body      model.ThirdPartyFirstLoginReq  true  "First login"
// @Success      200      {object}  model.Success
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/third-party/first-login [post]
func (h *ThirdPartyHandler) FirstLogin(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	profile := &model.ThirdPartyFirstLoginReq{}

	if err := c.BodyParser(profile); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.FirstLogin(ctx, id, *profile)
	return utils.FiberResponse(c, &data)
}

// GetProfile godoc
// @Summary      Get profile
// @Description  Returns a profile
// @Tags         third-party
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  model.ThirdPartyGetProfileRes
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/third-party/profile [get]
func (h *ThirdPartyHandler) GetProfile(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	data := h.service.GetProfile(ctx, id)
	return utils.FiberResponse(c, &data)
}

// GetMyCars godoc
// @Summary      Get my cars
// @Description  Returns my cars
// @Tags         third-party
// @Produce      json
// @Security     BearerAuth
// @Success      200      {array}  model.GetCarsResponse
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/third-party/profile/my-cars [get]
func (h *ThirdPartyHandler) GetMyCars(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	data := h.service.GetMyCars(ctx, id)
	return utils.FiberResponse(c, &data)
}

// OnSale godoc
// @Summary      On sale
// @Description  Returns on sale cars
// @Tags         third-party
// @Produce      json
// @Security     BearerAuth
// @Success      200      {array}  model.GetCarsResponse
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/third-party/profile/on-sale [get]
func (h *ThirdPartyHandler) OnSale(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	data := h.service.OnSale(ctx, id)
	return utils.FiberResponse(c, &data)
}

// GetRegistrationData godoc
// @Summary      Get registration data
// @Description  Returns registration data
// @Tags         third-party
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  model.ThirdPartyGetRegistrationDataRes
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      404      {object}  model.ResultMessage
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/third-party/registration-data [get]
func (h *ThirdPartyHandler) GetRegistrationData(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetRegistrationData(ctx)
	return utils.FiberResponse(c, &data)
}

// AvatarImages godoc
// @Summary      Upload avatar images
// @Description  Uploads avatar images
// @Tags         third-party
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        avatar_image  formData  file    true   "Avatar image required"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure      403     {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/third-party/profile/images [post]
func (h *ThirdPartyHandler) AvatarImages(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	form, _ := c.MultipartForm()
	data := h.service.CreateAvatarImages(ctx, form, id)
	return utils.FiberResponse(c, &data)
}

// BannerImage godoc
// @Summary      Upload banner images
// @Description  Uploads banner images
// @Tags         third-party
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        banner_image  formData  file    true   "Banner image required"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure      403     {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/third-party/profile/banner [post]
func (h *ThirdPartyHandler) BannerImage(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	form, _ := c.MultipartForm()
	data := h.service.CreateBannerImage(ctx, form, id)
	return utils.FiberResponse(c, &data)
}

// CreateDealerCar godoc
// @Summary      Create a dealer car
// @Description  Creates a new car for the authenticated dealer
// @Tags         dealer
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        car  body      model.ThirdPartyCreateCarRequest  true  "Car data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/dealer/car [post]
func (h *ThirdPartyHandler) CreateDealerCar(c *fiber.Ctx) error {
	var car model.ThirdPartyCreateCarRequest
	dealerID := c.Locals("id").(int)
	ctx := c.Context()

	if err := c.BodyParser(&car); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.CreateDealerCar(ctx, &car, dealerID)
	return utils.FiberResponse(c, &data)
}

// UpdateDealerCar godoc
// @Summary      Update a dealer car
// @Description  Updates an existing car for the authenticated dealer
// @Tags         dealer
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Car ID"
// @Param        car  body      model.UpdateCarRequest  true  "Car data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/dealer/car/{id} [post]
func (h *ThirdPartyHandler) UpdateDealerCar(c *fiber.Ctx) error {
	var car model.UpdateCarRequest
	dealerID := c.Locals("id").(int)
	ctx := c.Context()

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	if err := c.BodyParser(&car); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	car.ID = id
	data := h.service.UpdateDealerCar(ctx, &car, dealerID)
	return utils.FiberResponse(c, &data)
}

// CreateCarImages godoc
// @Summary      Upload car images
// @Description  Uploads images for a car (max 10 files)
// @Tags         dealer
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id      path      int     true   "Car ID"
// @Param        images  formData  file    true   "Car images (max 10)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/third-party/dealer/car/{car_id}/images [post]
func (h *ThirdPartyHandler) CreateDealerCarImages(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid dealer car ID"),
		})
	}

	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the dealer car images"),
		})
	}

	images := form.File["images"]

	if len(images) > 10 {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 10 dealer car images"),
		})
	}

	paths, status, err := files.SaveFiles(images, config.ENV.STATIC_PATH+"cars/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: status,
			Error:  err,
		})
	}

	data := h.service.CreateDealerCarImages(ctx, id, paths)
	return utils.FiberResponse(c, data)
}

// CreateCarVideos godoc
// @Summary      Upload car videos
// @Description  Uploads videos for a car (max 1 files)
// @Tags         dealer
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id      path      int     true   "Car ID"
// @Param        videos  formData  file    true   "Car videos (max 10)"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure	 	 403  	 {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/third-party/dealer/car/{car_id}/videos [post]
func (h *ThirdPartyHandler) CreateDealerCarVideos(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("invalid dealer car ID"),
		})
	}

	form, _ := c.MultipartForm()

	if form == nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("didn't upload the dealer car videos"),
		})
	}

	videos := form.File["videos"]

	if len(videos) > 1 {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("must load maximum 1 dealer car video(s)"),
		})
	}

	// path, err := pkg.SaveVideos(videos[0], config.ENV.STATIC_PATH+"cars/"+idStr+"/videos") // if have ffmpeg on server
	path, err := files.SaveOriginal(videos[0], config.ENV.STATIC_PATH+"dealer-cars/"+idStr+"/videos")

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.CreateDealerCarVideos(ctx, id, path)
	return utils.FiberResponse(c, data)
}

// StatusDealer godoc
// @Summary      Change dealer car status
// @Description  Changes the status of a dealer car (sell/dont-sell)
// @Tags         dealer
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/dealer/car/{id}/sell [post]
// @Router       /api/v1/third-party/dealer/car/{id}/dont-sell [post]
func (h *ThirdPartyHandler) StatusDealer(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	dealerID := c.Locals("id").(int)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}

	// Determine if this is sell or dont-sell based on the path
	path := string(c.Context().Path())

	if strings.Contains(path, "dont-sell") {
		data := h.service.DealerDontSell(ctx, &id, &dealerID)
		return utils.FiberResponse(c, &data)
	}

	data := h.service.DealerSell(ctx, &id, &dealerID)
	return utils.FiberResponse(c, &data)
}

// DeleteDealerCar godoc
// @Summary      Delete dealer car
// @Description  Deletes a dealer car
// @Tags         dealer
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Car ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/dealer/car/{id} [delete]
func (h *ThirdPartyHandler) DeleteDealerCar(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("car id must be integer"),
		})
	}
	data := h.service.DeleteDealerCar(ctx, id)
	return utils.FiberResponse(c, &data)
}

// GetLogistDestinations godoc
// @Summary      Get logist destinations
// @Description  Returns a list of logist destinations (routes)
// @Tags         logist
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.LogistDestinationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/logist/destinations [get]
func (h *ThirdPartyHandler) GetLogistDestinations(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetLogistDestinations(ctx)
	return utils.FiberResponse(c, &data)
}

// CreateLogistDestination godoc
// @Summary      Create logist destination
// @Description  Creates a new logist destination (route)
// @Tags         logist
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        destination  body      model.CreateLogistDestinationRequest  true  "Destination data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/logist/destinations [post]
func (h *ThirdPartyHandler) CreateLogistDestination(c *fiber.Ctx) error {
	ctx := c.Context()
	var req model.CreateLogistDestinationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.CreateLogistDestination(ctx, req)
	return utils.FiberResponse(c, &data)
}

// DeleteLogistDestination godoc
// @Summary      Delete logist destination
// @Description  Deletes a logist destination (route) by ID
// @Tags         logist
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Destination ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/third-party/logist/destinations/{id} [delete]
func (h *ThirdPartyHandler) DeleteLogistDestination(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("destination id must be integer"),
		})
	}

	data := h.service.DeleteLogistDestination(ctx, id)
	return utils.FiberResponse(c, &data)
}
