package Controller

import (
	"gx_myfood/Common"
	"gx_myfood/Models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orders []Models.Order
	client_id, _ := c.Get("client_id")
	if err := Models.GetOrders(&orders, client_id.(string)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	for i := range orders {
		productos_pedidos := []Models.ProductOrder{}
		for j := range orders[i].Productos {
			productos_pedidos = append(productos_pedidos, Models.ProductOrder{
				Cantidad: orders[i].Productos[j].Cantidad,
				Producto: orders[i].Productos[j],
			})
			orders[i].Productos[j].Cantidad = 0
		}
		for k := range productos_pedidos {
			productos_pedidos[k].Producto.Cantidad = 0
		}
		orders[i].ProductosPedidos = productos_pedidos
		orders[i].Fecha = strings.Split(orders[i].Fecha, "T")[0]
	}

	c.JSON(http.StatusOK, gin.H{"orders": &orders})
}

func AddOrder(c *gin.Context) {
	var products []Models.ProductToSave
	var productsOrder []Models.ProductOrderToSave
	var newOrder Models.OrderRequest
	client_id, _ := c.Get("client_id")
	c.BindJSON(&products)

	newOrder.Id_pedido = Common.GenerateRandomIdWithLength(10)
	newOrder.Id_cliente = client_id.(string)
	newOrder.Fecha = time.Now().Format("2006-01-02")
	newOrder.Estado = "PENDIENTE"

	for i := range products {
		productsOrder = append(productsOrder, Models.ProductOrderToSave{
			Id_producto: products[i].Id_producto,
			Id_pedido:   newOrder.Id_pedido,
			Cantidad:    products[i].Cantidad,
		})
	}

	if err := Models.AddOrder(&newOrder, productsOrder); err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	c.JSON(http.StatusOK, newOrder.Id_pedido)
}
