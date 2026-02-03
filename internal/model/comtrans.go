package model

import "time"

// RESPONSES
type GetComtransCategoriesResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type GetComtransBrandsResponse struct {
	Name       string  `json:"name"`
	ModelCount int     `json:"model_count"`
	Image      *string `json:"image"`
	ID         int     `json:"id"`
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
	CityID            int      `json:"city_id" validate:"required"`
	ModelID           int      `json:"comtran_model_id" validate:"required"`
	VinCode           string   `json:"vin_code" validate:"required"`
	Description       string   `json:"description"`
	PhoneNumbers      []string `json:"phone_numbers" validate:"required"`
	EngineID          int      `json:"engine_id" validate:"required"`
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
	New               *bool    `json:"new"`
	ComtranCategoryID int      `json:"comtran_category_id"`
	BrandID           int      `json:"comtran_brand_id"`
	CityID            int      `json:"city_id"`
	ModelID           int      `json:"comtran_model_id"`
	VinCode           string   `json:"vin_code"`
	Description       string   `json:"description"`
	PhoneNumbers      []string `json:"phone_numbers"`
	EngineID          int      `json:"engine_id"`
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
	RoleID   int               `json:"role_id"`
	Contacts map[string]string `json:"contacts"`
	Username string            `json:"username"`
	Avatar   string            `json:"avatar"`
	ID       int               `json:"id"`
}

type GetComtransResponse struct {
	Type      string     `json:"type"`
	Odometer  int        `json:"odometer"`
	ID        int        `json:"id"`
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
	Liked     *bool      `json:"liked"`
	MyComtran *bool      `json:"my_comtran"`
	OwnerName *string    `json:"owner_name"`
	City      *string    `json:"city"`
}

type GetComtranResponse struct {
	UpdatedAt        time.Time     `json:"updated_at"`
	CreatedAt        time.Time     `json:"created_at"`
	Images           []string      `json:"images"`
	Videos           []string      `json:"videos"`
	PhoneNumbers     []string      `json:"phone_numbers"`
	Owner            ComtransOwner `json:"owner"`
	VinCode          string        `json:"vin_code"`
	Status           string        `json:"status"`
	ComtranCategory  string        `json:"comtran_category"`
	ComtranBrand     string        `json:"comtran_brand"`
	ComtranModel     string        `json:"comtran_model"`
	EngineType       string        `json:"engine_type"`
	City             string        `json:"city"`
	Color            string        `json:"color"`
	ID               int           `json:"id"`
	Year             int           `json:"year"`
	Odometer         int           `json:"odometer"`
	Owners           int           `json:"owners"`
	TradeIn          int           `json:"trade_in"`
	Price            int           `json:"price"`
	UserRoleID       int           `json:"user_role_id"`
	ModerationStatus int           `json:"moderation_status"`
	Description      *string       `json:"description"`
	Power            *int          `json:"power"`
	Crash            *bool         `json:"crash"`
	Wheel            *bool         `json:"wheel"`
	Engine           *int          `json:"engine"`
	MyComtrans       bool          `json:"my_comtrans"`
}

type GetEditComtransResponse struct {
	UpdatedAt       time.Time     `json:"updated_at"`
	CreatedAt       time.Time     `json:"created_at"`
	Images          []ImageObject `json:"images"`
	Videos          []VideoObject `json:"videos"`
	PhoneNumbers    []string      `json:"phone_numbers"`
	Owner           ComtransOwner `json:"owner"`
	VinCode         string        `json:"vin_code"`
	Status          string        `json:"status"`
	ID              int           `json:"id"`
	Year            int           `json:"year"`
	Odometer        int           `json:"odometer"`
	Owners          int           `json:"owners"`
	TradeIn         int           `json:"trade_in"`
	Price           int           `json:"price"`
	ComtranCategory *Model        `json:"comtran_category"`
	ComtranBrand    *Model        `json:"comtran_brand"`
	ComtranModel    *Model        `json:"comtran_model"`
	EngineType      *Model        `json:"engine_type"`
	City            *City         `json:"city"`
	Color           *Color        `json:"color"`
	Engine          *int          `json:"engine"`
	Power           *int          `json:"power"`
	Description     *string       `json:"description"`
	Crash           *bool         `json:"crash"`
	Wheel           *bool         `json:"wheel"`
	New             *bool         `json:"new"`
	MyComtrans      bool          `json:"my_comtrans"`
}
