package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/pkg/auth"
)

type MotorcycleRepository struct {
	config *config.Config
	db     *pgxpool.Pool
}

func NewMotorcycleRepository(config *config.Config, db *pgxpool.Pool) *MotorcycleRepository {
	return &MotorcycleRepository{config, db}
}

func (r *MotorcycleRepository) GetMotorcycleCategories(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetMotorcycleCategoriesResponse, error) {
	data := make([]model.GetMotorcycleCategoriesResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM moto_categories
	`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category model.GetMotorcycleCategoriesResponse
		err = rows.Scan(&category.ID, &category.Name)

		if err != nil {
			return nil, err
		}

		data = append(data, category)
	}

	return data, nil
}

func (r *MotorcycleRepository) GetMotorcycleBrands(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetMotorcycleBrandsResponse, error) {
	data := make([]model.GetMotorcycleBrandsResponse, 0)
	q := `
		SELECT id, ` + nameColumn + `, $1 || logo as image, model_count FROM moto_brands
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand model.GetMotorcycleBrandsResponse
		err = rows.Scan(&brand.ID, &brand.Name, &brand.Image, &brand.ModelCount)

		if err != nil {
			return nil, err
		}

		data = append(data, brand)
	}

	return data, nil
}

func (r *MotorcycleRepository) GetNumberOfCycles(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetNumberOfCyclesResponse, error) {
	data := make([]model.GetNumberOfCyclesResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM number_of_cycles
		ORDER BY id ASC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var numberOfCycle model.GetNumberOfCyclesResponse
		err = rows.Scan(&numberOfCycle.ID, &numberOfCycle.Name)

		if err != nil {
			return nil, err
		}

		data = append(data, numberOfCycle)
	}

	return data, nil
}

func (r *MotorcycleRepository) GetMotorcycleModelsByBrandID(ctx *fasthttp.RequestCtx, brandID int, nameColumn string) ([]model.GetMotorcycleModelsResponse, error) {
	data := make([]model.GetMotorcycleModelsResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM moto_models
		WHERE moto_brand_id = $1
	`
	rows, err := r.db.Query(ctx, q, brandID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var model model.GetMotorcycleModelsResponse
		err = rows.Scan(&model.ID, &model.Name)

		if err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	return data, nil
}

func (r *MotorcycleRepository) GetMotoEngines(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetMotorcycleModelsResponse, error) {
	data := make([]model.GetMotorcycleModelsResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM moto_engines
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var engine model.GetMotorcycleModelsResponse
		err = rows.Scan(&engine.ID, &engine.Name)
		data = append(data, engine)
	}
	return data, nil
}

func (r *MotorcycleRepository) CreateMotorcycle(ctx *fasthttp.RequestCtx, req model.CreateMotorcycleRequest, userID int) (model.SuccessWithId, error) {
	data := model.SuccessWithId{}

	keys, values, args := auth.BuildParams(req)

	q := `
		INSERT INTO motorcycles ( 
			` + strings.Join(keys, ", ") + `,
			user_id
		) VALUES (
			` + strings.Join(values, ", ") + `,
			$` + strconv.Itoa(len(keys)+1) + `
		) returning id
	`
	var id int
	args = append(args, userID)
	err := r.db.QueryRow(ctx, q, args...).Scan(&id)

	if err != nil {
		return data, err
	}

	data.Message = "Motorcycle created successfully"
	data.Id = id

	return data, err
}

func (r *MotorcycleRepository) GetMotorcycles(ctx *fasthttp.RequestCtx, userID int, brands, models, regions, cities,
	transmissions, categories, engines, fuel_types, colors, dealers []string,
	year_from, year_to, price_from, price_to, tradeIn, owners, crash, odometer string, newQ, wheelQ *bool,
	limit, lastID int, nameColumn string) ([]model.GetMotorcyclesResponse, error) {
	data := make([]model.GetMotorcyclesResponse, 0)
	var qWhere string
	var qValues []any
	// $1 = IMAGE_BASE_URL, $2 = userID, $3 = lastID
	qValues = append(qValues, r.config.IMAGE_BASE_URL, userID, lastID)
	i := 3

	if len(brands) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mbs.id = ANY($%d)", i)
		qValues = append(qValues, brands)
	}

	if len(models) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mms.id = ANY($%d)", i)
		qValues = append(qValues, models)
	}

	if len(categories) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mcs.moto_category_id = ANY($%d)", i)
		qValues = append(qValues, categories)
	}

	if len(cities) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mcs.city_id = ANY($%d)", i)
		qValues = append(qValues, cities)
	}

	if len(engines) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mcs.engine_id = ANY($%d)", i)
		qValues = append(qValues, engines)
	}

	if len(colors) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mcs.color_id = ANY($%d)", i)
		qValues = append(qValues, colors)
	}

	if len(dealers) > 0 {
		i++
		qWhere += fmt.Sprintf(" AND mcs.user_id = ANY($%d)", i)
		qValues = append(qValues, dealers)
	}

	if year_from != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.year >= $%d", i)
		qValues = append(qValues, year_from)
	}

	if year_to != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.year <= $%d", i)
		qValues = append(qValues, year_to)
	}

	if tradeIn != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.trade_in = $%d", i)
		qValues = append(qValues, tradeIn)
	}

	if owners != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.owners = $%d", i)
		qValues = append(qValues, owners)
	}

	if crash != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.crash = $%d", i)
		qValues = append(qValues, crash)
	}

	if price_from != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.price >= $%d", i)
		qValues = append(qValues, price_from)
	}

	if price_to != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.price <= $%d", i)
		qValues = append(qValues, price_to)
	}

	if newQ != nil {
		i++
		qWhere += fmt.Sprintf(" AND mcs.new = $%d", i)
		qValues = append(qValues, newQ)
	}

	if wheelQ != nil {
		i++
		qWhere += fmt.Sprintf(" AND mcs.wheel = $%d", i)
		qValues = append(qValues, wheelQ)
	}

	if odometer != "" {
		i++
		qWhere += fmt.Sprintf(" AND mcs.odometer <= $%d", i)
		qValues = append(qValues, odometer)
	}

	// Add limit parameter
	i++
	limitPlaceholder := fmt.Sprintf("$%d", i)
	qValues = append(qValues, limit)

	q := `
		SELECT 
			mcs.id,
			'moto' as type,
			mbs.` + nameColumn + ` as moto_brand,
			mms.` + nameColumn + ` as moto_model,
			mcs.year,
			mcs.price,
			mcs.created_at,
			images.images,
			mcs.new,
			mcs.status,
			mcs.trade_in,
			mcs.crash,
			mcs.odometer,
			mcs.view_count,
			CASE
				WHEN mcs.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_moto,
			CASE
				WHEN u.role_id = 2 THEN u.username
				ELSE NULL
			END AS owner_name,
			cs.` + nameColumn + ` as city,
			CASE 
				WHEN ul.moto_id IS NOT NULL THEN true
				ELSE false
			END AS liked
		FROM motorcycles mcs
		LEFT JOIN moto_brands mbs ON mbs.id = mcs.moto_brand_id
		left join user_moto_likes ul on ul.moto_id = vs.id AND ul.user_id = $2
		LEFT JOIN moto_models mms ON mms.id = mcs.moto_model_id
		LEFT JOIN users u ON u.id = mcs.user_id
		LEFT JOIN cities cs on cs.id = mcs.city_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $1 || image as image
				FROM moto_images
				WHERE moto_id = mcs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		WHERE mcs.status = 3 AND (mcs.moderation_status = 1 OR mcs.moderation_status = 2) AND mcs.id > $3
		` + qWhere + `
		ORDER BY mcs.id DESC
		LIMIT ` + limitPlaceholder + `
	`

	rows, err := r.db.Query(ctx, q, qValues...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var motorcycle model.GetMotorcyclesResponse
		err = rows.Scan(
			&motorcycle.ID, &motorcycle.Type, &motorcycle.Brand, &motorcycle.Model, &motorcycle.Year, &motorcycle.Price,
			&motorcycle.CreatedAt, &motorcycle.Images,
			&motorcycle.New, &motorcycle.Status, &motorcycle.TradeIn,
			&motorcycle.Crash, &motorcycle.Odometer,
			&motorcycle.ViewCount, &motorcycle.MyMoto,
			&motorcycle.OwnerName, &motorcycle.City, &motorcycle.Liked)

		if err != nil {
			return nil, err
		}

		data = append(data, motorcycle)
	}

	return data, nil
}

func (r *MotorcycleRepository) CreateMotorcycleImages(ctx *fasthttp.RequestCtx, motorcycleID int, images []string) error {

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

func (r *MotorcycleRepository) CreateMotorcycleVideos(ctx *fasthttp.RequestCtx, motorcycleID int, video string) error {

	q := `
		INSERT INTO moto_videos (moto_id, video) VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, q, motorcycleID, video)
	if err != nil {
		return err
	}

	return err
}

func (r *MotorcycleRepository) DeleteMotorcycleImage(ctx *fasthttp.RequestCtx, motorcycleID int, imageID int) error {
	q := `
		DELETE FROM moto_images WHERE moto_id = $1 AND id = $2
	`

	_, err := r.db.Exec(ctx, q, motorcycleID, imageID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MotorcycleRepository) DeleteMotorcycleVideo(ctx *fasthttp.RequestCtx, motorcycleID int, videoID int) error {
	q := `
		DELETE FROM moto_videos WHERE moto_id = $1 AND id = $2
	`

	_, err := r.db.Exec(ctx, q, motorcycleID, videoID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MotorcycleRepository) GetMotorcycleByID(ctx *fasthttp.RequestCtx, motorcycleID, userID int, nameColumn string) (model.GetMotorcycleResponse, error) {
	var motorcycle model.GetMotorcycleResponse
	q := `
		WITH updated AS (
			UPDATE motorcycles
			SET view_count = COALESCE(view_count, 0) + 1
			WHERE id = $1
			RETURNING *
		)
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
			nocs.` + nameColumn + ` as number_of_cycles,
			mcs.odometer,
			mcs.crash,
			mcs.wheel,
			mcs.owners,
			mcs.vin_code,
			mcs.description,
			mcs.phone_numbers,
			mcs.price,
			mcs.trade_in,
			mcs.status::text,
			mcs.updated_at,
			mcs.created_at,
			mocs.` + nameColumn + ` as moto_category,
			mbs.` + nameColumn + ` as moto_brand,
			mms.` + nameColumn + ` as moto_model,
			meng.` + nameColumn + ` as engine_type,
			cs.name as city,
			cls.` + nameColumn + ` as color,
			CASE
				WHEN mcs.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_moto,
			images.images,
			videos.videos
		from updated mcs
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
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $3 || image as image
				FROM moto_images
				WHERE moto_id = mcs.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT $3 || video as video
				FROM moto_videos
				WHERE moto_id = mcs.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE mcs.id = $1;
	`

	err := r.db.QueryRow(ctx, q, motorcycleID, userID, r.config.IMAGE_BASE_URL).Scan(
		&motorcycle.ID, &motorcycle.Owner, &motorcycle.Engine, &motorcycle.Power, &motorcycle.Year,
		&motorcycle.NumberOfCycles, &motorcycle.Odometer, &motorcycle.Crash, &motorcycle.Wheel,
		&motorcycle.Owners, &motorcycle.VinCode, &motorcycle.Description, &motorcycle.PhoneNumbers,
		&motorcycle.Price, &motorcycle.TradeIn, &motorcycle.Status,
		&motorcycle.UpdatedAt, &motorcycle.CreatedAt, &motorcycle.MotoCategory, &motorcycle.MotoBrand,
		&motorcycle.MotoModel, &motorcycle.EngineType, &motorcycle.City, &motorcycle.Color, &motorcycle.MyMoto,
		&motorcycle.Images, &motorcycle.Videos)

	return motorcycle, err
}

func (r *MotorcycleRepository) GetEditMotorcycleByID(ctx *fasthttp.RequestCtx, motorcycleID, userID int, nameColumn string) (model.GetEditMotorcycleResponse, error) {
	var motorcycle model.GetEditMotorcycleResponse
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

	err := r.db.QueryRow(ctx, q, motorcycleID, userID, r.config.IMAGE_BASE_URL).Scan(
		&motorcycle.ID, &motorcycle.Owner, &motorcycle.Engine, &motorcycle.Power, &motorcycle.Year,
		&motorcycle.NumberOfCycles, &motorcycle.Odometer, &motorcycle.Crash, &motorcycle.Wheel, &motorcycle.New,
		&motorcycle.Owners, &motorcycle.VinCode, &motorcycle.Description, &motorcycle.PhoneNumbers,
		&motorcycle.Price, &motorcycle.TradeIn, &motorcycle.Status,
		&motorcycle.UpdatedAt, &motorcycle.CreatedAt, &motorcycle.MotoCategory, &motorcycle.MotoBrand,
		&motorcycle.MotoModel, &motorcycle.EngineType, &motorcycle.City, &motorcycle.Color, &motorcycle.MyMoto,
		&motorcycle.Images, &motorcycle.Videos)

	return motorcycle, err
}

func (r *MotorcycleRepository) BuyMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, userID int) error {
	q := `
		UPDATE motorcycles 
		SET status = 2,
			user_id = $1
		WHERE id = $2 AND status = 3 -- 3 is on sale
	`

	_, err := r.db.Exec(ctx, q, userID, motorcycleID)
	return err
}

func (r *MotorcycleRepository) DontSellMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, userID int) error {
	q := `
		UPDATE motorcycles 
		SET status = 2 -- 2 is not sale
		WHERE id = $1 AND status = 3 -- 3 is on sale
			AND user_id = $2
	`

	_, err := r.db.Exec(ctx, q, motorcycleID, userID)
	return err
}

func (r *MotorcycleRepository) SellMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID, userID int) error {
	q := `
		UPDATE motorcycles 
		SET status = 3 -- 3 is on sale
		WHERE id = $1 AND status = 2 -- 2 is not sale 
			AND user_id = $2
	`

	_, err := r.db.Exec(ctx, q, motorcycleID, userID)
	return err
}

func (r *MotorcycleRepository) CancelMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID *int) error {
	q := `
		DELETE FROM motorcycles WHERE id = $1
	`
	_, err := r.db.Exec(ctx, q, *motorcycleID)
	return err
}

func (r *MotorcycleRepository) UpdateMotorcycle(ctx *fasthttp.RequestCtx, motorcycle *model.UpdateMotorcycleRequest, userID int) error {
	// First check if the motorcycle belongs to the user
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM motorcycles WHERE id = $1 AND user_id = $2)`
	err := r.db.QueryRow(ctx, checkQuery, motorcycle.ID, userID).Scan(&exists)

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

	updateArgs = append(updateArgs, userID)

	_, err = r.db.Exec(ctx, q, updateArgs...)
	return err
}

func (r *MotorcycleRepository) DeleteMotorcycle(ctx *fasthttp.RequestCtx, motorcycleID int) error {
	q := `
		DELETE FROM motorcycles WHERE id = $1
	`

	_, err := r.db.Exec(ctx, q, motorcycleID)
	return err
}
