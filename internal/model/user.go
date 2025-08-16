package model

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type DeleteCarImageRequest struct {
	Image string `json:"image" validate:"required"`
}

type CreateCarRequest struct {
	// new
	CityID  int `json:"city_id" validate:"required"`
	BrandID int `json:"brand_id" validate:"required"`
	ModelID int `json:"model_id" validate:"required"`
	// BodyTypeID     int      `json:"body_type_id" validate:"required"`
	Wheel          *bool    `json:"wheel" validate:"required"` // left true, right false
	ModificationID int      `json:"modification_id" validate:"required"`
	Year           int      `json:"year" validate:"required"`
	Odometer       int      `json:"odometer" validate:"required"`
	VinCode        string   `json:"vin_code" validate:"required"`
	PhoneNumbers   []string `json:"phone_numbers" validate:"required"`
	Price          int      `json:"price" validate:"required"`
	New            bool     `json:"new"`
	Crash          bool     `json:"crash"`
	Description    string   `json:"description"`
	ColorID        int      `json:"color_id" validate:"required"`
	Owners         int      `json:"owners" validate:"required"`
	TradeIn        int      `json:"trade_in" validate:"required"`

	//
	// OwnershipTypeId int    `json:"ownership_type_id"`
	// Credit          bool   `json:"credit"`
	// DoorCount       int    `json:"door_count"`
	// InteriorColorID int `json:"interior_color_id"`
	// Negotiable      bool `json:"negotiable"`
	// ModificationID  int  `json:"modification_id"`
	// MileageKM       int    `json:"mileage_km"`
	// GenerationID int `json:"generation_id" validate:"required"`
}

type DeleteCarVideoRequest struct {
	Video string `json:"video" validate:"required"`
}

type UpdateCarRequest struct {
	ID             int      `json:"id" validate:"required"`
	CityID         int      `json:"city_id" validate:"required"`
	BrandID        int      `json:"brand_id" validate:"required"`
	ModificationID int      `json:"modification_id" validate:"required"`
	ModelID        int      `json:"model_id" validate:"required"`
	Wheel          *bool    `json:"wheel" validate:"required"` // left true, right false
	Year           int      `json:"year" validate:"required"`
	Odometer       int      `json:"odometer" validate:"required"`
	VinCode        string   `json:"vin_code" validate:"required"`
	PhoneNumbers   []string `json:"phone_numbers" validate:"required"`
	Price          int      `json:"price" validate:"required"`
	New            bool     `json:"new"`
	Crash          bool     `json:"crash"`
	ColorID        int      `json:"color_id" validate:"required"`
	Owners         int      `json:"owners" validate:"required"`
	Description    string   `json:"description"`
	TradeIn        int      `json:"trade_in" validate:"required"`
}

type UpdateProfileRequest struct {
	DrivingExperience int    `json:"driving_experience"`
	Notification      bool   `json:"notification"`
	Username          string `json:"username" validate:"required,min=3,max=20"`
	Google            string `json:"google"`
	Birthday          string `json:"birthday"`
	AboutMe           string `json:"about_me"`
	// todo: add city
}
