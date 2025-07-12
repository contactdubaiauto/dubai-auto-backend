package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xuri/excelize/v2"
)

func ExcelMigrate(db *pgxpool.Pool) error {
	// err := Marks("docs/1.marks.xlsx", db)
	// fmt.Println(err)
	// err = Models("docs/2.models.xlsx", db)
	// fmt.Println(err)
	// err = Generations("docs/3.generations.xlsx", db)
	// fmt.Println(err)
	// err = Configurations("docs/4.configurations.xlsx", db)
	// fmt.Println(err)
	// err = Complectations("docs/5.complectations.xlsx", db)
	// fmt.Println(err)
	return nil
}

func Marks(filePath string, db *pgxpool.Pool) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	if sheetName == "" {
		return fmt.Errorf("no sheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)

	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	q := `
		insert into marks_exel(
			id, inner_id, name, cyrillic_name, numeric_id, logo, big_logo, year_from, year_to
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	for i := range rows {
		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(rows[i][0])
		numeric_id, _ := strconv.Atoi(rows[i][4])
		year_from, _ := strconv.Atoi(rows[i][7])
		year_to, _ := strconv.Atoi(rows[i][8])
		_, err = db.Exec(context.Background(), q,
			id, rows[i][1], rows[i][2], rows[i][3],
			numeric_id, rows[i][5], rows[i][6], year_from, year_to)

		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func Models(filePath string, db *pgxpool.Pool) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	if sheetName == "" {
		return fmt.Errorf("no sheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)

	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	q := `
		insert into models_exel(
			id, inner_id, name, cyrillic_name, year_from, year_to, autoru_mark_id, autoru_mark_inner_id
		) values ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	for i := range rows {

		if i == 0 {
			continue
		}
		id, _ := strconv.Atoi(rows[i][0])
		autoru_mark_id, _ := strconv.Atoi(rows[i][6])
		year_from, _ := strconv.Atoi(rows[i][4])
		year_to, _ := strconv.Atoi(rows[i][5])

		_, err = db.Exec(context.Background(), q,
			id, rows[i][1], rows[i][2], rows[i][3],
			year_from, year_to, autoru_mark_id, rows[i][7])

		if err != nil {
			fmt.Println(rows[i])
			fmt.Println(err)
		}
	}
	return nil
}

func Generations(filePath string, db *pgxpool.Pool) error {
	f, err := excelize.OpenFile(filePath)

	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	if sheetName == "" {
		return fmt.Errorf("no sheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)

	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	q := `
		insert into generations_exel (
			id, inner_id, name, year_from, year_to, photo_main, 
			photo_main_minicard, photo_mobile, autoru_model_id, 
			autoru_mark_id, autoru_mark_inner_id, autoru_model_inner_id
		) values (
		 	$1, $2, $3, $4, $5, $6, 
		 	$7, $8, $9,
			$10, $11, $12)
	`

	for i := range rows {

		if i == 0 {
			continue
		}
		if i == 479 {
			break
		}
		id, _ := strconv.Atoi(rows[i][0])
		inner_id, _ := strconv.Atoi(rows[i][1])
		autoru_model_id, _ := strconv.Atoi(rows[i][8])
		autoru_mark_id, _ := strconv.Atoi(rows[i][9])
		year_from, _ := strconv.Atoi(rows[i][3])
		year_to, _ := strconv.Atoi(rows[i][4])

		_, err = db.Exec(context.Background(), q,
			id, inner_id, rows[i][2], year_from, year_to, rows[i][5],
			rows[i][6], rows[i][7], autoru_model_id,
			autoru_mark_id, rows[i][10], rows[i][11])

		if err != nil {
			fmt.Println(rows[i][11])
			fmt.Println(autoru_model_id)
			fmt.Println(err)
		}
	}
	return nil
}

func Configurations(filePath string, db *pgxpool.Pool) error {
	f, err := excelize.OpenFile(filePath)

	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	if sheetName == "" {
		return fmt.Errorf("no sheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)

	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	q := `
		insert into configurations_exel (
			id, inner_id, name, autoru_body_type, body_type, photo_main, 
			photo_main_minicard, photo_mobile, autoru_generation_id, 
			autoru_mark_id, autoru_mark_inner_id, autoru_generation_inner_id
		) values (
		 	$1, $2, $3, $4, $5, $6, 
		 	$7, $8, $9,
			$10, $11, $12)
	`

	for i := range rows {

		if i == 0 {
			continue
		}

		id, _ := strconv.Atoi(rows[i][0])
		inner_id, _ := strconv.Atoi(rows[i][1])
		autoru_generation_id, _ := strconv.Atoi(rows[i][8])
		autoru_generation_inner_id, _ := strconv.Atoi(rows[i][11])
		autoru_mark_id, _ := strconv.Atoi(rows[i][9])

		_, err = db.Exec(context.Background(), q,
			id, inner_id, rows[i][2], rows[i][3], rows[i][4], rows[i][5],
			rows[i][6], rows[i][7], autoru_generation_id,
			autoru_mark_id, rows[i][10], autoru_generation_inner_id)

		if err != nil {
			fmt.Println("98s7fdyh8")
			fmt.Println(rows[i])
			fmt.Println(err)
		}
	}
	return nil
}

func Complectations(filePath string, db *pgxpool.Pool) error {
	f, err := excelize.OpenFile(filePath)

	if err != nil {
		return fmt.Errorf("failed to open Excel file: %w", err)
	}
	defer f.Close()

	sheetName := f.GetSheetName(0)

	if sheetName == "" {
		return fmt.Errorf("no sheets found in Excel file")
	}

	rows, err := f.GetRows(sheetName)

	if err != nil {
		return fmt.Errorf("failed to get rows: %w", err)
	}

	q := `
		insert into complectations_exel (
			id, inner_id, name, full_name, main_fields, packages, 
			specifications, autoru_configuration_id, 
			autoru_mark_id, autoru_mark_inner_id, autoru_configuration_inner_id
		) values (
		 	$1, $2, $3, $4, $5, $6, 
		 	$7, $8, 
			$9, $10, $11)
	`

	for i := range rows {

		if i == 0 {
			continue
		}

		id, _ := strconv.Atoi(rows[i][0])
		inner_id, _ := strconv.Atoi(rows[i][1])
		autoru_configuration_id, _ := strconv.Atoi(rows[i][7])
		autoru_configuration_inner_id, _ := strconv.Atoi(rows[i][10])
		autoru_mark_id, _ := strconv.Atoi(rows[i][8])

		_, err = db.Exec(context.Background(), q,
			id, inner_id, rows[i][2], rows[i][3], rows[i][4], rows[i][5],
			rows[i][6], autoru_configuration_id,
			autoru_mark_id, rows[i][9], autoru_configuration_inner_id)

		if err != nil {
			fmt.Println("74787y87")
			fmt.Println(rows[i][8])
			fmt.Println(err)
		}
	}
	return nil
}
