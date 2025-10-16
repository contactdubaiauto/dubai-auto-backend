package model

type CreateNameRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

type CreateBodyTypeRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Image string `json:"image"`
}

// Brand requests
type CreateBrandRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=255"`
	Popular bool   `json:"popular"`
}

// Model requests
type CreateModelRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=255"`
	Popular bool   `json:"popular"`
}

type UpdateModelRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=255"`
	BrandID int    `json:"brand_id" validate:"required"`
	Popular bool   `json:"popular"`
}

// Transmission requests
type CreateTransmissionRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

// Engine requests
type CreateEngineRequest struct {
	Value string `json:"value" validate:"required,min=1,max=255"`
}

// Drivetrain requests
type CreateDrivetrainRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
}

// Fuel Type requests
type CreateFuelTypeRequest struct {
	Name string `json:"name" validate:"required,min=2,max=255"`
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

type AdminLoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Generation requests
type CreateGenerationRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=255"`
	Image     string `json:"image"`
	ModelID   int    `json:"model_id" validate:"required"`
	StartYear int    `json:"start_year" validate:"required"`
	EndYear   int    `json:"end_year" validate:"required"`
	Wheel     bool   `json:"wheel"`
}

type UpdateGenerationRequest struct {
	Name      string `json:"name" validate:"required,min=2,max=255"`
	Image     string `json:"image"`
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
	Name  string `json:"name" validate:"required,min=2,max=255"`
	Image string `json:"image" validate:"required"`
}

type UpdateColorRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=255"`
	Image string `json:"image" validate:"required"`
}

// Moto Category requests
type CreateMotoCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type UpdateMotoCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// Moto Brand requests
type CreateMotoBrandRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	MotoCategoryID int    `json:"moto_category_id" validate:"required"`
}

type UpdateMotoBrandRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	Image          string `json:"image" validate:"required"`
	MotoCategoryID int    `json:"moto_category_id" validate:"required"`
}

// Moto Model requests
type CreateMotoModelRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	MotoBrandID int    `json:"moto_brand_id" validate:"required"`
}

type UpdateMotoModelRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	MotoBrandID int    `json:"moto_brand_id" validate:"required"`
}

// Moto Parameter requests
type CreateMotoParameterRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	MotoCategoryID int    `json:"moto_category_id" validate:"required"`
}

type UpdateMotoParameterRequest struct {
	Name           string `json:"name" validate:"required,min=2,max=100"`
	MotoCategoryID int    `json:"moto_category_id" validate:"required"`
}

// Moto Parameter Value requests
type CreateMotoParameterValueRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type UpdateMotoParameterValueRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
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
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type UpdateComtransCategoryRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// Comtrans Brand requests
type CreateComtransBrandRequest struct {
	Name               string `json:"name" validate:"required,min=2,max=100"`
	ComtransCategoryID int    `json:"comtrans_category_id" validate:"required"`
}

type UpdateComtransBrandRequest struct {
	Name               string `json:"name" validate:"required,min=2,max=100"`
	ComtransCategoryID int    `json:"comtrans_category_id" validate:"required"`
}

// Comtrans Model requests
type CreateComtransModelRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	ComtransBrandID int    `json:"comtrans_brand_id" validate:"required"`
}

type UpdateComtransModelRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=100"`
	ComtransBrandID int    `json:"comtrans_brand_id" validate:"required"`
}

// Comtrans Parameter requests
type CreateComtransParameterRequest struct {
	Name               string `json:"name" validate:"required,min=2,max=100"`
	ComtransCategoryID int    `json:"comtrans_category_id" validate:"required"`
}

type UpdateComtransParameterRequest struct {
	Name               string `json:"name" validate:"required,min=2,max=100"`
	ComtransCategoryID int    `json:"comtrans_category_id" validate:"required"`
}

// Comtrans Parameter Value requests
type CreateComtransParameterValueRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type UpdateComtransParameterValueRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
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
	AboutUs     string `json:"about_us" validate:"max=300"`
	Email       string `json:"email" validate:"required"`
	Message     string `json:"message"`
	Whatsapp    string `json:"whatsapp"`
	Telegram    string `json:"telegram"`
	Phone       string `json:"phone" validate:"required"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
}

type ThirdPartyFirstLoginReq struct {
	Message string `json:"message" validate:"required,max=300"`
}

// Logist Destination requests
type CreateLogistDestinationRequest struct {
	FromID int `json:"from_id" validate:"required"`
	ToID   int `json:"to_id" validate:"required"`
}
