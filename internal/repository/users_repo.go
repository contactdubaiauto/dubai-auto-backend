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
			vs.phone_number, 
			vs.view_count
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
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.InteriorColor, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.CreditPrice, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumber, &car.ViewCount,
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
			vs.phone_number, 
			vs.view_count, 
			true as my_car
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
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.InteriorColor, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.CreditPrice, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumber, &car.ViewCount, &car.MyCar,
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

func (r *UserRepository) Delete(ctx *context.Context, carID *int) error {
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

func (r *UserRepository) GetBrands(ctx *context.Context, text string) (model.GetBrandsResponse, error) {
	var brand model.GetBrandsResponse
	q := `
		with popular as (
			SELECT 
				json_agg(
					json_build_object(
						'id', id, 
						'name', name, 
						'logo', logo, 
						'car_count', car_count 
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
						'car_count', car_count 
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
	rows, err := r.db.Query(*ctx, q, text)

	if err != nil {
		return brand, err
	}

	defer rows.Close()

	if err := rows.Scan(&brand.PopularBrands, &brand.AllBrands); err != nil {
		return brand, err
	}
	return brand, err
}

func (r *UserRepository) GetModifications(ctx *context.Context, generationID, bodyTypeID, fuelTypeID, drivetrainID, transmissionID int) ([]*model.GetModificationsResponse, error) {
	q := `
		SELECT id, name FROM generation_modifications 
		WHERE 
			generation_id = $1 and body_type_id = $2 and fuel_type_id = $3 and drivetrain_id = $4 and transmission_id = $5
	`
	rows, err := r.db.Query(*ctx, q, generationID, bodyTypeID, fuelTypeID, drivetrainID, transmissionID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var modifications = make([]*model.GetModificationsResponse, 0)

	for rows.Next() {
		var modification model.GetModificationsResponse
		if err := rows.Scan(&modification.ID, &modification.Name); err != nil {
			return nil, err
		}
		modifications = append(modifications, &modification)
	}
	return modifications, err
}

func (r *UserRepository) GetModelsByBrandID(ctx *context.Context, brandID int64, text string) (model.GetModelsResponse, error) {
	responseModel := model.GetModelsResponse{}
	q := `
		with popular as (
			SELECT 
				json_agg(
					json_build_object(
						'id', id, 
						'name', name, 
						'car_count', car_count 
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
						'car_count', car_count 
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

func (r *UserRepository) GetGenerationsByModelID(ctx *context.Context, modelID int64) ([]model.Generation, error) {
	q := `
		with fts as (
			select
				gft.generation_id,
				json_agg(
					json_build_object(
						'id', ft.id,
						'name', ft.name
					)
				) as fuel_types
			from generation_fuel_types gft
			left join fuel_types ft on gft.fuel_type_id = ft.id
			group by gft.generation_id
		), bts as (
			select
				gbts.generation_id,
				json_agg(
					json_build_object(
						'id', bts.id,
						'name', bts.name,
						'image', bts.image
					)
				) as body_types
			from generation_body_types gbts 
			left join body_types bts on gbts.body_type_id = bts.id
			group by gbts.generation_id
		), dts as (
			select
				gds.generation_id,
				json_agg(
					json_build_object(
						'id', ds.id,
						'name', ds.name
					)
				) as drivetrains
			from generation_drivetrains gds
			left join drivetrains ds on gds.drivetrain_id = ds.id
			group by gds.generation_id
		), ts as (
			select
				gts.generation_id,
				json_agg(
					json_build_object(
						'id', t.id,
						'name', t.name
					)
				) as transmissions
			from generation_transmissions gts
			left join transmissions t on gts.transmission_id = t.id
			group by gts.generation_id
		)

		SELECT 
			gs.id, 
			gs.name, 
			gs.image, 
			gs.start_year, 
			gs.end_year,
			fts.fuel_types,
			bts.body_types,
			dts.drivetrains,
			ts.transmissions
		FROM generations gs
		left join fts on gs.id = fts.generation_id
		left join bts on gs.id = bts.generation_id
		left join dts on gs.id = dts.generation_id
		left join ts on gs.id = ts.generation_id
		WHERE gs.model_id = $1;
	`
	rows, err := r.db.Query(*ctx, q, modelID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	generations := make([]model.Generation, 0)

	for rows.Next() {
		var generation model.Generation
		if err := rows.Scan(&generation.ID, &generation.Name,
			&generation.Image, &generation.StartYear, &generation.EndYear,
			&generation.FuelTypes, &generation.BodyTypes, &generation.Drivetrains, &generation.Transmissions,
		); err != nil {
			return nil, err
		}
		generations = append(generations, generation)
	}
	return generations, err
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

func (r *UserRepository) GetCars(ctx *context.Context, userID int,
	brands, models, regions, cities, generations, transmissions,
	engines, drivetrains, body_types, fuel_types, ownership_types []string, year_from, year_to, exchange, credit,
	right_hand_drive, price_from, price_to string) ([]model.GetCarsResponse, error) {
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
	if len(generations) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND gs.id = ANY($%d)", i)
		qValues = append(qValues, generations)
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
	if right_hand_drive != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.right_hand_drive = $%d", i)
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

	fmt.Println("qWhere")
	fmt.Println(qWhere)
	fmt.Println(qValues...)
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
			vs.phone_number,
			CASE
				WHEN vs.user_id = $1 THEN TRUE
				ELSE FALSE
			END AS my_car
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
		left join generations gs on gs.id = vs.generation_id
		left join lateral (
			select 
				json_agg(image) as images
			from images 
			where vehicle_id = vs.id
		) images on true
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
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.InteriorColor, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Exchange, &car.Credit, &car.New, &car.CreditPrice, &car.Status, &car.CreatedAt,
			&car.UpdatedAt, &car.Images, &car.PhoneNumber, &car.MyCar,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *UserRepository) GetCarByID(ctx *context.Context, carID, userID int) (model.GetCarsResponse, error) {
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
			vs.phone_number,
			CASE
				WHEN vs.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_car
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
		where vs.id = $1

	`

	err := r.db.QueryRow(*ctx, q, carID, userID).Scan(
		&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.InteriorColor, &car.Model, &car.Transmission, &car.Engine,
		&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
		&car.Exchange, &car.Credit, &car.New, &car.CreditPrice, &car.Status, &car.CreatedAt,
		&car.UpdatedAt, &car.Images, &car.PhoneNumber, &car.MyCar,
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
