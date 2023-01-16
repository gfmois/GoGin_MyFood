package Models

import (
	"gx_myfood/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetCategories(categories *[]Category) (err error) {
	if err := Config.DB.Preload("Productos").Find(&categories).Error; err != nil {
		return err
	}

	return nil
}

func GetCategoryById(category *Category, id_categoria string) (err error) {
	if err := Config.DB.Preload("Productos").Find(category, "id_categoria = ?", id_categoria).Error; err != nil {
		return err
	}

	return nil
}
