package model

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type CreateCarRequest struct {
	UserID         int64  `json:"user_id"`
	BrandID        int64  `json:"brand_id" binding:"required"`
	RegionID       int64  `json:"region_id"`
	CityID         int64  `json:"city_id"`
	ModelID        int64  `json:"model_id" binding:"required"`
	TransmissionID int64  `json:"transmission_id"`
	EngineID       int64  `json:"engine_id"`
	DriveID        int64  `json:"drive_id"`
	BodyTypeID     int64  `json:"body_type_id" binding:"required"`
	FuelTypeID     int64  `json:"fuel_type_id"`
	Ownership      int64  `json:"ownership"`
	Year           int64  `json:"year" binding:"required"`
	Exchange       bool   `json:"exchange"`
	Credit         bool   `json:"credit"`
	Milage         int64  `json:"milage"`
	VinCode        string `json:"vin_code"`
	DoorCount      int64  `json:"door_count"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Price          int64  `json:"price" binding:"required"`
	New            bool   `json:"new"`
	Color          string `json:"color"`
	CreditPrice    int64  `json:"credit_price"`
}
