package model

import "time"

// RESPONSES
type GetMotorcycleCategoriesResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type GetNumberOfCyclesResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// type GetMotorcycleParameterValuesResponse struct {
// 	Name string `json:"name"`
// 	ID   int    `json:"id"`
// }

// type GetMotorcycleParametersResponse struct {
// 	Values []GetMotorcycleParameterValuesResponse `json:"values"`
// 	Name   string                                 `json:"name"`
// 	ID     int                                    `json:"id"`
// }

type GetMotorcycleBrandsResponse struct {
	Name       string  `json:"name"`
	ID         int     `json:"id"`
	Image      *string `json:"image"`
	ModelCount int     `json:"model_count"`
}

type GetMotorcycleModelsResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// REQUESTS
// type CreateMotorcycleParameterRequest struct {
// 	ParameterID int `json:"parameter_id" validate:"required"`
// 	ValueID     int `json:"value_id" validate:"required"`
// }

type CreateMotorcycleRequest struct {
	Crash            *bool    `json:"crash"`
	Wheel            *bool    `json:"wheel"`
	MotoCategoryID   int      `json:"moto_category_id" validate:"required"`
	BrandID          int      `json:"moto_brand_id" validate:"required"`
	ModelID          int      `json:"moto_model_id" validate:"required"`
	VinCode          string   `json:"vin_code" validate:"required"`
	Description      string   `json:"description"`
	PhoneNumbers     []string `json:"phone_numbers" validate:"required"`
	EngineID         int      `json:"engine_id" validate:"required"`
	CityID           int      `json:"city_id" validate:"required"`
	ColorID          int      `json:"color_id" validate:"required"`
	Engine           int      `json:"engine"`
	Power            int      `json:"power"`
	Year             int      `json:"year" validate:"required"`
	NumberOfCyclesID int      `json:"number_of_cycles_id" validate:"required"`
	TradeIn          int      `json:"trade_in" validate:"required"`
	Odometer         int      `json:"odometer"`
	Owners           int      `json:"owners"`
	Price            int      `json:"price" validate:"required"`
}

type UpdateMotorcycleRequest struct {
	ID               int      `json:"id" validate:"required"`
	Crash            *bool    `json:"crash"`
	Wheel            *bool    `json:"wheel"`
	New              *bool    `json:"new"`
	MotoCategoryID   int      `json:"moto_category_id"`
	BrandID          int      `json:"moto_brand_id"`
	ModelID          int      `json:"moto_model_id"`
	VinCode          string   `json:"vin_code"`
	Description      string   `json:"description"`
	PhoneNumbers     []string `json:"phone_numbers"`
	EngineID         int      `json:"engine_id"`
	CityID           int      `json:"city_id"`
	ColorID          int      `json:"color_id"`
	Engine           int      `json:"engine"`
	Power            int      `json:"power"`
	Year             int      `json:"year"`
	NumberOfCyclesID int      `json:"number_of_cycles_id"`
	TradeIn          int      `json:"trade_in"`
	Odometer         int      `json:"odometer"`
	Owners           int      `json:"owners"`
	Price            int      `json:"price"`
}

// Owner represents the motorcycle owner information
type MotorcycleOwner struct {
	Contacts map[string]string `json:"contacts"`
	Username string            `json:"username"`
	Avatar   string            `json:"avatar"`
	ID       int               `json:"id"`
	RoleID   int               `json:"role_id"`
}

// // MotorcycleParameter represents a motorcycle parameter with its value
// type MotorcycleParameter struct {
// 	Parameter        string `json:"parameter"`
// 	ParameterValue   string `json:"parameter_value"`
// 	ParameterID      int    `json:"parameter_id"`
// 	ParameterValueID int    `json:"parameter_value_id"`
// }

type GetMotorcycleResponse struct {
	UpdatedAt        time.Time       `json:"updated_at"`
	CreatedAt        time.Time       `json:"created_at"`
	Images           []string        `json:"images"`
	Videos           []string        `json:"videos"`
	PhoneNumbers     []string        `json:"phone_numbers"`
	Owner            MotorcycleOwner `json:"owner"`
	VinCode          string          `json:"vin_code"`
	Status           string          `json:"status"`
	MotoCategory     string          `json:"moto_category"`
	MotoBrand        string          `json:"moto_brand"`
	MotoModel        string          `json:"moto_model"`
	EngineType       string          `json:"engine_type"`
	City             string          `json:"city"`
	Color            string          `json:"color"`
	NumberOfCycles   string          `json:"number_of_cycles"`
	ID               int             `json:"id"`
	Year             int             `json:"year"`
	Odometer         int             `json:"odometer"`
	Owners           int             `json:"owners"`
	Price            int             `json:"price"`
	ModerationStatus int             `json:"moderation_status"`
	UserRoleID       int             `json:"user_role_id"`
	TradeIn          int             `json:"trade_in"`
	Description      *string         `json:"description"`
	Crash            *bool           `json:"crash"`
	Wheel            *bool           `json:"wheel"`
	Power            *int            `json:"power"`
	Engine           *int            `json:"engine"`
	MyMoto           bool            `json:"my_moto"`
}

type GetMotorcyclesResponse struct {
	Type      string     `json:"type"`
	CreatedAt *time.Time `json:"created_at"`
	Images    *[]string  `json:"images"`
	Model     *string    `json:"model"`
	Brand     *string    `json:"brand"`
	Status    *int       `json:"status"`
	TradeIn   *int       `json:"trade_in"`
	Year      *int       `json:"year"`
	Price     *int       `json:"price"`
	ViewCount *int       `json:"view_count"`
	New       *bool      `json:"new"`
	Crash     *bool      `json:"crash"`
	MyMoto    *bool      `json:"my_moto"`
	Odometer  *int       `json:"odometer"`
	OwnerName *string    `json:"owner_name"`
	City      *string    `json:"city"`
	ID        int        `json:"id"`
}

type GetEditMotorcycleResponse struct {
	UpdatedAt      time.Time       `json:"updated_at"`
	CreatedAt      time.Time       `json:"created_at"`
	Images         []ImageObject   `json:"images"`
	Videos         []VideoObject   `json:"videos"`
	PhoneNumbers   []string        `json:"phone_numbers"`
	Owner          MotorcycleOwner `json:"owner"`
	VinCode        string          `json:"vin_code"`
	Status         string          `json:"status"`
	ID             int             `json:"id"`
	Year           int             `json:"year"`
	Odometer       int             `json:"odometer"`
	Owners         int             `json:"owners"`
	TradeIn        int             `json:"trade_in"`
	Engine         *int            `json:"engine"`
	Power          *int            `json:"power"`
	Price          *int            `json:"price"`
	Description    *string         `json:"description"`
	MotoCategory   *Model          `json:"moto_category"`
	MotoBrand      *Model          `json:"moto_brand"`
	MotoModel      *Model          `json:"moto_model"`
	EngineType     *Model          `json:"engine_type"`
	NumberOfCycles *Model          `json:"number_of_cycles"`
	City           *City           `json:"city"`
	Color          *Color          `json:"color"`
	Crash          *bool           `json:"crash"`
	Wheel          *bool           `json:"wheel"`
	New            *bool           `json:"new"`
	MyMoto         bool            `json:"my_moto"`
}
