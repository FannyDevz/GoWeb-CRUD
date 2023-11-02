package categoriesmodel

import (
	"goweb/config"
	"goweb/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Type, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func Add(category entities.Category) bool {

	result, err := config.DB.Exec(
		"INSERT INTO categories (name, type, created_at, updated_at) VALUES (?, ?, ?, ?)",
		category.Name,
		category.Type,
		category.CreatedAt,
		category.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}
	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return LastInsertId > 0
}

func GetById(id string) entities.Category {
	var category entities.Category
	row := config.DB.QueryRow("SELECT * FROM categories WHERE id=?", id)
	if err := row.Scan(&category.Id, &category.Name, &category.Type, &category.CreatedAt, &category.UpdatedAt); err != nil {
		panic(err)
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec("UPDATE categories SET name=?, type=?, updated_at=? WHERE id=?", category.Name, category.Type, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}
