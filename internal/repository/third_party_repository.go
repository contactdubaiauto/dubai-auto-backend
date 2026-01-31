package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/pkg/auth"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
)

type ThirdPartyRepository struct {
	config *config.Config
	db     *pgxpool.Pool
}

func NewThirdPartyRepository(config *config.Config, db *pgxpool.Pool) *ThirdPartyRepository {
	return &ThirdPartyRepository{config, db}
}

func (r *ThirdPartyRepository) Profile(ctx *fasthttp.RequestCtx, id int, profile model.ThirdPartyProfileReq) model.Response {
	contactsJSON, err := json.Marshal(profile.Contacts)
	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	q := `
		update profiles set
			about_me = $1,
			contacts = $2,
			address = $3,
			coordinates = $4,
			message = $5
		where user_id = $6
	`
	_, err = r.db.Exec(ctx, q, profile.AboutUs, contactsJSON,
		profile.Address, profile.Coordinates, profile.Message, id)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	q = `
		update users set
			phone = $1,
			username = $2
		where id = $3
	`
	_, err = r.db.Exec(ctx, q, profile.Phone, profile.Username, id)

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

func (r *ThirdPartyRepository) GetProfile(ctx *fasthttp.RequestCtx, id int, nameColumn string) (model.ThirdPartyGetProfileRes, error) {
	q := `
		with ds as (
			select
				uds.user_id,
				json_agg(
					json_build_object(
						'from_country', json_build_object(
							'id', fc.id,
							'name', fc.` + nameColumn + `,
							'flag', $2 || fc.flag
						),
						'to_country', json_build_object(
							'id', tc.id,
							'name', tc.` + nameColumn + `,
							'flag', $2 || tc.flag
						)
					)
				) as destinations
			from user_destinations uds
			left join countries fc on fc.id = uds.from_id
			left join countries tc on tc.id = uds.to_id
			where uds.user_id = $1
			group by uds.user_id
		)
		select
			p.user_id,
			p.about_me,
			p.contacts,
			p.address,
			p.coordinates,
			$2 || p.avatar,
			$2 || p.banner,
			p.company_name,
			p.message,
			p.vat_number,
			company_types.` + nameColumn + ` as company_type,
			activity_fields.` + nameColumn + ` as activity_field,
			p.created_at,
			ds.destinations
		from profiles p
		left join ds on ds.user_id = p.user_id
		left join company_types on company_types.id = p.company_type_id
		left join activity_fields on activity_fields.id = p.activity_field_id
		where p.user_id = $1
	`
	var profile model.ThirdPartyGetProfileRes
	err := r.db.QueryRow(ctx, q, id, r.config.IMAGE_BASE_URL).Scan(
		&profile.UserID,
		&profile.AboutUs, &profile.Contacts,
		&profile.Address,
		&profile.Coordinates, &profile.Avatar,
		&profile.Banner,
		&profile.CompanyName, &profile.Message,
		&profile.VATNumber, &profile.CompanyType,
		&profile.ActivityField,
		&profile.Registered,
		&profile.Destinations,
	)

	if err != nil {
		return profile, err
	}

	q = `
		select 
			email,
			phone,
			role_id
		from users 
		where id = $1
	`
	err = r.db.QueryRow(ctx, q, id).Scan(&profile.Email, &profile.Phone, &profile.RoleID)
	return profile, err
}

func (r *ThirdPartyRepository) GetMyCars(ctx *fasthttp.RequestCtx, userID, limit, lastID int, nameColumn string) ([]model.GetMyCarsResponse, error) {

	cars := make([]model.GetMyCarsResponse, 0)
	q := `
		with vs as (        
			select 
				vs.id,
				'car' as type,
				bs.` + nameColumn + ` as brand,
				ms.` + nameColumn + ` as model,
				cs.` + nameColumn + ` as city,
				vs.year,
				vs.price,
				vs.status,
				vs.created_at,
				images.images,
				vs.view_count,
				true as my_car,
				vs.crash,
				vs.moderation_status
			from vehicles vs
			left join brands bs on vs.brand_id = bs.id
			left join models ms on vs.model_id = ms.id
			left join cities cs on cs.id = vs.city_id
			LEFT JOIN LATERAL (
				SELECT json_agg(img.image) AS images
				FROM (
					SELECT $2 || image as image
					FROM images
					WHERE vehicle_id = vs.id
					ORDER BY created_at DESC
				) img
			) images ON true
			where vs.user_id = $1 and vs.status = 2
			order by vs.id desc
		),
		cms as (
			select
				cm.id,
				'comtran' as type,
				cbs.` + nameColumn + ` as brand,
				cms.` + nameColumn + ` as model,
				cs.` + nameColumn + ` as city,
				cm.year,
				cm.price,
				cm.status,
				cm.created_at,
				images.images,
				cm.view_count,
				true as my_car,
				cm.crash,
				cm.moderation_status
			from comtrans cm
			left join com_brands cbs on cbs.id = cm.comtran_brand_id
			left join com_models cms on cms.id = cm.comtran_model_id
			left join cities cs on cs.id = cm.city_id
			LEFT JOIN LATERAL (
				SELECT json_agg(img.image) AS images
				FROM (
					SELECT $2 || image as image
					FROM comtran_images
					WHERE comtran_id = cm.id
					ORDER BY created_at DESC
				) img
			) images ON true
			where cm.user_id = $1 and cm.status = 2
		),
		mts as (
			select
				mt.id,
				'motorcycle' as type,
				mbs.` + nameColumn + ` as brand,
				mms.` + nameColumn + ` as model,
				cs.` + nameColumn + ` as city,
				mt.year,
				mt.price,
				mt.status,
				mt.created_at,
				mt.view_count,
				images.images,
				true as my_car,
				mt.crash,
				mt.moderation_status
			from motorcycles mt
			left join moto_brands mbs on mbs.id = mt.moto_brand_id
			left join moto_models mms on mms.id = mt.moto_model_id
			left join cities cs on cs.id = mt.city_id
			LEFT JOIN LATERAL (
				SELECT json_agg(img.image) AS images
				FROM (
					SELECT $2 || image as image
					FROM moto_images
					WHERE moto_id = mt.id
					ORDER BY created_at DESC
				) img
			) images ON true
			where mt.user_id = $1 and mt.status = 2
		)
		-- Union all three CTEs
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash, moderation_status, city
		from vs
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash, moderation_status, city
		from cms
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash, moderation_status, city
		from mts
		order by created_at desc;

	`
	rows, err := r.db.Query(ctx, q, userID, r.config.IMAGE_BASE_URL)

	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.GetMyCarsResponse
		if err := rows.Scan(
			&car.ID,
			&car.Type,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.Price,
			&car.Status,
			&car.CreatedAt,
			&car.ViewCount,
			&car.Images,
			&car.MyCar,
			&car.Crash,
			&car.ModerationStatus,
			&car.City,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *ThirdPartyRepository) OnSale(ctx *fasthttp.RequestCtx, userID int, nameColumn string) ([]model.GetMyCarsResponse, error) {

	cars := make([]model.GetMyCarsResponse, 0)
	q := `
		with vs as (        
			select 
				vs.id,
				'car' as type,
				bs.` + nameColumn + ` as brand,
				ms.` + nameColumn + ` as model,
				cs.` + nameColumn + ` as city,
				vs.year,
				vs.price,
				vs.status,
				vs.created_at,
				images.images,
				vs.view_count,
				true as my_car,
				vs.crash,
				vs.moderation_status
			from vehicles vs
			left join brands bs on vs.brand_id = bs.id
			left join models ms on vs.model_id = ms.id
			left join cities cs on cs.id = vs.city_id
			LEFT JOIN LATERAL (
				SELECT json_agg(img.image) AS images
				FROM (
					SELECT $2 || image as image
					FROM images
					WHERE vehicle_id = vs.id
					ORDER BY created_at DESC
				) img
			) images ON true
			where vs.user_id = $1 and (status = 3 or status = 1)
			order by vs.id desc
		),
		cms as (
			select
				cm.id,
				'comtran' as type,
				cbs.` + nameColumn + ` as brand,
				cms.` + nameColumn + ` as model,
				cs.` + nameColumn + ` as city,
				cm.year,
				cm.price,
				cm.status,
				cm.created_at,
				images.images,
				cm.view_count,
				true as my_car,
				cm.crash,
				cm.moderation_status
			from comtrans cm
			left join com_brands cbs on cbs.id = cm.comtran_brand_id
			left join com_models cms on cms.id = cm.comtran_model_id
			left join cities cs on cs.id = cm.city_id
			LEFT JOIN LATERAL (
				SELECT json_agg(img.image) AS images
				FROM (
					SELECT $2 || image as image
					FROM comtran_images
					WHERE comtran_id = cm.id
					ORDER BY created_at DESC
				) img
			) images ON true
			where cm.user_id = $1 and (status = 3 or status = 1)
		),
		mts as (
			select
				mt.id,
				'motorcycle' as type,
				mbs.` + nameColumn + ` as brand,
				mms.` + nameColumn + ` as model,
				cs.` + nameColumn + ` as city,
				mt.year,
				mt.price,
				mt.status,
				mt.created_at,
				mt.view_count,
				images.images,
				true as my_car,
				mt.crash,
				mt.moderation_status
			from motorcycles mt
			left join moto_brands mbs on mbs.id = mt.moto_brand_id
			left join moto_models mms on mms.id = mt.moto_model_id
			left join cities cs on cs.id = mt.city_id
			LEFT JOIN LATERAL (
				SELECT json_agg(img.image) AS images
				FROM (
					SELECT $2 || image as image
					FROM moto_images
					WHERE moto_id = mt.id
					ORDER BY created_at DESC
				) img
			) images ON true
			where mt.user_id = $1 and (status = 3 or status = 1)
		)
		-- Union all three CTEs
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash, moderation_status, city
		from vs
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash, moderation_status, city
		from cms
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash, moderation_status, city
		from mts
		order by created_at desc;

	`
	rows, err := r.db.Query(ctx, q, userID, r.config.IMAGE_BASE_URL)

	if err != nil {
		return cars, err
	}
	defer rows.Close()

	for rows.Next() {
		var car model.GetMyCarsResponse
		if err := rows.Scan(
			&car.ID,
			&car.Type,
			&car.Brand,
			&car.Model,
			&car.Year,
			&car.Price,
			&car.Status,
			&car.CreatedAt,
			&car.ViewCount,
			&car.Images,
			&car.MyCar,
			&car.Crash,
			&car.ModerationStatus,
			&car.City,
		); err != nil {
			return cars, err
		}
		cars = append(cars, car)
	}

	return cars, err
}

func (r *ThirdPartyRepository) GetRegistrationData(ctx *fasthttp.RequestCtx, nameColumn string) model.Response {

	q := `
		select id, ` + nameColumn + ` from company_types
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
		select id, ` + nameColumn + ` from activity_fields
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

func (r *ThirdPartyRepository) DeleteAvatarImages(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		update profiles set avatar = null where user_id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *ThirdPartyRepository) CreateBannerImage(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `
		update profiles set banner = $1 where user_id = $2
	`
	_, err := r.db.Exec(ctx, q, paths[0], id)

	return err
}

func (r *ThirdPartyRepository) DeleteBannerImage(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		update profiles set banner = null where user_id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
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

	// shun icinde eger false bolsa crash we beyleki boolean params build edilenok
	keys, _, args := auth.BuildParams(car)

	var updateFields []string
	var updateArgs []any
	updateArgs = append(updateArgs, car.ID)
	paramIndex := 2

	if !car.Crash {
		updateFields = append(updateFields, "crash = $"+strconv.Itoa(paramIndex))
		updateArgs = append(updateArgs, false)
		paramIndex++
	}

	if !car.New {
		updateFields = append(updateFields, "new = $"+strconv.Itoa(paramIndex))
		updateArgs = append(updateArgs, false)
		paramIndex++
	}

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

func (r *ThirdPartyRepository) GetEditDealerCarByID(ctx *fasthttp.RequestCtx, carID, dealerID int, nameColumn string) (model.GetEditCarsResponse, error) {
	car := model.GetEditCarsResponse{}

	q := `
		select 
			vs.id,
			json_build_object(
				'id', bs.id,
				'name', bs.` + nameColumn + `,
				'logo', $3 || bs.logo,
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
				'name', ms.` + nameColumn + `
			) as model,
			json_build_object(
				'id', mfs.id,
				'engine', es.` + nameColumn + `,
				'fuel_type', fts.` + nameColumn + `,
				'drivetrain', ds.` + nameColumn + `,
				'transmission', ts.` + nameColumn + `
			) as modification,
			json_build_object(
				'id', cls.id,
				'name', cls.` + nameColumn + `,
				'image', $3 || cls.image
			) as color,
			json_build_object(
				'id', bts.id,
				'name', bts.` + nameColumn + `,
				'image', $3 || bts.image
			) as body_type,
			json_build_object(
				'id', gs.id,
				'name', gs.` + nameColumn + `,
				'image', $3 || gs.image,
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
				SELECT $3 || image as image
				FROM images
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT $3 || video as video
				FROM videos
				WHERE vehicle_id = vs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		where vs.id = $1 and vs.user_id = $2;
	`
	err := r.db.QueryRow(ctx, q, carID, dealerID, r.config.IMAGE_BASE_URL).Scan(
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

// Dealer Motorcycle repository methods

func (r *ThirdPartyRepository) CreateDealerMotorcycle(ctx *fasthttp.RequestCtx, motorcycle *model.CreateMotorcycleRequest, dealerID int) (int, error) {
	keys, values, args := auth.BuildParams(motorcycle)

	q := `
		INSERT INTO motorcycles 
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

func (r *ThirdPartyRepository) UpdateDealerMotorcycle(ctx *fasthttp.RequestCtx, motorcycle *model.UpdateMotorcycleRequest, dealerID int) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM motorcycles WHERE id = $1 AND user_id = $2)`
	err := r.db.QueryRow(ctx, checkQuery, motorcycle.ID, dealerID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("motorcycle not found or access denied")
	}

	keys, _, args := auth.BuildParams(motorcycle)

	var updateFields []string
	var updateArgs []any
	updateArgs = append(updateArgs, motorcycle.ID)
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
		UPDATE motorcycles 
		SET ` + strings.Join(updateFields, ", ") + `, updated_at = NOW()
		WHERE id = $1 AND user_id = $` + fmt.Sprintf("%d", paramIndex)

	updateArgs = append(updateArgs, dealerID)

	_, err = r.db.Exec(ctx, q, updateArgs...)
	return err
}

func (r *ThirdPartyRepository) GetEditDealerMotorcycleByID(ctx *fasthttp.RequestCtx, motorcycleID, dealerID int, nameColumn string) (model.GetEditMotorcycleResponse, error) {
	motorcycle := model.GetEditMotorcycleResponse{}

	q := `
		select 
			mcs.id,
			json_build_object(
				'id', u.id,
				'username', u.username,
				'avatar', $3 || pf.avatar,
				'contacts', pf.contacts,
				'role_id', u.role_id
			) as owner,
			mcs.engine,
			mcs.power,
			mcs.year,
			json_build_object(
				'id', nocs.id,
				'name', nocs.` + nameColumn + `
			) as number_of_cycles,
			mcs.odometer,
			mcs.crash,
			mcs.wheel,
			mcs.new,
			mcs.owners,
			mcs.vin_code,
			mcs.description,
			mcs.phone_numbers,
			mcs.price,
			mcs.trade_in,
			mcs.status::text,
			mcs.updated_at,
			mcs.created_at,
			json_build_object(
				'id', mocs.id,
				'name', mocs.` + nameColumn + `
			) as moto_category,
			json_build_object(
				'id', mbs.id,
				'name', mbs.` + nameColumn + `
			) as moto_brand,
			json_build_object(
				'id', mms.id,
				'name', mms.` + nameColumn + `
			) as moto_model,
			json_build_object(
				'id', meng.id,
				'name', meng.` + nameColumn + `
			) as engine_type,
			json_build_object(
				'id', cs.id,
				'name', cs.name
			) as city,
			json_build_object(
				'id', cls.id,
				'name', cls.` + nameColumn + `,
				'image', $3 || cls.image
			) as color,
			CASE
				WHEN mcs.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_moto,
			images.images,
			videos.videos
		from motorcycles mcs
		left join users u on u.id = mcs.user_id
		left join profiles pf on pf.user_id = mcs.user_id
		left join moto_categories mocs on mocs.id = mcs.moto_category_id
		left join moto_brands mbs on mbs.id = mcs.moto_brand_id
		left join moto_models mms on mms.id = mcs.moto_model_id
		left join moto_engines meng on meng.id = mcs.engine_id
		left join cities cs on cs.id = mcs.city_id
		left join colors cls on cls.id = mcs.color_id
		left join number_of_cycles nocs on nocs.id = mcs.number_of_cycles_id
		LEFT JOIN LATERAL (
			SELECT json_agg(json_build_object('image', img.image, 'id', img.id)) AS images
			FROM (
				SELECT $3 || image as image, id
				FROM moto_images
				WHERE moto_id = mcs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(json_build_object('video', v.video, 'id', v.id)) AS videos
			FROM (
				SELECT $3 || video as video, id
				FROM moto_videos
				WHERE moto_id = mcs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE mcs.id = $1 AND mcs.user_id = $2;
	`
	err := r.db.QueryRow(ctx, q, motorcycleID, dealerID, r.config.IMAGE_BASE_URL).Scan(
		&motorcycle.ID, &motorcycle.Owner, &motorcycle.Engine, &motorcycle.Power, &motorcycle.Year,
		&motorcycle.NumberOfCycles, &motorcycle.Odometer, &motorcycle.Crash, &motorcycle.Wheel, &motorcycle.New,
		&motorcycle.Owners, &motorcycle.VinCode, &motorcycle.Description, &motorcycle.PhoneNumbers,
		&motorcycle.Price, &motorcycle.TradeIn, &motorcycle.Status,
		&motorcycle.UpdatedAt, &motorcycle.CreatedAt, &motorcycle.MotoCategory, &motorcycle.MotoBrand,
		&motorcycle.MotoModel, &motorcycle.EngineType, &motorcycle.City, &motorcycle.Color, &motorcycle.MyMoto,
		&motorcycle.Images, &motorcycle.Videos)

	return motorcycle, err
}

func (r *ThirdPartyRepository) CreateDealerMotorcycleImages(ctx *fasthttp.RequestCtx, motorcycleID int, images []string) error {
	if len(images) == 0 {
		return nil
	}

	q := `
		INSERT INTO moto_images (moto_id, image) VALUES ($1, $2)
	`

	for i := range images {
		_, err := r.db.Exec(ctx, q, motorcycleID, images[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ThirdPartyRepository) CreateDealerMotorcycleVideos(ctx *fasthttp.RequestCtx, motorcycleID int, video string) error {
	q := `
		INSERT INTO moto_videos (moto_id, video) VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, q, motorcycleID, video)
	return err
}

func (r *ThirdPartyRepository) DealerDontSellMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, dealerID *int) error {
	q := `
		update motorcycles 
			set status = 2 -- 2 is not sale
		where id = $1 and status = 3 -- 3 is on sale
			and user_id = $2
	`

	_, err := r.db.Exec(ctx, q, *motorcycleID, *dealerID)
	return err
}

func (r *ThirdPartyRepository) DealerSellMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, dealerID *int) error {
	q := `
		update motorcycles 
			set status = 3 -- 3 is on sale
		where id = $1 and status = 2 -- 2 is not sale 
			and user_id = $2
	`
	_, err := r.db.Exec(ctx, q, *motorcycleID, *dealerID)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerMotorcycleImage(ctx *fasthttp.RequestCtx, motorcycleID int, imagePath string) error {
	q := `DELETE FROM moto_images WHERE moto_id = $1 AND image = $2`
	_, err := r.db.Exec(ctx, q, motorcycleID, imagePath)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerMotorcycleVideo(ctx *fasthttp.RequestCtx, motorcycleID int, videoPath string) error {
	q := `DELETE FROM moto_videos WHERE moto_id = $1 AND video = $2`
	_, err := r.db.Exec(ctx, q, motorcycleID, videoPath)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerMotorcycle(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		delete from motorcycles where id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Dealer Comtrans repository methods

func (r *ThirdPartyRepository) CreateDealerComtrans(ctx *fasthttp.RequestCtx, comtrans *model.CreateComtransRequest, dealerID int) (int, error) {
	keys, values, args := auth.BuildParams(comtrans)

	q := `
		INSERT INTO comtrans 
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

func (r *ThirdPartyRepository) UpdateDealerComtrans(ctx *fasthttp.RequestCtx, comtrans *model.UpdateComtransRequest, dealerID int) error {
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM comtrans WHERE id = $1 AND user_id = $2)`
	err := r.db.QueryRow(ctx, checkQuery, comtrans.ID, dealerID).Scan(&exists)

	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("comtrans not found or access denied")
	}

	keys, _, args := auth.BuildParams(comtrans)

	var updateFields []string
	var updateArgs []any
	updateArgs = append(updateArgs, comtrans.ID)
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
		UPDATE comtrans 
		SET ` + strings.Join(updateFields, ", ") + `, updated_at = NOW()
		WHERE id = $1 AND user_id = $` + fmt.Sprintf("%d", paramIndex)

	updateArgs = append(updateArgs, dealerID)

	_, err = r.db.Exec(ctx, q, updateArgs...)
	return err
}

func (r *ThirdPartyRepository) GetEditDealerComtransByID(ctx *fasthttp.RequestCtx, comtransID, dealerID int, nameColumn string) (model.GetEditComtransResponse, error) {
	comtrans := model.GetEditComtransResponse{}

	q := `
		select 
			cts.id,
			json_build_object(
				'id', u.id,
				'username', u.username,
				'avatar', $3 || pf.avatar,
				'contacts', pf.contacts,
				'role_id', u.role_id
			) as owner,
			cts.engine,
			cts.power,
			cts.year,
			cts.odometer,
			cts.crash,
			cts.wheel,
			cts.new,
			cts.owners,
			cts.vin_code,
			cts.description,
			cts.phone_numbers,
			cts.price,
			cts.trade_in,
			cts.status,
			cts.updated_at,
			cts.created_at,
			json_build_object(
				'id', cocs.id,
				'name', cocs.` + nameColumn + `
			) as comtran_category,
			json_build_object(
				'id', cbs.id,
				'name', cbs.` + nameColumn + `
			) as comtran_brand,
			json_build_object(
				'id', cms.id,
				'name', cms.` + nameColumn + `
			) as comtran_model,
			json_build_object(
				'id', ces.id,
				'name', ces.` + nameColumn + `
			) as engine_type,
			json_build_object(
				'id', cs.id,
				'name', cs.name
			) as city,
			json_build_object(
				'id', cls.id,
				'name', cls.` + nameColumn + `,
				'image', $3 || cls.image
			) as color,
			CASE
				WHEN cts.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_comtrans,
			images.images,
			videos.videos
		from comtrans cts
		left join users u on u.id = cts.user_id
		left join profiles pf on pf.user_id = cts.user_id
		left join com_categories cocs on cocs.id = cts.comtran_category_id
		left join com_brands cbs on cbs.id = cts.comtran_brand_id
		left join com_models cms on cms.id = cts.comtran_model_id
		left join com_engines ces on ces.id = cts.engine_id
		left join cities cs on cs.id = cts.city_id
		left join colors cls on cls.id = cts.color_id
		LEFT JOIN LATERAL (
			SELECT json_agg(json_build_object('image', img.image, 'id', img.id)) AS images
			FROM (
				SELECT $3 || image as image, id
				FROM comtran_images
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(json_build_object('video', v.video, 'id', v.id)) AS videos
			FROM (
				SELECT $3 || video as video, id
				FROM comtran_videos
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE cts.id = $1 AND cts.user_id = $2;
	`

	err := r.db.QueryRow(ctx, q, comtransID, dealerID, r.config.IMAGE_BASE_URL).Scan(
		&comtrans.ID, &comtrans.Owner, &comtrans.Engine, &comtrans.Power, &comtrans.Year,
		&comtrans.Odometer, &comtrans.Crash, &comtrans.Wheel, &comtrans.New, &comtrans.Owners, &comtrans.VinCode, &comtrans.Description,
		&comtrans.PhoneNumbers, &comtrans.Price, &comtrans.TradeIn, &comtrans.Status,
		&comtrans.UpdatedAt, &comtrans.CreatedAt, &comtrans.ComtranCategory, &comtrans.ComtranBrand,
		&comtrans.ComtranModel, &comtrans.EngineType, &comtrans.City, &comtrans.Color, &comtrans.MyComtrans,
		&comtrans.Images, &comtrans.Videos)

	return comtrans, err
}

func (r *ThirdPartyRepository) CreateDealerComtransImages(ctx *fasthttp.RequestCtx, comtransID int, images []string) error {
	if len(images) == 0 {
		return nil
	}

	q := `
		INSERT INTO comtran_images (comtran_id, image) VALUES ($1, $2)
	`

	for i := range images {
		_, err := r.db.Exec(ctx, q, comtransID, images[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ThirdPartyRepository) CreateDealerComtransVideos(ctx *fasthttp.RequestCtx, comtransID int, video string) error {
	q := `
		INSERT INTO comtran_videos (comtran_id, video) VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, q, comtransID, video)
	return err
}

func (r *ThirdPartyRepository) DealerDontSellComtrans(ctx *fasthttp.RequestCtx, comtransID, dealerID *int) error {
	q := `
		update comtrans 
			set status = 2 -- 2 is not sale
		where id = $1 and status = 3 -- 3 is on sale
			and user_id = $2
	`

	_, err := r.db.Exec(ctx, q, *comtransID, *dealerID)
	return err
}

func (r *ThirdPartyRepository) DealerSellComtrans(ctx *fasthttp.RequestCtx, comtransID, dealerID *int) error {
	q := `
		update comtrans 
			set status = 3 -- 3 is on sale
		where id = $1 and status = 2 -- 2 is not sale 
			and user_id = $2
	`
	_, err := r.db.Exec(ctx, q, *comtransID, *dealerID)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerComtransImage(ctx *fasthttp.RequestCtx, comtransID int, imagePath string) error {
	q := `DELETE FROM comtran_images WHERE comtran_id = $1 AND image = $2`
	_, err := r.db.Exec(ctx, q, comtransID, imagePath)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerComtransVideo(ctx *fasthttp.RequestCtx, comtransID int, videoPath string) error {
	q := `DELETE FROM comtran_videos WHERE comtran_id = $1 AND video = $2`
	_, err := r.db.Exec(ctx, q, comtransID, videoPath)
	return err
}

func (r *ThirdPartyRepository) DeleteDealerComtrans(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		delete from comtrans where id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *ThirdPartyRepository) GetLogistDestinations(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.LogistDestinationResponse, error) {
	q := `
		SELECT 
			r.id,
			r.created_at,
			json_build_object(
				'id', cf.id,
				'name', cf.` + nameColumn + `,
				'flag', $1 || cf.flag as flag
			) as from_country,
			json_build_object(
				'id', ct.id,
				'name', ct.` + nameColumn + `,
				'flag', $1 || ct.flag as flag
			) as to_country
		FROM user_destinations r
		LEFT JOIN countries cf ON r.from_id = cf.id
		LEFT JOIN countries ct ON r.to_id = ct.id
		ORDER BY r.created_at DESC
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
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

func (r *ThirdPartyRepository) CreateLogistDestination(ctx *fasthttp.RequestCtx, userID int, req model.CreateLogistDestinationRequest) (int, error) {
	q := `
		INSERT INTO user_destinations (user_id, from_id, to_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, from_id, to_id) DO NOTHING
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, q, userID, req.FromID, req.ToID).Scan(&id)
	return id, err
}

func (r *ThirdPartyRepository) DeleteLogistDestination(ctx *fasthttp.RequestCtx, userID int, id int) error {
	q := `
		DELETE FROM user_destinations WHERE user_id = $1 AND id = $2
	`
	_, err := r.db.Exec(ctx, q, userID, id)
	return err
}

func (r *ThirdPartyRepository) DeleteAllLogistDestinations(ctx *fasthttp.RequestCtx, userID int) error {
	q := `
		DELETE FROM user_destinations WHERE user_id = $1
	`
	_, err := r.db.Exec(ctx, q, userID)
	return err
}
