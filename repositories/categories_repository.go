package repositories

import (
	"database/sql"
	"errors"
	"tugas/models"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (repo *CategoriesRepository) GetAll() ([]models.Barang, error) {
	query := "SELECT id, category_name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]models.Barang, 0)
	for rows.Next() {
		var p models.Barang
		err := rows.Scan(&p.ID, &p.CategoryName, &p.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, p)
	}

	return categories, nil
}

func (repo *CategoriesRepository) GetByID(id int) (*models.Barang, error) {
	query := "SELECT id, category_name, description FROM categories WHERE id = $1"

	var p models.Barang
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.CategoryName, &p.Description)

	if err == sql.ErrNoRows {
		return nil, errors.New("Kategori tidak ditemukan.")
	}

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *CategoriesRepository) Create(categories *models.Barang) error {
	query := "INSERT INTO categories (category_name, description) VALUES ($1, $2) RETURNING ID"
	err := repo.db.QueryRow(query, categories.CategoryName, categories.Description).Scan(&categories.ID)

	return err
}

func (repo *CategoriesRepository) Update(category *models.Barang) error {
	query := "UPDATE categories SET category_name = $1, description = $2 WHERE id = $3"
	result, err := repo.db.Exec(query, category.CategoryName, category.Description, category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Categories tidak ditemukan")
	}

	return nil
}

func (repo *CategoriesRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("Categories tidak ditemukan")
	}

	return err
}
