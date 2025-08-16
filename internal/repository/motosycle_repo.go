package repository

import (
	"dubai-auto/internal/model"
	"dubai-auto/pkg/auth"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"
)

type MotorcycleRepository struct {
	db *pgxpool.Pool
}

func NewMotorcycleRepository(db *pgxpool.Pool) *MotorcycleRepository {
	return &MotorcycleRepository{db}
}

func (r *MotorcycleRepository) GetMotorcycleCategories(ctx *fasthttp.RequestCtx) (data []model.GetMotorcycleCategoriesResponse, err error) {
	q := `
		SELECT id, name FROM moto_categories
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

func (r *MotorcycleRepository) GetMotorcycleParameters(ctx *fasthttp.RequestCtx, categoryID string) (data []model.GetMotorcycleParametersResponse, err error) {
	q := `
		SELECT 
			moto_parameters.id,
			moto_parameters.name,
			json_agg(
				json_build_object(
					'id', moto_parameter_values.id,
					'name', moto_parameter_values.name
				)
			) as values
		FROM moto_parameters
		LEFT JOIN moto_parameter_values ON moto_parameters.id = moto_parameter_values.moto_parameter_id
		WHERE moto_parameters.moto_category_id = $1
		GROUP BY moto_parameters.id, moto_parameters.name
	`

	rows, err := r.db.Query(ctx, q, categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var parameter model.GetMotorcycleParametersResponse
		err = rows.Scan(&parameter.ID, &parameter.Name, &parameter.Values)

		if err != nil {
			return nil, err
		}

		data = append(data, parameter)
	}

	return data, nil
}

func (r *MotorcycleRepository) GetMotorcycleBrands(ctx *fasthttp.RequestCtx, categoryID string) (data []model.GetMotorcycleBrandsResponse, err error) {
	q := `
		SELECT id, name, image FROM moto_brands
		WHERE moto_category_id = $1
	`

	rows, err := r.db.Query(ctx, q, categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand model.GetMotorcycleBrandsResponse
		err = rows.Scan(&brand.ID, &brand.Name, &brand.Image)

		if err != nil {
			return nil, err
		}

		data = append(data, brand)
	}

	return data, nil
}

func (r *MotorcycleRepository) GetMotorcycleModelsByBrandID(ctx *fasthttp.RequestCtx, categoryID string, brandID string) (data []model.GetMotorcycleModelsResponse, err error) {
	q := `
		SELECT id, name FROM moto_models
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

func (r *MotorcycleRepository) CreateMotorcycle(ctx *fasthttp.RequestCtx, req model.CreateMotorcycleRequest, userID int) (data model.SuccessWithId, err error) {
	parameters := req.Parameters
	req.Parameters = nil

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
	fmt.Println("user_id", userID)
	var id int
	args = append(args, userID)
	err = r.db.QueryRow(ctx, q, args...).Scan(&id)

	if err != nil {
		return data, err
	}

	for i := range parameters {

		q := `
			INSERT INTO motorcycle_parameters (motorcycle_id, moto_parameter_id, moto_parameter_value_id)
			VALUES ($1, $2, $3)
		`
		_, err = r.db.Exec(ctx, q, id, parameters[i].ParameterID, parameters[i].ValueID)

		if err != nil {
			return data, err
		}
	}

	data.Message = "Motorcycle created successfully"
	data.Id = id

	return data, err
}
