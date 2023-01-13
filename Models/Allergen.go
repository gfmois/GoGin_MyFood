package Models

import (
	"gx_myfood/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllergens(allergen *[]Allergen) (err error) {
	if err := Config.DB.Find(allergen).Error; err != nil {
		return err
	}

	return nil
}

func GetAllergeno(allergen *Allergen, id_allergen string) (err error) {
	if err := Config.DB.Find(allergen, "id_alergeno = ?", id_allergen).Error; err != nil {
		return err
	}

	return nil
}
