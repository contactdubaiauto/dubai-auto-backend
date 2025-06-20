package model

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type CreateCarRequest struct {
	UserID          int64  `json:"user_id"`
	BrandID         int64  `json:"brand_id" binding:"required"`
	RegionID        int64  `json:"region_id"`
	CityID          int64  `json:"city_id"`
	ModelID         int64  `json:"model_id" binding:"required"`
	TransmissionID  int64  `json:"transmission_id"`
	EngineID        int64  `json:"engine_id"`
	DrivetrainID    int64  `json:"drivetrain_id"`
	BodyTypeID      int64  `json:"body_type_id" binding:"required"`
	FuelTypeID      int64  `json:"fuel_type_id"`
	OwnershipTypeId int64  `json:"ownership_type_id"`
	Year            int64  `json:"year" binding:"required"`
	Exchange        bool   `json:"exchange"`
	Credit          bool   `json:"credit"`
	Odometer        int64  `json:"odometer"`
	VinCode         string `json:"vin_code"`
	DoorCount       int64  `json:"door_count"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	Price           int64  `json:"price" binding:"required"`
	New             bool   `json:"new"`
	ColorID         int    `json:"color_id"`
	InteriorColorID int64  `json:"interior_color_id"`
	Crash           bool   `json:"crash"`
	Negotiable      bool   `json:"negotiable"`
	CreditPrice     int64  `json:"credit_price"`
	RightHandDrive  bool   `json:"right_hand_drive"`
}
