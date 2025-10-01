package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
)

type AdminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{db}
}

// Application CRUD operations
func (r *AdminRepository) GetApplications(ctx *fasthttp.RequestCtx) ([]model.AdminApplicationResponse, error) {
	applications := make([]model.AdminApplicationResponse, 0)
	q := `
			SELECT 
				id, 
				company_name, 
				licence_issue_date, 
				licence_expiry_date, 
				username, 
				email, 
				phone, 
				status, 
				created_at 
			FROM temp_users 
			ORDER BY id DESC
		`
	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return applications, err
	}

	defer rows.Close()

	for rows.Next() {
		var application model.AdminApplicationResponse

		if err := rows.Scan(&application.ID, &application.CompanyName, &application.LicenceIssueDate,
			&application.LicenceExpiryDate, &application.FullName, &application.Email,
			&application.Phone, &application.Status, &application.CreatedAt); err != nil {
			return applications, err
		}

		applications = append(applications, application)
	}

	return applications, err
}

func (r *AdminRepository) GetApplication(ctx *fasthttp.RequestCtx, id int) (model.AdminApplicationResponse, error) {

	q := `SELECT id, company_name, licence_issue_date, licence_expiry_date, username, email, phone, status, created_at FROM temp_users WHERE id = $1`
	var application model.AdminApplicationResponse
	rows, err := r.db.Query(ctx, q, id)

	if err != nil {
		return model.AdminApplicationResponse{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var application model.AdminApplicationResponse
		if err := rows.Scan(&application.ID, &application.CompanyName, &application.LicenceIssueDate, &application.LicenceExpiryDate, &application.FullName, &application.Email, &application.Phone, &application.Status, &application.CreatedAt); err != nil {
			return model.AdminApplicationResponse{}, err
		}
	}

	return application, nil

}

func (r *AdminRepository) AcceptApplication(ctx *fasthttp.RequestCtx, id int) error {
	q := `UPDATE temp_users SET status = 1 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *AdminRepository) RejectApplication(ctx *fasthttp.RequestCtx, id int) error {
	q := `UPDATE temp_users SET status = 2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Cities CRUD operations
func (r *AdminRepository) GetCities(ctx *fasthttp.RequestCtx) ([]model.AdminCityResponse, error) {
	cities := make([]model.AdminCityResponse, 0)
	q := `SELECT id, name, created_at FROM cities ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return cities, err
	}

	defer rows.Close()

	for rows.Next() {
		var city model.AdminCityResponse
		if err := rows.Scan(&city.ID, &city.Name, &city.CreatedAt); err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}

	return cities, err
}

func (r *AdminRepository) CreateCity(ctx *fasthttp.RequestCtx, req *model.CreateNameRequest) (int, error) {
	q := `INSERT INTO cities (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateCity(ctx *fasthttp.RequestCtx, id int, req *model.CreateNameRequest) error {
	q := `UPDATE cities SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteCity(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM cities WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Brands CRUD operations
func (r *AdminRepository) GetBrands(ctx *fasthttp.RequestCtx) ([]model.AdminBrandResponse, error) {
	brands := make([]model.AdminBrandResponse, 0)
	q := `SELECT id, name, logo, model_count, popular, updated_at FROM brands ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return brands, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand model.AdminBrandResponse
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.Logo, &brand.ModelCount, &brand.Popular, &brand.UpdatedAt); err != nil {
			return brands, err
		}
		brands = append(brands, brand)
	}

	return brands, err
}

func (r *AdminRepository) CreateBrand(ctx *fasthttp.RequestCtx, req *model.CreateBrandRequest) (int, error) {
	q := `INSERT INTO brands (name, popular, updated_at) VALUES ($1, $2, NOW()) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Popular).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateBrand(ctx *fasthttp.RequestCtx, id int, req *model.CreateBrandRequest) error {
	q := `UPDATE brands SET name = $2, popular = $3, updated_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Popular)
	return err
}

func (r *AdminRepository) DeleteBrand(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM brands WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Models CRUD operations
func (r *AdminRepository) GetModels(ctx *fasthttp.RequestCtx, brand_id int) ([]model.AdminModelResponse, error) {
	models := make([]model.AdminModelResponse, 0)
	q := `
		SELECT m.id, m.name, m.brand_id, b.name as brand_name, m.popular, m.updated_at 
		FROM models m
		LEFT JOIN brands b ON m.brand_id = b.id
		WHERE m.brand_id = $1
		ORDER BY m.id DESC
	`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return models, err
	}

	defer rows.Close()

	for rows.Next() {
		var modelItem model.AdminModelResponse
		if err := rows.Scan(&modelItem.ID, &modelItem.Name, &modelItem.BrandID, &modelItem.BrandName, &modelItem.Popular, &modelItem.UpdatedAt); err != nil {
			return models, err
		}
		models = append(models, modelItem)
	}

	return models, err
}

func (r *AdminRepository) CreateModel(ctx *fasthttp.RequestCtx, brand_id int, req *model.CreateModelRequest) (int, error) {
	q := `INSERT INTO models (name, brand_id, popular, updated_at) VALUES ($1, $2, $3, NOW()) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, brand_id, req.Popular).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateModelRequest) error {
	q := `UPDATE models SET name = $2, brand_id = $3, popular = $4, updated_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.BrandID, req.Popular)
	return err
}

func (r *AdminRepository) DeleteModel(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM models WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Body Types CRUD operations
func (r *AdminRepository) GetBodyTypes(ctx *fasthttp.RequestCtx) ([]model.AdminBodyTypeResponse, error) {
	bodyTypes := make([]model.AdminBodyTypeResponse, 0)
	q := `SELECT id, name, image, created_at FROM body_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return bodyTypes, err
	}

	defer rows.Close()

	for rows.Next() {
		var bodyType model.AdminBodyTypeResponse
		if err := rows.Scan(&bodyType.ID, &bodyType.Name, &bodyType.Image, &bodyType.CreatedAt); err != nil {
			return bodyTypes, err
		}
		bodyTypes = append(bodyTypes, bodyType)
	}

	return bodyTypes, err
}

func (r *AdminRepository) CreateBodyType(ctx *fasthttp.RequestCtx, req *model.CreateBodyTypeRequest) (int, error) {
	q := `INSERT INTO body_types (name, image) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Image).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateBodyTypeImage(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `INSERT INTO body_types_images (body_type_id, image) VALUES ($1, $2) RETURNING id`
	_, err := r.db.Exec(ctx, q, id, paths)
	return err
}

func (r *AdminRepository) DeleteBodyTypeImage(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM body_types_images WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *AdminRepository) UpdateBodyType(ctx *fasthttp.RequestCtx, id int, req *model.CreateBodyTypeRequest) error {
	q := `UPDATE body_types SET name = $2, image = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Image)
	return err
}

func (r *AdminRepository) DeleteBodyType(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM body_types WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Regions CRUD operations
func (r *AdminRepository) GetRegions(ctx *fasthttp.RequestCtx, cityID int) ([]model.AdminCityResponse, error) {
	regions := make([]model.AdminCityResponse, 0)
	q := `SELECT id, name, created_at FROM regions where city_id = $1`

	rows, err := r.db.Query(ctx, q, cityID)

	if err != nil {
		return regions, err
	}

	defer rows.Close()

	for rows.Next() {
		var region model.AdminCityResponse
		if err := rows.Scan(&region.ID, &region.Name, &region.CreatedAt); err != nil {
			return regions, err
		}
		regions = append(regions, region)
	}

	return regions, err
}

func (r *AdminRepository) CreateRegion(ctx *fasthttp.RequestCtx, city_id int, req *model.CreateNameRequest) (int, error) {
	q := `INSERT INTO regions (name, city_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, city_id).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateRegion(ctx *fasthttp.RequestCtx, id int, req *model.CreateNameRequest) error {
	q := `UPDATE regions SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteRegion(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM regions WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Transmissions CRUD operations
func (r *AdminRepository) GetTransmissions(ctx *fasthttp.RequestCtx) ([]model.AdminTransmissionResponse, error) {
	transmissions := make([]model.AdminTransmissionResponse, 0)
	q := `SELECT id, name, created_at FROM transmissions ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return transmissions, err
	}
	defer rows.Close()

	for rows.Next() {
		var transmission model.AdminTransmissionResponse
		if err := rows.Scan(&transmission.ID, &transmission.Name, &transmission.CreatedAt); err != nil {
			return transmissions, err
		}
		transmissions = append(transmissions, transmission)
	}

	return transmissions, err
}

func (r *AdminRepository) CreateTransmission(ctx *fasthttp.RequestCtx, req *model.CreateTransmissionRequest) (int, error) {
	q := `INSERT INTO transmissions (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateTransmission(ctx *fasthttp.RequestCtx, id int, req *model.CreateTransmissionRequest) error {
	q := `UPDATE transmissions SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteTransmission(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM transmissions WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Engines CRUD operations
func (r *AdminRepository) GetEngines(ctx *fasthttp.RequestCtx) ([]model.AdminEngineResponse, error) {
	engines := make([]model.AdminEngineResponse, 0)
	q := `SELECT id, value, created_at FROM engines ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return engines, err
	}
	defer rows.Close()

	for rows.Next() {
		var engine model.AdminEngineResponse
		if err := rows.Scan(&engine.ID, &engine.Value, &engine.CreatedAt); err != nil {
			return engines, err
		}
		engines = append(engines, engine)
	}

	return engines, err
}

func (r *AdminRepository) CreateEngine(ctx *fasthttp.RequestCtx, req *model.CreateEngineRequest) (int, error) {
	q := `INSERT INTO engines (value) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Value).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateEngine(ctx *fasthttp.RequestCtx, id int, req *model.CreateEngineRequest) error {
	q := `UPDATE engines SET value = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Value)
	return err
}

func (r *AdminRepository) DeleteEngine(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM engines WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Drivetrains CRUD operations
func (r *AdminRepository) GetDrivetrains(ctx *fasthttp.RequestCtx) ([]model.AdminDrivetrainResponse, error) {
	drivetrains := make([]model.AdminDrivetrainResponse, 0)
	q := `SELECT id, name, created_at FROM drivetrains ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return drivetrains, err
	}
	defer rows.Close()

	for rows.Next() {
		var drivetrain model.AdminDrivetrainResponse
		if err := rows.Scan(&drivetrain.ID, &drivetrain.Name, &drivetrain.CreatedAt); err != nil {
			return drivetrains, err
		}
		drivetrains = append(drivetrains, drivetrain)
	}

	return drivetrains, err
}

func (r *AdminRepository) CreateDrivetrain(ctx *fasthttp.RequestCtx, req *model.CreateDrivetrainRequest) (int, error) {
	q := `INSERT INTO drivetrains (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateDrivetrain(ctx *fasthttp.RequestCtx, id int, req *model.CreateDrivetrainRequest) error {
	q := `UPDATE drivetrains SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteDrivetrain(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM drivetrains WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Fuel Types CRUD operations
func (r *AdminRepository) GetFuelTypes(ctx *fasthttp.RequestCtx) ([]model.AdminFuelTypeResponse, error) {
	fuelTypes := make([]model.AdminFuelTypeResponse, 0)
	q := `SELECT id, name, created_at FROM fuel_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return fuelTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var fuelType model.AdminFuelTypeResponse
		if err := rows.Scan(&fuelType.ID, &fuelType.Name, &fuelType.CreatedAt); err != nil {
			return fuelTypes, err
		}
		fuelTypes = append(fuelTypes, fuelType)
	}

	return fuelTypes, err
}

func (r *AdminRepository) CreateFuelType(ctx *fasthttp.RequestCtx, req *model.CreateFuelTypeRequest) (int, error) {
	q := `INSERT INTO fuel_types (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateFuelType(ctx *fasthttp.RequestCtx, id int, req *model.CreateFuelTypeRequest) error {
	q := `UPDATE fuel_types SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteFuelType(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM fuel_types WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Service Types CRUD operations
func (r *AdminRepository) GetServiceTypes(ctx *fasthttp.RequestCtx) ([]model.AdminServiceTypeResponse, error) {
	serviceTypes := make([]model.AdminServiceTypeResponse, 0)
	q := `SELECT id, name, created_at FROM service_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return serviceTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var serviceType model.AdminServiceTypeResponse
		if err := rows.Scan(&serviceType.ID, &serviceType.Name, &serviceType.CreatedAt); err != nil {
			return serviceTypes, err
		}
		serviceTypes = append(serviceTypes, serviceType)
	}

	return serviceTypes, err
}

func (r *AdminRepository) CreateServiceType(ctx *fasthttp.RequestCtx, req *model.CreateServiceTypeRequest) (int, error) {
	q := `INSERT INTO service_types (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateServiceType(ctx *fasthttp.RequestCtx, id int, req *model.CreateServiceTypeRequest) error {
	q := `UPDATE service_types SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteServiceType(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM service_types WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Services CRUD operations
func (r *AdminRepository) GetServices(ctx *fasthttp.RequestCtx) ([]model.AdminServiceResponse, error) {
	services := make([]model.AdminServiceResponse, 0)
	q := `
		SELECT s.id, s.name, s.service_type_id, st.name as service_type_name, s.created_at 
		FROM services s
		LEFT JOIN service_types st ON s.service_type_id = st.id
		ORDER BY s.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return services, err
	}
	defer rows.Close()

	for rows.Next() {
		var service model.AdminServiceResponse
		if err := rows.Scan(&service.ID, &service.Name, &service.ServiceTypeID, &service.ServiceTypeName, &service.CreatedAt); err != nil {
			return services, err
		}
		services = append(services, service)
	}

	return services, err
}

func (r *AdminRepository) CreateService(ctx *fasthttp.RequestCtx, req *model.CreateServiceRequest) (int, error) {
	q := `INSERT INTO services (name, service_type_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.ServiceTypeID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateService(ctx *fasthttp.RequestCtx, id int, req *model.CreateServiceRequest) error {
	q := `UPDATE services SET name = $2, service_type_id = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.ServiceTypeID)
	return err
}

func (r *AdminRepository) DeleteService(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM services WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Generations CRUD operations
func (r *AdminRepository) GetGenerations(ctx *fasthttp.RequestCtx) ([]model.AdminGenerationResponse, error) {
	generations := make([]model.AdminGenerationResponse, 0)
	q := `
		SELECT g.id, g.name, g.model_id, m.name as model_name, g.start_year, g.end_year, g.wheel, g.image, g.created_at 
		FROM generations g
		LEFT JOIN models m ON g.model_id = m.id
		ORDER BY g.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return generations, err
	}
	defer rows.Close()

	for rows.Next() {
		var generation model.AdminGenerationResponse
		if err := rows.Scan(&generation.ID, &generation.Name, &generation.ModelID, &generation.ModelName,
			&generation.StartYear, &generation.EndYear, &generation.Wheel, &generation.Image, &generation.CreatedAt); err != nil {
			return generations, err
		}
		generations = append(generations, generation)
	}

	return generations, err
}

func (r *AdminRepository) CreateGeneration(ctx *fasthttp.RequestCtx, req *model.CreateGenerationRequest) (int, error) {
	q := `INSERT INTO generations (name, model_id, start_year, end_year, wheel, image) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.ModelID, req.StartYear, req.EndYear, req.Wheel, req.Image).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateGeneration(ctx *fasthttp.RequestCtx, id int, req *model.UpdateGenerationRequest) error {
	q := `UPDATE generations SET name = $2, model_id = $3, start_year = $4, end_year = $5, wheel = $6, image = $7 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.ModelID, req.StartYear, req.EndYear, req.Wheel, req.Image)
	return err
}

func (r *AdminRepository) DeleteGeneration(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM generations WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Generation Modifications CRUD operations
func (r *AdminRepository) GetGenerationModifications(ctx *fasthttp.RequestCtx, generationID int) ([]model.AdminGenerationModificationResponse, error) {
	modifications := make([]model.AdminGenerationModificationResponse, 0)
	q := `
		SELECT gm.id, gm.generation_id, gm.body_type_id, bt.name as body_type_name,
			   gm.engine_id, e.value as engine_value, gm.fuel_type_id, ft.name as fuel_type_name,
			   gm.drivetrain_id, dt.name as drivetrain_name, gm.transmission_id, t.name as transmission_name
		FROM generation_modifications gm
		LEFT JOIN body_types bt ON gm.body_type_id = bt.id
		LEFT JOIN engines e ON gm.engine_id = e.id
		LEFT JOIN fuel_types ft ON gm.fuel_type_id = ft.id
		LEFT JOIN drivetrains dt ON gm.drivetrain_id = dt.id
		LEFT JOIN transmissions t ON gm.transmission_id = t.id
		WHERE gm.generation_id = $1
		ORDER BY gm.id DESC
	`

	rows, err := r.db.Query(ctx, q, generationID)
	if err != nil {
		return modifications, err
	}
	defer rows.Close()

	for rows.Next() {
		var modification model.AdminGenerationModificationResponse
		if err := rows.Scan(&modification.ID, &modification.GenerationID, &modification.BodyTypeID, &modification.BodyTypeName,
			&modification.EngineID, &modification.EngineValue, &modification.FuelTypeID, &modification.FuelTypeName,
			&modification.DrivetrainID, &modification.DrivetrainName, &modification.TransmissionID, &modification.TransmissionName); err != nil {
			return modifications, err
		}
		modifications = append(modifications, modification)
	}

	return modifications, err
}

func (r *AdminRepository) CreateGenerationModification(ctx *fasthttp.RequestCtx, generationID int, req *model.CreateGenerationModificationRequest) (int, error) {
	q := `INSERT INTO generation_modifications (generation_id, body_type_id, engine_id, fuel_type_id, drivetrain_id, transmission_id) 
		  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, generationID, req.BodyTypeID, req.EngineID, req.FuelTypeID, req.DrivetrainID, req.TransmissionID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateGenerationModification(ctx *fasthttp.RequestCtx, generationID int, id int, req *model.UpdateGenerationModificationRequest) error {
	q := `UPDATE generation_modifications SET body_type_id = $3, engine_id = $4, fuel_type_id = $5, drivetrain_id = $6, transmission_id = $7 
		  WHERE generation_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, generationID, id, req.BodyTypeID, req.EngineID, req.FuelTypeID, req.DrivetrainID, req.TransmissionID)
	return err
}

func (r *AdminRepository) DeleteGenerationModification(ctx *fasthttp.RequestCtx, generationID int, id int) error {
	q := `DELETE FROM generation_modifications WHERE generation_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, generationID, id)
	return err
}

// Configurations CRUD operations
func (r *AdminRepository) GetConfigurations(ctx *fasthttp.RequestCtx) ([]model.AdminConfigurationResponse, error) {
	configurations := make([]model.AdminConfigurationResponse, 0)
	q := `
		SELECT c.id, c.body_type_id, bt.name as body_type_name, c.generation_id
		FROM configurations c
		LEFT JOIN body_types bt ON c.body_type_id = bt.id
		ORDER BY c.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return configurations, err
	}
	defer rows.Close()

	for rows.Next() {
		var configuration model.AdminConfigurationResponse
		if err := rows.Scan(&configuration.ID, &configuration.BodyTypeID, &configuration.BodyTypeName, &configuration.GenerationID); err != nil {
			return configurations, err
		}
		configurations = append(configurations, configuration)
	}

	return configurations, err
}

func (r *AdminRepository) CreateConfiguration(ctx *fasthttp.RequestCtx, req *model.CreateConfigurationRequest) (int, error) {
	q := `INSERT INTO configurations (body_type_id, generation_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.BodyTypeID, req.GenerationID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateConfiguration(ctx *fasthttp.RequestCtx, id int, req *model.UpdateConfigurationRequest) error {
	q := `UPDATE configurations SET body_type_id = $2, generation_id = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.BodyTypeID, req.GenerationID)
	return err
}

func (r *AdminRepository) DeleteConfiguration(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM configurations WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Colors CRUD operations
func (r *AdminRepository) GetColors(ctx *fasthttp.RequestCtx) ([]model.AdminColorResponse, error) {
	colors := make([]model.AdminColorResponse, 0)
	q := `SELECT id, name, image, created_at FROM colors ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return colors, err
	}
	defer rows.Close()

	for rows.Next() {
		var color model.AdminColorResponse
		if err := rows.Scan(&color.ID, &color.Name, &color.Image, &color.CreatedAt); err != nil {
			return colors, err
		}
		colors = append(colors, color)
	}

	return colors, err
}

func (r *AdminRepository) CreateColor(ctx *fasthttp.RequestCtx, req *model.CreateColorRequest) (int, error) {
	q := `INSERT INTO colors (name, image) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Image).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateColor(ctx *fasthttp.RequestCtx, id int, req *model.UpdateColorRequest) error {
	q := `UPDATE colors SET name = $2, image = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Image)
	return err
}

func (r *AdminRepository) DeleteColor(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM colors WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Categories CRUD operations
func (r *AdminRepository) GetMotoCategories(ctx *fasthttp.RequestCtx) ([]model.AdminMotoCategoryResponse, error) {
	motoCategories := make([]model.AdminMotoCategoryResponse, 0)
	q := `SELECT id, name, created_at FROM moto_categories ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return motoCategories, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoCategory model.AdminMotoCategoryResponse
		if err := rows.Scan(&motoCategory.ID, &motoCategory.Name, &motoCategory.CreatedAt); err != nil {
			return motoCategories, err
		}
		motoCategories = append(motoCategories, motoCategory)
	}

	return motoCategories, err
}

func (r *AdminRepository) CreateMotoCategory(ctx *fasthttp.RequestCtx, req *model.CreateMotoCategoryRequest) (int, error) {
	q := `INSERT INTO moto_categories (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoCategory(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoCategoryRequest) error {
	q := `UPDATE moto_categories SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteMotoCategory(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM moto_categories WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Brands CRUD operations
func (r *AdminRepository) GetMotoBrands(ctx *fasthttp.RequestCtx) ([]model.AdminMotoBrandResponse, error) {
	motoBrands := make([]model.AdminMotoBrandResponse, 0)
	q := `
		SELECT mb.id, mb.name, mb.image, mb.moto_category_id, mc.name as moto_category_name, mb.created_at
		FROM moto_brands mb
		LEFT JOIN moto_categories mc ON mb.moto_category_id = mc.id
		ORDER BY mb.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return motoBrands, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoBrand model.AdminMotoBrandResponse
		if err := rows.Scan(&motoBrand.ID, &motoBrand.Name, &motoBrand.Image, &motoBrand.MotoCategoryID,
			&motoBrand.MotoCategoryName, &motoBrand.CreatedAt); err != nil {
			return motoBrands, err
		}
		motoBrands = append(motoBrands, motoBrand)
	}

	return motoBrands, err
}

func (r *AdminRepository) CreateMotoBrand(ctx *fasthttp.RequestCtx, req *model.CreateMotoBrandRequest) (int, error) {
	q := `INSERT INTO moto_brands (name, image, moto_category_id) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Image, req.MotoCategoryID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoBrandRequest) error {
	q := `UPDATE moto_brands SET name = $2, image = $3, moto_category_id = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Image, req.MotoCategoryID)
	return err
}

func (r *AdminRepository) DeleteMotoBrand(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM moto_brands WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Models CRUD operations
func (r *AdminRepository) GetMotoModels(ctx *fasthttp.RequestCtx) ([]model.AdminMotoModelResponse, error) {
	motoModels := make([]model.AdminMotoModelResponse, 0)
	q := `
		SELECT mm.id, mm.name, mm.moto_brand_id, mb.name as moto_brand_name, mm.created_at
		FROM moto_models mm
		LEFT JOIN moto_brands mb ON mm.moto_brand_id = mb.id
		ORDER BY mm.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return motoModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoModel model.AdminMotoModelResponse
		if err := rows.Scan(&motoModel.ID, &motoModel.Name, &motoModel.MotoBrandID,
			&motoModel.MotoBrandName, &motoModel.CreatedAt); err != nil {
			return motoModels, err
		}
		motoModels = append(motoModels, motoModel)
	}

	return motoModels, err
}

func (r *AdminRepository) CreateMotoModel(ctx *fasthttp.RequestCtx, req *model.CreateMotoModelRequest) (int, error) {
	q := `INSERT INTO moto_models (name, moto_brand_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.MotoBrandID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoModelRequest) error {
	q := `UPDATE moto_models SET name = $2, moto_brand_id = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.MotoBrandID)
	return err
}

func (r *AdminRepository) DeleteMotoModel(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM moto_models WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Parameters CRUD operations
func (r *AdminRepository) GetMotoParameters(ctx *fasthttp.RequestCtx) ([]model.AdminMotoParameterResponse, error) {
	motoParameters := make([]model.AdminMotoParameterResponse, 0)
	q := `
		SELECT mp.id, mp.name, mp.moto_category_id, mc.name as moto_category_name, mp.created_at
		FROM moto_parameters mp
		LEFT JOIN moto_categories mc ON mp.moto_category_id = mc.id
		ORDER BY mp.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return motoParameters, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoParameter model.AdminMotoParameterResponse
		if err := rows.Scan(&motoParameter.ID, &motoParameter.Name, &motoParameter.MotoCategoryID,
			&motoParameter.MotoCategoryName, &motoParameter.CreatedAt); err != nil {
			return motoParameters, err
		}
		motoParameters = append(motoParameters, motoParameter)
	}

	return motoParameters, err
}

func (r *AdminRepository) CreateMotoParameter(ctx *fasthttp.RequestCtx, req *model.CreateMotoParameterRequest) (int, error) {
	q := `INSERT INTO moto_parameters (name, moto_category_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.MotoCategoryID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoParameter(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoParameterRequest) error {
	q := `UPDATE moto_parameters SET name = $2, moto_category_id = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.MotoCategoryID)
	return err
}

func (r *AdminRepository) DeleteMotoParameter(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM moto_parameters WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Parameter Values CRUD operations
func (r *AdminRepository) GetMotoParameterValues(ctx *fasthttp.RequestCtx, motoParamID int) ([]model.AdminMotoParameterValueResponse, error) {
	motoParameterValues := make([]model.AdminMotoParameterValueResponse, 0)
	q := `SELECT id, name, moto_parameter_id, created_at FROM moto_parameter_values WHERE moto_parameter_id = $1 ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q, motoParamID)
	if err != nil {
		return motoParameterValues, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoParameterValue model.AdminMotoParameterValueResponse
		if err := rows.Scan(&motoParameterValue.ID, &motoParameterValue.Name, &motoParameterValue.MotoParameterID,
			&motoParameterValue.CreatedAt); err != nil {
			return motoParameterValues, err
		}
		motoParameterValues = append(motoParameterValues, motoParameterValue)
	}

	return motoParameterValues, err
}

func (r *AdminRepository) CreateMotoParameterValue(ctx *fasthttp.RequestCtx, motoParamID int, req *model.CreateMotoParameterValueRequest) (int, error) {
	q := `INSERT INTO moto_parameter_values (name, moto_parameter_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, motoParamID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoParameterValue(ctx *fasthttp.RequestCtx, motoParamID int, id int, req *model.UpdateMotoParameterValueRequest) error {
	q := `UPDATE moto_parameter_values SET name = $3 WHERE moto_parameter_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, motoParamID, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteMotoParameterValue(ctx *fasthttp.RequestCtx, motoParamID int, id int) error {
	q := `DELETE FROM moto_parameter_values WHERE moto_parameter_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, motoParamID, id)
	return err
}

// Moto Category Parameters CRUD operations
func (r *AdminRepository) GetMotoCategoryParameters(ctx *fasthttp.RequestCtx, categoryID int) ([]model.AdminMotoCategoryParameterResponse, error) {
	motoCategoryParameters := make([]model.AdminMotoCategoryParameterResponse, 0)
	q := `
		SELECT mcp.id, mcp.moto_category_id, mcp.moto_parameter_id, mp.name as moto_parameter_name, mcp.created_at
		FROM moto_category_parameters mcp
		LEFT JOIN moto_parameters mp ON mcp.moto_parameter_id = mp.id
		WHERE mcp.moto_category_id = $1
		ORDER BY mcp.id DESC
	`

	rows, err := r.db.Query(ctx, q, categoryID)
	if err != nil {
		return motoCategoryParameters, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoCategoryParameter model.AdminMotoCategoryParameterResponse
		if err := rows.Scan(&motoCategoryParameter.ID, &motoCategoryParameter.MotoCategoryID, &motoCategoryParameter.MotoParameterID,
			&motoCategoryParameter.MotoParameterName, &motoCategoryParameter.CreatedAt); err != nil {
			return motoCategoryParameters, err
		}
		motoCategoryParameters = append(motoCategoryParameters, motoCategoryParameter)
	}

	return motoCategoryParameters, err
}

func (r *AdminRepository) CreateMotoCategoryParameter(ctx *fasthttp.RequestCtx, categoryID int, req *model.CreateMotoCategoryParameterRequest) (int, error) {
	q := `INSERT INTO moto_category_parameters (moto_category_id, moto_parameter_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, categoryID, req.MotoParameterID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoCategoryParameter(ctx *fasthttp.RequestCtx, categoryID int, id int, req *model.UpdateMotoCategoryParameterRequest) error {
	q := `UPDATE moto_category_parameters SET moto_parameter_id = $3 WHERE moto_category_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, categoryID, id, req.MotoParameterID)
	return err
}

func (r *AdminRepository) DeleteMotoCategoryParameter(ctx *fasthttp.RequestCtx, categoryID int, id int) error {
	q := `DELETE FROM moto_category_parameters WHERE moto_category_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, categoryID, id)
	return err
}

// Comtrans Categories CRUD operations
func (r *AdminRepository) GetComtransCategories(ctx *fasthttp.RequestCtx) ([]model.AdminComtransCategoryResponse, error) {
	comtransCategories := make([]model.AdminComtransCategoryResponse, 0)
	q := `SELECT id, name, created_at FROM com_categories ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return comtransCategories, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransCategory model.AdminComtransCategoryResponse
		if err := rows.Scan(&comtransCategory.ID, &comtransCategory.Name, &comtransCategory.CreatedAt); err != nil {
			return comtransCategories, err
		}
		comtransCategories = append(comtransCategories, comtransCategory)
	}

	return comtransCategories, err
}

func (r *AdminRepository) CreateComtransCategory(ctx *fasthttp.RequestCtx, req *model.CreateComtransCategoryRequest) (int, error) {
	q := `INSERT INTO com_categories (name) VALUES ($1) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransCategory(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransCategoryRequest) error {
	q := `UPDATE com_categories SET name = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteComtransCategory(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_categories WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Category Parameters CRUD operations
func (r *AdminRepository) GetComtransCategoryParameters(ctx *fasthttp.RequestCtx, categoryID int) ([]model.AdminComtransCategoryParameterResponse, error) {
	comtransCategoryParameters := make([]model.AdminComtransCategoryParameterResponse, 0)
	q := `
		SELECT ccp.comtran_category_id, ccp.comtran_parameter_id, cp.name as comtrans_parameter_name, ccp.created_at
		FROM com_category_parameters ccp
		LEFT JOIN com_parameters cp ON ccp.comtran_parameter_id = cp.id
		WHERE ccp.comtran_category_id = $1
		ORDER BY ccp.created_at DESC
	`

	rows, err := r.db.Query(ctx, q, categoryID)
	if err != nil {
		return comtransCategoryParameters, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransCategoryParameter model.AdminComtransCategoryParameterResponse
		if err := rows.Scan(&comtransCategoryParameter.ComtransCategoryID, &comtransCategoryParameter.ComtransParameterID,
			&comtransCategoryParameter.ComtransParameterName, &comtransCategoryParameter.CreatedAt); err != nil {
			return comtransCategoryParameters, err
		}
		comtransCategoryParameters = append(comtransCategoryParameters, comtransCategoryParameter)
	}

	return comtransCategoryParameters, err
}

func (r *AdminRepository) CreateComtransCategoryParameter(ctx *fasthttp.RequestCtx, categoryID int, req *model.CreateComtransCategoryParameterRequest) (int, error) {
	q := `INSERT INTO com_category_parameters (comtran_category_id, comtran_parameter_id) VALUES ($1, $2) RETURNING comtran_parameter_id`
	var id int
	err := r.db.QueryRow(ctx, q, categoryID, req.ComtransParameterID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransCategoryParameter(ctx *fasthttp.RequestCtx, categoryID int, id int, req *model.UpdateComtransCategoryParameterRequest) error {
	q := `UPDATE com_category_parameters SET comtran_parameter_id = $3 WHERE comtran_category_id = $1 AND comtran_parameter_id = $2`
	_, err := r.db.Exec(ctx, q, categoryID, id, req.ComtransParameterID)
	return err
}

func (r *AdminRepository) DeleteComtransCategoryParameter(ctx *fasthttp.RequestCtx, categoryID int, id int) error {
	q := `DELETE FROM com_category_parameters WHERE comtran_category_id = $1 AND comtran_parameter_id = $2`
	_, err := r.db.Exec(ctx, q, categoryID, id)
	return err
}

// Comtrans Brands CRUD operations
func (r *AdminRepository) GetComtransBrands(ctx *fasthttp.RequestCtx) ([]model.AdminComtransBrandResponse, error) {
	comtransBrands := make([]model.AdminComtransBrandResponse, 0)
	q := `
		SELECT cb.id, cb.name, cb.image, cb.comtran_category_id, cc.name as comtrans_category_name, cb.created_at
		FROM com_brands cb
		LEFT JOIN com_categories cc ON cb.comtran_category_id = cc.id
		ORDER BY cb.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return comtransBrands, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransBrand model.AdminComtransBrandResponse
		if err := rows.Scan(&comtransBrand.ID, &comtransBrand.Name, &comtransBrand.Image, &comtransBrand.ComtransCategoryID,
			&comtransBrand.ComtransCategoryName, &comtransBrand.CreatedAt); err != nil {
			return comtransBrands, err
		}
		comtransBrands = append(comtransBrands, comtransBrand)
	}

	return comtransBrands, err
}

func (r *AdminRepository) CreateComtransBrand(ctx *fasthttp.RequestCtx, req *model.CreateComtransBrandRequest) (int, error) {
	q := `INSERT INTO com_brands (name, image, comtran_category_id) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.Image, req.ComtransCategoryID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransBrandRequest) error {
	q := `UPDATE com_brands SET name = $2, image = $3, comtran_category_id = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.Image, req.ComtransCategoryID)
	return err
}

func (r *AdminRepository) DeleteComtransBrand(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_brands WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Models CRUD operations
func (r *AdminRepository) GetComtransModels(ctx *fasthttp.RequestCtx) ([]model.AdminComtransModelResponse, error) {
	comtransModels := make([]model.AdminComtransModelResponse, 0)
	q := `
		SELECT cm.id, cm.name, cm.comtran_brand_id, cb.name as comtrans_brand_name, cm.created_at
		FROM com_models cm
		LEFT JOIN com_brands cb ON cm.comtran_brand_id = cb.id
		ORDER BY cm.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return comtransModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransModel model.AdminComtransModelResponse
		if err := rows.Scan(&comtransModel.ID, &comtransModel.Name, &comtransModel.ComtransBrandID,
			&comtransModel.ComtransBrandName, &comtransModel.CreatedAt); err != nil {
			return comtransModels, err
		}
		comtransModels = append(comtransModels, comtransModel)
	}

	return comtransModels, err
}

func (r *AdminRepository) CreateComtransModel(ctx *fasthttp.RequestCtx, req *model.CreateComtransModelRequest) (int, error) {
	q := `INSERT INTO com_models (name, comtran_brand_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.ComtransBrandID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransModelRequest) error {
	q := `UPDATE com_models SET name = $2, comtran_brand_id = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.ComtransBrandID)
	return err
}

func (r *AdminRepository) DeleteComtransModel(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_models WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Parameters CRUD operations
func (r *AdminRepository) GetComtransParameters(ctx *fasthttp.RequestCtx) ([]model.AdminComtransParameterResponse, error) {
	comtransParameters := make([]model.AdminComtransParameterResponse, 0)
	q := `
		SELECT cp.id, cp.name, cp.comtran_category_id, cc.name as comtrans_category_name, cp.created_at
		FROM com_parameters cp
		LEFT JOIN com_categories cc ON cp.comtran_category_id = cc.id
		ORDER BY cp.id DESC
	`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return comtransParameters, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransParameter model.AdminComtransParameterResponse
		if err := rows.Scan(&comtransParameter.ID, &comtransParameter.Name, &comtransParameter.ComtransCategoryID,
			&comtransParameter.ComtransCategoryName, &comtransParameter.CreatedAt); err != nil {
			return comtransParameters, err
		}
		comtransParameters = append(comtransParameters, comtransParameter)
	}

	return comtransParameters, err
}

func (r *AdminRepository) CreateComtransParameter(ctx *fasthttp.RequestCtx, req *model.CreateComtransParameterRequest) (int, error) {
	q := `INSERT INTO com_parameters (name, comtran_category_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.ComtransCategoryID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransParameter(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransParameterRequest) error {
	q := `UPDATE com_parameters SET name = $2, comtran_category_id = $3 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.ComtransCategoryID)
	return err
}

func (r *AdminRepository) DeleteComtransParameter(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_parameters WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Parameter Values CRUD operations
func (r *AdminRepository) GetComtransParameterValues(ctx *fasthttp.RequestCtx, parameterID int) ([]model.AdminComtransParameterValueResponse, error) {
	comtransParameterValues := make([]model.AdminComtransParameterValueResponse, 0)
	q := `SELECT id, name, comtran_parameter_id, created_at FROM com_parameter_values WHERE comtran_parameter_id = $1 ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q, parameterID)
	if err != nil {
		return comtransParameterValues, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransParameterValue model.AdminComtransParameterValueResponse
		if err := rows.Scan(&comtransParameterValue.ID, &comtransParameterValue.Name, &comtransParameterValue.ComtransParameterID,
			&comtransParameterValue.CreatedAt); err != nil {
			return comtransParameterValues, err
		}
		comtransParameterValues = append(comtransParameterValues, comtransParameterValue)
	}

	return comtransParameterValues, err
}

func (r *AdminRepository) CreateComtransParameterValue(ctx *fasthttp.RequestCtx, parameterID int, req *model.CreateComtransParameterValueRequest) (int, error) {
	q := `INSERT INTO com_parameter_values (name, comtran_parameter_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, parameterID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransParameterValue(ctx *fasthttp.RequestCtx, parameterID int, id int, req *model.UpdateComtransParameterValueRequest) error {
	q := `UPDATE com_parameter_values SET name = $3 WHERE comtran_parameter_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, parameterID, id, req.Name)
	return err
}

func (r *AdminRepository) DeleteComtransParameterValue(ctx *fasthttp.RequestCtx, parameterID int, id int) error {
	q := `DELETE FROM com_parameter_values WHERE comtran_parameter_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, parameterID, id)
	return err
}
