package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"empty/internal/model"
	"empty/internal/service"
	"empty/internal/utils"
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
		Status: 200,
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
