package Models

type Reserve struct {
	Id_reserva   string `json:"id_reserva"`
	Id_cliente   string `json:"id_cliente"`
	Fecha        string `json:"fecha"`
	Tipo         string `json:"tipo"`
	N_comensales int    `json:"n_comensales"`
	Estado       string `json:"estado"`
}

type ReserveInput struct {
	Id_reserva   string `json:"id_reserva"`
	Id_cliente   string `json:"id_cliente"`
	Fecha        string `json:"fecha"`
	Tipo         string `json:"servicio"`
	N_comensales int    `json:"n_comensales"`
	Estado       string `json:"estado"`
}

func (b *Reserve) TableName() string {
	return "reservas"
}
