package repository

import (
	"dubai-auto/internal/model"
	"net/http"

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
			coordinates = $5 
		where user_id = $6
	`
	_, err := r.db.Exec(ctx, q, profile.AboutUs, profile.Whatsapp,
		profile.Telegram, profile.Address, profile.Coordinates, id)

	if err != nil {
		return model.Response{Error: err, Status: http.StatusInternalServerError}
	}

	return model.Response{Data: model.Success{Message: "Profile updated successfully"}}
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
			created_at
		from profiles where user_id = $1
	`
	var profile model.ThirdPartyGetProfileRes
	err := r.db.QueryRow(ctx, q, id).Scan(
		&profile.AboutUs, &profile.Whatsapp,
		&profile.Telegram, &profile.Address,
		&profile.Coordinates, &profile.Avatar, &profile.Banner, &profile.Registered)

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
