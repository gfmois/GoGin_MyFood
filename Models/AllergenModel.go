package Models

type Allergen struct {
	Id_alergeno string `json:"id_alergeno"`
	Nombre      string `json:"nombre"`
	Imagen      string `json:"imagen"`
}

func (b *Allergen) TableName() string {
	return "alergenos"
}
