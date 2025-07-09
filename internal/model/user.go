package model

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type CreateCarRequest struct {
	UserID         int `json:"user_id"`
	BrandID        int `json:"brand_id" binding:"required"`
	RegionID       int `json:"region_id"`
	CityID         int `json:"city_id"`
	ModelID        int `json:"model_id" binding:"required"`
	TransmissionID int `json:"transmission_id"`
	EngineID       int `json:"engine_id"`
	DrivetrainID   int `json:"drivetrain_id"`
	BodyTypeID     int `json:"body_type_id" binding:"required"`
	FuelTypeID     int `json:"fuel_type_id"`
	// OwnershipTypeId int    `json:"ownership_type_id"`
	Year            int    `json:"year" binding:"required"`
	Exchange        bool   `json:"exchange"`
	Credit          bool   `json:"credit"`
	Odometer        int    `json:"odometer"`
	VinCode         string `json:"vin_code"`
	DoorCount       int    `json:"door_count"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Price           int    `json:"price" binding:"required"`
	New             bool   `json:"new"`
	ColorID         int    `json:"color_id"`
	InteriorColorID int    `json:"interior_color_id"`
	Crash           bool   `json:"crash"`
	Negotiable      bool   `json:"negotiable"`
	CreditPrice     int    `json:"credit_price"`
	RightHandDrive  bool   `json:"right_hand_drive"`
	ModificationID  int    `json:"modification_id"`
	// MileageKM       int    `json:"mileage_km"`
	GenerationID int `json:"generation_id" binding:"required"`
}
