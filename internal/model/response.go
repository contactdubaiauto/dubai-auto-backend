package model

import "time"

type GetBrandsResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	CarCount int64  `json:"car_count"`
}

type Model struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type BodyType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Transmission struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Engine struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

type Drive struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type FuelType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetCarsResponse struct {
	ID           int        `json:"id"`
	Brand        *string    `json:"brand"`
	Region       *string    `json:"region"`
	City         *string    `json:"city"`
	Model        string     `json:"model"`
	Transmission *string    `json:"transmission"`
	Engine       *string    `json:"engine"`
	Drive        *string    `json:"drive"`
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
	CreditPrice  *int       `json:"credit_price"`
	Status       *int       `json:"status"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	Images       *[]string  `json:"images"`
	PhoneNumber  *string    `json:"phone_number"`
}
