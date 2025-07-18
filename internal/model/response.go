package model

import "time"

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
	Mileage      *int       `json:"mileage"`
	VinCode      *string    `json:"vin_code"`
	Exchange     *bool      `json:"exchange"`
	Credit       *bool      `json:"credit"`
	New          *bool      `json:"new"`
	Color        *string    `json:"color"`
	Status       *int       `json:"status"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Images       *[]string  `json:"images"`
	Videos       *[]string  `json:"videos"`
	PhoneNumbers *[]string  `json:"phone_numbers"`
	ViewCount    int        `json:"view_count"`
	MyCar        *bool      `json:"my_car"`
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
	Exchange     *bool              `json:"exchange"`
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
	Description  string             `json:"description"`
	Crash        *bool              `json:"crash"`
}
