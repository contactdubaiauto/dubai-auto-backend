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
			created_at,
			message
		from profiles where user_id = $1
	`
	var profile model.ThirdPartyGetProfileRes
	err := r.db.QueryRow(ctx, q, id).Scan(
		&profile.AboutUs, &profile.Whatsapp,
		&profile.Telegram, &profile.Address,
		&profile.Coordinates, &profile.Avatar, &profile.Banner, &profile.Registered, &profile.Message)

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

func (r *ThirdPartyRepository) CreateBannerImage(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `
		update profiles set banner = $1 where user_id = $2
	`
	_, err := r.db.Exec(ctx, q, paths[0], id)

	return err
}

func (r *ThirdPartyRepository) CreateDealerCar(ctx *fasthttp.RequestCtx, car *model.CreateCarRequest, dealerID int) (int, error) {
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
