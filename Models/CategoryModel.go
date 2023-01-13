package Models

type Category struct {
	Id_cateogoria string `json:"id_categoria"`
	Slug          string `json:"slug"`
	Nombre        string `json:"nombre"`
	Icono         string `json:"icono"`
}

func (b *Category) TableName() string {
	return "categorias"
}
