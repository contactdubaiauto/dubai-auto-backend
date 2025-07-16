package model

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type DeleteCarImageRequest struct {
	Image string `json:"image" binding:"required"`
}

type CreateCarRequest struct {
	// new
	UserID  int `json:"user_id"`
	CityID  int `json:"city_id" binding:"required"`
	BrandID int `json:"brand_id" binding:"required"`
	ModelID int `json:"model_id" binding:"required"`
	// BodyTypeID     int      `json:"body_type_id" binding:"required"`
	Wheel          *bool    `json:"wheel" binding:"required"` // left true, right false
	ModificationID int      `json:"modification_id" binding:"required"`
	Year           int      `json:"year" binding:"required"`
	Odometer       int      `json:"odometer" binding:"required"`
	VinCode        string   `json:"vin_code" binding:"required"`
	PhoneNumbers   []string `json:"phone_numbers" binding:"required"`
	Price          int      `json:"price" binding:"required"`
	Exchange       bool     `json:"exchange"`
	New            bool     `json:"new"`
	Crash          bool     `json:"crash"`
	ColorID        int      `json:"color_id" binding:"required"`
	Owners         int      `json:"owners" binding:"required"`
	TradeIn        int      `json:"trade_in" binding:"required"`

	//
	// OwnershipTypeId int    `json:"ownership_type_id"`
	// Credit          bool   `json:"credit"`
	// DoorCount       int    `json:"door_count"`
	// InteriorColorID int `json:"interior_color_id"`
	// Negotiable      bool `json:"negotiable"`
	// ModificationID  int  `json:"modification_id"`
	// MileageKM       int    `json:"mileage_km"`
	// GenerationID int `json:"generation_id" binding:"required"`
}

type UpdateCarRequest struct {
	ID             int      `json:"id" binding:"required"`
	CityID         int      `json:"city_id" binding:"required"`
	BrandID        int      `json:"brand_id" binding:"required"`
	ModificationID int      `json:"modification_id" binding:"required"`
	ModelID        int      `json:"model_id" binding:"required"`
	Wheel          *bool    `json:"wheel" binding:"required"` // left true, right false
	Year           int      `json:"year" binding:"required"`
	Odometer       int      `json:"odometer" binding:"required"`
	VinCode        string   `json:"vin_code" binding:"required"`
	PhoneNumbers   []string `json:"phone_numbers" binding:"required"`
	Price          int      `json:"price" binding:"required"`
	Exchange       bool     `json:"exchange"`
	New            bool     `json:"new"`
	Crash          bool     `json:"crash"`
	ColorID        int      `json:"color_id" binding:"required"`
	Owners         int      `json:"owners" binding:"required"`
	TradeIn        int      `json:"trade_in" binding:"required"`
}

type UpdateProfileRequest struct {
	DrivingExperience int    `json:"driving_experience"`
	Notification      bool   `json:"notification"`
	Username          string `json:"username" binding:"required,min=3,max=20"`
	Google            string `json:"google"`
	Birthday          string `json:"birthday"`
	AboutMe           string `json:"about_me"`
}
