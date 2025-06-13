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

func (h *UserHandler) GetBrands(c *gin.Context) {
	text := c.Query("text")
	ctx := c.Request.Context()
	brands, err := h.UserService.GetBrands(&ctx, text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Status: 200,
		Data:   brands,
	})
}

func (h *UserHandler) GetModelsByBrandID(c *gin.Context) {
	brandID := c.Param("id")
	text := c.Query("text")
	brandIDInt, err := strconv.ParseInt(brandID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	models, err := h.UserService.GetModelsByBrandID(&ctx, brandIDInt, text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Data: models,
	})
}

func (h *UserHandler) GetBodyTypes(c *gin.Context) {
	ctx := c.Request.Context()
	bodyTypes, err := h.UserService.GetBodyTypes(&ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Data: bodyTypes,
	})
}

func (h *UserHandler) GetTransmissions(c *gin.Context) {
	ctx := c.Request.Context()
	transmissions, err := h.UserService.GetTransmissions(&ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: transmissions,
	})
}

func (h *UserHandler) GetEngines(c *gin.Context) {
	ctx := c.Request.Context()
	engines, err := h.UserService.GetEngines(&ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: engines,
	})
}

func (h *UserHandler) GetDrivetrains(c *gin.Context) {
	ctx := c.Request.Context()
	drivetrains, err := h.UserService.GetDrivetrains(&ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Data: drivetrains,
	})
}

func (h *UserHandler) GetFuelTypes(c *gin.Context) {
	ctx := c.Request.Context()
	fuelTypes, err := h.UserService.GetFuelTypes(&ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: fuelTypes,
	})
}

func (h *UserHandler) GetCars(c *gin.Context) {
	ctx := c.Request.Context()
	cars, err := h.UserService.GetCars(&ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{Data: cars})
}

func (h *UserHandler) CreateCar(c *gin.Context) {
	var car model.CreateCarRequest
	userID := c.MustGet("id").(int)
	car.UserID = int64(userID)
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.UserService.CreateCar(&ctx, &car)

	utils.GinResponse(c, data)
}

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
	image := form.File["image"]

	if len(image) != 1 {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("must load just 1 file"),
		})
		return
	}

	paths, status, err := pkg.SaveFiles(image, config.ENV.STATIC_PATH+"cars/"+strconv.Itoa(id), config.ENV.DEFAULT_IMAGE_WIDTHS)

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
