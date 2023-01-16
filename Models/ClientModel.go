package Models

type Client struct {
	Id_cliente string `gorm:"primary_key" json:"id_cliente"`
	Nombre     string `json:"nombre"`
	Email      string `json:"email"`
	Telefono   string `json:"telefono"`
	Contraseña string `json:"contraseña"`
	Avatar     string `json:"avatar"`
}

func (c *Client) TableName() string {
	return "clientes"
}
