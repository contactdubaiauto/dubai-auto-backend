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

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)

	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return &user, err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepository) GetBrands(ctx context.Context, text string) ([]*model.GetBrandsResponse, error) {
	q := `
		SELECT id, name, logo, car_count FROM brands WHERE name ILIKE '%' || $1 || '%'
	`
	rows, err := r.db.Query(ctx, q, text)

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

func (r *UserRepository) GetModelsByBrandID(ctx context.Context, brandID int64, text string) ([]model.Model, error) {
	q := `
			SELECT id, name FROM models WHERE brand_id = $1 AND name ILIKE '%' || $2 || '%'
		`
	rows, err := r.db.Query(ctx, q, brandID, text)

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

func (r *UserRepository) GetBodyTypes(ctx context.Context) ([]model.BodyType, error) {
	q := `
		SELECT id, name FROM body_types
	`

	rows, err := r.db.Query(ctx, q)

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

func (r *UserRepository) GetTransmissions(ctx context.Context) ([]model.Transmission, error) {
	q := `
		SELECT id, name FROM transmissions
	`

	rows, err := r.db.Query(ctx, q)

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

func (r *UserRepository) GetEngines(ctx context.Context) ([]model.Engine, error) {
	q := `
		SELECT id, value FROM engines
	`

	rows, err := r.db.Query(ctx, q)

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

func (r *UserRepository) GetDrives(ctx context.Context) ([]model.Drive, error) {
	q := `
		SELECT id, name FROM drives
	`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	drives := make([]model.Drive, 0)

	for rows.Next() {
		var drive model.Drive
		if err := rows.Scan(&drive.ID, &drive.Name); err != nil {
			return nil, err
		}
		drives = append(drives, drive)
	}
	return drives, err
}

func (r *UserRepository) GetFuelTypes(ctx context.Context) ([]model.FuelType, error) {
	q := `
		SELECT id, name FROM fuel_types
	`

	rows, err := r.db.Query(ctx, q)

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

func (r *UserRepository) GetCars(ctx context.Context) ([]model.GetCarsResponse, error) {
	q := `
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.mileage,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.color,
			vs.credit_price,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_number
		from vehicles vs
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on vs.transmission_id = ts.id
		left join engines es on vs.engine_id = es.id
		left join drives ds on vs.drive_id = ds.id
		left join body_types bts on vs.body_type_id = bts.id
		left join fuel_types fts on vs.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true;

	`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	cars := make([]model.GetCarsResponse, 0)

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Model, &car.Transmission, &car.Engine,
			&car.Drive, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.Color, &car.CreditPrice, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumber,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *UserRepository) CreateCar(ctx context.Context, car *model.CreateCarRequest) (int, error) {

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
	err := r.db.QueryRow(ctx, q, args...).Scan(&id)

	return id, err
}
