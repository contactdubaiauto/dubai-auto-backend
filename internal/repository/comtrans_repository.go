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

type ComtransRepository struct {
	config *config.Config
	db     *pgxpool.Pool
}

func NewComtransRepository(config *config.Config, db *pgxpool.Pool) *ComtransRepository {
	return &ComtransRepository{config, db}
}

func (r *ComtransRepository) GetComtransCategories(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetComtransCategoriesResponse, error) {
	data := make([]model.GetComtransCategoriesResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM com_categories
	`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category model.GetComtransCategoriesResponse
		err = rows.Scan(&category.ID, &category.Name)

		if err != nil {
			return nil, err
		}

		data = append(data, category)
	}

	return data, nil
}

func (r *ComtransRepository) GetComtransBrands(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetComtransBrandsResponse, error) {
	data := make([]model.GetComtransBrandsResponse, 0)
	q := `
		SELECT id, ` + nameColumn + `, $1 || logo, model_count FROM com_brands
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand model.GetComtransBrandsResponse
		err = rows.Scan(&brand.ID, &brand.Name, &brand.Image, &brand.ModelCount)

		if err != nil {
			return nil, err
		}

		data = append(data, brand)
	}

	return data, nil
}

func (r *ComtransRepository) GetComtransEngines(ctx *fasthttp.RequestCtx, nameColumn string) ([]model.GetComtransModelsResponse, error) {
	data := make([]model.GetComtransModelsResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM com_engines
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var engine model.GetComtransModelsResponse
		err = rows.Scan(&engine.ID, &engine.Name)
	}
	return data, nil
}

func (r *ComtransRepository) GetComtransModelsByBrandID(ctx *fasthttp.RequestCtx, brandID, nameColumn string) ([]model.GetComtransModelsResponse, error) {
	data := make([]model.GetComtransModelsResponse, 0)
	q := `
		SELECT id, ` + nameColumn + ` FROM com_models
		WHERE comtran_brand_id = $1
	`

	rows, err := r.db.Query(ctx, q, brandID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var model model.GetComtransModelsResponse
		err = rows.Scan(&model.ID, &model.Name)

		if err != nil {
			return nil, err
		}

		data = append(data, model)
	}

	return data, nil
}

func (r *ComtransRepository) CreateComtrans(ctx *fasthttp.RequestCtx, req model.CreateComtransRequest, userID int) (model.SuccessWithId, error) {
	data := model.SuccessWithId{}

	keys, values, args := auth.BuildParams(req)

	q := `
		INSERT INTO comtrans ( 
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

	data.Message = "Commercial transport created successfully"
	data.Id = id

	return data, err
}

func (r *ComtransRepository) GetComtrans(ctx *fasthttp.RequestCtx, userID int, targetUserID string, brands, models, regions, cities,
	generations, transmissions, engines, drivetrains, body_types, fuel_types, ownership_types, colors, dealers []string,
	year_from, year_to, credit, price_from, price_to, tradeIn, owners, crash, odometer string,
	new, wheel *bool, limit, lastID int, nameColumn string) ([]model.GetComtransResponse, error) {
	data := make([]model.GetComtransResponse, 0)
	var qWhere string
	var qValues []any
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

	if len(generations) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND gms.generation_id = ANY($%d)", i)
		qValues = append(qValues, generations)
	}

	if len(colors) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.color_id = ANY($%d)", i)
		qValues = append(qValues, colors)
	}

	if len(dealers) > 0 {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.user_id = ANY($%d)", i)
		qValues = append(qValues, dealers)
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

	if tradeIn != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.trade_in = $%d", i)
		qValues = append(qValues, tradeIn)
	}

	if owners != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.owners = $%d", i)
		qValues = append(qValues, owners)
	}

	if crash != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.crash = $%d", i)
		qValues = append(qValues, crash)
	}

	if credit != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.credit = $%d", i)
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

	if new != nil {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.new = $%d", i)
		qValues = append(qValues, new)
	}

	if wheel != nil {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.wheel = $%d", i)
		qValues = append(qValues, wheel)
	}

	if odometer != "" {
		i += 1
		qWhere += fmt.Sprintf(" AND vs.odometer <= $%d", i)
		qValues = append(qValues, odometer)
	}

	q := `
		select 
			cts.id,
			'comtran' as type,
			cbs.` + nameColumn + ` as brand,
			cms.` + nameColumn + ` as model,
			cts.year,
			cts.price,
			cts.created_at,
			images.images,
			cts.new,
			cts.status,
			cts.trade_in,
			cts.crash,
			cts.view_count,
			CASE
				WHEN cts.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_comtrans,
			cts.odometer,
			CASE
				WHEN u.role_id = 2 THEN u.username
				ELSE NULL
			END AS owner_name
		from comtrans cts
		left join com_brands cbs on cbs.id = cts.comtran_brand_id
		left join com_models cms on cms.id = cts.comtran_model_id
		left join users u on u.id = cts.user_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $1 || image as image
				FROM comtran_images
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) img
		) images ON true;

	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL, userID)
	if err != nil {
		return data, err
	}

	defer rows.Close()
	fmt.Println("userID", userID)
	for rows.Next() {
		var com model.GetComtransResponse
		err = rows.Scan(
			&com.ID, &com.Type, &com.Brand, &com.Model, &com.Year, &com.Price,
			&com.CreatedAt, &com.Images,
			&com.New, &com.Status, &com.TradeIn, &com.Crash,
			&com.ViewCount, &com.MyComtran, &com.Odometer, &com.OwnerName)

		if err != nil {
			return data, err
		}

		data = append(data, com)
	}

	return data, err
}

func (r *ComtransRepository) CreateComtransImages(ctx *fasthttp.RequestCtx, comtransID int, images []string) error {

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

func (r *ComtransRepository) CreateComtransVideos(ctx *fasthttp.RequestCtx, comtransID int, video string) error {

	q := `
		INSERT INTO comtran_videos (comtran_id, video) VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, q, comtransID, video)
	if err != nil {
		return err
	}

	return err
}

func (r *ComtransRepository) DeleteComtransImage(ctx *fasthttp.RequestCtx, comtransID int, imageID int) error {
	q := `
		DELETE FROM comtran_images WHERE comtran_id = $1 AND id = $2
	`

	_, err := r.db.Exec(ctx, q, comtransID, imageID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ComtransRepository) DeleteComtransVideo(ctx *fasthttp.RequestCtx, comtransID int, videoID int) error {
	q := `
		DELETE FROM comtran_videos WHERE comtran_id = $1 AND id = $2
	`

	_, err := r.db.Exec(ctx, q, comtransID, videoID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ComtransRepository) GetComtransByID(ctx *fasthttp.RequestCtx, comtransID, userID int, nameColumn string) (model.GetComtranResponse, error) {
	var comtrans model.GetComtranResponse
	q := `
		WITH updated AS (
			UPDATE comtrans
			SET view_count = COALESCE(view_count, 0) + 1
			WHERE id = $1
			RETURNING *
		)
		select 
			cts.id,
			json_build_object(
				'id', pf.user_id,
				'username', pf.username,
				'avatar', $3 || pf.avatar,
				'contacts', pf.contacts
			) as owner,
			cts.engine,
			cts.power,
			cts.year,
			cts.odometer,
			cts.crash,
			cts.owners,
			cts.vin_code,
			cts.description,
			cts.phone_numbers,
			cts.price,
			cts.trade_in,
			cts.status,
			cts.updated_at,
			cts.created_at,
			cocs.` + nameColumn + ` as comtran_category,
			cbs.` + nameColumn + ` as comtran_brand,
			cms.` + nameColumn + ` as comtran_model,
			ces.` + nameColumn + ` as engine_type,
			cs.name as city,
			cls.` + nameColumn + ` as color,
			CASE
				WHEN cts.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_comtrans,
			images.images,
			videos.videos
		from updated cts
		left join profiles pf on pf.user_id = cts.user_id
		left join com_categories cocs on cocs.id = cts.comtran_category_id
		left join com_brands cbs on cbs.id = cts.comtran_brand_id
		left join com_models cms on cms.id = cts.comtran_model_id
		left join com_engines ces on ces.id = cts.engine_id
		left join cities cs on cs.id = cts.city_id
		left join colors cls on cls.id = cts.color_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $3 || image as image
				FROM comtran_images
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT $3 || video as video
				FROM comtran_videos
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE cts.id = $1;
	`

	err := r.db.QueryRow(ctx, q, comtransID, userID, r.config.IMAGE_BASE_URL).Scan(
		&comtrans.ID, &comtrans.Owner, &comtrans.Engine, &comtrans.Power, &comtrans.Year,
		&comtrans.Odometer, &comtrans.Crash, &comtrans.Owners, &comtrans.VinCode, &comtrans.Description,
		&comtrans.PhoneNumbers, &comtrans.Price, &comtrans.TradeIn, &comtrans.Status,
		&comtrans.UpdatedAt, &comtrans.CreatedAt, &comtrans.ComtranCategory, &comtrans.ComtranBrand,
		&comtrans.ComtranModel, &comtrans.EngineType, &comtrans.City, &comtrans.Color, &comtrans.MyComtrans,
		&comtrans.Images, &comtrans.Videos)

	return comtrans, err
}

func (r *ComtransRepository) GetEditComtransByID(ctx *fasthttp.RequestCtx, comtransID, userID int, nameColumn string) (model.GetEditComtransResponse, error) {
	var comtrans model.GetEditComtransResponse
	q := `
		select 
			cts.id,
			json_build_object(
				'id', pf.user_id,
				'username', pf.username,
				'avatar', $3 || pf.avatar,
				'contacts', pf.contacts
			) as owner,
			cts.engine,
			cts.power,
			cts.year,
			cts.odometer,
			cts.crash,
			cts.owners,
			cts.vin_code,
			cts.description,
			cts.phone_numbers,
			cts.price,
			cts.trade_in,
			cts.status,
			cts.updated_at,
			cts.created_at,
			cocs.` + nameColumn + ` as comtran_category,
			cbs.` + nameColumn + ` as comtran_brand,
			cms.` + nameColumn + ` as comtran_model,
			ces.` + nameColumn + ` as engine_type,
			cs.name as city,
			cls.` + nameColumn + ` as color,
			CASE
				WHEN cts.user_id = $2 THEN TRUE
				ELSE FALSE
			END AS my_comtrans,
			images.images,
			videos.videos
		from comtrans cts
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

	err := r.db.QueryRow(ctx, q, comtransID, userID, r.config.IMAGE_BASE_URL).Scan(
		&comtrans.ID, &comtrans.Owner, &comtrans.Engine, &comtrans.Power, &comtrans.Year,
		&comtrans.Odometer, &comtrans.Crash, &comtrans.Owners, &comtrans.VinCode, &comtrans.Description,
		&comtrans.PhoneNumbers, &comtrans.Price, &comtrans.TradeIn, &comtrans.Status,
		&comtrans.UpdatedAt, &comtrans.CreatedAt, &comtrans.ComtranCategory, &comtrans.ComtranBrand,
		&comtrans.ComtranModel, &comtrans.EngineType, &comtrans.City, &comtrans.Color, &comtrans.MyComtrans,
		&comtrans.Images, &comtrans.Videos)

	return comtrans, err
}

func (r *ComtransRepository) BuyComtrans(ctx *fasthttp.RequestCtx, comtransID, userID int) error {
	q := `
		UPDATE comtrans 
		SET status = 2,
			user_id = $1
		WHERE id = $2 AND status = 3 -- 3 is on sale
	`

	_, err := r.db.Exec(ctx, q, userID, comtransID)
	return err
}

func (r *ComtransRepository) DontSellComtrans(ctx *fasthttp.RequestCtx, comtransID, userID int) error {
	q := `
		UPDATE comtrans 
		SET status = 2 -- 2 is not sale
		WHERE id = $1 AND status = 3 -- 3 is on sale
			AND user_id = $2
	`

	_, err := r.db.Exec(ctx, q, comtransID, userID)
	return err
}

func (r *ComtransRepository) SellComtrans(ctx *fasthttp.RequestCtx, comtransID, userID int) error {
	q := `
		UPDATE comtrans 
		SET status = 3 -- 3 is on sale
		WHERE id = $1 AND status = 2 -- 2 is not sale 
			AND user_id = $2
	`

	_, err := r.db.Exec(ctx, q, comtransID, userID)
	return err
}

func (r *ComtransRepository) CancelComtrans(ctx *fasthttp.RequestCtx, comtransID *int) error {
	q := `
		DELETE FROM comtrans WHERE id = $1
	`
	_, err := r.db.Exec(ctx, q, *comtransID)
	return err
}

func (r *ComtransRepository) UpdateComtrans(ctx *fasthttp.RequestCtx, comtrans *model.UpdateComtransRequest, userID int) error {
	// First check if the comtrans belongs to the user
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM comtrans WHERE id = $1 AND user_id = $2)`
	err := r.db.QueryRow(ctx, checkQuery, comtrans.ID, userID).Scan(&exists)

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
		if key != "id" && key != "user_id" && key != "parameters" {
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

	updateArgs = append(updateArgs, userID)

	_, err = r.db.Exec(ctx, q, updateArgs...)
	return err
}

func (r *ComtransRepository) DeleteComtrans(ctx *fasthttp.RequestCtx, comtransID int) error {
	q := `
		DELETE FROM comtrans WHERE id = $1
	`

	_, err := r.db.Exec(ctx, q, comtransID)
	return err
}
