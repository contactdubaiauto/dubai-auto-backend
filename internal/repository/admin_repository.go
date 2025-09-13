package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
)

type AdminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{db}
}

// Cities CRUD operations
func (r *AdminRepository) GetCities(ctx *fasthttp.RequestCtx) ([]model.AdminCityResponse, error) {
	cities := make([]model.AdminCityResponse, 0)
	q := `SELECT id, name, created_at FROM cities ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return cities, err
	}
	defer rows.Close()

	for rows.Next() {
		var city model.AdminCityResponse
		if err := rows.Scan(&city.ID, &city.Name, &city.CreatedAt); err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}
	return cities, err
}

func (r *AdminRepository) CreateCity(ctx *fasthttp.RequestCtx, req *model.CreateCityRequest) (int, error) {
	q := `INSERT INTO cities (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateCity(ctx *fasthttp.RequestCtx, id int, req *model.UpdateCityRequest) error {
	q := `UPDATE cities SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteCity(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM cities WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Brands CRUD operations
func (r *AdminRepository) GetBrands(ctx *fasthttp.RequestCtx) ([]model.AdminBrandResponse, error) {
	brands := make([]model.AdminBrandResponse, 0)
	q := `SELECT id, name, logo, model_count, popular, updated_at FROM brands ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return brands, err
	}
	defer rows.Close()

	for rows.Next() {
		var brand model.AdminBrandResponse
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.Logo, &brand.ModelCount, &brand.Popular, &brand.UpdatedAt); err != nil {
			return brands, err
		}
		brands = append(brands, brand)
	}
	return brands, err
}

func (r *AdminRepository) CreateBrand(ctx *fasthttp.RequestCtx, req *model.CreateBrandRequest) (int, error) {
	q := `INSERT INTO brands (name, popular, updated_at) VALUES ($1, $2, NOW()) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Popular).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateBrand(ctx *fasthttp.RequestCtx, id int, req *model.CreateBrandRequest) error {
	q := `UPDATE brands SET name = $2, popular = $3, updated_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Popular)
	return err
}

func (r *AdminRepository) DeleteBrand(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM brands WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Models CRUD operations
func (r *AdminRepository) GetModels(ctx *fasthttp.RequestCtx, brand_id int) ([]model.AdminModelResponse, error) {
	models := make([]model.AdminModelResponse, 0)
	q := `
		SELECT m.id, m.name, m.brand_id, b.name as brand_name, m.popular, m.updated_at 
		FROM models m
		LEFT JOIN brands b ON m.brand_id = b.id
		WHERE m.brand_id = $1
		ORDER BY m.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return models, err
	}
	defer rows.Close()

	for rows.Next() {
		var modelItem model.AdminModelResponse
		if err := rows.Scan(&modelItem.ID, &modelItem.Name, &modelItem.BrandID, &modelItem.BrandName, &modelItem.Popular, &modelItem.UpdatedAt); err != nil {
			return models, err
		}
		models = append(models, modelItem)
	}
	return models, err
}

func (r *AdminRepository) CreateModel(ctx *fasthttp.RequestCtx, brand_id int, req *model.CreateModelRequest) (int, error) {
	q := `INSERT INTO models (name, brand_id, popular, updated_at) VALUES ($1, $2, $3, NOW()) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, brand_id, req.Popular).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateModelRequest) error {
	q := `UPDATE models SET name = $2, brand_id = $3, popular = $4, updated_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.BrandID, req.Popular)
	return err
}

func (r *AdminRepository) DeleteModel(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM models WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Body Types CRUD operations
func (r *AdminRepository) GetBodyTypes(ctx *fasthttp.RequestCtx) ([]model.AdminBodyTypeResponse, error) {
	bodyTypes := make([]model.AdminBodyTypeResponse, 0)
	q := `SELECT id, name, image, created_at FROM body_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return bodyTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var bodyType model.AdminBodyTypeResponse
		if err := rows.Scan(&bodyType.ID, &bodyType.Name, &bodyType.Image, &bodyType.CreatedAt); err != nil {
			return bodyTypes, err
		}
		bodyTypes = append(bodyTypes, bodyType)
	}
	return bodyTypes, err
}

func (r *AdminRepository) CreateBodyType(ctx *fasthttp.RequestCtx, req *model.CreateBodyTypeRequest) (int, error) {
	q := `INSERT INTO body_types (name, image) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Image).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateBodyTypeImage(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `INSERT INTO body_types_images (body_type_id, image) VALUES ($1, $2) RETURNING id`
	_, err := r.db.Exec(ctx, q, id, paths)
	return err
}

func (r *AdminRepository) DeleteBodyTypeImage(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM body_types_images WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *AdminRepository) UpdateBodyType(ctx *fasthttp.RequestCtx, id int, req *model.CreateBodyTypeRequest) error {
	q := `UPDATE body_types SET name = $2, image = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Image)
	return err
}

func (r *AdminRepository) DeleteBodyType(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM body_types WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}
