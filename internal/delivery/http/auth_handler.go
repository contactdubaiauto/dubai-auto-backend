package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

// DeleteAccount godoc
// @Summary      Delete user account
// @Description  Deletes the authenticated user's account and related data
// @Tags         auth
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  pkg.ErrorResponse
// @Failure      403  {object}  pkg.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/auth/account/{id} [delete]
func (h *AuthHandler) DeleteAccount(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		utils.GinResponse(c, &model.Response{
			Status: 400,
			Error:  errors.New("user id must be integer"),
		})
		return
	}
	// Optionally, you can check if the user is deleting their own account:
	authUserID := c.MustGet("id").(int)
	if userID != authUserID {
		utils.GinResponse(c, &model.Response{
			Status: 403,
			Error:  errors.New("forbidden: cannot delete another user's account"),
		})
		return
	}

	data := h.service.DeleteAccount(&ctx, userID)
	utils.GinResponse(c, data)
}

// UserEmail confirmation godoc
// @Summary      User email confirmation
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserEmailConfirmationRequest  true  "User email confirmation credentials"
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

// UserPhone confirmation godoc
// @Summary      User phone confirmation
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserPhoneConfirmationRequest  true  "User phone confirmation credentials"
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

// UserLoginEmail godoc
// @Summary      User login email
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserLoginEmail  true  "User login email credentials"
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

// UserLoginPhone godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserLoginPhone  true  "User login phone credentials"
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
