package Models

type ProductOrder struct {
	Producto ProductToOrder `json:"producto"`
	Cantidad int            `json:"cantidad"`
}

type ProductOrderToSave struct {
	Id_pedido   string `json:"id_pedido"`
	Id_producto string `json:"id_producto"`
	Cantidad    int    `json:"cantidad"`
}

func (o *ProductOrderToSave) TableName() string {
	return "productos_pedidos"
}
