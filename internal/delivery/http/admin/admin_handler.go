package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
)

type AdminHandler struct {
	service   *service.AdminService
	validator *auth.Validator
}

func NewAdminHandler(service *service.AdminService, validator *auth.Validator) *AdminHandler {
	return &AdminHandler{service, validator}
}

// Users handlers

// CreateAdmin godoc
// @Summary      Create an admin
// @Description  Creates an admin
// @Tags         admin-users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user  body      model.CreateUserRequest  true  "User"
// @Success      200   {object}  model.SuccessWithId
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/users [post]
func (h *AdminHandler) CreateUser(c *fiber.Ctx) error {
	user := &model.CreateUserRequest{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.CreateUser(ctx, user)
	return utils.FiberResponse(c, data)
}

// GetAdmins godoc
// @Summary      Get all users
// @Description  Returns a list of all users
// @Tags         admin-users
// @Produce      json
// @Security     BearerAuth
// @Param        role_id   query      int  true  "Role ID"
// @Success      200   {array}   model.UserResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/users [get]
func (h *AdminHandler) GetUsers(c *fiber.Ctx) error {
	ctx := c.Context()
	qRoleID := c.Query("role_id")
	data := h.service.GetUsers(ctx, qRoleID)
	return utils.FiberResponse(c, data)
}

// GetAdmin godoc
// @Summary      Get a user by ID
// @Description  Returns a single user by ID
// @Tags         admin-users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/users/{id} [get]
func (h *AdminHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("user id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetUser(ctx, id)
	return utils.FiberResponse(c, data)
}

// UpdateAdmin godoc
// @Summary      Update a user
// @Description  Updates an admin by ID
// @Tags         admin-users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Param        user  body      model.UpdateUserRequest  true  "User"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/users/{id} [put]
func (h *AdminHandler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("user id must be integer"),
		})
	}

	user := &model.UpdateUserRequest{}

	if err := c.BodyParser(user); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.UpdateUser(ctx, id, user)
	return utils.FiberResponse(c, data)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID
// @Tags         admin-users
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/users/{id} [delete]
func (h *AdminHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("user id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteUser(ctx, id)
	return utils.FiberResponse(c, data)
}

// SendUserNotifications godoc
// @Summary      Send global notifications via FCM
// @Description  Sends FCM notifications to all users with the given role_id. Request body: role_id (1 user, 2 dealer, 3 logistic, 4 broker, 5 car service), title, description.
// @Tags         admin-users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body      model.SendNotificationRequest  true  "role_id, title, description"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/users/notifications [post]
func (h *AdminHandler) SendUserNotifications(c *fiber.Ctx) error {
	var req model.SendNotificationRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}
	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}
	ctx := c.Context()
	data := h.service.SendUserNotifications(ctx, &req)
	return utils.FiberResponse(c, data)
}

// Profile handlers

// GetProfile godoc
// @Summary      Get profile
// @Description  Returns a profile
// @Tags         admin-profile
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  model.AdminProfileResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/profile [get]
func (h *AdminHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Locals("id").(int)
	ctx := c.Context()
	data := h.service.GetProfile(ctx, id)
	return utils.FiberResponse(c, data)
}

// Applications handlers

// GetApplications godoc
// @Summary      Get all applications
// @Description  Returns a list of all applications
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        role   query      int  true  "Role ID (2: Dealer, 3: Logist, 4: Broker, 5: Car Service)"
// @Param        status   query      int  true  "Status ID (1: Pending, 2: Approved, 3: Rejected)"
// @Param        limit   query      string  false  "Limit"
// @Param        last_id   query      string  false  "Last item ID"
// @Param        search   query      string  false  "Search"
// @Success      200  {array}  model.AdminApplicationResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications [get]
func (h *AdminHandler) GetApplications(c *fiber.Ctx) error {
	qRole := c.Query("role")
	qStatus := c.Query("status")
	limit := c.Query("limit")
	lastID := c.Query("last_id")
	search := c.Query("search")
	ctx := c.Context()
	data := h.service.GetApplications(ctx, qRole, qStatus, limit, lastID, search)
	return utils.FiberResponse(c, data)
}

// Send Application godoc
// @Summary      Send application
// @Description  Sends an application to the database
// @Tags         admin-applications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        application  body      model.UserApplication  true  "Application"
// @Success      200   {object}  model.SuccessWithId
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      404   {object}  model.ResultMessage
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/applications [post]
func (h *AdminHandler) CreateApplication(c *fiber.Ctx) error {
	application := &model.UserApplication{}

	if err := c.BodyParser(application); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	if err := h.validator.Validate(application); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.CreateApplication(c.Context(), *application)
	return utils.FiberResponse(c, data)
}

// GetApplication godoc
// @Summary      Get an application
// @Description  Returns an application by ID
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Application ID"
// @Param        status   query      int  true  "Status ID (1: Pending, 2: Approved, 3: Rejected)"
// @Success      200  {object}  model.AdminApplicationByIDResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications/{id} [get]
func (h *AdminHandler) GetApplication(c *fiber.Ctx) error {
	idStr := c.Params("id")
	qStatus := c.Query("status")
	ctx := c.Context()
	data := h.service.GetApplication(ctx, idStr, qStatus)
	return utils.FiberResponse(c, data)
}

// ApplicationDocuments godoc
// @Summary      Application documents
// @Description  Sends application documents to the database
// @Tags         admin-applications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
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
// @Router       /api/v1/admin/applications/{id}/documents [post]
func (h *AdminHandler) CreateApplicationDocuments(c *fiber.Ctx) error {
	ctx := c.Context()
	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("application userID must be integer"),
		})
	}
	licence, err := c.FormFile("licence")

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	memorandum, err := c.FormFile("memorandum")

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	copyOfID, err := c.FormFile("copy_of_id")

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.CreateApplicationDocuments(ctx, userID, licence, memorandum, copyOfID)
	return utils.FiberResponse(c, data)
}

// AcceptApplication godoc
// @Summary      Accept an application
// @Description  Accepts an application by ID
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Application ID"
// @Param        req  body      model.AcceptApplicationRequest  true  "Application request"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications/{id}/accept [post]
func (h *AdminHandler) AcceptApplication(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	var req model.AcceptApplicationRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("application id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.AcceptApplication(ctx, id, req)
	return utils.FiberResponse(c, data)
}

// RejectApplication godoc
// @Summary      Reject an application
// @Description  Rejects an application by ID
// @Tags         admin-applications
// @Produce      json
// @Security     BearerAuth
// @Param        status   query      int  true  "Status ID (1: Pending, 2: Approved, 3: Rejected)"
// @Param        message  query      string  true  "reasoning Message"
// @Param        id   path      int  true  "Application ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/applications/{id}/reject [post]
func (h *AdminHandler) RejectApplication(c *fiber.Ctx) error {
	idStr := c.Params("id")
	qStatus := c.Query("status")
	qMessage := c.Query("message")
	ctx := c.Context()
	data := h.service.RejectApplication(ctx, idStr, qStatus, qMessage)
	return utils.FiberResponse(c, data)
}

// Cities handlers

// GetCities godoc
// @Summary      Get all cities
// @Description  Returns a list of all cities
// @Tags         admin-cities
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminCityResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cities [get]
func (h *AdminHandler) GetCities(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetCities(ctx)
	return utils.FiberResponse(c, data)
}

// CreateCity godoc
// @Summary      Create a new city
// @Description  Creates a new city
// @Tags         admin-cities
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city  body      model.CreateNameRequest  true  "City data"
// @Success      200   {object}  model.SuccessWithId
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/cities [post]
func (h *AdminHandler) CreateCity(c *fiber.Ctx) error {
	var req model.CreateNameRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateCity(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateCity godoc
// @Summary      Update a city
// @Description  Updates an existing city
// @Tags         admin-cities
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path      int                      true  "City ID"
// @Param        city  body      model.CreateNameRequest  true  "City data"
// @Success      200   {object}  model.Success
// @Failure      400   {object}  model.ResultMessage
// @Failure      401   {object}  auth.ErrorResponse
// @Failure      403   {object}  auth.ErrorResponse
// @Failure      500   {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{id} [put]
func (h *AdminHandler) UpdateCity(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}

	var req model.CreateNameRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.UpdateCity(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteCity godoc
// @Summary      Delete a city
// @Description  Deletes a city by ID
// @Tags         admin-cities
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "City ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{id} [delete]
func (h *AdminHandler) DeleteCity(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteCity(ctx, id)
	return utils.FiberResponse(c, data)
}

// Company Types handlers

// GetCompanyTypes godoc
// @Summary      Get all company types
// @Description  Returns a list of all company types
// @Tags         admin-company-types
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.CompanyType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/company-types [get]
func (h *AdminHandler) GetCompanyTypes(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetCompanyTypes(ctx)
	return utils.FiberResponse(c, data)
}

// GetCompanyType godoc
// @Summary      Get company type by ID
// @Description  Returns a company type by ID
// @Tags         admin-company-types
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Company type ID"
// @Success      200  {object}  model.CompanyType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/company-types/{id} [get]
func (h *AdminHandler) GetCompanyType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("company type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetCompanyType(ctx, id)
	return utils.FiberResponse(c, data)
}

// CreateCompanyType godoc
// @Summary      Create a company type
// @Description  Creates a new company type
// @Tags         admin-company-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        companyType  body      model.CreateCompanyTypeRequest  true  "Company type data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/company-types [post]
func (h *AdminHandler) CreateCompanyType(c *fiber.Ctx) error {
	var req model.CreateCompanyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateCompanyType(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateCompanyType godoc
// @Summary      Update a company type
// @Description  Updates an existing company type
// @Tags         admin-company-types
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id           path      int                      true  "Company type ID"
// @Param        companyType  body      model.CreateCompanyTypeRequest  true  "Company type data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/company-types/{id} [put]
func (h *AdminHandler) UpdateCompanyType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("company type id must be integer"),
		})
	}

	var req model.CreateCompanyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.UpdateCompanyType(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteCompanyType godoc
// @Summary      Delete a company type
// @Description  Deletes a company type by ID
// @Tags         admin-company-types
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Company type ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/company-types/{id} [delete]
func (h *AdminHandler) DeleteCompanyType(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("company type id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteCompanyType(ctx, id)
	return utils.FiberResponse(c, data)
}

// Activity Fields handlers

// GetActivityFields godoc
// @Summary      Get all activity fields
// @Description  Returns a list of all activity fields
// @Tags         admin-activity-fields
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.CompanyType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/activity-fields [get]
func (h *AdminHandler) GetActivityFields(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetActivityFields(ctx)
	return utils.FiberResponse(c, data)
}

// GetActivityField godoc
// @Summary      Get an activity field by ID
// @Description  Returns an activity field by ID
// @Tags         admin-activity-fields
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Activity field ID"
// @Success      200  {object}  model.CompanyType
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/activity-fields/{id} [get]
func (h *AdminHandler) GetActivityField(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("activity field id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetActivityField(ctx, id)
	return utils.FiberResponse(c, data)
}

// CreateActivityField godoc
// @Summary      Create an activity field
// @Description  Creates a new activity field
// @Tags         admin-activity-fields
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        activityField  body      model.CreateCompanyTypeRequest  true  "Activity field data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/activity-fields [post]
func (h *AdminHandler) CreateActivityField(c *fiber.Ctx) error {
	var req model.CreateCompanyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateActivityField(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateActivityField godoc
// @Summary      Update an activity field
// @Description  Updates an existing activity field
// @Tags         admin-activity-fields
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      int                      true  "Activity field ID"
// @Param        activityField  body      model.CreateCompanyTypeRequest  true  "Activity field data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/activity-fields/{id} [put]
func (h *AdminHandler) UpdateActivityField(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("activity field id must be integer"),
		})
	}

	var req model.CreateCompanyTypeRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.UpdateActivityField(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteActivityField godoc
// @Summary      Delete an activity field
// @Description  Deletes an activity field by ID
// @Tags         admin-activity-fields
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Activity field ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/activity-fields/{id} [delete]
func (h *AdminHandler) DeleteActivityField(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("activity field id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteActivityField(ctx, id)
	return utils.FiberResponse(c, data)
}

// Regions handlers

// GetRegions godoc
// @Summary      Get all regions
// @Description  Returns a list of all regions
// @Tags         admin-regions
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Success      200  {array}  model.AdminCityResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions [get]
func (h *AdminHandler) GetRegions(c *fiber.Ctx) error {
	cityIdStr := c.Params("city_id")
	cityId, err := strconv.Atoi(cityIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}
	ctx := c.Context()
	data := h.service.GetRegions(ctx, cityId)
	return utils.FiberResponse(c, data)
}

// CreateRegion godoc
// @Summary      Create a new region
// @Description  Creates a new region
// @Tags         admin-regions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Param        region  body      model.CreateNameRequest  true  "Region data"
// @Success      200    {object}  model.SuccessWithId
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions [post]
func (h *AdminHandler) CreateRegion(c *fiber.Ctx) error {
	cityIdStr := c.Params("city_id")
	cityId, err := strconv.Atoi(cityIdStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("city id must be integer"),
		})
	}

	var req model.CreateNameRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	ctx := c.Context()
	data := h.service.CreateRegion(ctx, cityId, &req)
	return utils.FiberResponse(c, data)
}

// UpdateRegion godoc
// @Summary      Update a region
// @Description  Updates an existing region
// @Tags         admin-regions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Param        id        path      int  true  "Region ID"
// @Param        region    body      model.CreateNameRequest  true  "Region data"
// @Success      200       {object}  model.Success
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions/{id} [put]
func (h *AdminHandler) UpdateRegion(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("region id must be integer"),
		})
	}

	var req model.CreateNameRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	ctx := c.Context()
	data := h.service.UpdateRegion(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteRegion godoc
// @Summary      Delete a region
// @Description  Deletes an existing region
// @Tags         admin-regions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        city_id   path      int  true  "City ID"
// @Param        id        path      int  true  "Region ID"
// @Success      200       {object}  model.Success
// @Failure      400       {object}  model.ResultMessage
// @Failure      401       {object}  auth.ErrorResponse
// @Failure      403       {object}  auth.ErrorResponse
// @Failure      500       {object}  model.ResultMessage
// @Router       /api/v1/admin/cities/{city_id}/regions/{id} [delete]
func (h *AdminHandler) DeleteRegion(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("region id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteRegion(ctx, id)
	return utils.FiberResponse(c, data)
}

// Transmission handlers

// GetTransmissions godoc
// @Summary      Get all transmissions
// @Description  Returns a list of all transmissions
// @Tags         admin-transmissions
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminTransmissionResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/transmissions [get]
func (h *AdminHandler) GetTransmissions(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetTransmissions(ctx)
	return utils.FiberResponse(c, data)
}

// Color handlers

// GetColors godoc
// @Summary      Get all colors
// @Description  Returns a list of all colors
// @Tags         admin-colors
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminColorResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors [get]
func (h *AdminHandler) GetColors(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetColors(ctx)
	return utils.FiberResponse(c, data)
}

// CreateColor godoc
// @Summary      Create a color
// @Description  Creates a new color
// @Tags         admin-colors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        color  body      model.CreateColorRequest  true  "Color data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors [post]
func (h *AdminHandler) CreateColor(c *fiber.Ctx) error {
	var req model.CreateColorRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.CreateColor(ctx, &req)
	return utils.FiberResponse(c, data)
}

// CreateColorImage godoc
// @Summary      Create a color image
// @Description  Creates a new color image
// @Tags         admin-colors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                   true  "Color ID"
// @Param        image  formData  file    true   "Color image"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors/{id}/images [post]
func (h *AdminHandler) CreateColorImage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	ctx := c.Context()

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("color id must be integer"),
		})
	}

	form, _ := c.MultipartForm()
	data := h.service.CreateColorImage(ctx, form, id)
	return utils.FiberResponse(c, data)
}

// UpdateColor godoc
// @Summary      Update a color
// @Description  Updates a color by ID
// @Tags         admin-colors
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id     path      int                   true  "Color ID"
// @Param        color  body      model.UpdateColorRequest  true  "Color data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors/{id} [put]
func (h *AdminHandler) UpdateColor(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("color id must be integer"),
		})
	}

	var req model.UpdateColorRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.UpdateColor(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteColor godoc
// @Summary      Delete a color
// @Description  Deletes a color by ID
// @Tags         admin-colors
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Color ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/colors/{id} [delete]
func (h *AdminHandler) DeleteColor(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("color id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteColor(ctx, id)
	return utils.FiberResponse(c, data)
}

// Countries CRUD operations

// GetCountries godoc
// @Summary      Get all countries
// @Description  Returns a list of all countries
// @Tags         admin-countries
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminCountryResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/countries [get]
func (h *AdminHandler) GetCountries(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetCountries(ctx)
	return utils.FiberResponse(c, data)
}

// CreateCountry godoc
// @Summary      Create a new country
// @Description  Creates a new country
// @Tags         admin-countries
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        country  body      model.CreateNameRequest  true  "Country data"
// @Success      200      {object}  model.SuccessWithId
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/admin/countries [post]
func (h *AdminHandler) CreateCountry(c *fiber.Ctx) error {
	var req model.CreateNameRequest
	ctx := c.Context()

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	data := h.service.CreateCountry(ctx, &req)
	return utils.FiberResponse(c, data)
}

// CreateCountryImage godoc
// @Summary      Upload country flag image
// @Description  Uploads a flag image for a country
// @Tags         admin-countries
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id     path      int   true   "Country ID"
// @Param        image  formData  file  true   "Country flag image"
// @Success      200    {object}  model.SuccessWithId
// @Failure      400    {object}  model.ResultMessage
// @Failure      401    {object}  auth.ErrorResponse
// @Failure      403    {object}  auth.ErrorResponse
// @Failure      500    {object}  model.ResultMessage
// @Router       /api/v1/admin/countries/{id}/images [post]
func (h *AdminHandler) CreateCountryImage(c *fiber.Ctx) error {
	ctx := c.Context()
	form, _ := c.MultipartForm()
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("country id must be integer"),
		})
	}

	data := h.service.CreateCountryImage(ctx, form, id)
	return utils.FiberResponse(c, data)
}

// UpdateCountry godoc
// @Summary      Update a country
// @Description  Updates an existing country
// @Tags         admin-countries
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int                      true  "Country ID"
// @Param        country  body      model.CreateNameRequest  true  "Country data"
// @Success      200      {object}  model.Success
// @Failure      400      {object}  model.ResultMessage
// @Failure      401      {object}  auth.ErrorResponse
// @Failure      403      {object}  auth.ErrorResponse
// @Failure      500      {object}  model.ResultMessage
// @Router       /api/v1/admin/countries/{id} [put]
func (h *AdminHandler) UpdateCountry(c *fiber.Ctx) error {
	var req model.CreateNameRequest
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("country id must be integer"),
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	if err := h.validator.Validate(req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request data: " + err.Error()),
		})
	}

	ctx := c.Context()
	data := h.service.UpdateCountry(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteCountry godoc
// @Summary      Delete a country
// @Description  Deletes a country by ID
// @Tags         admin-countries
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Country ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/countries/{id} [delete]
func (h *AdminHandler) DeleteCountry(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("country id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteCountry(ctx, id)
	return utils.FiberResponse(c, data)
}

// Report handlers

// GetReports godoc
// @Summary      Get all reports
// @Description  Returns a list of all reports
// @Tags         admin-reports
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   model.GetReportsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/reports [get]
func (h *AdminHandler) GetReports(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetReports(ctx)
	return utils.FiberResponse(c, data)
}

// GetReportByID godoc
// @Summary      Get report by ID
// @Description  Returns a single report by ID
// @Tags         admin-reports
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Report ID"
// @Success      200  {object}  model.GetReportsResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/reports/{id} [get]
func (h *AdminHandler) GetReportByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("report id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.GetReportByID(ctx, id)
	return utils.FiberResponse(c, data)
}

// UpdateReport godoc
// @Summary      Update a report
// @Description  Updates a report by ID (typically to change status)
// @Tags         admin-reports
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path      int                   true  "Report ID"
// @Param        report  body      model.UpdateReportRequest  true  "Report data"
// @Success      200     {object}  model.Success
// @Failure      400     {object}  model.ResultMessage
// @Failure      401     {object}  auth.ErrorResponse
// @Failure      403     {object}  auth.ErrorResponse
// @Failure      404     {object}  model.ResultMessage
// @Failure      500     {object}  model.ResultMessage
// @Router       /api/v1/admin/reports/{id} [put]
func (h *AdminHandler) UpdateReport(c *fiber.Ctx) error {
	idStr := c.Params("id")

	var req model.UpdateReportRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	data := h.service.UpdateReport(c.Context(), idStr, &req)
	return utils.FiberResponse(c, data)
}

// DeleteReport godoc
// @Summary      Delete a report
// @Description  Deletes a report by ID
// @Tags         admin-reports
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Report ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      404  {object}  model.ResultMessage
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/reports/{id} [delete]
func (h *AdminHandler) DeleteReport(c *fiber.Ctx) error {
	idStr := c.Params("id")
	data := h.service.DeleteReport(c.Context(), idStr)
	return utils.FiberResponse(c, data)
}

// Number of cycles handlers

// GetNumberOfCycles godoc
// @Summary      Get all number of cycles
// @Description  Returns a list of all number of cycles
// @Tags         admin-number-of-cycles
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  model.AdminNumberOfCycleResponse
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/number-of-cycles [get]
func (h *AdminHandler) GetNumberOfCycles(c *fiber.Ctx) error {
	ctx := c.Context()
	data := h.service.GetNumberOfCycles(ctx)
	return utils.FiberResponse(c, data)
}

// CreateNumberOfCycle godoc
// @Summary      Create a number of cycle
// @Description  Creates a new number of cycle
// @Tags         admin-number-of-cycles
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        number_of_cycle  body      model.CreateNumberOfCycleRequest  true  "Number of cycle data"
// @Success      200  {object}  model.SuccessWithId
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/number-of-cycles [post]
func (h *AdminHandler) CreateNumberOfCycle(c *fiber.Ctx) error {
	var req model.CreateNumberOfCycleRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.CreateNumberOfCycle(ctx, &req)
	return utils.FiberResponse(c, data)
}

// UpdateNumberOfCycle godoc
// @Summary      Update a number of cycle
// @Description  Updates a number of cycle by ID
// @Tags         admin-number-of-cycles
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id                path      int                             true  "Number of cycle ID"
// @Param        number_of_cycle   body      model.CreateNumberOfCycleRequest  true  "Number of cycle data"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/number-of-cycles/{id} [put]
func (h *AdminHandler) UpdateNumberOfCycle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("number of cycle id must be integer"),
		})
	}

	var req model.CreateNumberOfCycleRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("invalid request body"),
		})
	}

	if err := h.validator.Validate(&req); err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  err,
		})
	}

	ctx := c.Context()
	data := h.service.UpdateNumberOfCycle(ctx, id, &req)
	return utils.FiberResponse(c, data)
}

// DeleteNumberOfCycle godoc
// @Summary      Delete a number of cycle
// @Description  Deletes a number of cycle by ID
// @Tags         admin-number-of-cycles
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Number of cycle ID"
// @Success      200  {object}  model.Success
// @Failure      400  {object}  model.ResultMessage
// @Failure      401  {object}  auth.ErrorResponse
// @Failure      403  {object}  auth.ErrorResponse
// @Failure      500  {object}  model.ResultMessage
// @Router       /api/v1/admin/number-of-cycles/{id} [delete]
func (h *AdminHandler) DeleteNumberOfCycle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return utils.FiberResponse(c, model.Response{
			Status: 400,
			Error:  errors.New("number of cycle id must be integer"),
		})
	}

	ctx := c.Context()
	data := h.service.DeleteNumberOfCycle(ctx, id)
	return utils.FiberResponse(c, data)
}
