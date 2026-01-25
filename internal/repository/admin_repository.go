package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"

	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
)

type AdminRepository struct {
	config *config.Config
	db     *pgxpool.Pool
}

func NewAdminRepository(config *config.Config, db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{config, db}
}

// Admin CRUD operations
func (r *AdminRepository) CreateUser(ctx *fasthttp.RequestCtx, req *model.CreateUserRequest) (int, error) {
	var id int
	q := `
		INSERT INTO users (username, email, password, permissions, role_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := r.db.QueryRow(ctx, q, req.Username, req.Email, req.Password, req.Permissions, req.RoleID).Scan(&id)
	return id, err
}

func (r *AdminRepository) GetUsers(ctx *fasthttp.RequestCtx, qRoleID int) ([]model.UserResponse, error) {
	users := make([]model.UserResponse, 0)
	q := `
		SELECT 
			id, 
			username, 
			email, 
			permissions, 
			status, 
			created_at::text, 
			updated_at::text
		FROM users
		WHERE role_id = $1
		ORDER BY id DESC
	`
	rows, err := r.db.Query(ctx, q, qRoleID)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.UserResponse
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Permissions, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *AdminRepository) GetUser(ctx *fasthttp.RequestCtx, id int) (model.UserResponse, error) {
	user := model.UserResponse{}
	q := `
		SELECT 
			id, 
			username, 
			email, 
			permissions, 
			status, 
			created_at::text, 
			updated_at::text
		FROM users 
		WHERE id = $1 AND role_id = 0
	`
	err := r.db.QueryRow(ctx, q, id).Scan(&user.ID, &user.Username, &user.Email, &user.Permissions, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (r *AdminRepository) UpdateUser(ctx *fasthttp.RequestCtx, id int, req *model.UpdateUserRequest) error {
	updates := []string{}
	args := []any{}
	argPos := 1

	if req.Username != "" {
		updates = append(updates, fmt.Sprintf("username = $%d", argPos))
		args = append(args, req.Username)
		argPos++
	}

	if req.Email != "" {
		updates = append(updates, fmt.Sprintf("email = $%d", argPos))
		args = append(args, req.Email)
		argPos++
	}

	if req.Password != "" {
		updates = append(updates, fmt.Sprintf("password = $%d", argPos))
		args = append(args, req.Password)
		argPos++
	}

	if req.Status != 0 {
		updates = append(updates, fmt.Sprintf("status = $%d", argPos))
		args = append(args, req.Status)
		argPos++
	}

	if req.RoleID != 0 {
		updates = append(updates, fmt.Sprintf("role_id = $%d", argPos))
		args = append(args, req.RoleID)
		argPos++
	}

	if req.Permissions != nil {
		updates = append(updates, fmt.Sprintf("permissions = $%d", argPos))
		args = append(args, req.Permissions)
		argPos++
	}

	if len(updates) == 0 {
		return nil // No updates to perform
	}

	updates = append(updates, "updated_at = now()")
	args = append(args, id)

	updateStr := ""
	for i, update := range updates {
		if i > 0 {
			updateStr += ", "
		}
		updateStr += update
	}

	q := fmt.Sprintf(`
		UPDATE users 
		SET %s
		WHERE id = $%d AND role_id = 0
	`, updateStr, len(args))

	_, err := r.db.Exec(ctx, q, args...)
	return err
}

func (r *AdminRepository) DeleteUser(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM users WHERE id = $1 AND role_id = 0`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Profile CRUD operations
func (r *AdminRepository) GetProfile(ctx *fasthttp.RequestCtx, id int) (model.AdminProfileResponse, error) {
	profile := model.AdminProfileResponse{}
	q := `
		SELECT 
			id, 
			username, 
			email,
			permissions
		FROM users 
		WHERE role_id = 0
		AND id = $1`
	err := r.db.QueryRow(ctx, q, id).Scan(&profile.ID, &profile.Username, &profile.Email, &profile.Permissions)

	return profile, err
}

// Application CRUD operations
func (r *AdminRepository) GetApplications(ctx *fasthttp.RequestCtx, qRole, qStatus, limit, lastID int, search string) ([]model.AdminApplicationResponse, error) {
	applications := make([]model.AdminApplicationResponse, 0)
	q := ``
	qWhere := ``
	var searchArgs []interface{}
	searchParamIndex := 4 // After qRole, lastID, limit

	if search != "" {
		// Use parameterized query to prevent SQL injection
		searchPattern := "'%' || $" + strconv.Itoa(searchParamIndex) + " || '%'"
		qWhere = fmt.Sprintf(" AND (u.username ILIKE %s OR p.company_name ILIKE %s OR u.email ILIKE %s OR u.phone ILIKE %s)",
			searchPattern, searchPattern, searchPattern, searchPattern)
		searchArgs = []interface{}{search}
	}

	switch qStatus {
	case model.APPLICATION_STATUS_APPROVED:
		// select from users table
		q = `
			SELECT 
				u.id,
				p.company_name, 
				d.licence_issue_date, 
				d.licence_expiry_date, 
				u.username, 
				u.email, 
				u.phone, 
				u.status, 
				u.created_at
			FROM users u
			left join profiles p on p.user_id = u.id
			left join documents d on d.id = p.documents_id
			WHERE role_id = $1 and u.id > $2` + qWhere + `
			ORDER BY id DESC
			limit $3
		`
	default:
		q = `
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
			WHERE role_id = $1 and id > $2` + qWhere + `
			ORDER BY id DESC
			limit $3
		`
	}

	// Combine all query arguments: qRole, lastID, limit, and search (if present)
	args := []interface{}{qRole, lastID, limit}
	args = append(args, searchArgs...)
	rows, err := r.db.Query(ctx, q, args...)

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

func (r *AdminRepository) CreateApplication(ctx *fasthttp.RequestCtx, req model.UserApplication) (int, error) {
	tx, err := r.db.Begin(ctx)

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	q := `
		insert into users (email, password, username, role_id, phone)
		values ($1, $2, $3, $4, $5) 
		ON CONFLICT (email) DO UPDATE
		SET password = EXCLUDED.password, created_at = now(), updated_at = now(), role_id = EXCLUDED.role_id, phone = EXCLUDED.phone
		RETURNING id
	`
	var userID int
	err = tx.QueryRow(ctx, q, req.Email, "password", req.FullName, req.RoleID, req.Phone).Scan(&userID)

	if err != nil {
		return 0, err
	}

	// insert to documents table

	q = `
		insert into documents (
			copy_of_id_path,
			memorandum_path,
			licence_path,
			licence_issue_date,
			licence_expiry_date
		)
		values ($1, $2, $3, $4, $5)
		returning id
	`
	var documentID int
	err = tx.QueryRow(ctx, q, "req.CopyOfIDPath", "req.MemorandumPath", "req.LicencePath", req.LicenceIssueDate, req.LicenceExpiryDate).Scan(&documentID)

	if err != nil {
		return 0, err
	}

	q = `
		insert into profiles (
			user_id, 
			username, 
			company_name, 
			registered_by,
			address,
			company_type_id,
			activity_field_id,
			vat_number,
			documents_id
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		on conflict (user_id)
		do update set
			username = EXCLUDED.username,
			company_name = EXCLUDED.company_name,
			registered_by = EXCLUDED.registered_by,
			address = EXCLUDED.address,
			company_type_id = EXCLUDED.company_type_id,
			activity_field_id = EXCLUDED.activity_field_id,
			vat_number = EXCLUDED.vat_number,
			documents_id = EXCLUDED.documents_id
	`
	_, err = tx.Exec(ctx, q,
		userID, req.FullName, req.CompanyName, "application",
		req.Address, req.CompanyTypeID, req.ActivityFieldID,
		req.VATNumber, documentID)

	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	return userID, err
}

func (r *AdminRepository) GetApplication(ctx *fasthttp.RequestCtx, id int, qStatus int) (model.AdminApplicationByIDResponse, error) {
	q := ``

	switch qStatus {
	case model.APPLICATION_STATUS_APPROVED:
		// select from users table
		q = `
			SELECT 
				u.id, 
				p.company_name, 
				ds.licence_issue_date, 
				ds.licence_expiry_date, 
				u.username, 
				u.email, 
				u.phone, 
				u.status, 
				u.created_at,
				$2 || ds.copy_of_id_path as copy_of_id_url,
				$2 || ds.memorandum_path as memorandum_url,
				$2 || ds.licence_path as licence_url,
				p.address,
				ct.name as company_type,
				af.name as activity_field,
				p.vat_number
			FROM users u
			left join profiles p on p.user_id = u.id
			left join documents ds on ds.id = p.documents_id
			left join company_types ct on ct.id = p.company_type_id
			left join activity_fields af on af.id = p.activity_field_id
			WHERE u.id = $1
		`
	default:
		q = `
			SELECT 
				tu.id, 
				tu.company_name, 
				tu.licence_issue_date, 
				tu.licence_expiry_date, 
				tu.username, 
				tu.email, 
				tu.phone, 
				tu.status, 
				tu.created_at,
				$2 || ds.copy_of_id_path as copy_of_id_url,
				$2 || ds.memorandum_path as memorandum_url,
				$2 || ds.licence_path as licence_url,
				tu.address,
				ct.name as company_type,
				af.name as activity_field,
				tu.vat_number
			FROM temp_users tu
			left join documents ds on ds.id = tu.documents_id
			left join company_types ct on ct.id = tu.company_type_id
			left join activity_fields af on af.id = tu.activity_field_id
			WHERE tu.id = $1
		`
	}

	var application model.AdminApplicationByIDResponse
	err := r.db.QueryRow(ctx, q, id, r.config.IMAGE_BASE_URL).Scan(
		&application.ID, &application.CompanyName,
		&application.LicenceIssueDate, &application.LicenceExpiryDate,
		&application.FullName, &application.Email, &application.Phone,
		&application.Status, &application.CreatedAt, &application.CopyOFIDURL,
		&application.MemorandumURL, &application.LicenceURL, &application.Address,
		&application.CompanyType, &application.ActivityField, &application.VATNumber)

	return application, err
}

func (r *AdminRepository) CreateApplicationDocuments(ctx *fasthttp.RequestCtx, userID int, documents model.UserApplicationDocuments) error {
	q := `
		select documents_id from profiles where user_id = $1
	`
	var documentsID int
	err := r.db.QueryRow(ctx, q, userID).Scan(&documentsID)

	if err != nil {
		return err
	}

	q = `
		update documents set
			licence_path = $1,
			memorandum_path = $2,
			copy_of_id_path = $3
		where id = $4
	`
	_, err = r.db.Exec(ctx, q, documents.Licence, documents.Memorandum, documents.CopyOfID, documentsID)
	return err
}

func (r *AdminRepository) AcceptApplication(ctx *fasthttp.RequestCtx, id int, password string) (string, error) {
	tx, err := r.db.Begin(ctx)

	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	tempUser := model.TempUser{}
	q := `
		select 
			company_name,
			company_type_id,
			activity_field_id,
			vat_number,
			address,
			licence_issue_date,
			licence_expiry_date,
			documents_id,
			email,
			username,
			role_id,
			phone
		from temp_users
		where id = $1
	`
	err = tx.QueryRow(ctx, q, id).Scan(
		&tempUser.CompanyName, &tempUser.CompanyTypeID, &tempUser.ActivityFieldID,
		&tempUser.VATNumber, &tempUser.Address, &tempUser.LicenceIssueDate,
		&tempUser.LicenceExpiryDate, &tempUser.DocumentsID, &tempUser.Email, &tempUser.FullName,
		&tempUser.RoleID, &tempUser.Phone)

	if err != nil {
		return "", err
	}

	q = `
		insert into users (email, password, username, role_id, phone)
		values ($1, $2, $3, $4, $5) 
		ON CONFLICT (email) DO UPDATE
		SET password = EXCLUDED.password, created_at = now(), updated_at = now(), role_id = EXCLUDED.role_id, phone = EXCLUDED.phone
		RETURNING id, phone
	`
	var userID int
	var phone string
	fmt.Println("tempUser.Phone", tempUser.Phone)
	err = tx.QueryRow(ctx, q, tempUser.Email, password, tempUser.FullName, tempUser.RoleID, tempUser.Phone).Scan(&userID, &phone)

	if err != nil {
		return "", err
	}

	fmt.Println("scanned phone", phone)

	q = `
		insert into profiles (
			user_id, 
			username, 
			company_name, 
			registered_by,
			address,
			company_type_id,
			activity_field_id,
			vat_number,
			documents_id
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		on conflict (user_id)
		do update set
			username = EXCLUDED.username,
			company_name = EXCLUDED.company_name,
			registered_by = EXCLUDED.registered_by,
			address = EXCLUDED.address,
			company_type_id = EXCLUDED.company_type_id,
			activity_field_id = EXCLUDED.activity_field_id,
			vat_number = EXCLUDED.vat_number,
			documents_id = EXCLUDED.documents_id
	`
	_, err = tx.Exec(ctx, q,
		userID, tempUser.FullName, tempUser.CompanyName, "application",
		tempUser.Address, tempUser.CompanyTypeID, tempUser.ActivityFieldID,
		tempUser.VATNumber, tempUser.DocumentsID)

	if err != nil {
		return "", err
	}

	q = `
		update documents set
			licence_issue_date = $1,
			licence_expiry_date = $2
		where id = $3
	`
	_, err = tx.Exec(ctx, q, tempUser.LicenceIssueDate, tempUser.LicenceExpiryDate, tempUser.DocumentsID)

	if err != nil {
		return "", err
	}

	q = `
		delete from temp_users where id = $1
	`
	_, err = tx.Exec(ctx, q, id)

	if err != nil {
		return "", err
	}

	err = tx.Commit(ctx)
	return *tempUser.Email, err
}

func (r *AdminRepository) RejectApplication(ctx *fasthttp.RequestCtx, id int, qStatus int) (string, error) {
	q := `delete from temp_users WHERE id = $1 returning email`
	var email string

	if qStatus == model.APPLICATION_STATUS_APPROVED {
		q = `delete from users where id = $1 returning email`
	}

	err := r.db.QueryRow(ctx, q, id).Scan(&email)
	return email, err
}

// Cities CRUD operations
func (r *AdminRepository) GetCities(ctx *fasthttp.RequestCtx) ([]model.AdminCityResponse, error) {
	cities := make([]model.AdminCityResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM cities ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return cities, err
	}

	defer rows.Close()

	for rows.Next() {
		var city model.AdminCityResponse

		if err := rows.Scan(&city.ID, &city.Name, &city.NameRu, &city.NameAe, &city.CreatedAt); err != nil {
			return cities, err
		}
		cities = append(cities, city)
	}

	return cities, err
}

func (r *AdminRepository) CreateCity(ctx *fasthttp.RequestCtx, req *model.CreateNameRequest) (int, error) {
	q := `INSERT INTO cities (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateCity(ctx *fasthttp.RequestCtx, id int, req *model.CreateNameRequest) error {
	q := `UPDATE cities SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteCity(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM cities WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Company Types CRUD operations
func (r *AdminRepository) GetCompanyTypes(ctx *fasthttp.RequestCtx) ([]model.CompanyType, error) {
	companyTypes := make([]model.CompanyType, 0)
	q := `SELECT id, name, name_ru, name_ae FROM company_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return companyTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.CompanyType
		if err := rows.Scan(&item.ID, &item.Name, &item.NameRu, &item.NameAe); err != nil {
			return companyTypes, err
		}
		companyTypes = append(companyTypes, item)
	}

	return companyTypes, err
}

func (r *AdminRepository) GetCompanyType(ctx *fasthttp.RequestCtx, id int) (*model.CompanyType, error) {
	q := `SELECT id, name, name_ru, name_ae FROM company_types WHERE id = $1`
	var item model.CompanyType
	err := r.db.QueryRow(ctx, q, id).Scan(&item.ID, &item.Name, &item.NameRu, &item.NameAe)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *AdminRepository) CreateCompanyType(ctx *fasthttp.RequestCtx, req *model.CreateCompanyTypeRequest) (int, error) {
	q := `INSERT INTO company_types (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateCompanyType(ctx *fasthttp.RequestCtx, id int, req *model.CreateCompanyTypeRequest) error {
	q := `UPDATE company_types SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteCompanyType(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM company_types WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Activity Fields CRUD operations
func (r *AdminRepository) GetActivityFields(ctx *fasthttp.RequestCtx) ([]model.CompanyType, error) {
	items := make([]model.CompanyType, 0)
	q := `
		SELECT 
			id, 
			name, 
			name_ru, 
			name_ae
		FROM activity_fields 
		ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.CompanyType
		if err := rows.Scan(&item.ID, &item.Name, &item.NameRu, &item.NameAe); err != nil {
			return items, err
		}
		items = append(items, item)
	}

	return items, err
}

func (r *AdminRepository) GetActivityField(ctx *fasthttp.RequestCtx, id int) (*model.CompanyType, error) {
	q := `SELECT id, name, name_ru, name_ae FROM activity_fields WHERE id = $1`
	var item model.CompanyType
	err := r.db.QueryRow(ctx, q, id).Scan(&item.ID, &item.Name, &item.NameRu, &item.NameAe)
	return &item, err
}

func (r *AdminRepository) CreateActivityField(ctx *fasthttp.RequestCtx, req *model.CreateCompanyTypeRequest) (int, error) {
	q := `INSERT INTO activity_fields (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateActivityField(ctx *fasthttp.RequestCtx, id int, req *model.CreateCompanyTypeRequest) error {
	q := `UPDATE activity_fields SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteActivityField(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM activity_fields WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Brands CRUD operations
func (r *AdminRepository) GetBrands(ctx *fasthttp.RequestCtx) ([]model.AdminBrandResponse, error) {
	brands := make([]model.AdminBrandResponse, 0)
	q := `
			SELECT 
				id, 
				name, 
				name_ru,
				name_ae,
				$1 || logo, 
				model_count, 
				popular, 
				updated_at 
			FROM brands 
			ORDER BY id DESC
		`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)

	if err != nil {
		return brands, err
	}

	defer rows.Close()

	for rows.Next() {
		var brand model.AdminBrandResponse
		if err := rows.Scan(&brand.ID, &brand.Name, &brand.NameRu, &brand.NameAe, &brand.Logo, &brand.ModelCount, &brand.Popular, &brand.UpdatedAt); err != nil {
			return brands, err
		}
		brands = append(brands, brand)
	}

	return brands, err
}

func (r *AdminRepository) CreateBrand(ctx *fasthttp.RequestCtx, req *model.CreateBrandRequest) (int, error) {
	q := `INSERT INTO brands (name, name_ru, name_ae, popular, updated_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, req.Popular).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateBrandImage(ctx *fasthttp.RequestCtx, id int, path string) error {
	q := `UPDATE brands SET logo = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, path)
	return err
}

func (r *AdminRepository) UpdateBrand(ctx *fasthttp.RequestCtx, id int, req *model.CreateBrandRequest) error {
	q := `UPDATE brands SET name = $2, name_ru = $3, name_ae = $4, popular = $5, updated_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe, req.Popular)
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
		SELECT m.id, m.name, m.name_ru, m.name_ae, m.brand_id, b.name as brand_name, b.name_ru as brand_name_ru, m.popular, m.updated_at 
		FROM models m
		LEFT JOIN brands b ON m.brand_id = b.id
		WHERE m.brand_id = $1
		ORDER BY m.id DESC
	`

	rows, err := r.db.Query(ctx, q, brand_id)

	if err != nil {
		return models, err
	}

	defer rows.Close()

	for rows.Next() {
		var modelItem model.AdminModelResponse
		if err := rows.Scan(
			&modelItem.ID, &modelItem.Name, &modelItem.NameRu, &modelItem.NameAe,
			&modelItem.BrandID, &modelItem.BrandName, &modelItem.BrandNameRu,
			&modelItem.Popular, &modelItem.UpdatedAt); err != nil {
			return models, err
		}
		models = append(models, modelItem)
	}

	return models, err
}

func (r *AdminRepository) CreateModel(ctx *fasthttp.RequestCtx, brand_id int, req *model.CreateModelRequest) (int, error) {
	q := `INSERT INTO models (name, name_ru, name_ae, brand_id, popular, updated_at) VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, brand_id, req.Popular).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateModelRequest) error {
	q := `UPDATE models SET name = $2, name_ru = $3, name_ae = $4, brand_id = $5, popular = $6, updated_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe, req.BrandID, req.Popular)
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
	q := `SELECT id, name, name_ru, name_ae, $1 || image, created_at FROM body_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)

	if err != nil {
		return bodyTypes, err
	}

	defer rows.Close()

	for rows.Next() {
		var bodyType model.AdminBodyTypeResponse
		if err := rows.Scan(&bodyType.ID, &bodyType.Name, &bodyType.NameRu, &bodyType.NameAe, &bodyType.Image, &bodyType.CreatedAt); err != nil {
			return bodyTypes, err
		}
		bodyTypes = append(bodyTypes, bodyType)
	}

	return bodyTypes, err
}

func (r *AdminRepository) CreateBodyType(ctx *fasthttp.RequestCtx, req *model.CreateBodyTypeRequest) (int, error) {
	q := `INSERT INTO body_types (name, name_ru, name_ae, image) VALUES ($1, $2, $3, 'not_uploaded') RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateBodyTypeImage(ctx *fasthttp.RequestCtx, id int, path string) error {
	q := `
		UPDATE body_types 
		SET image = $2 
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, q, id, path)
	return err
}

func (r *AdminRepository) DeleteBodyTypeImage(ctx *fasthttp.RequestCtx, id int) error {
	q := `
		UPDATE body_types 
		SET image = 'NULL' 
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *AdminRepository) UpdateBodyType(ctx *fasthttp.RequestCtx, id int, req *model.CreateBodyTypeRequest) error {
	q := `UPDATE body_types SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
	q := `SELECT id, name, name_ru, name_ae, created_at FROM regions where city_id = $1`

	rows, err := r.db.Query(ctx, q, cityID)

	if err != nil {
		return regions, err
	}

	defer rows.Close()

	for rows.Next() {
		var region model.AdminCityResponse
		if err := rows.Scan(&region.ID, &region.Name, &region.NameRu, &region.NameAe, &region.CreatedAt); err != nil {
			return regions, err
		}
		regions = append(regions, region)
	}

	return regions, err
}

func (r *AdminRepository) CreateRegion(ctx *fasthttp.RequestCtx, city_id int, req *model.CreateNameRequest) (int, error) {
	q := `INSERT INTO regions (name, name_ru, name_ae, city_id) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, city_id).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateRegion(ctx *fasthttp.RequestCtx, id int, req *model.CreateNameRequest) error {
	q := `UPDATE regions SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
	q := `SELECT id, name, name_ru, name_ae, created_at FROM transmissions ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return transmissions, err
	}
	defer rows.Close()

	for rows.Next() {
		var transmission model.AdminTransmissionResponse
		if err := rows.Scan(&transmission.ID, &transmission.Name, &transmission.NameRu, &transmission.NameAe, &transmission.CreatedAt); err != nil {
			return transmissions, err
		}
		transmissions = append(transmissions, transmission)
	}

	return transmissions, err
}

func (r *AdminRepository) CreateTransmission(ctx *fasthttp.RequestCtx, req *model.CreateTransmissionRequest) (int, error) {
	q := `INSERT INTO transmissions (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateTransmission(ctx *fasthttp.RequestCtx, id int, req *model.CreateTransmissionRequest) error {
	q := `UPDATE transmissions SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
	q := `SELECT id, name, name_ru, name_ae, created_at FROM engines ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return engines, err
	}
	defer rows.Close()

	for rows.Next() {
		var engine model.AdminEngineResponse
		if err := rows.Scan(&engine.ID, &engine.Name, &engine.NameRu, &engine.NameAe, &engine.CreatedAt); err != nil {
			return engines, err
		}
		engines = append(engines, engine)
	}

	return engines, err
}

func (r *AdminRepository) CreateEngine(ctx *fasthttp.RequestCtx, req *model.CreateEngineRequest) (int, error) {
	q := `INSERT INTO engines (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateEngine(ctx *fasthttp.RequestCtx, id int, req *model.CreateEngineRequest) error {
	q := `UPDATE engines SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteEngine(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM engines WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Engines CRUD operations
func (r *AdminRepository) GetComtransEngines(ctx *fasthttp.RequestCtx) ([]model.AdminComtransEngineResponse, error) {
	engines := make([]model.AdminComtransEngineResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM com_engines ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return engines, err
	}
	defer rows.Close()

	for rows.Next() {
		var engine model.AdminComtransEngineResponse
		if err := rows.Scan(&engine.ID, &engine.Name, &engine.NameRu, &engine.NameAe, &engine.CreatedAt); err != nil {
			return engines, err
		}
		engines = append(engines, engine)
	}

	return engines, err
}

func (r *AdminRepository) CreateComtransEngine(ctx *fasthttp.RequestCtx, req *model.CreateComtransEngineRequest) (int, error) {
	q := `INSERT INTO com_engines (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) DeleteComtransEngine(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_engines WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Engines CRUD operations
func (r *AdminRepository) GetMotoEngines(ctx *fasthttp.RequestCtx) ([]model.AdminMotoEngineResponse, error) {
	engines := make([]model.AdminMotoEngineResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM moto_engines ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return engines, err
	}
	defer rows.Close()

	for rows.Next() {
		var engine model.AdminMotoEngineResponse
		if err := rows.Scan(&engine.ID, &engine.Name, &engine.NameRu, &engine.NameAe, &engine.CreatedAt); err != nil {
			return engines, err
		}
		engines = append(engines, engine)
	}

	return engines, err
}

func (r *AdminRepository) CreateMotoEngine(ctx *fasthttp.RequestCtx, req *model.CreateMotoEngineRequest) (int, error) {
	q := `INSERT INTO moto_engines (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) DeleteMotoEngine(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM moto_engines WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Drivetrains CRUD operations
func (r *AdminRepository) GetDrivetrains(ctx *fasthttp.RequestCtx) ([]model.AdminDrivetrainResponse, error) {
	drivetrains := make([]model.AdminDrivetrainResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM drivetrains ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return drivetrains, err
	}
	defer rows.Close()

	for rows.Next() {
		var drivetrain model.AdminDrivetrainResponse
		if err := rows.Scan(&drivetrain.ID, &drivetrain.Name, &drivetrain.NameRu, &drivetrain.NameAe, &drivetrain.CreatedAt); err != nil {
			return drivetrains, err
		}
		drivetrains = append(drivetrains, drivetrain)
	}

	return drivetrains, err
}

func (r *AdminRepository) CreateDrivetrain(ctx *fasthttp.RequestCtx, req *model.CreateDrivetrainRequest) (int, error) {
	q := `INSERT INTO drivetrains (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateDrivetrain(ctx *fasthttp.RequestCtx, id int, req *model.CreateDrivetrainRequest) error {
	q := `UPDATE drivetrains SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
	q := `SELECT id, name, name_ru, name_ae, created_at FROM fuel_types ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return fuelTypes, err
	}
	defer rows.Close()

	for rows.Next() {
		var fuelType model.AdminFuelTypeResponse
		if err := rows.Scan(&fuelType.ID, &fuelType.Name, &fuelType.NameRu, &fuelType.NameAe, &fuelType.CreatedAt); err != nil {
			return fuelTypes, err
		}
		fuelTypes = append(fuelTypes, fuelType)
	}

	return fuelTypes, err
}

func (r *AdminRepository) CreateFuelType(ctx *fasthttp.RequestCtx, req *model.CreateFuelTypeRequest) (int, error) {
	q := `INSERT INTO fuel_types (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateFuelType(ctx *fasthttp.RequestCtx, id int, req *model.CreateFuelTypeRequest) error {
	q := `UPDATE fuel_types SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteFuelType(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM fuel_types WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Generations CRUD operations
func (r *AdminRepository) GetGenerations(ctx *fasthttp.RequestCtx) ([]model.AdminGenerationResponse, error) {
	generations := make([]model.AdminGenerationResponse, 0)
	q := `
		SELECT g.id, g.name, g.name_ru, g.name_ae, g.model_id, m.name as model_name, m.name_ru as model_name_ru, g.start_year, g.end_year, g.wheel, $1 || g.image, g.created_at 
		FROM generations g
		LEFT JOIN models m ON g.model_id = m.id
		ORDER BY g.id DESC
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
	if err != nil {
		return generations, err
	}
	defer rows.Close()

	for rows.Next() {
		var generation model.AdminGenerationResponse
		if err := rows.Scan(
			&generation.ID, &generation.Name, &generation.NameRu, &generation.NameAe,
			&generation.ModelID, &generation.ModelName, &generation.ModelNameRu,
			&generation.StartYear, &generation.EndYear, &generation.Wheel,
			&generation.Image, &generation.CreatedAt); err != nil {
			return generations, err
		}
		generations = append(generations, generation)
	}

	return generations, err
}

// ValidateModelBelongsToBrand checks if a model belongs to a specific brand
func (r *AdminRepository) ValidateModelBelongsToBrand(ctx *fasthttp.RequestCtx, modelId, brandId int) (bool, error) {
	var count int
	q := `SELECT COUNT(*) FROM models WHERE id = $1 AND brand_id = $2`

	err := r.db.QueryRow(ctx, q, modelId, brandId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *AdminRepository) GetGenerationsByModel(ctx *fasthttp.RequestCtx, modelId int) ([]model.AdminGenerationResponse, error) {
	generations := make([]model.AdminGenerationResponse, 0)
	q := `
		SELECT 
			g.id, g.name, g.name_ru, g.name_ae,
			g.model_id, 
			m.name as model_name, 
			m.name_ru as model_name_ru,
			g.start_year, 
			g.end_year, 
			g.wheel, 
			$2 || g.image, 
			g.created_at 
		FROM generations g
		LEFT JOIN models m ON g.model_id = m.id
		WHERE g.model_id = $1
		ORDER BY g.id DESC
	`
	rows, err := r.db.Query(ctx, q, modelId, r.config.IMAGE_BASE_URL)

	if err != nil {
		return generations, err
	}

	defer rows.Close()

	for rows.Next() {
		var generation model.AdminGenerationResponse
		if err := rows.Scan(&generation.ID, &generation.Name, &generation.NameRu, &generation.NameAe, &generation.ModelID, &generation.ModelName, &generation.ModelNameRu,
			&generation.StartYear, &generation.EndYear, &generation.Wheel, &generation.Image, &generation.CreatedAt); err != nil {
			return generations, err
		}
		generations = append(generations, generation)
	}

	return generations, err
}

func (r *AdminRepository) CreateGeneration(ctx *fasthttp.RequestCtx, req *model.CreateGenerationRequest) (int, error) {
	q := `
		INSERT INTO generations 
			(name, name_ru, name_ae, model_id, start_year, end_year, wheel, image) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, req.ModelID, req.StartYear, req.EndYear, req.Wheel, req.Image).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateGeneration(ctx *fasthttp.RequestCtx, id int, req *model.UpdateGenerationRequest) error {
	q := `
		UPDATE generations 
			SET name = $2, name_ru = $3, name_ae = $4, model_id = $5, 
			start_year = $6, end_year = $7, wheel = $8
		WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe, req.ModelID, req.StartYear, req.EndYear, req.Wheel)
	return err
}

func (r *AdminRepository) DeleteGeneration(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM generations WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

func (r *AdminRepository) CreateGenerationImage(ctx *fasthttp.RequestCtx, id int, paths []string) error {
	q := `update generations set image = $2 where id = $1`
	_, err := r.db.Exec(ctx, q, id, paths[0])
	return err
}

// Generation Modifications CRUD operations
func (r *AdminRepository) GetGenerationModifications(ctx *fasthttp.RequestCtx, generationID int) ([]model.AdminGenerationModificationResponse, error) {
	modifications := make([]model.AdminGenerationModificationResponse, 0)
	q := `
		SELECT 
			gm.id, gm.generation_id,
			gm.body_type_id, bt.name as body_type_name, 
			bt.name_ru as body_type_name_ru,
			gm.engine_id, e.name as engine_name, e.name_ru as engine_name_ru,
			gm.fuel_type_id, ft.name as fuel_type_name,
			ft.name_ru as fuel_type_name_ru, gm.drivetrain_id, 
			dt.name as drivetrain_name, dt.name_ru as drivetrain_name_ru, 
			gm.transmission_id, t.name as transmission_name, 
			t.name_ru as transmission_name_ru
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
		if err := rows.Scan(
			&modification.ID, &modification.GenerationID,
			&modification.BodyTypeID, &modification.BodyTypeName,
			&modification.BodyTypeNameRu,
			&modification.EngineID, &modification.EngineName, &modification.EngineNameRu,
			&modification.FuelTypeID, &modification.FuelTypeName,
			&modification.FuelTypeNameRu,
			&modification.DrivetrainID, &modification.DrivetrainName,
			&modification.DrivetrainNameRu,
			&modification.TransmissionID, &modification.TransmissionName,
			&modification.TransmissionNameRu,
		); err != nil {
			return modifications, err
		}
		modifications = append(modifications, modification)
	}

	return modifications, err
}

func (r *AdminRepository) CreateGenerationModification(ctx *fasthttp.RequestCtx, generationID int, req *model.CreateGenerationModificationRequest) (int, error) {
	q := `
		INSERT INTO generation_modifications (
				generation_id, body_type_id, 
				engine_id, fuel_type_id, 
				drivetrain_id, transmission_id
		) VALUES (
		 	$1, $2, $3, $4, $5, $6
		) RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, q, generationID, req.BodyTypeID, req.EngineID, req.FuelTypeID, req.DrivetrainID, req.TransmissionID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateGenerationModification(
	ctx *fasthttp.RequestCtx, generationID int, id int, req *model.UpdateGenerationModificationRequest) error {
	q := `
		UPDATE generation_modifications 
			SET body_type_id = $3, engine_id = $4, 
			fuel_type_id = $5, drivetrain_id = $6, 
			transmission_id = $7 
		  WHERE 
		  	generation_id = $1 AND id = $2
	`
	_, err := r.db.Exec(ctx, q,
		generationID, id, req.BodyTypeID,
		req.EngineID, req.FuelTypeID,
		req.DrivetrainID, req.TransmissionID,
	)
	return err
}

func (r *AdminRepository) DeleteGenerationModification(ctx *fasthttp.RequestCtx, generationID int, id int) error {
	q := `
		DELETE FROM 
			generation_modifications 
		WHERE 
			generation_id = $1 AND id = $2`
	_, err := r.db.Exec(ctx, q, generationID, id)
	return err
}

// Colors CRUD operations
func (r *AdminRepository) GetColors(ctx *fasthttp.RequestCtx) ([]model.AdminColorResponse, error) {
	colors := make([]model.AdminColorResponse, 0)
	q := `
		SELECT 
			id, name, name_ru, name_ae, 
			$1 || image, created_at 
		FROM colors ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
	if err != nil {
		return colors, err
	}
	defer rows.Close()

	for rows.Next() {
		var color model.AdminColorResponse
		err := rows.Scan(
			&color.ID, &color.Name,
			&color.NameRu, &color.NameAe, &color.Image,
			&color.CreatedAt,
		)
		if err != nil {
			return colors, err
		}
		colors = append(colors, color)
	}

	return colors, err
}

func (r *AdminRepository) CreateColor(ctx *fasthttp.RequestCtx, req *model.CreateColorRequest) (int, error) {
	q := `INSERT INTO colors (name, name_ru, name_ae, image) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, "req.Image").Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateColorImage(ctx *fasthttp.RequestCtx, id int, path string) error {
	q := `UPDATE colors SET image = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, path)
	return err
}

func (r *AdminRepository) UpdateColor(ctx *fasthttp.RequestCtx, id int, req *model.UpdateColorRequest) error {
	q := `UPDATE colors SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteColor(ctx *fasthttp.RequestCtx, id int) error {
	// todo: return image path if exist
	q := `DELETE FROM colors WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Moto Categories CRUD operations
func (r *AdminRepository) GetMotoCategories(ctx *fasthttp.RequestCtx) ([]model.AdminMotoCategoryResponse, error) {
	motoCategories := make([]model.AdminMotoCategoryResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM moto_categories ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return motoCategories, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoCategory model.AdminMotoCategoryResponse
		if err := rows.Scan(&motoCategory.ID, &motoCategory.Name, &motoCategory.NameRu, &motoCategory.NameAe, &motoCategory.CreatedAt); err != nil {
			return motoCategories, err
		}
		motoCategories = append(motoCategories, motoCategory)
	}

	return motoCategories, err
}

func (r *AdminRepository) CreateMotoCategory(ctx *fasthttp.RequestCtx, req *model.CreateMotoCategoryRequest) (int, error) {
	q := `INSERT INTO moto_categories (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoCategory(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoCategoryRequest) error {
	q := `UPDATE moto_categories SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
		SELECT mb.id, mb.name, mb.name_ru, mb.name_ae, $1 || mb.logo, mb.created_at
		FROM moto_brands mb
		ORDER BY mb.id DESC
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
	if err != nil {
		return motoBrands, err
	}
	defer rows.Close()

	for rows.Next() {
		var motoBrand model.AdminMotoBrandResponse
		if err := rows.Scan(&motoBrand.ID, &motoBrand.Name, &motoBrand.NameRu, &motoBrand.NameAe, &motoBrand.Image, &motoBrand.CreatedAt); err != nil {
			return motoBrands, err
		}
		motoBrands = append(motoBrands, motoBrand)
	}

	return motoBrands, err
}

func (r *AdminRepository) GetMotoModelsByBrandID(ctx *fasthttp.RequestCtx, id int) ([]model.AdminMotoModelResponse, error) {
	motoModels := make([]model.AdminMotoModelResponse, 0)
	q := `
		SELECT mm.id, mm.name, mm.name_ru, mm.name_ae, mm.moto_brand_id, mb.name as moto_brand_name, mb.name_ru as moto_brand_name_ru, mm.created_at
		FROM moto_models mm
		LEFT JOIN moto_brands mb ON mm.moto_brand_id = mb.id
		WHERE mm.moto_brand_id = $1
		ORDER BY mm.id DESC`

	rows, err := r.db.Query(ctx, q, id)

	if err != nil {
		return motoModels, err
	}

	defer rows.Close()

	for rows.Next() {
		var motoModel model.AdminMotoModelResponse
		if err := rows.Scan(&motoModel.ID, &motoModel.Name, &motoModel.NameRu, &motoModel.NameAe, &motoModel.MotoBrandID,
			&motoModel.MotoBrandName, &motoModel.MotoBrandNameRu, &motoModel.CreatedAt); err != nil {
			return motoModels, err
		}
		motoModels = append(motoModels, motoModel)
	}

	return motoModels, err
}

func (r *AdminRepository) CreateMotoBrand(ctx *fasthttp.RequestCtx, req *model.CreateMotoBrandRequest) (int, error) {
	q := `INSERT INTO moto_brands (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateMotoBrandImage(ctx *fasthttp.RequestCtx, id int, path string) error {
	q := `UPDATE moto_brands SET logo = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, path)
	return err
}

func (r *AdminRepository) UpdateMotoBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoBrandRequest) error {
	q := `UPDATE moto_brands SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
		SELECT mm.id, mm.name, mm.name_ru, mm.name_ae, mm.moto_brand_id, mb.name as moto_brand_name, mb.name_ru as moto_brand_name_ru, mm.created_at
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
		if err := rows.Scan(&motoModel.ID, &motoModel.Name, &motoModel.NameRu, &motoModel.NameAe, &motoModel.MotoBrandID,
			&motoModel.MotoBrandName, &motoModel.MotoBrandNameRu, &motoModel.CreatedAt); err != nil {
			return motoModels, err
		}
		motoModels = append(motoModels, motoModel)
	}

	return motoModels, err
}

func (r *AdminRepository) CreateMotoModel(ctx *fasthttp.RequestCtx, req *model.CreateMotoModelRequest) (int, error) {
	q := `INSERT INTO moto_models (name, name_ru, name_ae, moto_brand_id) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, req.MotoBrandID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateMotoModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateMotoModelRequest) error {
	q := `UPDATE moto_models SET name = $2, name_ru = $3, name_ae = $4, moto_brand_id = $5 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe, req.MotoBrandID)
	return err
}

func (r *AdminRepository) DeleteMotoModel(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM moto_models WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Categories CRUD operations
func (r *AdminRepository) GetComtransCategories(ctx *fasthttp.RequestCtx) ([]model.AdminComtransCategoryResponse, error) {
	comtransCategories := make([]model.AdminComtransCategoryResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM com_categories ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return comtransCategories, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransCategory model.AdminComtransCategoryResponse
		if err := rows.Scan(&comtransCategory.ID, &comtransCategory.Name, &comtransCategory.NameRu, &comtransCategory.NameAe, &comtransCategory.CreatedAt); err != nil {
			return comtransCategories, err
		}
		comtransCategories = append(comtransCategories, comtransCategory)
	}

	return comtransCategories, err
}

func (r *AdminRepository) CreateComtransCategory(ctx *fasthttp.RequestCtx, req *model.CreateComtransCategoryRequest) (int, error) {
	q := `INSERT INTO com_categories (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransCategory(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransCategoryRequest) error {
	q := `UPDATE com_categories SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteComtransCategory(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_categories WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Comtrans Brands CRUD operations
func (r *AdminRepository) GetComtransBrands(ctx *fasthttp.RequestCtx) ([]model.AdminComtransBrandResponse, error) {
	comtransBrands := make([]model.AdminComtransBrandResponse, 0)
	q := `
		SELECT cb.id, cb.name, cb.name_ru, cb.name_ae, $1 || cb.logo, cb.created_at
		FROM com_brands cb
		ORDER BY cb.id DESC
	`

	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
	if err != nil {
		return comtransBrands, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransBrand model.AdminComtransBrandResponse
		if err := rows.Scan(&comtransBrand.ID, &comtransBrand.Name, &comtransBrand.NameRu, &comtransBrand.NameAe, &comtransBrand.Image, &comtransBrand.CreatedAt); err != nil {
			return comtransBrands, err
		}
		comtransBrands = append(comtransBrands, comtransBrand)
	}

	return comtransBrands, err
}

// Comtrans Models CRUD operations
func (r *AdminRepository) GetComtransModelsByBrandID(ctx *fasthttp.RequestCtx, id int) ([]model.AdminComtransModelResponse, error) {
	comtransModels := make([]model.AdminComtransModelResponse, 0)
	q := `
		SELECT cm.id, cm.name, cm.name_ru, cm.name_ae, cm.comtran_brand_id, cb.name as comtrans_brand_name, cb.name_ru as comtrans_brand_name_ru, cm.created_at
		FROM com_models cm
		LEFT JOIN com_brands cb ON cm.comtran_brand_id = cb.id
		WHERE cm.comtran_brand_id = $1
		ORDER BY cm.id DESC
	`

	rows, err := r.db.Query(ctx, q, id)
	if err != nil {
		return comtransModels, err
	}
	defer rows.Close()

	for rows.Next() {
		var comtransModel model.AdminComtransModelResponse
		if err := rows.Scan(&comtransModel.ID, &comtransModel.Name, &comtransModel.NameRu, &comtransModel.NameAe, &comtransModel.ComtransBrandID,
			&comtransModel.ComtransBrandName, &comtransModel.ComtransBrandNameRu, &comtransModel.CreatedAt); err != nil {
			return comtransModels, err
		}
		comtransModels = append(comtransModels, comtransModel)
	}

	return comtransModels, err
}

func (r *AdminRepository) CreateComtransBrand(ctx *fasthttp.RequestCtx, req *model.CreateComtransBrandRequest) (int, error) {
	q := `INSERT INTO com_brands (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateComtransBrandImage(ctx *fasthttp.RequestCtx, id int, path string) error {
	q := `UPDATE com_brands SET logo = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, path)
	return err
}

func (r *AdminRepository) UpdateComtransBrand(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransBrandRequest) error {
	q := `UPDATE com_brands SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
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
		SELECT cm.id, cm.name, cm.name_ru, cm.name_ae, cm.comtran_brand_id, cb.name as comtrans_brand_name, cb.name_ru as comtrans_brand_name_ru, cm.created_at
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
		if err := rows.Scan(&comtransModel.ID, &comtransModel.Name, &comtransModel.NameRu, &comtransModel.NameAe, &comtransModel.ComtransBrandID,
			&comtransModel.ComtransBrandName, &comtransModel.ComtransBrandNameRu, &comtransModel.CreatedAt); err != nil {
			return comtransModels, err
		}
		comtransModels = append(comtransModels, comtransModel)
	}

	return comtransModels, err
}

func (r *AdminRepository) CreateComtransModel(ctx *fasthttp.RequestCtx, req *model.CreateComtransModelRequest) (int, error) {
	q := `INSERT INTO com_models (name, name_ru, name_ae, comtran_brand_id) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, req.ComtransBrandID).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateComtransModel(ctx *fasthttp.RequestCtx, id int, req *model.UpdateComtransModelRequest) error {
	q := `UPDATE com_models SET name = $2, name_ru = $3, name_ae = $4, comtran_brand_id = $5 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe, req.ComtransBrandID)
	return err
}

func (r *AdminRepository) DeleteComtransModel(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM com_models WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Countries CRUD operations
func (r *AdminRepository) GetCountries(ctx *fasthttp.RequestCtx) ([]model.AdminCountryResponse, error) {
	q := `SELECT id, name, name_ru, name_ae, country_code, flag, created_at FROM countries ORDER BY id`
	rows, err := r.db.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var countries []model.AdminCountryResponse

	for rows.Next() {
		var country model.AdminCountryResponse

		if err := rows.Scan(&country.ID, &country.Name, &country.NameRu, &country.NameAe, &country.CountryCode, &country.Flag, &country.CreatedAt); err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}

	return countries, nil
}

func (r *AdminRepository) CreateCountry(ctx *fasthttp.RequestCtx, req *model.CreateNameRequest) (int, error) {
	q := `INSERT INTO countries (name, name_ru, name_ae, country_code, flag) VALUES ($1, $2, $3, $4, '') RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe, req.CountryCode).Scan(&id)
	return id, err
}

func (r *AdminRepository) CreateCountryImage(ctx *fasthttp.RequestCtx, id int, path string) error {
	q := `UPDATE countries SET flag = $2 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, path)
	return err
}

func (r *AdminRepository) UpdateCountry(ctx *fasthttp.RequestCtx, id int, req *model.CreateNameRequest) error {
	q := `UPDATE countries SET name = $2, name_ru = $3, name_ae = $4, country_code = $5 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe, req.CountryCode)
	return err
}

func (r *AdminRepository) DeleteCountry(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM countries WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Report CRUD operations
func (r *AdminRepository) GetReports(ctx *fasthttp.RequestCtx) ([]model.GetReportsResponse, error) {
	reports := make([]model.GetReportsResponse, 0)
	q := `
		SELECT 
			r.id,
			r.reported_user_id,
			r.report_type,
			r.report_description,
			r.report_status,
			r.created_at,
			json_build_object(
				'id', reporter_user.id,
				'username', reporter_profile.username,
				'avatar', CASE
					WHEN reporter_profile.avatar IS NULL OR reporter_profile.avatar = '' THEN NULL
					ELSE $1 || reporter_profile.avatar
				END,
				'role_id', reporter_user.role_id,
				'contacts', reporter_profile.contacts
			) as reporter,
			json_build_object(
				'id', reported_user.id,
				'username', reported_profile.username,
				'avatar', CASE
					WHEN reported_profile.avatar IS NULL OR reported_profile.avatar = '' THEN NULL
					ELSE $1 || reported_profile.avatar
				END,
				'role_id', reported_user.role_id,
				'contacts', reported_profile.contacts
			) as reported_user,
			r.item_type,
			r.item_id,
			CASE
				WHEN r.item_type = 'car' THEN (
					SELECT json_build_object(
						'id', v.id,
						'brand', b.name,
						'model', m.name,
						'price', v.price,
						'images', images.images
					)
					FROM vehicles v
					LEFT JOIN brands b ON v.brand_id = b.id
					LEFT JOIN models m ON v.model_id = m.id
					LEFT JOIN LATERAL (
						SELECT json_agg($1 || image) AS images
						FROM (
							SELECT image FROM images WHERE vehicle_id = v.id ORDER BY created_at DESC
						) img
					) images ON true
					WHERE v.id = r.item_id
				)
				WHEN r.item_type = 'moto' THEN (
					SELECT json_build_object(
						'id', m.id,
						'brand', mb.name,
						'model', mm.name,
						'price', m.price,
						'images', images.images
					)
					FROM motorcycles m
					LEFT JOIN moto_brands mb ON m.moto_brand_id = mb.id
					LEFT JOIN moto_models mm ON m.moto_model_id = mm.id
					LEFT JOIN LATERAL (
						SELECT json_agg($1 || image) AS images
						FROM (
							SELECT image FROM moto_images WHERE moto_id = m.id ORDER BY created_at DESC
						) img
					) images ON true
					WHERE m.id = r.item_id
				)
				WHEN r.item_type = 'comtran' THEN (
					SELECT json_build_object(
						'id', c.id,
						'brand', cb.name,
						'model', cm.name,
						'price', c.price,
						'images', images.images
					)
					FROM comtrans c
					LEFT JOIN com_brands cb ON c.comtran_brand_id = cb.id
					LEFT JOIN com_models cm ON c.comtran_model_id = cm.id
					LEFT JOIN LATERAL (
						SELECT json_agg($1 || image) AS images
						FROM (
							SELECT image FROM comtran_images WHERE comtran_id = c.id ORDER BY created_at DESC
						) img
					) images ON true
					WHERE c.id = r.item_id
				)
				ELSE NULL
			END as item
		FROM reports r
		LEFT JOIN users reporter_user ON reporter_user.id = r.user_id
		LEFT JOIN profiles reporter_profile ON reporter_profile.user_id = r.user_id
		LEFT JOIN users reported_user ON reported_user.id = r.reported_user_id
		LEFT JOIN profiles reported_profile ON reported_profile.user_id = r.reported_user_id
		ORDER BY r.created_at DESC
	`
	rows, err := r.db.Query(ctx, q, r.config.IMAGE_BASE_URL)
	if err != nil {
		return reports, err
	}
	defer rows.Close()

	for rows.Next() {
		var report model.GetReportsResponse
		if err := rows.Scan(&report.ID, &report.ReportedUserID, &report.ReportType, &report.ReportDescription, &report.ReportStatus, &report.CreatedAt, &report.Reporter, &report.ReportedUser, &report.ItemType, &report.ItemID, &report.Item); err != nil {
			return reports, err
		}
		reports = append(reports, report)
	}
	return reports, err
}

func (r *AdminRepository) GetReportByID(ctx *fasthttp.RequestCtx, id int) (model.GetReportsResponse, error) {
	report := model.GetReportsResponse{}
	q := `
		SELECT 
			r.id,
			r.reported_user_id,
			r.report_type,
			r.report_description,
			r.report_status,
			r.created_at,
			json_build_object(
				'id', reporter_user.id,
				'username', reporter_profile.username,
				'avatar', CASE
					WHEN reporter_profile.avatar IS NULL OR reporter_profile.avatar = '' THEN NULL
					ELSE $2 || reporter_profile.avatar
				END,
				'role_id', reporter_user.role_id,
				'contacts', reporter_profile.contacts
			) as reporter,
			json_build_object(
				'id', reported_user.id,
				'username', reported_profile.username,
				'avatar', CASE
					WHEN reported_profile.avatar IS NULL OR reported_profile.avatar = '' THEN NULL
					ELSE $2 || reported_profile.avatar
				END,
				'role_id', reported_user.role_id,
				'contacts', reported_profile.contacts
			) as reported_user,
			r.item_type,
			r.item_id,
			CASE
				WHEN r.item_type = 'car' THEN (
					SELECT json_build_object(
						'id', v.id,
						'brand', b.name,
						'model', m.name,
						'price', v.price,
						'images', images.images
					)
					FROM vehicles v
					LEFT JOIN brands b ON v.brand_id = b.id
					LEFT JOIN models m ON v.model_id = m.id
					LEFT JOIN LATERAL (
						SELECT json_agg($2 || image) AS images
						FROM (
							SELECT image FROM images WHERE vehicle_id = v.id ORDER BY created_at DESC
						) img
					) images ON true
					WHERE v.id = r.item_id
				)
				WHEN r.item_type = 'moto' THEN (
					SELECT json_build_object(
						'id', m.id,
						'brand', mb.name,
						'model', mm.name,
						'price', m.price,
						'images', images.images
					)
					FROM motorcycles m
					LEFT JOIN moto_brands mb ON m.moto_brand_id = mb.id
					LEFT JOIN moto_models mm ON m.moto_model_id = mm.id
					LEFT JOIN LATERAL (
						SELECT json_agg($2 || image) AS images
						FROM (
							SELECT image FROM moto_images WHERE moto_id = m.id ORDER BY created_at DESC
						) img
					) images ON true
					WHERE m.id = r.item_id
				)
				WHEN r.item_type = 'comtran' THEN (
					SELECT json_build_object(
						'id', c.id,
						'brand', cb.name,
						'model', cm.name,
						'price', c.price,
						'images', images.images
					)
					FROM comtrans c
					LEFT JOIN com_brands cb ON c.comtran_brand_id = cb.id
					LEFT JOIN com_models cm ON c.comtran_model_id = cm.id
					LEFT JOIN LATERAL (
						SELECT json_agg($2 || image) AS images
						FROM (
							SELECT image FROM comtran_images WHERE comtran_id = c.id ORDER BY created_at DESC
						) img
					) images ON true
					WHERE c.id = r.item_id
				)
				ELSE NULL
			END as item
		FROM reports r
		LEFT JOIN users reporter_user ON reporter_user.id = r.user_id
		LEFT JOIN profiles reporter_profile ON reporter_profile.user_id = r.user_id
		LEFT JOIN users reported_user ON reported_user.id = r.reported_user_id
		LEFT JOIN profiles reported_profile ON reported_profile.user_id = r.reported_user_id
		WHERE r.id = $1
	`
	err := r.db.QueryRow(ctx, q, id, r.config.IMAGE_BASE_URL).Scan(&report.ID, &report.ReportedUserID, &report.ReportType, &report.ReportDescription, &report.ReportStatus, &report.CreatedAt, &report.Reporter, &report.ReportedUser, &report.ItemType, &report.ItemID, &report.Item)
	return report, err
}

func (r *AdminRepository) UpdateReport(ctx *fasthttp.RequestCtx, id int, req *model.UpdateReportRequest) error {
	q := `
		UPDATE reports
		SET report_status = $1
		WHERE id = $2
	`
	_, err := r.db.Exec(ctx, q, req.ReportStatus, id)
	return err
}

func (r *AdminRepository) DeleteReport(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM reports WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// Number of cycles CRUD operations
func (r *AdminRepository) GetNumberOfCycles(ctx *fasthttp.RequestCtx) ([]model.AdminNumberOfCycleResponse, error) {
	numberOfCycles := make([]model.AdminNumberOfCycleResponse, 0)
	q := `SELECT id, name, name_ru, name_ae, created_at FROM number_of_cycles ORDER BY id DESC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return numberOfCycles, err
	}
	defer rows.Close()

	for rows.Next() {
		var numberOfCycle model.AdminNumberOfCycleResponse
		if err := rows.Scan(&numberOfCycle.ID, &numberOfCycle.Name, &numberOfCycle.NameRu, &numberOfCycle.NameAe, &numberOfCycle.CreatedAt); err != nil {
			return numberOfCycles, err
		}
		numberOfCycles = append(numberOfCycles, numberOfCycle)
	}

	return numberOfCycles, err
}

func (r *AdminRepository) CreateNumberOfCycle(ctx *fasthttp.RequestCtx, req *model.CreateNumberOfCycleRequest) (int, error) {
	q := `INSERT INTO number_of_cycles (name, name_ru, name_ae) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, q, req.Name, req.NameRu, req.NameAe).Scan(&id)
	return id, err
}

func (r *AdminRepository) UpdateNumberOfCycle(ctx *fasthttp.RequestCtx, id int, req *model.CreateNumberOfCycleRequest) error {
	q := `UPDATE number_of_cycles SET name = $2, name_ru = $3, name_ae = $4 WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id, req.Name, req.NameRu, req.NameAe)
	return err
}

func (r *AdminRepository) DeleteNumberOfCycle(ctx *fasthttp.RequestCtx, id int) error {
	q := `DELETE FROM number_of_cycles WHERE id = $1`
	_, err := r.db.Exec(ctx, q, id)
	return err
}

// GetDeviceTokensByRoleID returns FCM device tokens for all users with the given role_id.
func (r *AdminRepository) GetDeviceTokensByRoleID(ctx context.Context, roleID int) ([]string, error) {
	q := `
		SELECT ut.device_token
		FROM user_tokens ut
		INNER JOIN users u ON u.id = ut.user_id
		WHERE u.role_id = $1
		  AND ut.device_token IS NOT NULL
	`
	rows, err := r.db.Query(ctx, q, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []string
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	return tokens, rows.Err()
}
