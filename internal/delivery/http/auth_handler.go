package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) UserLogin(c *gin.Context) {
	user := &model.UserLogin{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.service.UserLogin(c.Request.Context(), user)

	utils.GinResponse(c, &data)
}

func (h *AuthHandler) UserRegister(c *gin.Context) {
	user := &model.UserRegister{}

	if err := c.ShouldBindJSON(user); err != nil {
		utils.GinResponse(c, &model.Response{Error: err, Status: http.StatusBadRequest})
		return
	}

	data := h.service.UserRegister(c.Request.Context(), user)

	utils.GinResponse(c, &data)
}
