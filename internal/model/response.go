package model

import "time"

type GetPriceRecommendationRequest struct {
	BrandID        string `json:"brand_id"`
	ModelID        string `json:"model_id"`
	Year           string `json:"year"`
	ModificationID string `json:"modification_id"`
	Odometer       string `json:"odometer"`
	CityID         string `json:"city_id"`
}

type GetPriceRecommendationResponse struct {
	MinPrice int `json:"min_price"`
	MaxPrice int `json:"max_price"`
	AvgPrice int `json:"avg_price"`
}

type Brand struct {
	ID         *int    `json:"id"`
	Name       *string `json:"name"`
	Logo       *string `json:"logo"`
	ModelCount *int    `json:"model_count"`
}

type GetBrandsResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	ModelCount int    `json:"model_count"`
}

type GetProfileResponse struct {
	ID                int        `json:"id"`
	Email             *string    `json:"email"`
	Phone             *string    `json:"phone"`
	DrivingExperience *int       `json:"driving_experience"`
	Notification      *bool      `json:"notification"`
	Username          *string    `json:"username"`
	Google            *string    `json:"google"`
	RegisteredBy      *string    `json:"registered_by"`
	Birthday          *time.Time `json:"birthday"`
	AboutMe           *string    `json:"about_me"`
}

type GetFilterBrandsResponse struct {
	PopularBrands []Brand `json:"popular_brands"`
	AllBrands     []Brand `json:"all_brands"`
}

type Region struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type GetCitiesResponse struct {
	ID      int      `json:"id"`
	Name    *string  `json:"name"`
	Regions []Region `json:"regions"`
}

type GetModificationsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type GetFilterModelsResponse struct {
	PopularModels []Model `json:"popular_models"`
	AllModels     []Model `json:"all_models"`
}

type GetYearsResponse struct {
	Years []*int `json:"years"`
}

type Modification struct {
	ID           *int    `json:"id"`
	Engine       *string `json:"engine"`
	FuelType     *string `json:"fuel_type"`
	Drivetrain   *string `json:"drivetrain"`
	Transmission *string `json:"transmission"`
}

type Generation struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Image         string         `json:"image"`
	StartYear     int            `json:"start_year"`
	EndYear       int            `json:"end_year"`
	Modifications []Modification `json:"modifications"`
}

type BodyType struct {
	ID    *int    `json:"id"`
	Name  *string `json:"name"`
	Image *string `json:"image"`
}

type Transmission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Engine struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type Drivetrain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FuelType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Color struct {
	ID    *int    `json:"id"`
	Name  *string `json:"name"`
	Image *string `json:"image"`
}

type Home struct {
	Popular []GetCarsResponse `json:"popular"`
}

type Owner struct {
	Id       *int    `json:"id"`
	Avatar   *string `json:"avatar"`
	Username *string `json:"username"`
}

type GetCarsResponse struct {
	ID           int        `json:"id"`
	Brand        *string    `json:"brand"`
	Region       *string    `json:"region"`
	City         *string    `json:"city"`
	Model        string     `json:"model"`
	Transmission *string    `json:"transmission"`
	Engine       *string    `json:"engine"`
	Drivetrain   *string    `json:"drivetrain"`
	BodyType     string     `json:"body_type"`
	FuelType     *string    `json:"fuel_type"`
	Year         int        `json:"year"`
	Price        int        `json:"price"`
	Mileage      *int       `json:"mileage"` // todo: change it to odometer
	VinCode      *string    `json:"vin_code"`
	Credit       *bool      `json:"credit"`
	New          *bool      `json:"new"`
	Crash        *bool      `json:"crash"`
	Color        *string    `json:"color"`
	Status       *int       `json:"status"`
	TradeIn      *int       `json:"trade_in"`
	Owners       *int       `json:"owners"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Images       *[]string  `json:"images"`
	Videos       *[]string  `json:"videos"`
	PhoneNumbers *[]string  `json:"phone_numbers"`
	ViewCount    int        `json:"view_count"`
	MyCar        *bool      `json:"my_car"`
	Liked        *bool      `json:"liked"`
	Owner        *Owner     `json:"owner"`
	Description  *string    `json:"description"`
}

type City struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type EditCarGeneration struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
}

type GetEditCarsResponse struct {
	ID           int                `json:"id"`
	Brand        *Brand             `json:"brand"`
	Region       *Region            `json:"region"`
	City         *City              `json:"city"`
	Model        *Model             `json:"model"`
	Modification *Modification      `json:"modification"`
	BodyType     *BodyType          `json:"body_type"`
	Generation   *EditCarGeneration `json:"generation"`
	Year         *int               `json:"year"`
	Price        *int               `json:"price"`
	Odometer     *int               `json:"odometer"`
	VinCode      *string            `json:"vin_code"`
	Credit       *bool              `json:"credit"`
	New          *bool              `json:"new"`
	Color        *Color             `json:"color"`
	Status       *int               `json:"status"`
	CreatedAt    *time.Time         `json:"created_at"`
	UpdatedAt    *time.Time         `json:"updated_at"`
	Images       *[]string          `json:"images"`
	Videos       *[]string          `json:"videos"`
	PhoneNumbers *[]string          `json:"phone_numbers"`
	ViewCount    *int               `json:"view_count"`
	MyCar        *bool              `json:"my_car"`
	Wheel        *bool              `json:"wheel"`
	TradeIN      *int               `json:"trade_id"`
	Owners       *int               `json:"owners"`
	Description  *string            `json:"description"`
	Crash        *bool              `json:"crash"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type AdminApplicationResponse struct {
	LicenceIssueDate  time.Time `json:"licence_issue_date"`
	LicenceExpiryDate time.Time `json:"licence_expiry_date"`
	CreatedAt         time.Time `json:"created_at"`
	CompanyName       string    `json:"company_name"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Status            string    `json:"status"`
	ID                int       `json:"id"`
}

// Admin response models
type AdminCityResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminBrandResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Logo       string    `json:"logo"`
	ModelCount int       `json:"model_count"`
	Popular    bool      `json:"popular"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AdminModelResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	BrandID   int       `json:"brand_id"`
	BrandName string    `json:"brand_name"`
	Popular   bool      `json:"popular"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminBodyTypeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminTransmissionResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminEngineResponse struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminDrivetrainResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminFuelTypeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminServiceTypeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminServiceResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	ServiceTypeID   int       `json:"service_type_id"`
	ServiceTypeName string    `json:"service_type_name"`
	CreatedAt       time.Time `json:"created_at"`
}

type AdminGenerationResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	ModelID   int       `json:"model_id"`
	ModelName string    `json:"model_name"`
	StartYear int       `json:"start_year"`
	EndYear   int       `json:"end_year"`
	Wheel     bool      `json:"wheel"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminGenerationModificationResponse struct {
	ID               int    `json:"id"`
	GenerationID     int    `json:"generation_id"`
	BodyTypeID       int    `json:"body_type_id"`
	BodyTypeName     string `json:"body_type_name"`
	EngineID         int    `json:"engine_id"`
	EngineValue      string `json:"engine_value"`
	FuelTypeID       int    `json:"fuel_type_id"`
	FuelTypeName     string `json:"fuel_type_name"`
	DrivetrainID     int    `json:"drivetrain_id"`
	DrivetrainName   string `json:"drivetrain_name"`
	TransmissionID   int    `json:"transmission_id"`
	TransmissionName string `json:"transmission_name"`
}

type AdminConfigurationResponse struct {
	ID           int    `json:"id"`
	BodyTypeID   int    `json:"body_type_id"`
	BodyTypeName string `json:"body_type_name"`
	GenerationID int    `json:"generation_id"`
}

type AdminColorResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminMotoCategoryResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminMotoBrandResponse struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Image            string    `json:"image"`
	MotoCategoryID   int       `json:"moto_category_id"`
	MotoCategoryName string    `json:"moto_category_name"`
	CreatedAt        time.Time `json:"created_at"`
}

type AdminMotoModelResponse struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	MotoBrandID   int       `json:"moto_brand_id"`
	MotoBrandName string    `json:"moto_brand_name"`
	CreatedAt     time.Time `json:"created_at"`
}

type AdminMotoParameterResponse struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	MotoCategoryID   int       `json:"moto_category_id"`
	MotoCategoryName string    `json:"moto_category_name"`
	CreatedAt        time.Time `json:"created_at"`
}

type AdminMotoParameterValueResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	MotoParameterID int       `json:"moto_parameter_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type AdminMotoCategoryParameterResponse struct {
	ID                int       `json:"id"`
	MotoCategoryID    int       `json:"moto_category_id"`
	MotoParameterID   int       `json:"moto_parameter_id"`
	MotoParameterName string    `json:"moto_parameter_name"`
	CreatedAt         time.Time `json:"created_at"`
}

type AdminComtransCategoryResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminComtransBrandResponse struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	Image                string    `json:"image"`
	ComtransCategoryID   int       `json:"comtrans_category_id"`
	ComtransCategoryName string    `json:"comtrans_category_name"`
	CreatedAt            time.Time `json:"created_at"`
}

type AdminComtransModelResponse struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	ComtransBrandID   int       `json:"comtrans_brand_id"`
	ComtransBrandName string    `json:"comtrans_brand_name"`
	CreatedAt         time.Time `json:"created_at"`
}

type AdminComtransParameterResponse struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	ComtransCategoryID   int       `json:"comtrans_category_id"`
	ComtransCategoryName string    `json:"comtrans_category_name"`
	CreatedAt            time.Time `json:"created_at"`
}

type AdminComtransParameterValueResponse struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	ComtransParameterID int       `json:"comtrans_parameter_id"`
	CreatedAt           time.Time `json:"created_at"`
}

type AdminComtransCategoryParameterResponse struct {
	ID                    int       `json:"id"`
	ComtransCategoryID    int       `json:"comtrans_category_id"`
	ComtransParameterID   int       `json:"comtrans_parameter_id"`
	ComtransParameterName string    `json:"comtrans_parameter_name"`
	CreatedAt             time.Time `json:"created_at"`
}
