package Models

import (
	"encoding/json"
)

type NullableInt struct {
	Value int
	Valid bool
}

func (ni NullableInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Value)
}

type Product struct {
	Id_producto  string      `gorm:"primary_key" json:"id_producto"`
	Imagen       string      `json:"imagen"`
	Nombre       string      `json:"nombre"`
	Precio       float32     `json:"precio"`
	Slug         string      `json:"slug"`
	Cantidad     NullableInt `json:"cantidad"`
	Categorias   []Category  `gorm:"many2many:categorias_productos;association_foreignkey:id_categoria;foreignkey:id_producto;association_jointable_foreignkey:id_categoria;jointable_foreignkey:id_producto;" json:"-"`
	C_categorias []string    `json:"categorias"`
	Alergenos    []Allergen  `gorm:"many2many:productos_alergenos;association_foreignkey:id_alergeno;foreignkey:id_producto;association_jointable_foreignkey:id_alergeno;jointable_foreignkey:id_producto;" json:"-"`
	C_alergenos  []string    `json:"alergenos"`
}

type ProductToOrder struct {
	Id_producto  string     `gorm:"primary_key" json:"id_producto"`
	Imagen       string     `json:"imagen"`
	Nombre       string     `json:"nombre"`
	Precio       float32    `json:"precio"`
	Slug         string     `json:"slug"`
	Cantidad     int        `json:"cantidad"`
	Categorias   []Category `gorm:"many2many:categorias_productos;association_foreignkey:id_categoria;foreignkey:id_producto;association_jointable_foreignkey:id_categoria;jointable_foreignkey:id_producto;" json:"-"`
	C_categorias []string   `json:"categorias"`
	Alergenos    []Allergen `gorm:"many2many:productos_alergenos;association_foreignkey:id_alergeno;foreignkey:id_producto;association_jointable_foreignkey:id_alergeno;jointable_foreignkey:id_producto;" json:"-"`
	C_alergenos  []string   `json:"alergenos"`
}

type ProductToSave struct {
	Id_producto string  `gorm:"primary_key" json:"id_producto"`
	Imagen      string  `json:"imagen"`
	Nombre      string  `json:"nombre"`
	Precio      float32 `json:"precio"`
	Slug        string  `json:"slug"`
	Cantidad    int     `json:"cantidad"`
}

func (p *Product) TableName() string {
	return "productos"
}

func (p *ProductToOrder) TableName() string {
	return "productos"
}

func (p *ProductToSave) TableName() string {
	return "productos"
}
