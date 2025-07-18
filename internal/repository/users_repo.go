package repository

import (
	"context"
	"fmt"
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

func (r *UserRepository) GetMyCars(ctx *context.Context, userID *int) ([]model.GetCarsResponse, error) {
	cars := make([]model.GetCarsResponse, 0)
	q := `
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_numbers, 
			vs.view_count,
			true as my_car
		from vehicles vs
		left join generation_modifications gms on gms.id = vs.modification_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on gms.transmission_id = ts.id
		left join engines es on gms.engine_id = es.id
		left join drivetrains ds on gms.drivetrain_id = ds.id
		left join body_types bts on gms.body_type_id = bts.id
		left join fuel_types fts on gms.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		where vs.user_id = $1 and status = 2
		order by vs.id desc
	`
	rows, err := r.db.Query(*ctx, q, *userID)

	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumbers, &car.ViewCount, &car.MyCar,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}
	return cars, err
}

func (r *UserRepository) OnSale(ctx *context.Context, userID *int) ([]model.GetCarsResponse, error) {
	cars := make([]model.GetCarsResponse, 0)
	q := `
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_numbers, 
			vs.view_count, 
			true as my_car
		from vehicles vs
		left join generation_modifications gms on gms.id = vs.modification_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on gms.transmission_id = ts.id
		left join engines es on gms.engine_id = es.id
		left join drivetrains ds on gms.drivetrain_id = ds.id
		left join body_types bts on gms.body_type_id = bts.id
		left join fuel_types fts on gms.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		where vs.user_id = $1 and status = 3
		order by vs.id desc
	`
	rows, err := r.db.Query(*ctx, q, *userID)

	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumbers, &car.ViewCount, &car.MyCar,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}
	return cars, err
}

func (r *UserRepository) Cancel(ctx *context.Context, carID *int) error {
	q := `
		delete from vehicles where id = $1
	`
	_, err := r.db.Exec(*ctx, q, *carID)
	return err
}

func (r *UserRepository) DeleteCar(ctx *context.Context, carID *int) error {
	q := `
		delete from vehicles where id = $1
	`
	_, err := r.db.Exec(*ctx, q, *carID)
	return err
}

func (r *UserRepository) DontSell(ctx *context.Context, carID, userID *int) error {
	q := `
		update vehicles 
			set status = 2 -- 2 is not sale
		where id = $1 and status = 3 -- 3 is on sale
			and user_id = $2
	`

	_, err := r.db.Exec(*ctx, q, *carID, *userID)
	return err
}

func (r *UserRepository) Sell(ctx *context.Context, carID, userID *int) error {
	q := `
		update vehicles 
			set status = 3 -- 3 is on sale
		where id = $1 and status = 2 -- 2 is not sale 
			and user_id = $2
	`
	_, err := r.db.Exec(*ctx, q, *carID, *userID)
	return err
}

func (r *UserRepository) GetBrands(ctx *context.Context, text string) ([]*model.GetBrandsResponse, error) {
	q := `
		SELECT id, name, logo, model_count FROM brands WHERE name ILIKE '%' || $1 || '%'
	`
	rows, err := r.db.Query(*ctx, q, text)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var brands = make([]*model.GetBrandsResponse, 0)

	for rows.Next() {
		var brand model.GetBrandsResponse
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.Logo, &brand.ModelCount); err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}
	return brands, err
}

func (r *UserRepository) GetProfile(ctx *context.Context, userID int) (model.GetProfileResponse, error) {
	q := `
		select 
			us.id,
			us.email,
			us.phone,
			ps.driving_experience,
			ps.notification,
			ps.username,
			ps.google,
			ps.birthday,
			ps.about_me,
			ps.registered_by
		from users us
		left join profiles as ps on ps.user_id = us.id
		where us.id = $1;

	`
	var pf model.GetProfileResponse
	err := r.db.QueryRow(*ctx, q, userID).Scan(&pf.ID, &pf.Email, &pf.Phone,
		&pf.DrivingExperience, &pf.Notification, &pf.Username, &pf.Google, &pf.Birthday, &pf.AboutMe, &pf.RegisteredBy)

	return pf, err
}

func (r *UserRepository) UpdateProfile(ctx *context.Context, userID int, profile *model.UpdateProfileRequest) error {
	// Parse birthday string to time.Time if it's not empty
	// var birthdayTime *time.Time
	// if profile.Birthday != "" {
	// 	parsedTime, err := time.Parse("2006-01-02", profile.Birthday)
	// 	if err != nil {
	// 		return fmt.Errorf("invalid birthday format: %v", err)
	// 	}
	// 	birthdayTime = &parsedTime
	// }

	q := `
		UPDATE profiles 
		SET 
			driving_experience = $1,
			notification = $2,
			username = $3,
			google = $4,
			birthday = $5,
			about_me = $6,
			last_active_date = NOW()
		WHERE user_id = $7
	`
	_, err := r.db.Exec(*ctx, q,
		profile.DrivingExperience,
		profile.Notification,
		profile.Username,
		profile.Google,
		profile.Birthday,
		profile.AboutMe,
		userID)

	return err
}

func (r *UserRepository) GetFilterBrands(ctx *context.Context, text string) (model.GetFilterBrandsResponse, error) {
	var brand model.GetFilterBrandsResponse
	q := `
		with popular as (
			SELECT 
				json_agg(
					json_build_object(
						'id', id, 
						'name', name, 
						'logo', logo, 
						'model_count', model_count 
					)
				) as popular_brands
			FROM brands 
			WHERE name ILIKE '%' || $1 || '%' and popular = true
		), all_brands as (
			SELECT 
				json_agg(
					json_build_object(
						'id', id, 
						'name', name, 
						'logo', logo, 
						'model_count', model_count 
					)
				) as all_brands
			FROM brands 
			WHERE name ILIKE '%' || $1 || '%'
		)
		select 
			pp.popular_brands,
			ab.all_brands
		from popular as pp
		left join all_brands as ab on true;

	`
	err := r.db.QueryRow(*ctx, q, text).Scan(&brand.PopularBrands, &brand.AllBrands)

	return brand, err
}

func (r *UserRepository) GetCities(ctx *context.Context, text string) ([]model.GetCitiesResponse, error) {
	var cities = make([]model.GetCitiesResponse, 0)
	var city model.GetCitiesResponse
	q := `
		select 
			c.id, 
			c.name,
			json_agg(
				json_build_object(
					'id', r.id,
					'name', r.name
				)
			) as regions
		from cities c
		left join regions r on r.city_id = c.id
		where c.name ILIKE '%' || $1 || '%'
		group by c.id, c.name;
	`
	rows, err := r.db.Query(*ctx, q, text)

	if err != nil {
		return cities, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&city.ID, &city.Name, &city.Regions)

		if err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}
	return cities, err
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

func (r *UserRepository) GetFilterModelsByBrandID(ctx *context.Context, brandID int64, text string) (model.GetFilterModelsResponse, error) {
	responseModel := model.GetFilterModelsResponse{}
	q := `
		with popular as (
			SELECT 
				json_agg(
					json_build_object(
						'id', id, 
						'name', name, 
						'model_count', model_count 
					)
				) as popular_models
			FROM models 
			models WHERE brand_id = $1 AND name ILIKE '%' || $2 || '%' and popular = true
		), all_models as (
			SELECT 
				json_agg(
					json_build_object(
						'id', id, 
						'name', name, 
						'model_count', model_count 
					)
				) as all_models
			FROM models 
			models WHERE brand_id = $1 AND name ILIKE '%' || $2 || '%'
		)
		select 
			pp.popular_models,
			ms.all_models
		from popular as pp
		left join all_models as ms on true;
		`
	err := r.db.QueryRow(*ctx, q, brandID, text).Scan(&responseModel.PopularModels, &responseModel.AllModels)

	if err != nil {
		return responseModel, err
	}

	return responseModel, err
}

func (r *UserRepository) GetGenerationsByModelID(ctx *context.Context, modelID int, wheel bool, year, bodyTypeID string) ([]model.Generation, error) {
	q := `
		with gms as (
			select 
				json_agg(
					json_build_object(
						'id', gms.id,
						'engine', es.value, 
						'fuel_type', fts.name, 
						'drivetrain', ds.name, 
						'transmission', ts.name
					)
				) as modifications,
				gms.generation_id
			from generation_modifications gms
			left join engines es on es.id = gms.engine_id
			left join fuel_types fts on fts.id = gms.fuel_type_id
			left join drivetrains ds on ds.id = gms.drivetrain_id
			left join transmissions ts on ts.id = gms.transmission_id
			where gms.generation_id = any (
				select 
					id 
				from generations 
				where 
					model_id = $1 and start_year <= $2 and end_year >= $2
					and body_type_id = $3 and wheel = $4
			)
			group by gms.generation_id 
		)
		select
			gs.id,
			gs.name,
			gs.image,
			gs.start_year,
			gs.end_year,
			gms.modifications
		from gms
		left join generations gs on gs.id = gms.generation_id;
	`
	rows, err := r.db.Query(*ctx, q, modelID, year, bodyTypeID, wheel)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	generations := make([]model.Generation, 0)

	for rows.Next() {
		var generation model.Generation
		if err = rows.Scan(&generation.ID, &generation.Name,
			&generation.Image, &generation.StartYear, &generation.EndYear,
			&generation.Modifications,
		); err != nil {
			return nil, err
		}
		generations = append(generations, generation)
	}
	return generations, err
}

// todo: after remove the array response, return an object
func (r *UserRepository) GetYearsByModelID(ctx *context.Context, modelID int64, wheel bool) ([]*int, error) {
	q := `
		SELECT 
			array_agg(y ORDER BY y) AS years
		FROM (
			SELECT DISTINCT generate_series(start_year, end_year) AS y
			FROM generations
			WHERE model_id = $1 AND wheel = $2
		) AS years_series;
	`
	var years []*int
	err := r.db.QueryRow(*ctx, q, modelID, wheel).Scan(&years)

	return years, err
}

func (r *UserRepository) GetBodysByModelID(ctx *context.Context, modelID int, wheel bool, year string) ([]model.BodyType, error) {
	q := `
		select DISTINCT ON (bts.id)
			bts.id,
			bts.name,
			bts.image
		from generation_modifications gms
		left join body_types bts on bts.id = gms.body_type_id
		where gms.generation_id in (
			select 
				gs.id 
			from generations gs 
			WHERE gs.start_year <= $1 AND gs.end_year >= $1 
					and gs.model_id = $2 and gs.wheel = $3
		)
	`

	rows, err := r.db.Query(*ctx, q, year, modelID, wheel)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	bodyTypes := make([]model.BodyType, 0)

	for rows.Next() {
		var bodyType model.BodyType

		if err := rows.Scan(&bodyType.ID, &bodyType.Name, &bodyType.Image); err != nil {
			return nil, err
		}

		bodyTypes = append(bodyTypes, bodyType)
	}

	return bodyTypes, err
}

func (r *UserRepository) GetBodyTypes(ctx *context.Context) ([]model.BodyType, error) {
	q := `
		SELECT id, name, image FROM body_types
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	bodyTypes := make([]model.BodyType, 0)

	for rows.Next() {
		var bodyType model.BodyType

		if err := rows.Scan(&bodyType.ID, &bodyType.Name, &bodyType.Image); err != nil {
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
		SELECT id, name, image FROM colors
	`

	rows, err := r.db.Query(*ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	colors := make([]model.Color, 0)

	for rows.Next() {
		var color model.Color

		if err := rows.Scan(&color.ID, &color.Name, &color.Image); err != nil {
			return nil, err
		}
		colors = append(colors, color)
	}
	return colors, err
}

func (r *UserRepository) GetHome(ctx *context.Context, userID int) (model.Home, error) {
	home := model.Home{}
	cars := make([]model.GetCarsResponse, 0)

	q := `
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			vs.phone_numbers,
			vs.view_count,
			CASE
				WHEN vs.user_id = $1 THEN TRUE
				ELSE FALSE
			END AS my_car
		from vehicles vs
		left join generation_modifications gms on gms.id = vs.modification_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on gms.transmission_id = ts.id
		left join engines es on gms.engine_id = es.id
		left join drivetrains ds on gms.drivetrain_id = ds.id
		left join body_types bts on gms.body_type_id = bts.id
		left join fuel_types fts on gms.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		where vs.status = 3
		order by vs.id desc limit 4
	`

	rows, err := r.db.Query(*ctx, q, userID)

	if err != nil {
		return home, err
	}

	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumbers, &car.ViewCount, &car.MyCar,
		); err != nil {
			return home, err
		}
		cars = append(cars, car)
	}
	home.Popular = cars
	return home, nil
}

func (r *UserRepository) GetCars(ctx *context.Context, userID int,
	brands, models, regions, cities, generations, transmissions,
	engines, drivetrains, body_types, fuel_types, ownership_types []string, year_from, year_to, exchange, credit,
	price_from, price_to string) ([]model.GetCarsResponse, error) {
	var qWhere string
	var qValues []interface{}
	qValues = append(qValues, userID)
	var i = 1

	if len(brands) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND bs.id = ANY($%d)", i)
		qValues = append(qValues, brands)
	}
	if len(models) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND ms.id = ANY($%d)", i)
		qValues = append(qValues, models)
	}
	if len(regions) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND rs.id = ANY($%d)", i)
		qValues = append(qValues, regions)
	}
	if len(cities) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND cs.id = ANY($%d)", i)
		qValues = append(qValues, cities)
	}
	if len(transmissions) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND ts.id = ANY($%d)", i)
		qValues = append(qValues, transmissions)
	}
	if len(engines) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND es.id = ANY($%d)", i)
		qValues = append(qValues, engines)
	}
	if len(drivetrains) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND ds.id = ANY($%d)", i)
		qValues = append(qValues, drivetrains)
	}
	if len(body_types) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND bts.id = ANY($%d)", i)
		qValues = append(qValues, body_types)
	}
	if len(fuel_types) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND fts.id = ANY($%d)", i)
		qValues = append(qValues, fuel_types)
	}
	if len(ownership_types) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.ownership_type_id = ANY($%d) ", i)
		qValues = append(qValues, ownership_types)
	}
	if year_from != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.year >= $%d", i)
		qValues = append(qValues, year_from)
	}
	if year_to != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.year <= $%d", i)
		qValues = append(qValues, year_to)
	}
	if exchange != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.exchange = $%d", i)
		qValues = append(qValues, true)
	}
	if credit != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.credit = $%d", i)
		qValues = append(qValues, true)
	}
	if price_from != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.price >= $%d", i)
		qValues = append(qValues, price_from)
	}
	if price_to != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.price <= $%d", i)
		qValues = append(qValues, price_to)
	}

	cars := make([]model.GetCarsResponse, 0)
	q := `
		select 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			videos,
			vs.phone_numbers,
			vs.view_count,
			CASE
				WHEN vs.user_id = $1 THEN TRUE
				ELSE FALSE
			END AS my_car
		from vehicles vs
		left join generation_modifications gms on gms.id = vs.modification_id
		left join colors cls on vs.color_id = cls.id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join transmissions ts on gms.transmission_id = ts.id
		left join engines es on gms.engine_id = es.id
		left join drivetrains ds on gms.drivetrain_id = ds.id
		left join body_types bts on gms.body_type_id = bts.id
		left join fuel_types fts on gms.fuel_type_id = fts.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		left join lateral (
			select 
				json_agg(video) as videos
			from videos 
			where vehicle_id = vs.id
		) videos on true
		where vs.status = 3
		` + qWhere + `
		order by vs.id desc
	`

	rows, err := r.db.Query(*ctx, q, qValues...)

	if err != nil {
		return cars, err
	}

	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.Videos, &car.PhoneNumbers, &car.ViewCount, &car.MyCar,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *UserRepository) GetCarByID(ctx *context.Context, carID, userID int) (model.GetCarsResponse, error) {
	car := model.GetCarsResponse{}
	fmt.Println("sd8f9j98j")
	q := `
		WITH updated AS (
			UPDATE vehicles
			SET view_count = view_count + 1
			WHERE id = $1
			RETURNING *
		)
		SELECT 
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			ms.name as model,
			ts.name as transmission,
			es.value as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.exchange,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.updated_at,
			images,
			videos,
			vs.phone_numbers,
			vs.view_count,
			CASE
				WHEN vs.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_car,
			json_build_object(
				'id', pf.user_id,
				'username', pf.username,
				'avatar', pf.avatar
			) as owner,
			 vs.description
		FROM updated vs
		LEFT JOIN generation_modifications gms ON gms.id = vs.modification_id
		LEFT JOIN profiles pf on pf.user_id = vs.user_id
		LEFT JOIN colors cls ON vs.color_id = cls.id
		LEFT JOIN brands bs ON vs.brand_id = bs.id
		LEFT JOIN regions rs ON vs.region_id = rs.id
		LEFT JOIN cities cs ON vs.city_id = cs.id
		LEFT JOIN models ms ON vs.model_id = ms.id
		LEFT JOIN transmissions ts ON gms.transmission_id = ts.id
		LEFT JOIN engines es ON gms.engine_id = es.id
		LEFT JOIN drivetrains ds ON gms.drivetrain_id = ds.id
		LEFT JOIN body_types bts ON gms.body_type_id = bts.id
		LEFT JOIN fuel_types fts ON gms.fuel_type_id = fts.id
		LEFT JOIN LATERAL (
			SELECT json_agg(image) AS images
			FROM images
			WHERE vehicle_id = vs.id
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(video) AS videos
			FROM videos
			WHERE vehicle_id = vs.id
		) videos ON true
		WHERE vs.id = $1;
	`

	err := r.db.QueryRow(*ctx, q, carID, userID).Scan(
		&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
		&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
		&car.Exchange, &car.Credit, &car.New, &car.Status, &car.CreatedAt,
		&car.UpdatedAt, &car.Images, &car.Videos, &car.PhoneNumbers, &car.ViewCount, &car.MyCar, &car.Owner, &car.Description,
	)

	return car, err
}

func (r *UserRepository) GetEditCarByID(ctx *context.Context, carID, userID int) (model.GetEditCarsResponse, error) {
	car := model.GetEditCarsResponse{}

	q := `
		select 
			vs.id,
			json_build_object(
				'id', bs.id,
				'name', bs.name,
				'logo', bs.logo,
				'model_count', bs.model_count
			) as brand,
			json_build_object(
				'id', rs.id,
				'name', rs.name
			) as region,
			json_build_object(
				'id', cs.id,
				'name', cs.name
			) as city,
			json_build_object(
				'id', ms.id,
				'name', ms.name
			) as model,
			json_build_object(
				'id', mfs.id,
				'engine', es.value,
				'fuel_type', fts.name,
				'drivetrain', ds.name,
				'transmission', ts.name
			) as modification,
			json_build_object(
				'id', cls.id,
				'name', cls.name,
				'image', cls.image
			) as color,
			json_build_object(
				'id', bts.id,
				'name', bts.name,
				'image', bts.image
			) as body_type,
			json_build_object(
				'id', gs.id,
				'name', gs.name,
				'image', gs.image,
				'start_year', gs.start_year,
				'end_year', gs.end_year
			) as generation,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.exchange,
			vs.wheel,
			vs.trade_in,
			vs.crash,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			images,
			videos,
			vs.phone_numbers,
			vs.view_count,
			vs.description,
			CASE
				WHEN vs.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_car,
			vs.owners
		from vehicles vs
		left join colors cls on vs.color_id = cls.id
		left join generation_modifications mfs on mfs.id = vs.modification_id
		left join generations gs on gs.id = mfs.generation_id
		left join body_types bts on bts.id = mfs.body_type_id
		left join engines es on es.id = mfs.engine_id
		left join transmissions ts on ts.id = mfs.transmission_id
		left join drivetrains ds on ds.id = mfs.drivetrain_id
		left join fuel_types fts on fts.id = mfs.fuel_type_id
		left join brands bs on vs.brand_id = bs.id
		left join regions rs on vs.region_id = rs.id
		left join cities cs on vs.city_id = cs.id
		left join models ms on vs.model_id = ms.id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
		left join lateral (
			select 
				json_agg(video) as videos
			from videos 
			where vehicle_id = vs.id
		) videos on true
		where vs.id = $1;
	`
	err := r.db.QueryRow(*ctx, q, carID, userID).Scan(
		&car.ID, &car.Brand, &car.Region, &car.City, &car.Model, &car.Modification,
		&car.Color, &car.BodyType, &car.Generation, &car.Year, &car.Price, &car.Odometer, &car.VinCode,
		&car.Exchange, &car.Wheel, &car.TradeIN, &car.Crash,
		&car.Credit, &car.New, &car.Status, &car.CreatedAt, &car.Images, &car.Videos, &car.PhoneNumbers,
		&car.ViewCount, &car.Description, &car.MyCar, &car.Owners,
	)

	return car, err
}

func (r *UserRepository) BuyCar(ctx *context.Context, carID, userID int) error {

	q := `
		update vehicles 
			set status = 2,
				user_id = $1
		where id = $2 and status = 3 -- 3 is on sale
	`

	_, err := r.db.Exec(*ctx, q, userID, carID)

	return err
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

func (r *UserRepository) UpdateCar(ctx *context.Context, car *model.UpdateCarRequest, userID int) error {
	// First check if the car belongs to the user
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM vehicles WHERE id = $1 AND user_id = $2)`
	err := r.db.QueryRow(*ctx, checkQuery, car.ID, userID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("car not found or access denied")
	}

	keys, _, args := pkg.BuildParams(car)

	var updateFields []string
	var updateArgs []interface{}
	updateArgs = append(updateArgs, car.ID)

	paramIndex := 2
	for i, key := range keys {
		if key != "id" && key != "user_id" {
			updateFields = append(updateFields, fmt.Sprintf("%s = $%d", key, paramIndex))
			updateArgs = append(updateArgs, args[i])
			paramIndex++
		}
	}

	if len(updateFields) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	q := `
		UPDATE vehicles 
		SET ` + strings.Join(updateFields, ", ") + `, updated_at = NOW()
		WHERE id = $1 AND user_id = $` + fmt.Sprintf("%d", paramIndex)

	updateArgs = append(updateArgs, userID)

	_, err = r.db.Exec(*ctx, q, updateArgs...)
	return err
}

func (r *UserRepository) CreateCarImages(ctx *context.Context, carID int, images []string) error {

	if len(images) == 0 {
		return nil
	}

	q := `
		INSERT INTO images (vehicle_id, image) VALUES ($1, $2)
	`

	for i := range images {
		_, err := r.db.Exec(*ctx, q, carID, images[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepository) CreateCarVideos(ctx *context.Context, carID int, video string) error {

	q := `
		INSERT INTO videos (vehicle_id, video) VALUES ($1, $2)
	`

	_, err := r.db.Exec(*ctx, q, carID, video)
	if err != nil {
		return err
	}

	return err
}

func (r *UserRepository) DeleteCarImage(ctx *context.Context, carID int, imagePath string) error {
	q := `DELETE FROM images WHERE vehicle_id = $1 AND image = $2`
	_, err := r.db.Exec(*ctx, q, carID, imagePath)
	return err
}

func (r *UserRepository) DeleteCarVideo(ctx *context.Context, carID int, videoPath string) error {
	q := `DELETE FROM videos WHERE vehicle_id = $1 AND video = $2`
	_, err := r.db.Exec(*ctx, q, carID, videoPath)
	return err
}
