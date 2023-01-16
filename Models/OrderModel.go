package Models

type Order struct {
	Id_pedido        string           `gorm:"primary_key" json:"id_pedido"`
	Id_cliente       string           `json:"id_cliente"`
	Fecha            string           `json:"fecha"`
	Estado           string           `json:"estado"`
	Productos        []ProductToOrder `gorm:"many2many:productos_pedidos;association_foreignkey:id_producto;foreignkey:id_pedido;association_jointable_foreignkey:id_producto;jointable_foreignkey:id_pedido;" json:"-"`
	ProductosPedidos []ProductOrder   `json:"productos_pedidos"`
}

type OrderRequest struct {
	Id_pedido  string          `gorm:"primary_key" json:"id_pedido"`
	Id_cliente string          `json:"id_cliente"`
	Fecha      string          `json:"fecha"`
	Estado     string          `json:"estado"`
	Productos  []ProductToSave `gorm:"many2many:productos_pedidos;association_foreignkey:id_producto;foreignkey:id_pedido;association_jointable_foreignkey:id_producto;jointable_foreignkey:id_pedido;" json:"productos"`
}

func (o *OrderRequest) TableName() string {
	return "pedidos"
}

func (o *Order) TableName() string {
	return "pedidos"
}
