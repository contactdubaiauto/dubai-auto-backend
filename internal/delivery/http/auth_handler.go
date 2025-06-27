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

// UserLogin godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserLoginRequest  true  "User login credentials"
// @Success      200   {object}  model.LoginResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure      403   {object}  pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-login [post]
func (h *AuthHandler) UserLogin(c *gin.Context) {
	user := &model.UserLoginRequest{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.service.UserLogin(c.Request.Context(), user)

	utils.GinResponse(c, &data)
}

// UserRegister godoc
// @Summary      User registration
// @Description  Registers a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserRegisterRequest  true  "User registration data"
// @Success      200   {object}  model.LoginResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure      403   {object}  pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-register [post]
func (h *AuthHandler) UserRegister(c *gin.Context) {
	user := &model.UserRegisterRequest{}

	if err := c.ShouldBindJSON(user); err != nil {
		utils.GinResponse(c, &model.Response{Error: err, Status: http.StatusBadRequest})
		return
	}

	data := h.service.UserRegister(c.Request.Context(), user)

	utils.GinResponse(c, &data)
}
