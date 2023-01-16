package Models

import (
	"gx_myfood/Config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetProducts(products *[]Product) (err error) {
	if err := Config.DB.Preload("Categorias").Preload("Alergenos").Find(&products).Error; err != nil {
		return err
	}

	return nil
}

func GetFilteredProducts(products *[]Product, categorias []string, orden string, rango1 string, rango2 string, paginacion string) (err error) {
	intRango1, _ := strconv.Atoi(rango1)
	intRango2, _ := strconv.Atoi(rango2)
	var intPaginacion int
	if paginacion != "" {
		intPaginacion, _ = strconv.Atoi(paginacion)
		intPaginacion = (intPaginacion - 1) * 12
	}
	query := Config.DB.Preload("Categorias").Preload("Alergenos").Table("productos").Where("precio BETWEEN ? AND ?", intRango1, intRango2).Order("precio " + orden)
	if paginacion != "" {
		query = query.Limit(12).Offset(intPaginacion)
	}
	if len(categorias) != 0 {
		query = query.Joins("INNER JOIN categorias_productos ON productos.id_producto = categorias_productos.id_producto").Where("categorias_productos.id_categoria IN(?)", categorias)
	}
	err = query.Find(&products).Error
	return err
}

func GetProductBySlug(product *Product, slug_producto string) (err error) {
	if err := Config.DB.Preload("Categorias").Preload("Alergenos").Find(product, "slug = ?", slug_producto).Error; err != nil {
		return err
	}

	return nil
}

func SearchProducts(products *[]Product, producto string) (err error) {
	if err := Config.DB.Preload("Categorias").Preload("Alergenos").Where("nombre LIKE ?", "%"+producto+"%").Find(products).Error; err != nil {
		return err
	}
	return nil
}
