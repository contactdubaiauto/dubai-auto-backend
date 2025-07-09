package model

import "time"

type Brands struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	CarCount int    `json:"car_count"`
}

type GetBrandsResponse struct {
	PopularBrands []Brands `json:"popular_brands"`
	AllBrands     []Brands `json:"all_brands"`
}

type GetModificationsResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CarCount int    `json:"car_count"`
}

type GetModelsResponse struct {
	PopularModels []Model `json:"popular_models"`
	AllModels     []Model `json:"all_models"`
}

type Generation struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Image         string          `json:"image"`
	StartYear     int             `json:"start_year"`
	EndYear       int             `json:"end_year"`
	FuelTypes     []*FuelType     `json:"fuel_types"`
	BodyTypes     []*BodyType     `json:"body_types"`
	Drivetrains   []*Drivetrain   `json:"drivetrains"`
	Transmissions []*Transmission `json:"transmissions"`
}

type BodyType struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
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
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	HexCode *string `json:"hex_code"`
}

type GetCarsResponse struct {
	ID            int        `json:"id"`
	Brand         *string    `json:"brand"`
	Region        *string    `json:"region"`
	City          *string    `json:"city"`
	Model         string     `json:"model"`
	Transmission  *string    `json:"transmission"`
	Engine        *string    `json:"engine"`
	Drivetrain    *string    `json:"drivetrain"`
	BodyType      string     `json:"body_type"`
	FuelType      *string    `json:"fuel_type"`
	Year          int        `json:"year"`
	Price         int        `json:"price"`
	Mileage       *int       `json:"mileage"`
	VinCode       *string    `json:"vin_code"`
	Exchange      *bool      `json:"exchange"`
	Credit        *bool      `json:"credit"`
	New           *bool      `json:"new"`
	Color         *string    `json:"color"`
	InteriorColor *string    `json:"interior_color"`
	CreditPrice   *int       `json:"credit_price"`
	Status        *int       `json:"status"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	Images        *[]string  `json:"images"`
	PhoneNumber   *string    `json:"phone_number"`
	ViewCount     int        `json:"view_count"`
	MyCar         *bool      `json:"my_car"`
}
