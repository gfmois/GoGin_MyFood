package Models

type Product struct {
	Id_producto string `json:"id_producto"`
	Slug        string `json:"slug"`
	Nombre      string `json:"nombre"`
	Precio      string `json:"precio"`
	Imagen      string `json:"imagen"`
}

func (b *Product) TableName() string {
	return "productos"
}
