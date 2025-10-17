package repository

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"dubai-auto/internal/model"
	"dubai-auto/pkg/auth"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
)

type ThirdPartyRepository struct {
	db *pgxpool.Pool
}

func NewThirdPartyRepository(db *pgxpool.Pool) *ThirdPartyRepository {
	return &ThirdPartyRepository{db}
}

func (r *ThirdPartyRepository) Profile(ctx *fasthttp.RequestCtx, id int, profile model.ThirdPartyProfileReq) model.Response {
	q := `
		update profiles set
			about_me = $1,
			whatsapp = $2,
			telegram = $3,
			address = $4,
			coordinates = $5,
			message = $6

		where user_id = $7
	`
	_, err := r.db.Exec(ctx, q, profile.AboutUs, profile.Whatsapp,
		profile.Telegram, profile.Address, profile.Coordinates, profile.Message, id)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "Profile updated successfully"}}
}

func (r *ThirdPartyRepository) FirstLogin(ctx *fasthttp.RequestCtx, id int, profile model.ThirdPartyFirstLoginReq) model.Response {
	q := `
		update profiles set
			message = $1
		where user_id = $2
	`
	_, err := r.db.Exec(ctx, q, profile.Message, id)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "First login updated successfully"}}
}

func (r *ThirdPartyRepository) GetProfile(ctx *fasthttp.RequestCtx, id int) model.Response {
	q := `
		select
			about_me,
			whatsapp,
			telegram,
			address,
			coordinates,
			avatar,
			banner,
			company_name,
			message,
			vat_number,
			company_types.name as company_type,
			activity_fields.name as activity_field,
			profiles.created_at
		from profiles 
		left join company_types on company_types.id = profiles.company_type_id
		left join activity_fields on activity_fields.id = profiles.activity_field_id
		where user_id = $1
	`
	var profile model.ThirdPartyGetProfileRes
	err := r.db.QueryRow(ctx, q, id).Scan(
		&profile.AboutUs, &profile.Whatsapp,
		&profile.Telegram, &profile.Address,
		&profile.Coordinates, &profile.Avatar,
		&profile.Banner,
		&profile.CompanyName, &profile.Message,
		&profile.VATNumber, &profile.CompanyType,
		&profile.ActivityField,
		&profile.Registered,
	)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	q = `
		select 
			email,
			phone
		from users 
		where id = $1
	`
	err = r.db.QueryRow(ctx, q, id).Scan(&profile.Email, &profile.Phone)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: profile}
}

func (r *ThirdPartyRepository) GetMyCars(ctx *fasthttp.RequestCtx, userID int) ([]model.GetCarsResponse, error) {
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
			true as my_car,
			vs.description,
			CASE 
				WHEN ul.vehicle_id IS NOT NULL THEN true
				ELSE false
			END AS liked
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
		left join user_likes ul on ul.vehicle_id = vs.id AND ul.user_id = $1
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT image
				FROM images
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT video
				FROM videos
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		where vs.user_id = $1 and status = 2
		order by vs.id desc
	`
	rows, err := r.db.Query(ctx, q, userID)

	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Credit, &car.New, &car.Status, &car.CreatedAt, &car.TradeIn, &car.Owners,
			&car.UpdatedAt, &car.Images, &car.Videos, &car.PhoneNumbers, &car.ViewCount, &car.MyCar, &car.Description, &car.Liked,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}
	return cars, err
}

func (r *ThirdPartyRepository) OnSale(ctx *fasthttp.RequestCtx, userID int) ([]model.GetCarsResponse, error) {
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
			vs.credit,
			vs.new,
			vs.status,
			vs.created_at,
			vs.trade_in,
			vs.owners,
			vs.updated_at,
			images,
			videos,
			vs.phone_numbers, 
			vs.view_count, 
			true as my_car,
			vs.description,
			CASE 
				WHEN ul.vehicle_id IS NOT NULL THEN true
				ELSE false
			END AS liked
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
		left join user_likes ul on ul.vehicle_id = vs.id AND ul.user_id = $1
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT image
				FROM images
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT video
				FROM videos
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		where vs.user_id = $1 and status = 3
		order by vs.id desc
	`
	rows, err := r.db.Query(ctx, q, userID)

	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.GetCarsResponse
		if err := rows.Scan(
			&car.ID, &car.Brand, &car.Region, &car.City, &car.Color, &car.Model, &car.Transmission, &car.Engine,
			&car.Drivetrain, &car.BodyType, &car.FuelType, &car.Year, &car.Price, &car.Mileage, &car.VinCode,
			&car.Credit, &car.New, &car.Status, &car.CreatedAt, &car.TradeIn, &car.Owners,
			&car.UpdatedAt, &car.Images, &car.Videos, &car.PhoneNumbers, &car.ViewCount, &car.MyCar, &car.Description, &car.Liked,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}
	return cars, err
}

func (r *ThirdPartyRepository) GetRegistrationData(ctx *fasthttp.RequestCtx) model.Response {
	q := `
		select id, name from company_types
	`
	var companyTypes []model.Model
	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	defer rows.Close()

	for rows.Next() {
		var companyType model.Model
		err = rows.Scan(&companyType.ID, &companyType.Name)

		if err != nil {
			return model.Response{Error: err, Status: http.StatusInternalServerError}
		}

		companyTypes = append(companyTypes, companyType)
	}

	q = `
		select id, name from activity_fields
	`
	var activityFields []model.Model
	rows, err = r.db.Query(ctx, q)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	defer rows.Close()

	for rows.Next() {
		var activityField model.Model
		err = rows.Scan(&activityField.ID, &activityField.Name)

		if err != nil {
			return model.Response{Error: err, Status: http.StatusInternalServerError}
		}

		activityFields = append(activityFields, activityField)
	}

	return model.Response{Data: model.ThirdPartyGetRegistrationDataRes{
		CompanyTypes:   companyTypes,
		ActivityFields: activityFields,
	}}
}

func (r *ThirdPartyRepository) CreateAvatarImages(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `
		update profiles set avatar = $1 where user_id = $2
	`
	_, err := r.db.Exec(ctx, q, paths[0], id)

	return err
}

func (r *ThirdPartyRepository) CreateBannerImage(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `
		update profiles set banner = $1 where user_id = $2
	`
	_, err := r.db.Exec(ctx, q, paths[0], id)

	return err
}

func (r *ThirdPartyRepository) CreateDealerCar(ctx *fasthttp.RequestCtx, car *model.ThirdPartyCreateCarRequest, dealerID int) (int, error) {
	keys, values, args := auth.BuildParams(car)

	q := `
		INSERT INTO vehicles 
			(
				` + strings.Join(keys, ", ") + `
				, user_id
			) VALUES (
				` + strings.Join(values, ", ") + `,
				$` + strconv.Itoa(len(keys)+1) + `
			) RETURNING id
	`
	var id int
	args = append(args, dealerID)
	err := r.db.QueryRow(ctx, q, args...).Scan(&id)

	return id, err
}

func (r *ThirdPartyRepository) UpdateDealerCar(ctx *fasthttp.RequestCtx, car *model.UpdateCarRequest, dealerID int) error {
	// First check if the car belongs to the dealer
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM vehicles WHERE id = $1 AND user_id = $2)`
	err := r.db.QueryRow(ctx, checkQuery, car.ID, dealerID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("car not found or access denied")
	}

	keys, _, args := auth.BuildParams(car)

	var updateFields []string
	var updateArgs []any
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

	updateArgs = append(updateArgs, dealerID)

	_, err = r.db.Exec(ctx, q, updateArgs...)
	return err
}

func (r *ThirdPartyRepository) GetEditDealerCarByID(ctx *fasthttp.RequestCtx, carID, dealerID int) (model.GetEditCarsResponse, error) {
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
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT image
				FROM images
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT video
				FROM videos
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		where vs.id = $1 and vs.user_id = $2;
	`
	err := r.db.QueryRow(ctx, q, carID, dealerID).Scan(
		&car.ID, &car.Brand, &car.Region, &car.City, &car.Model, &car.Modification,
		&car.Color, &car.BodyType, &car.Generation, &car.Year, &car.Price, &car.Odometer, &car.VinCode,
		&car.Wheel, &car.TradeIN, &car.Crash,
		&car.Credit, &car.New, &car.Status, &car.CreatedAt, &car.Images, &car.Videos, &car.PhoneNumbers,
		&car.ViewCount, &car.Description, &car.MyCar, &car.Owners,
	)

	return car, err
}

func (r *ThirdPartyRepository) CreateDealerCarImages(ctx *fasthttp.RequestCtx, carID int, images []string) error {

	if len(images) == 0 {
		return nil
	}

	q := `
		INSERT INTO images (vehicle_id, image) VALUES ($1, $2)
	`

	for i := range images {
		_, err := r.db.Exec(ctx, q, carID, images[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ThirdPartyRepository) CreateDealerCarVideos(ctx *fasthttp.RequestCtx, carID int, video string) error {

	q := `
		INSERT INTO videos (vehicle_id, video) VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, q, carID, video)
	if err != nil {
		return err
	}

	return err
}

func (r *ThirdPartyRepository) DealerDontSell(ctx *fasthttp.RequestCtx, carID, dealerID *int) error {
	q := `
		update vehicles 
			set status = 2 -- 2 is not sale
		where id = $1 and status = 3 -- 3 is on sale
			and user_id = $2
	`

	_, err := r.db.Exec(ctx, q, *carID, *dealerID)
	return err
}

func (r *ThirdPartyRepository) DealerSell(ctx *fasthttp.RequestCtx, carID, dealerID *int) error {
	q := `
		update vehicles 
			set status = 3 -- 3 is on sale
		where id = $1 and status = 2 -- 2 is not sale 
			and user_id = $2
	`
	_, err := r.db.Exec(ctx, q, *carID, *dealerID)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerCarImage(ctx *fasthttp.RequestCtx, carID int, imagePath string) error {
	q := `DELETE FROM images WHERE vehicle_id = $1 AND image = $2`
	_, err := r.db.Exec(ctx, q, carID, imagePath)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerCarVideo(ctx *fasthttp.RequestCtx, carID int, videoPath string) error {
	q := `DELETE FROM videos WHERE vehicle_id = $1 AND video = $2`
	_, err := r.db.Exec(ctx, q, carID, videoPath)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerCar(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		delete from vehicles where id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *ThirdPartyRepository) GetLogistDestinations(ctx *fasthttp.RequestCtx) ([]model.LogistDestinationResponse, error) {
	q := `
		SELECT 
			r.id,
			r.created_at,
			json_build_object(
				'id', cf.id,
				'name', cf.name,
				'flag', cf.flag
			) as from_country,
			json_build_object(
				'id', ct.id,
				'name', ct.name,
				'flag', ct.flag
			) as to_country
		FROM routes r
		LEFT JOIN countries cf ON r.from_id = cf.id
		LEFT JOIN countries ct ON r.to_id = ct.id
		ORDER BY r.created_at DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	destinations := make([]model.LogistDestinationResponse, 0)
	for rows.Next() {
		var dest model.LogistDestinationResponse
		if err := rows.Scan(&dest.ID, &dest.CreatedAt, &dest.From, &dest.To); err != nil {
			return nil, err
		}
		destinations = append(destinations, dest)
	}

	return destinations, nil
}

func (r *ThirdPartyRepository) CreateLogistDestination(ctx *fasthttp.RequestCtx, req model.CreateLogistDestinationRequest) (int, error) {
	q := `
		INSERT INTO routes (from_id, to_id)
		VALUES ($1, $2)
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, q, req.FromID, req.ToID).Scan(&id)
	return id, err
}

func (r *ThirdPartyRepository) DeleteLogistDestination(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		DELETE FROM routes WHERE id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
	return err
}
