package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"

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
