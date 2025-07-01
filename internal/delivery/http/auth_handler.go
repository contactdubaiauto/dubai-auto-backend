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
// @Param        user  body      model.UserEmailConfirmationRequest  true  "User login credentials"
// @Success      200   {object}  model.LoginResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure      403   {object}  pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-email-confirmation [post]
func (h *AuthHandler) UserEmailConfirmation(c *gin.Context) {
	user := &model.UserEmailConfirmationRequest{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.service.UserEmailConfirmation(c.Request.Context(), user)

	utils.GinResponse(c, &data)
}

// UserLogin godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserPhoneConfirmationRequest  true  "User login credentials"
// @Success      200   {object}  model.LoginResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure      403   {object}  pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-phone-confirmation [post]
func (h *AuthHandler) UserPhoneConfirmation(c *gin.Context) {
	user := &model.UserPhoneConfirmationRequest{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.service.UserPhoneConfirmation(c.Request.Context(), user)

	utils.GinResponse(c, &data)
}

// UserLogin godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserLoginEmail  true  "User login credentials"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure      403   {object}  pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-login-email [post]
func (h *AuthHandler) UserLoginEmail(c *gin.Context) {
	user := &model.UserLoginEmail{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.service.UserLoginEmail(c.Request.Context(), user)
	utils.GinResponse(c, &data)
}

// UserLogin godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserLoginPhone  true  "User login credentials"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  pkg.ErrorResponse
// @Failure      403   {object}  pkg.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-login-phone [post]
func (h *AuthHandler) UserLoginPhone(c *gin.Context) {
	user := &model.UserLoginPhone{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := h.service.UserLoginPhone(c.Request.Context(), user)
	utils.GinResponse(c, &data)
}
