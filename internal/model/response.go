package model

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
