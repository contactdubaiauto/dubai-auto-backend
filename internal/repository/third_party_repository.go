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
				vs.year,
				vs.price,
				vs.status,
				vs.created_at,
				images.images,
				vs.view_count,
				true as my_car,
				vs.crash
			from vehicles vs
			left join brands bs on vs.brand_id = bs.id
			left join models ms on vs.model_id = ms.id
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
				cm.year,
				cm.price,
				cm.status,
				cm.created_at,
				images.images,
				cm.view_count,
				true as my_car,
				cm.crash
			from comtrans cm
			left join com_brands cbs on cbs.id = cm.comtran_brand_id
			left join com_models cms on cms.id = cm.comtran_model_id
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
				mt.year,
				mt.price,
				mt.status,
				mt.created_at,
				mt.view_count,
				images.images,
				true as my_car,
				mt.crash
			from motorcycles mt
			left join moto_brands mbs on mbs.id = mt.moto_brand_id
			left join moto_models mms on mms.id = mt.moto_model_id
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
			crash 
		from vs
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash 
		from cms
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash 
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
				vs.year,
				vs.price,
				vs.status,
				vs.created_at,
				images.images,
				vs.view_count,
				true as my_car,
				vs.crash
			from vehicles vs
			left join brands bs on vs.brand_id = bs.id
			left join models ms on vs.model_id = ms.id
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
				cm.year,
				cm.price,
				cm.status,
				cm.created_at,
				images.images,
				cm.view_count,
				true as my_car,
				cm.crash
			from comtrans cm
			left join com_brands cbs on cbs.id = cm.comtran_brand_id
			left join com_models cms on cms.id = cm.comtran_model_id
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
				mt.year,
				mt.price,
				mt.status,
				mt.created_at,
				mt.view_count,
				images.images,
				true as my_car,
				mt.crash
			from motorcycles mt
			left join moto_brands mbs on mbs.id = mt.moto_brand_id
			left join moto_models mms on mms.id = mt.moto_model_id
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
			crash 
		from vs
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash 
		from cms
		union all
		select 
			id, type, brand, model, 
			year, price,
			status, created_at, 
			view_count, images, my_car, 
			crash 
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
