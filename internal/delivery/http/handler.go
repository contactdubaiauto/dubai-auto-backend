package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserService.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: model.Success{
			Message: "User created successfully",
			Id:      int(user.ID),
		},
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.UserService.GetUserByID(c.Request.Context(), idInt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Status: 200,
		Data:   user,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Status: 200,
		Data:   users,
	})
}

func (h *UserHandler) GetBrands(c *gin.Context) {
	text := c.Query("text")
	brands, err := h.UserService.GetBrands(c.Request.Context(), text)

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

	models, err := h.UserService.GetModelsByBrandID(c.Request.Context(), brandIDInt, text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Data: models,
	})
}

func (h *UserHandler) GetBodyTypes(c *gin.Context) {
	bodyTypes, err := h.UserService.GetBodyTypes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Data: bodyTypes,
	})
}

func (h *UserHandler) GetTransmissions(c *gin.Context) {
	transmissions, err := h.UserService.GetTransmissions(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: transmissions,
	})
}

func (h *UserHandler) GetEngines(c *gin.Context) {
	engines, err := h.UserService.GetEngines(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: engines,
	})
}

func (h *UserHandler) GetDrives(c *gin.Context) {
	drives, err := h.UserService.GetDrives(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: drives,
	})
}

func (h *UserHandler) GetFuelTypes(c *gin.Context) {
	fuelTypes, err := h.UserService.GetFuelTypes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.GinResponse(c, &model.Response{
		Data: fuelTypes,
	})
}

func (h *UserHandler) GetCars(c *gin.Context) {
	cars, err := h.UserService.GetCars(c.Request.Context())

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

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.UserService.CreateCar(c.Request.Context(), &car)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.GinResponse(c, &model.Response{
		Data: model.Success{
			Message: "Car created successfully",
			Id:      id,
		},
	})
}

// func (h *UserHandler) CreateCarImages(c *gin.Context) {
// 	carID := c.Param("id")
// 	carIDInt, err := strconv.ParseInt(carID, 10, 64)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	images, err := h.UserService.CreateCarImages(c.Request.Context(), carIDInt)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	utils.GinResponse(c, &model.Response{
// 		Data: images,
// 	})
// }
