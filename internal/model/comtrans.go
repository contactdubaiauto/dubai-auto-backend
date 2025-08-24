package model

import "time"

// RESPONSES
type GetComtransCategoriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetComtransParameterValuesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetComtransParametersResponse struct {
	ID     int                                  `json:"id"`
	Name   string                               `json:"name"`
	Values []GetComtransParameterValuesResponse `json:"values"`
}

type GetComtransBrandsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type GetComtransModelsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// REQUESTS
type CreateComtransParameterRequest struct {
	ParameterID int `json:"parameter_id" validate:"required"`
	ValueID     int `json:"value_id" validate:"required"`
}

type CreateComtransRequest struct {
	ComtranCategoryID  string                           `json:"comtran_category_id" validate:"required"`
	BrandID            string                           `json:"comtran_brand_id" validate:"required"`
	ModelID            string                           `json:"comtran_model_id" validate:"required"`
	FuelTypeID         int                              `json:"fuel_type_id" validate:"required"`
	CityID             int                              `json:"city_id" validate:"required"`
	ColorID            int                              `json:"color_id" validate:"required"`
	Engine             int                              `json:"engine"`
	Power              int                              `json:"power"`
	Year               int                              `json:"year" validate:"required"`
	NumberOfCycles     int                              `json:"number_of_cycles"`
	Odometer           int                              `json:"odometer"`
	Crash              *bool                            `json:"crash"`
	NotCleared         *bool                            `json:"not_cleared"`
	Owners             int                              `json:"owners"`
	DateOfPurchase     string                           `json:"date_of_purchase"`
	WarrantyDate       string                           `json:"warranty_date"`
	PTC                *bool                            `json:"ptc"`
	VinCode            string                           `json:"vin_code" validate:"required"`
	Certificate        string                           `json:"certificate"`
	Description        string                           `json:"description"`
	CanLookCoordinate  string                           `json:"can_look_coordinate"`
	PhoneNumber        string                           `json:"phone_number" validate:"required"`
	RefuseDealersCalls *bool                            `json:"refuse_dealers_calls"`
	OnlyChat           *bool                            `json:"only_chat"`
	ProtectSpam        *bool                            `json:"protect_spam"`
	VerifiedBuyers     *bool                            `json:"verified_buyers"`
	ContactPerson      string                           `json:"contact_person"`
	Email              string                           `json:"email"`
	Price              int                              `json:"price" validate:"required"`
	PriceType          string                           `json:"price_type" validate:"required,oneof=USD AED RUB EUR"`
	Parameters         []CreateComtransParameterRequest `json:"parameters"`
}

// Owner represents the comtrans owner information
type ComtransOwner struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// ComtransParameter represents a comtrans parameter with its value
type ComtransParameter struct {
	ParameterID      int    `json:"parameter_id"`
	ParameterValueID int    `json:"parameter_value_id"`
	Parameter        string `json:"parameter"`
	ParameterValue   string `json:"parameter_value"`
}

type GetComtransResponse struct {
	ID                 int                 `json:"id"`
	Owner              ComtransOwner       `json:"owner"`
	Engine             int                 `json:"engine"`
	Power              int                 `json:"power"`
	Year               int                 `json:"year"`
	NumberOfCycles     int                 `json:"number_of_cycles"`
	Odometer           int                 `json:"odometer"`
	Crash              *bool               `json:"crash"`
	NotCleared         *bool               `json:"not_cleared"`
	Owners             int                 `json:"owners"`
	DateOfPurchase     string              `json:"date_of_purchase"`
	WarrantyDate       string              `json:"warranty_date"`
	PTC                *bool               `json:"ptc"`
	VinCode            string              `json:"vin_code"`
	Certificate        string              `json:"certificate"`
	Description        string              `json:"description"`
	CanLookCoordinate  string              `json:"can_look_coordinate"`
	PhoneNumber        string              `json:"phone_number"`
	RefuseDealersCalls *bool               `json:"refuse_dealers_calls"`
	OnlyChat           *bool               `json:"only_chat"`
	ProtectSpam        *bool               `json:"protect_spam"`
	VerifiedBuyers     *bool               `json:"verified_buyers"`
	ContactPerson      string              `json:"contact_person"`
	Email              string              `json:"email"`
	Price              int                 `json:"price"`
	PriceType          string              `json:"price_type"`
	Status             string              `json:"status"`
	UpdatedAt          time.Time           `json:"updated_at"`
	CreatedAt          time.Time           `json:"created_at"`
	ComtranCategory    string              `json:"comtran_category"`
	ComtranBrand       string              `json:"comtran_brand"`
	ComtranModel       string              `json:"comtran_model"`
	FuelType           string              `json:"fuel_type"`
	City               string              `json:"city"`
	Color              string              `json:"color"`
	MyCar              bool                `json:"my_car"`
	Parameters         []ComtransParameter `json:"parameters"`
	Images             []string            `json:"images"`
	Videos             []string            `json:"videos"`
}
