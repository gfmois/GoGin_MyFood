package Models

type Allergen struct {
	Id_alergeno string    `json:"id_alergeno"`
	Nombre      string    `json:"nombre"`
	Icono       string    `json:"icono"`
	Productos   []Product `gorm:"many2many:productos_alergenos;association_foreignkey:id_producto;foreignkey:id_alergeno;association_jointable_foreignkey:id_producto;jointable_foreignkey:id_alergeno;" json:"-"`
	C_productos []string  `json:"productos"`
}

func (b *Allergen) TableName() string {
	return "alergenos"
}
