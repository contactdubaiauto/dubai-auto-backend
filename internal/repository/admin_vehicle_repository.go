package repository

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
	"dubai-auto/pkg/auth"
)

// Vehicles (admin) repository methods.

func (r *AdminRepository) GetVehicles(ctx *fasthttp.RequestCtx, limit, lastID int) ([]model.AdminVehicleListItem, error) {
	vehicles := make([]model.AdminVehicleListItem, 0)

	q := `
		SELECT
			vs.id,
			bs.name as brand,
			ms.name as model,
			vs.description,
			vs.price,
			vs.status,
			u.phone as user_phone,
			pf.username as user_name,
			CASE
				WHEN pf.avatar IS NULL OR pf.avatar = '' THEN NULL
				ELSE $1 || pf.avatar
			END AS user_avatar
		FROM vehicles vs
		LEFT JOIN brands bs ON vs.brand_id = bs.id
		LEFT JOIN models ms ON vs.model_id = ms.id
		LEFT JOIN users u ON u.id = vs.user_id
		LEFT JOIN profiles pf ON pf.user_id = vs.user_id
		WHERE vs.id > $2
		ORDER BY vs.id DESC
		LIMIT $3
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL, lastID, limit)
	if err != nil {
		return vehicles, err
	}
	defer rows.Close()

	for rows.Next() {
		var v model.AdminVehicleListItem
		if err := rows.Scan(
			&v.ID,
			&v.Brand,
			&v.Model,
			&v.Description,
			&v.Price,
			&v.Status,
			&v.UserPhone,
			&v.UserName,
			&v.UserAvatar,
		); err != nil {
			return vehicles, err
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, err
}

func (r *AdminRepository) GetVehicleByID(ctx *fasthttp.RequestCtx, vehicleID int) (model.GetCarsResponse, error) {
	car := model.GetCarsResponse{}

	// Similar to user GetCarByID, but does NOT increment view_count and does not depend on "my_car" / "liked".
	q := `
		SELECT
			vs.id,
			bs.name as brand,
			rs.name as region,
			cs.name as city,
			cls.name as color,
			ms.name as model,
			ts.name as transmission,
			es.name as engine,
			ds.name as drive,
			bts.name as body_type,
			fts.name as fuel_type,
			vs.year,
			vs.price,
			vs.odometer,
			vs.vin_code,
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.trade_in,
			vs.owners,
			vs.updated_at,
			images.images,
			videos.videos,
			vs.phone_numbers,
			vs.view_count,
			false AS my_car,
			json_build_object(
				'id', pf.user_id,
				'username', pf.username,
				'avatar', $2 || pf.avatar,
				'contacts', pf.contacts
			) as owner,
			vs.description,
			false AS liked
		FROM vehicles vs
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
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT '` + r.config.IMAGE_BASE_URL + `' || image as image
				FROM images
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT '` + r.config.IMAGE_BASE_URL + `' || video as video
				FROM videos
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE vs.id = $1;
	`

	err := r.db.QueryRow(ctx, q, vehicleID, r.config.IMAGE_BASE_URL).Scan(
		&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
		&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
		&car.Credit, &car.New, &car.Status, &car.CreatedAt, &car.TradeIn, &car.Owners,
		&car.UpdatedAt, &car.Images, &car.Videos, &car.PhoneNumbers, &car.ViewCount, &car.MyCar, &car.Owner, &car.Description, &car.Liked,
	)
	return car, err
}

func (r *AdminRepository) CreateVehicle(ctx *fasthttp.RequestCtx, req *model.AdminCreateVehicleRequest) (int, error) {
	keys, values, args := auth.BuildParams(req)
	if len(keys) == 0 {
		return 0, fmt.Errorf("invalid request data")
	}

	q := `
		INSERT INTO vehicles (` + strings.Join(keys, ", ") + `)
		VALUES (` + strings.Join(values, ", ") + `)
		RETURNING id
	`

	var id int
	err := r.db.QueryRow(ctx, q, args...).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateVehicleStatus(ctx *fasthttp.RequestCtx, vehicleID int, status int) error {
	q := `
		UPDATE vehicles
		SET status = $2, updated_at = now()
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, q, vehicleID, status)
	return err
}

func (r *AdminRepository) DeleteVehicle(ctx *fasthttp.RequestCtx, vehicleID int) error {
	q := `delete from vehicles where id = $1`
	_, err := r.db.Exec(ctx, q, vehicleID)
	return err
}
