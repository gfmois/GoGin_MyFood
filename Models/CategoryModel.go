package Models

type Category struct {
	Id_categoria string    `gorm:"primary_key" json:"id_categoria"`
	Slug         string    `json:"slug"`
	Nombre       string    `json:"nombre"`
	Icono        string    `json:"icono"`
	Productos    []Product `gorm:"many2many:categorias_productos;association_foreignkey:id_producto;foreignkey:id_categoria;association_jointable_foreignkey:id_producto;jointable_foreignkey:id_categoria;" json:"-"`
	C_productos  []string  `json:"productos"`
}

func (b *Category) TableName() string {
	return "categorias"
}
