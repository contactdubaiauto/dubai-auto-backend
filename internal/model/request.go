package model

type CreateCompanyTypeRequest struct {
	Name   string `json:"name" validate:"required"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type CreateUserRequest struct {
	Username    string   `json:"username" validate:"required,min=2,max=255"`
	Email       string   `json:"email" validate:"required,email"`
	Password    string   `json:"password" validate:"required,min=8"`
	RoleID      int      `json:"role_id" default:"0"`
	Permissions []string `json:"permissions" validate:"required"`
}

type UpdateUserRequest struct {
	Username    string   `json:"username" validate:"omitempty,min=2,max=255"`
	Email       string   `json:"email" validate:"omitempty,email"`
	Password    string   `json:"password" validate:"omitempty,min=8"`
	Status      int      `json:"status" validate:"omitempty"`
	RoleID      int      `json:"role_id" default:"1"`
	Permissions []string `json:"permissions" validate:"omitempty"`
}

// SendNotificationRequest is the admin request to send global FCM notifications by role.
// RoleID: 1 user, 2 dealer, 3 logistic, 4 broker, 5 car service.
type SendNotificationRequest struct {
	RoleID      int    `json:"role_id" validate:"required,max=5"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
}

// Vehicles (admin) requests
// NOTE: "vehicles" table is used for cars in this project.
type AdminCreateVehicleRequest struct {
	UserID         int      `json:"user_id" validate:"required"`
	PhoneNumbers   []string `json:"phone_numbers" validate:"required"`
	Wheel          *bool    `json:"wheel" validate:"required"` // true left, false right
	Description    string   `json:"description"`
	VinCode        string   `json:"vin_code" validate:"required"`
	CityID         int      `json:"city_id" validate:"required"`
	BrandID        int      `json:"brand_id" validate:"required"`
	ModelID        int      `json:"model_id" validate:"required"`
	ModificationID int      `json:"modification_id" validate:"required"`
	Year           int      `json:"year" validate:"required"`
	Odometer       int      `json:"odometer" validate:"required"`
	Price          int      `json:"price" validate:"required"`
	ColorID        int      `json:"color_id" validate:"required"`
	Owners         int      `json:"owners"`
	TradeIn        int      `json:"trade_in" validate:"required"`
	New            bool     `json:"new"`
	Crash          bool     `json:"crash"`
}

type AdminUpdateVehicleStatusRequest struct {
	Status int `json:"status" validate:"required"` // 1-pending, 2-not sale, 3-on sale
}

type CreateNameRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=255"`
	NameRu      string `json:"name_ru"`
	NameAe      string `json:"name_ae"`
	CountryCode string `json:"country_code"`
}

type CreateBodyTypeRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=50"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Brand requests
type CreateBrandRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=255"`
	NameRu  string `json:"name_ru"`
	NameAe  string `json:"name_ae"`
	Popular bool   `json:"popular"`
}

// Model requests
type CreateModelRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=255"`
	NameRu  string `json:"name_ru"`
	NameAe  string `json:"name_ae"`
	Popular bool   `json:"popular"`
}

type UpdateModelRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=255"`
	NameRu  string `json:"name_ru"`
	NameAe  string `json:"name_ae"`
	BrandID int    `json:"brand_id" validate:"required"`
	Popular bool   `json:"popular"`
}

// Transmission requests
type CreateTransmissionRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Number of cycles requests
type CreateNumberOfCycleRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Engine requests
type CreateEngineRequest struct {
	Name   string `json:"name" validate:"required,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Comtrans Engine requests
type CreateComtransEngineRequest struct {
	Name   string `json:"name" validate:"required,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Moto Engine requests
type CreateMotoEngineRequest struct {
	Name   string `json:"name" validate:"required,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Drivetrain requests
type CreateDrivetrainRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Fuel Type requests
type CreateFuelTypeRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Service Type requests
type CreateServiceTypeRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

// Service requests
type CreateServiceRequest struct {
	Name          string `json:"name" validate:"required,min=2,max=255"`
	ServiceTypeID int    `json:"service_type_id" validate:"required"`
}

type ThirdPartyLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserForgetPasswordReq struct {
	Email string `json:"email" binding:"required,email"`
}

type UserResetPasswordReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

type AdminLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Generation requests
type CreateGenerationRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=255"`
	NameRu    string `json:"name_ru"`
	NameAe    string `json:"name_ae"`
	Image     string `json:"image"`
	ModelID   int    `json:"model_id" validate:"required"`
	StartYear int    `json:"start_year" validate:"required"`
	EndYear   int    `json:"end_year" validate:"required"`
	Wheel     bool   `json:"wheel"`
}

type UpdateGenerationRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=255"`
	NameRu    string `json:"name_ru"`
	NameAe    string `json:"name_ae"`
	ModelID   int    `json:"model_id" validate:"required"`
	StartYear int    `json:"start_year" validate:"required"`
	EndYear   int    `json:"end_year" validate:"required"`
	Wheel     bool   `json:"wheel"`
}

// Generation Modification requests
type CreateGenerationModificationRequest struct {
	BodyTypeID     int `json:"body_type_id" validate:"required"`
	EngineID       int `json:"engine_id" validate:"required"`
	FuelTypeID     int `json:"fuel_type_id" validate:"required"`
	DrivetrainID   int `json:"drivetrain_id" validate:"required"`
	TransmissionID int `json:"transmission_id" validate:"required"`
}

type UpdateGenerationModificationRequest struct {
	BodyTypeID     int `json:"body_type_id" validate:"required"`
	EngineID       int `json:"engine_id" validate:"required"`
	FuelTypeID     int `json:"fuel_type_id" validate:"required"`
	DrivetrainID   int `json:"drivetrain_id" validate:"required"`
	TransmissionID int `json:"transmission_id" validate:"required"`
}

// Configuration requests
type CreateConfigurationRequest struct {
	BodyTypeID   int `json:"body_type_id" validate:"required"`
	GenerationID int `json:"generation_id" validate:"required"`
}

type UpdateConfigurationRequest struct {
	BodyTypeID   int `json:"body_type_id" validate:"required"`
	GenerationID int `json:"generation_id" validate:"required"`
}

// Color requests
type CreateColorRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateColorRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=255"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
	Image  string `json:"image" validate:"required"`
}

// Moto Category requests
type CreateMotoCategoryRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateMotoCategoryRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Moto Brand requests
type CreateMotoBrandRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateMotoBrandRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Moto Model requests
type CreateMotoModelRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	NameRu      string `json:"name_ru"`
	NameAe      string `json:"name_ae"`
	MotoBrandID int    `json:"moto_brand_id" validate:"required"`
}

type UpdateMotoModelRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	NameRu      string `json:"name_ru"`
	NameAe      string `json:"name_ae"`
	MotoBrandID int    `json:"moto_brand_id" validate:"required"`
}

// Moto Parameter requests
type CreateMotoParameterRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	NameRu         string `json:"name_ru"`
	NameAe         string `json:"name_ae"`
	MotoCategoryID int    `json:"moto_category_id" validate:"required"`
}

type UpdateMotoParameterRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	NameRu         string `json:"name_ru"`
	NameAe         string `json:"name_ae"`
	MotoCategoryID int    `json:"moto_category_id" validate:"required"`
}

// Moto Parameter Value requests
type CreateMotoParameterValueRequest struct {
	Name   string `json:"name" validate:"required,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateMotoParameterValueRequest struct {
	Name   string `json:"name" validate:"required,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Moto Category Parameter requests
type CreateMotoCategoryParameterRequest struct {
	MotoParameterID int `json:"moto_parameter_id" validate:"required"`
}

type UpdateMotoCategoryParameterRequest struct {
	MotoParameterID int `json:"moto_parameter_id" validate:"required"`
}

// Comtrans Category requests
type CreateComtransCategoryRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateComtransCategoryRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Comtrans Brand requests
type CreateComtransBrandRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateComtransBrandRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Comtrans Model requests
type CreateComtransModelRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	NameRu          string `json:"name_ru"`
	NameAe          string `json:"name_ae"`
	ComtransBrandID int    `json:"comtrans_brand_id" validate:"required"`
}

type UpdateComtransModelRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	NameRu          string `json:"name_ru"`
	NameAe          string `json:"name_ae"`
	ComtransBrandID int    `json:"comtrans_brand_id" validate:"required"`
}

// Comtrans Parameter requests
type CreateComtransParameterRequest struct {
	Name               string `json:"name" validate:"required,min=2,max=100"`
	NameRu             string `json:"name_ru"`
	NameAe             string `json:"name_ae"`
	ComtransCategoryID int    `json:"comtrans_category_id" validate:"required"`
}

type UpdateComtransParameterRequest struct {
	Name               string `json:"name" validate:"required,min=2,max=100"`
	NameRu             string `json:"name_ru"`
	NameAe             string `json:"name_ae"`
	ComtransCategoryID int    `json:"comtrans_category_id" validate:"required"`
}

// Comtrans Parameter Value requests
type CreateComtransParameterValueRequest struct {
	Name   string `json:"name" validate:"required,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

type UpdateComtransParameterValueRequest struct {
	Name   string `json:"name" validate:"required,max=100"`
	NameRu string `json:"name_ru"`
	NameAe string `json:"name_ae"`
}

// Comtrans Category Parameter requests
type CreateComtransCategoryParameterRequest struct {
	ComtransParameterID int `json:"comtrans_parameter_id" validate:"required"`
}

type UpdateComtransCategoryParameterRequest struct {
	ComtransParameterID int `json:"comtrans_parameter_id" validate:"required"`
}

type AcceptApplicationRequest struct {
	Password string `json:"password"`
}

type ThirdPartyProfileReq struct {
	AboutUs     string            `json:"about_us" validate:"max=300"`
	Message     string            `json:"message"`
	Contacts    map[string]string `json:"contacts"`
	Phone       string            `json:"phone" validate:"required"`
	Address     string            `json:"address"`
	Coordinates string            `json:"coordinates"`
	Username    string            `json:"username"`
}

type ThirdPartyFirstLoginReq struct {
	Message string `json:"message" validate:"required,max=300"`
}

// Logist Destination requests
type CreateLogistDestinationRequest struct {
	FromID int `json:"from_id" validate:"required"`
	ToID   int `json:"to_id" validate:"required"`
}

// Report requests
type CreateReportRequest struct {
	ReportedUserID    int    `json:"reported_user_id" validate:"required"`
	ReportType        string `json:"report_type" validate:"required,min=2,max=255"`
	ReportDescription string `json:"report_description" validate:"max=255"`
}

type CreateItemReportRequest struct {
	ReportedUserID    int    `json:"reported_user_id" validate:"required"`
	ReportType        string `json:"report_type" validate:"required,min=2,max=255"`
	ReportDescription string `json:"report_description" validate:"max=255"`
	ItemType          string `json:"item_type" validate:"required,oneof=car moto comtran"`
	ItemID            int    `json:"item_id" validate:"required"`
}

type UpdateReportRequest struct {
	ReportStatus int `json:"report_status" validate:"required,max=3"` // 1-pending, 2-resolved, 3-closed
}

// ModerateItemRequest is the request body for moderating items (vehicles, motorcycles, comtrans)
type ModerateItemRequest struct {
	ID          int    `json:"id" validate:"required"`
	Status      int    `json:"status" validate:"required,max=3"` // 1-pending, 2-accepted, 3-declined
	Title       string `json:"title"`                            // optional, used for declined notification
	Description string `json:"description"`                      // optional, used for declined notification
}
