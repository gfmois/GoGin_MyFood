package Models

import "gx_myfood/Config"

func GetOrders(orders *[]Order, id_client string) (err error) {
	if err := Config.DB.Preload("Productos").Where("id_cliente = ?", id_client).Find(&orders).Error; err != nil {
		return err
	}
	return nil
}

func AddOrder(order *OrderRequest, products []ProductOrderToSave) (err error) {
	if err := Config.DB.Create(&order).Error; err != nil {
		return err
	}

	for _, product := range products {
		Config.DB.Model(&order).Association("Productos").Append(product)
	}

	return nil
}
