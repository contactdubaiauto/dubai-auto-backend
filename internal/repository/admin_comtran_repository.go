package repository

import (
	"context"

	"dubai-auto/internal/model"
)

func (r *AdminRepository) GetComtrans(ctx context.Context, limit, lastID int) ([]model.AdminComtranListItem, error) {
	list := make([]model.AdminComtranListItem, 0)
	q := `
		SELECT
			cts.id,
			cbs.name AS brand,
			cms.name AS model,
			cts.description,
			cts.price,
			cts.status,
			u.phone AS user_phone,
			images.images,
			pf.username AS user_name,
			CASE
				WHEN pf.avatar IS NULL OR pf.avatar = '' THEN NULL
				ELSE $1 || pf.avatar
			END AS user_avatar
		FROM comtrans cts
		LEFT JOIN com_brands cbs ON cts.comtran_brand_id = cbs.id
		LEFT JOIN com_models cms ON cts.comtran_model_id = cms.id
		LEFT JOIN users u ON u.id = cts.user_id
		LEFT JOIN profiles pf ON pf.user_id = cts.user_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $1 || image AS image
				FROM comtran_images
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) img
		) images ON true
		WHERE cts.id > $2
		ORDER BY cts.id DESC
		LIMIT $3
	`
	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL, lastID, limit)
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var v model.AdminComtranListItem
		if err := rows.Scan(
			&v.ID,
			&v.Brand,
			&v.Model,
			&v.Description,
			&v.Price,
			&v.Status,
			&v.UserPhone,
			&v.Images,
			&v.UserName,
			&v.UserAvatar,
		); err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, err
}

func (r *AdminRepository) GetComtranByID(ctx context.Context, id int) (model.GetComtranResponse, error) {
	var com model.GetComtranResponse
	q := `
		SELECT
			cts.id,
			json_build_object(
				'id', pf.user_id,
				'username', pf.username,
				'avatar', $2 || pf.avatar,
				'contacts', pf.contacts
			) AS owner,
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
			cts.status::text,
			cts.updated_at,
			cts.created_at,
			cocs.name AS comtran_category,
			cbs.name AS comtran_brand,
			cms.name AS comtran_model,
			ces.name AS engine_type,
			cs.name AS city,
			cls.name AS color,
			false AS my_comtrans,
			images.images,
			videos.videos
		FROM comtrans cts
		LEFT JOIN profiles pf ON pf.user_id = cts.user_id
		LEFT JOIN com_categories cocs ON cocs.id = cts.comtran_category_id
		LEFT JOIN com_brands cbs ON cbs.id = cts.comtran_brand_id
		LEFT JOIN com_models cms ON cms.id = cts.comtran_model_id
		LEFT JOIN com_engines ces ON ces.id = cts.engine_id
		LEFT JOIN cities cs ON cs.id = cts.city_id
		LEFT JOIN colors cls ON cls.id = cts.color_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $2 || image AS image
				FROM comtran_images
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT $2 || video AS video
				FROM comtran_videos
				WHERE comtran_id = cts.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE cts.id = $1
	`
	err := r.db.QueryRow(ctx, q, id, r.config.IMAGE_BASE_URL).Scan(
		&com.ID, &com.Owner, &com.Engine, &com.Power, &com.Year,
		&com.Odometer, &com.Crash, &com.Owners, &com.VinCode, &com.Description,
		&com.PhoneNumbers, &com.Price, &com.TradeIn, &com.Status,
		&com.UpdatedAt, &com.CreatedAt, &com.ComtranCategory, &com.ComtranBrand,
		&com.ComtranModel, &com.EngineType, &com.City, &com.Color, &com.MyComtrans,
		&com.Images, &com.Videos,
	)
	return com, err
}

func (r *AdminRepository) DeleteComtran(ctx context.Context, id int) error {
	q := `DELETE FROM comtrans WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}
