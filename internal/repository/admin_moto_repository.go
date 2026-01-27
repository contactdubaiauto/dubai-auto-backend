package repository

import (
	"context"
	"fmt"

	"dubai-auto/internal/model"
)

func (r *AdminRepository) GetMotorcycles(ctx context.Context, limit, lastID, moderationStatus int) ([]model.AdminMotoListItem, error) {
	qWhere := ""

	if moderationStatus != 0 {
		qWhere = fmt.Sprintf(" AND m.moderation_status = %d ", moderationStatus)
	}

	list := make([]model.AdminMotoListItem, 0)
	q := `
		SELECT
			m.id,
			mb.name AS brand,
			mm.name AS model,
			m.description,
			m.price,
			m.status,
			u.phone AS user_phone,
			images.images,
			pf.username AS user_name,
			CASE
				WHEN pf.avatar IS NULL OR pf.avatar = '' THEN NULL
				ELSE $1 || pf.avatar
			END AS user_avatar,
			m.moderation_status,
			u.role_id
		FROM motorcycles m
		LEFT JOIN moto_brands mb ON m.moto_brand_id = mb.id
		LEFT JOIN moto_models mm ON m.moto_model_id = mm.id
		LEFT JOIN users u ON u.id = m.user_id
		LEFT JOIN profiles pf ON pf.user_id = m.user_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $1 || image AS image
				FROM moto_images
				WHERE moto_id = m.id
				ORDER BY created_at DESC
			) img
		) images ON true
		WHERE m.id > $2 ` + qWhere + `
		ORDER BY m.id DESC
		LIMIT $3
	`
	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL, lastID, limit)

	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var v model.AdminMotoListItem
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
			&v.ModerationStatus,
			&v.UserRoleID,
		); err != nil {
			return list, err
		}
		list = append(list, v)
	}
	return list, err
}

func (r *AdminRepository) GetMotorcycleByID(ctx context.Context, id int) (model.GetMotorcycleResponse, error) {
	var m model.GetMotorcycleResponse
	q := `
		SELECT
			m.id,
			json_build_object(
				'id', u.id,
				'username', u.username,
				'avatar', $2 || pf.avatar,
				'contacts', pf.contacts,
				'role_id', u.role_id
			) as owner,
			m.engine,
			m.power,
			m.year,
			noc.name AS number_of_cycles,
			m.odometer,
			m.crash,
			m.wheel,
			m.owners,
			m.vin_code,
			m.description,
			m.phone_numbers,
			m.price,
			m.trade_in,
			m.status::text,
			m.updated_at,
			m.created_at,
			mc.name AS moto_category,
			mb.name AS moto_brand,
			mm.name AS moto_model,
			me.name AS engine_type,
			c.name AS city,
			cl.name AS color,
			false AS my_moto,
			images.images,
			videos.videos,
			m.moderation_status,
			u.role_id
		FROM motorcycles m
		LEFT JOIN users u ON u.id = m.user_id
		LEFT JOIN profiles pf ON pf.user_id = m.user_id
		LEFT JOIN moto_categories mc ON mc.id = m.moto_category_id
		LEFT JOIN moto_brands mb ON mb.id = m.moto_brand_id
		LEFT JOIN moto_models mm ON mm.id = m.moto_model_id
		LEFT JOIN moto_engines me ON me.id = m.engine_id
		LEFT JOIN cities c ON c.id = m.city_id
		LEFT JOIN colors cl ON cl.id = m.color_id
		LEFT JOIN number_of_cycles noc ON noc.id = m.number_of_cycles_id
		LEFT JOIN LATERAL (
			SELECT json_agg(img.image) AS images
			FROM (
				SELECT $2 || image AS image
				FROM moto_images
				WHERE moto_id = m.id
				ORDER BY created_at DESC
			) img
		) images ON true
		LEFT JOIN LATERAL (
			SELECT json_agg(v.video) AS videos
			FROM (
				SELECT $2 || video AS video
				FROM moto_videos
				WHERE moto_id = m.id
				ORDER BY created_at DESC
			) v
		) videos ON true
		WHERE m.id = $1
	`
	err := r.db.QueryRow(ctx, q, id, r.config.IMAGE_BASE_URL).Scan(
		&m.ID, &m.Owner, &m.Engine, &m.Power, &m.Year,
		&m.NumberOfCycles, &m.Odometer, &m.Crash, &m.Wheel,
		&m.Owners, &m.VinCode, &m.Description, &m.PhoneNumbers,
		&m.Price, &m.TradeIn, &m.Status,
		&m.UpdatedAt, &m.CreatedAt, &m.MotoCategory, &m.MotoBrand,
		&m.MotoModel, &m.EngineType, &m.City, &m.Color, &m.MyMoto,
		&m.Images, &m.Videos,
		&m.ModerationStatus,
		&m.UserRoleID,
	)
	return m, err
}

func (r *AdminRepository) DeleteMotorcycle(ctx context.Context, id int) error {
	q := `DELETE FROM motorcycles WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// ModerateMotorcycle updates the moderation_status of a motorcycle and returns the user_id
func (r *AdminRepository) ModerateMotorcycle(ctx context.Context, id int, status int) (int, error) {
	q := `
		UPDATE motorcycles
		SET moderation_status = $2, updated_at = now()
		WHERE id = $1
		RETURNING user_id
	`
	var userID int
	err := r.db.QueryRow(ctx, q, id, status).Scan(&userID)
	return userID, err
}
