package model

import "time"

// RESPONSES
type GetComtransCategoriesResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type GetComtransBrandsResponse struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	ID    int    `json:"id"`
}

type GetComtransModelsResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type CreateComtransRequest struct {
	Crash             *bool    `json:"crash"`
	Wheel             *bool    `json:"wheel"`
	ComtranCategoryID int      `json:"comtran_category_id" validate:"required"`
	BrandID           int      `json:"comtran_brand_id" validate:"required"`
	ModelID           int      `json:"comtran_model_id" validate:"required"`
	VinCode           string   `json:"vin_code" validate:"required"`
	Description       string   `json:"description"`
	PhoneNumbers      []string `json:"phone_numbers" validate:"required"`
	FuelTypeID        int      `json:"fuel_type_id" validate:"required"`
	ColorID           int      `json:"color_id" validate:"required"`
	Engine            int      `json:"engine"`
	Power             int      `json:"power"`
	Year              int      `json:"year" validate:"required"`
	TradeIn           int      `json:"trade_in" validate:"required"`
	Odometer          int      `json:"odometer"`
	Owners            int      `json:"owners"`
	Price             int      `json:"price" validate:"required"`
}

type UpdateComtransRequest struct {
	ID                int      `json:"id" validate:"required"`
	Crash             *bool    `json:"crash"`
	Wheel             *bool    `json:"wheel"`
	ComtranCategoryID int      `json:"comtran_category_id"`
	BrandID           int      `json:"comtran_brand_id"`
	ModelID           int      `json:"comtran_model_id"`
	VinCode           string   `json:"vin_code"`
	Description       string   `json:"description"`
	CanLookCoordinate string   `json:"can_look_coordinate"`
	PhoneNumbers      []string `json:"phone_numbers"`
	FuelTypeID        int      `json:"fuel_type_id"`
	CityID            int      `json:"city_id"`
	ColorID           int      `json:"color_id"`
	Engine            int      `json:"engine"`
	Power             int      `json:"power"`
	Year              int      `json:"year"`
	TradeIn           int      `json:"trade_in"`
	Odometer          int      `json:"odometer"`
	Owners            int      `json:"owners"`
	Price             int      `json:"price"`
}

// Owner represents the comtrans owner information
type ComtransOwner struct {
	Contacts map[string]string `json:"contacts"`
	Username string            `json:"username"`
	Avatar   string            `json:"avatar"`
	ID       int               `json:"id"`
}

type GetComtransResponse struct {
	UpdatedAt         time.Time     `json:"updated_at"`
	CreatedAt         time.Time     `json:"created_at"`
	Images            []string      `json:"images"`
	Videos            []string      `json:"videos"`
	PhoneNumbers      []string      `json:"phone_numbers"`
	Owner             ComtransOwner `json:"owner"`
	Crash             *bool         `json:"crash"`
	VinCode           string        `json:"vin_code"`
	Wheel             *bool         `json:"wheel"`
	Description       string        `json:"description"`
	CanLookCoordinate string        `json:"can_look_coordinate"`
	Status            string        `json:"status"`
	ComtranCategory   string        `json:"comtran_category"`
	ComtranBrand      string        `json:"comtran_brand"`
	ComtranModel      string        `json:"comtran_model"`
	FuelType          string        `json:"fuel_type"`
	City              string        `json:"city"`
	Color             string        `json:"color"`
	ID                int           `json:"id"`
	Engine            int           `json:"engine"`
	Power             int           `json:"power"`
	Year              int           `json:"year"`
	Odometer          int           `json:"odometer"`
	Owners            int           `json:"owners"`
	TradeIn           int           `json:"trade_in"`
	Price             int           `json:"price"`
	MyComtrans        bool          `json:"my_comtrans"`
}
