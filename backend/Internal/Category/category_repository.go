package category

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) saveCategory(c Category) (*Category, error) {
	var category Category
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO category "+
		"(name,nsfw) "+
		"VALUES ($1, $2)",
		c.Name, c.Nsfw).
		Scan(&category.Id, &category.Name, &category.Nsfw)

	if err != nil {
		err = fmt.Errorf("failed query, could not save category: - %w", err)
	}
	return &category, nil
}

func (repo *Repository) getCategoryById(id uint64) (*Category, error) {
	var category Category
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM category "+
		"WHERE id = $1", id).Scan(
		&category.Id, &category.Name, &category.Nsfw)

	if err != nil {
		err = fmt.Errorf("failed query, could not get category with ID: - %w", err)
		return nil, err
	}

	return &category, nil
}

func (repo *Repository) getAllCategories() ([]Category, error) {
	rows, err := repo.DB.Query(context.Background(), "SELECT * FROM category")

	if err != nil {
		err = fmt.Errorf("failed query, couldn't get all categories: - %w", err)
		return nil, err
	}

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[Category])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, couldn't get rows or parse them: - %w", err)
		return nil, err
	}

	return categories, nil
}

func (repo *Repository) updateCategory(c Category) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE category "+
		"SET name = $2, nsfw = $3 WHERE id = $1", c.Id, c.Name, c.Nsfw)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't update category: - %w", err)
		return err
	}

	return nil
}

func (repo *Repository) deleteCategoryById(id uint64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM category WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't delete category: - %w", err)
		return err
	}

	return nil
}

func (repo *Repository) getCategoryByName(name string) (*Category, error) {
	var category Category
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM category WHERE name = $1", name).
		Scan(&category.Id, &category.Name, &category.Nsfw)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't get an specific category: - %w", err)
		return nil, err
	}

	return &category, nil
}
