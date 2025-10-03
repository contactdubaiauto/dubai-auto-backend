package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
)

type AuthHandler struct {
	service   *service.AuthService
	validator *auth.Validator
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service, auth.NewValidator()}
}

// UserLoginGoogle godoc
// @Summary      User login google
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserLoginGoogle  true  "User login google credentials"
// @Success      200   {object}  model.LoginFiberResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-login-google [post]
func (h *AuthHandler) UserLoginGoogle(c *fiber.Ctx) error {
	user := &model.UserLoginGoogle{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	if err := h.validator.Validate(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.UserLoginGoogle(c.Context(), user.TokenID)
	return utils.FiberResponse(c, &data)
}

// Send Application godoc
// @Summary      Send application
// @Description  Sends an application to the database
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        application  body      model.UserApplication  true  "Application"
// @Success      200   {object}  model.LoginFiberResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/send-application [post]
func (h *AuthHandler) Application(c *fiber.Ctx) error {
	application := &model.UserApplication{}

	if err := c.BodyParser(application); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	if err := h.validator.Validate(application); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.Application(c.Context(), *application)
	return utils.FiberResponse(c, &data)
}

// ApplicationDocuments godoc
// @Summary      Application documents
// @Description  Sends application documents to the database
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Application ID"
// @Param        licence  	 formData  file    true   "A PDF document file"
// @Param        memorandum  formData  file    true   "A PDF document file"
// @Param        copy_of_id  formData  file    true   "A PDF document file"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/send-application-document [post]
func (h *AuthHandler) ApplicationDocuments(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Locals("id").(int)
	licence, err := c.FormFile("licence")

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	memorandum, err := c.FormFile("memorandum")

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	copyOfID, err := c.FormFile("copy_of_id")

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.ApplicationDocuments(ctx, id, licence, memorandum, copyOfID)
	return utils.FiberResponse(c, data)
}

// UserEmail confirmation godoc
// @Summary      User email confirmation
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserEmailConfirmationRequest  true  "User email confirmation credentials"
// @Success      200   {object}  model.LoginFiberResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-email-confirmation [post]
func (h *AuthHandler) UserEmailConfirmation(c *fiber.Ctx) error {
	user := &model.UserEmailConfirmationRequest{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.UserEmailConfirmation(c.Context(), user)

	return utils.FiberResponse(c, &data)
}

// UserPhone confirmation godoc
// @Summary      User phone confirmation
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserPhoneConfirmationRequest  true  "User phone confirmation credentials"
// @Success      200   {object}  model.LoginFiberResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-phone-confirmation [post]
func (h *AuthHandler) UserPhoneConfirmation(c *fiber.Ctx) error {
	user := &model.UserPhoneConfirmationRequest{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.UserPhoneConfirmation(c.Context(), user)

	return utils.FiberResponse(c, &data)
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
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-login-email [post]
func (h *AuthHandler) UserLoginEmail(c *fiber.Ctx) error {
	user := &model.UserLoginEmail{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.UserLoginEmail(c.Context(), user)
	return utils.FiberResponse(c, &data)
}

// ThirdPartyLogin godoc
// @Summary      Third party login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      model.ThirdPartyLoginReq  true  "Third party login credentials"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/third-party-login [post]
func (h *AuthHandler) ThirdPartyLogin(c *fiber.Ctx) error {
	user := &model.ThirdPartyLoginReq{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.ThirdPartyLogin(c.Context(), user)
	return utils.FiberResponse(c, &data)
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
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/auth/user-login-phone [post]
func (h *AuthHandler) UserLoginPhone(c *fiber.Ctx) error {
	user := &model.UserLoginPhone{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.UserLoginPhone(c.Context(), user)
	return utils.FiberResponse(c, &data)
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
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/auth/account/{id} [delete]
func (h *AuthHandler) DeleteAccount(c *fiber.Ctx) error {
	ctx := c.Context()
	idStr := c.Params("id")
	userID, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, &model.Response{
			Status: 400,
			Error:  err,
		})
	}

	// Optionally, you can check if the user is deleting their own account:
	authUserID := c.Locals("id").(int)

	if userID != authUserID {
		return utils.FiberResponse(c, &model.Response{
			Status: 403,
			Error:  err,
		})
	}

	data := h.service.DeleteAccount(ctx, userID)
	return utils.FiberResponse(c, data)
}
