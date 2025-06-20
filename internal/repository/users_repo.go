package repository

import (
	"context"
	"strings"

	"dubai-auto/internal/model"
	"dubai-auto/pkg"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetBrands(ctx *context.Context, text string) ([]*model.GetBrandsResponse, error) {
	q := `
		SELECT id, name, logo, car_count FROM brands WHERE name ILIKE '%' || $1 || '%'
	`
	rows, err := r.db.Query(*ctx, q, text)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var brands []*model.GetBrandsResponse

	for rows.Next() {
		var brand model.GetBrandsResponse
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.Logo, &brand.CarCount); err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}
	return brands, err
}

func (r *UserRepository) GetModelsByBrandID(ctx *context.Context, brandID int64, text string) ([]model.Model, error) {
	q := `
			SELECT id, name FROM models WHERE brand_id = $1 AND name ILIKE '%' || $2 || '%'
		`
	rows, err := r.db.Query(*ctx, q, brandID, text)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	models := make([]model.Model, 0)

	for rows.Next() {
		var model model.Model

		if err := rows.Scan(&model.ID, &model.Name); err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, err
}

func (r *UserRepository) GetGenerationsByModelID(ctx *context.Context, modelID int64) ([]model.Generation, error) {
	q := `
		SELECT id, name, image, start_year, end_year FROM generations where model_id = $1;
	`
	rows, err := r.db.Query(*ctx, q, modelID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	generations := make([]model.Generation, 0)

	for rows.Next() {
		var generation model.Generation
		if err := rows.Scan(&generation.ID, &generation.Name, &generation.Image, &generation.StartYear, &generation.EndYear); err != nil {
			return nil, err
		}
		generations = append(generations, generation)
	}
	return generations, err
}

func (r *UserRepository) GetBodyTypes(ctx *context.Context) ([]model.BodyType, error) {
	q := `
		SELECT id, name FROM body_types
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	bodyTypes := make([]model.BodyType, 0)

	for rows.Next() {
		var bodyType model.BodyType
		if err := rows.Scan(&bodyType.ID, &bodyType.Name); err != nil {
			return nil, err
		}
		bodyTypes = append(bodyTypes, bodyType)
	}
	return bodyTypes, err
}

func (r *UserRepository) GetTransmissions(ctx *context.Context) ([]model.Transmission, error) {
	q := `
		SELECT id, name FROM transmissions
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	transmissions := make([]model.Transmission, 0)

	for rows.Next() {
		var transmission model.Transmission
		if err := rows.Scan(&transmission.ID, &transmission.Name); err != nil {
			return nil, err
		}
		transmissions = append(transmissions, transmission)
	}
	return transmissions, err
}

func (r *UserRepository) GetEngines(ctx *context.Context) ([]model.Engine, error) {
	q := `
		SELECT id, value FROM engines
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	engines := make([]model.Engine, 0)

	for rows.Next() {
		var engine model.Engine
		if err := rows.Scan(&engine.ID, &engine.Value); err != nil {
			return nil, err
		}
		engines = append(engines, engine)
	}
	return engines, err
}

func (r *UserRepository) GetDrivetrains(ctx *context.Context) ([]model.Drivetrain, error) {
	q := `
		SELECT id, name FROM drivetrains
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	drivetrains := make([]model.Drivetrain, 0)

	for rows.Next() {
		var drivetrain model.Drivetrain
		if err := rows.Scan(&drivetrain.ID, &drivetrain.Name); err != nil {
			return nil, err
		}
		drivetrains = append(drivetrains, drivetrain)
	}
	return drivetrains, err
}

func (r *UserRepository) GetFuelTypes(ctx *context.Context) ([]model.FuelType, error) {
	q := `
		SELECT id, name FROM fuel_types
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	fuelTypes := make([]model.FuelType, 0)

	for rows.Next() {
		var fuelType model.FuelType
		if err := rows.Scan(&fuelType.ID, &fuelType.Name); err != nil {
			return nil, err
		}
		fuelTypes = append(fuelTypes, fuelType)
	}
	return fuelTypes, err
}

func (r *UserRepository) GetColors(ctx *context.Context) ([]model.Color, error) {
	q := `
		SELECT id, name, hex_code FROM colors
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	colors := make([]model.Color, 0)

	for rows.Next() {
		var color model.Color

		if err := rows.Scan(&color.ID, &color.Name, &color.HexCode); err != nil {
			return nil, err
		}
		colors = append(colors, color)
	}
	return colors, err
}

func (r *UserRepository) GetCars(ctx *context.Context) ([]model.GetCarsResponse, error) {
	cars := make([]model.GetCarsResponse, 0)

	q := `
		
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			icls.name as interior_color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.mileage_km,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.credit_price,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_number
		from vehicles vs
		left join colors icls on icls.id = vs.interior_color_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on vs.transmission_id = ts.id
		left join engines es on vs.engine_id = es.id
		left join drivetrains ds on vs.drivetrain_id = ds.id
		left join body_types bts on vs.body_type_id = bts.id
		left join fuel_types fts on vs.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true

	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return cars, err
	}

	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.InteriorColor, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.CreditPrice, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumber,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *UserRepository) GetCarByID(ctx *context.Context, carID int) (model.GetCarsResponse, error) {
	car := model.GetCarsResponse{}

	q := `
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			icls.name as interior_color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.mileage_km,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.credit_price,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_number
		from vehicles vs
		left join colors icls on icls.id = vs.interior_color_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on vs.transmission_id = ts.id
		left join engines es on vs.engine_id = es.id
		left join drivetrains ds on vs.drivetrain_id = ds.id
		left join body_types bts on vs.body_type_id = bts.id
		left join fuel_types fts on vs.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		where vs.id = $1;

	`

	err := r.db.QueryRow(*ctx, q, carID).Scan(
		&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.InteriorColor, &car.Model, &car.Transmission, &car.Engine,
		&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
		&car.Exchange, &car.Credit, &car.New, &car.CreditPrice, &car.Status, &car.CreatedAt,
		&car.UpdatedAt, &car.Images, &car.PhoneNumber,
	)

	return car, err
}

func (r *UserRepository) CreateCar(ctx *context.Context, car *model.CreateCarRequest) (int, error) {

	keys, values, args := pkg.BuildParams(car)

	q := `
		INSERT INTO vehicles 
			(
				` + strings.Join(keys, ", ") + `
			) VALUES (
				` + strings.Join(values, ", ") + `
			) RETURNING id
	`
	var id int
	err := r.db.QueryRow(*ctx, q, args...).Scan(&id)

	return id, err
}

func (r *UserRepository) CreateCarImages(ctx *context.Context, carID int, images []string) error {
	if len(images) == 0 {
		return nil
	}

	q := `
		INSERT INTO images (vehicle_id, image) VALUES ($1, $2)
	`

	_, err := r.db.Exec(*ctx, q, carID, images[0])
	return err
}
