package model

// RESPONSES
type GetMotorcycleCategoriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetMotorcycleParameterValuesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetMotorcycleParametersResponse struct {
	ID     int                                    `json:"id"`
	Name   string                                 `json:"name"`
	Values []GetMotorcycleParameterValuesResponse `json:"values"`
}

type GetMotorcycleBrandsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type GetMotorcycleModelsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// REQUESTS
type CreateMotorcycleParameterRequest struct {
	ParameterID int `json:"parameter_id" validate:"required"`
	ValueID     int `json:"value_id" validate:"required"`
}

type CreateMotorcycleRequest struct {
	MotoCategoryID     string                             `json:"moto_category_id" validate:"required"`
	BrandID            string                             `json:"moto_brand_id" validate:"required"`
	ModelID            string                             `json:"moto_model_id" validate:"required"`
	FuelTypeID         int                                `json:"fuel_type_id" validate:"required"`
	CityID             int                                `json:"city_id" validate:"required"`
	ColorID            int                                `json:"color_id" validate:"required"`
	Engine             int                                `json:"engine"`
	Power              int                                `json:"power"`
	Year               int                                `json:"year" validate:"required"`
	NumberOfCycles     int                                `json:"number_of_cycles"`
	Odometer           int                                `json:"odometer"`
	Crash              *bool                              `json:"crash"`
	NotCleared         *bool                              `json:"not_cleared"`
	Owners             int                                `json:"owners"`
	DateOfPurchase     string                             `json:"date_of_purchase"`
	WarrantyDate       string                             `json:"warranty_date"`
	PTC                *bool                              `json:"ptc"`
	VinCode            string                             `json:"vin_code" validate:"required"`
	Certificate        string                             `json:"certificate"`
	Description        string                             `json:"description"`
	CanLookCoordinate  string                             `json:"can_look_coordinate"`
	PhoneNumber        string                             `json:"phone_number" validate:"required"`
	RefuseDealersCalls *bool                              `json:"refuse_dealers_calls"`
	OnlyChat           *bool                              `json:"only_chat"`
	ProtectSpam        *bool                              `json:"protect_spam"`
	VerifiedBuyers     *bool                              `json:"verified_buyers"`
	ContactPerson      string                             `json:"contact_person"`
	Email              string                             `json:"email"`
	Price              int                                `json:"price" validate:"required"`
	PriceType          string                             `json:"price_type" validate:"required,oneof=USD AED RUB EUR"`
	Parameters         []CreateMotorcycleParameterRequest `json:"parameters"`
}
