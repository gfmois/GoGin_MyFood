package Models

import (
	"gx_myfood/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetProducts(product *[]Product) (err error) {
	if err := Config.DB.Find(product).Error; err != nil {
		return err
	}

	return nil
}

func GetProductById(product *Product, id_producto string) (err error) {
	if err := Config.DB.Find(product, "id_producto = ?", id_producto).Error; err != nil {
		return err
	}

	return nil
}
