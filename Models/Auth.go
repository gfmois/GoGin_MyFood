package Models

import "gx_myfood/Config"

func Register(client *Client) (err error) {
	if err := Config.DB.Create(client).Error; err != nil {
		return err
	}
	return nil
}

func Login(client *Client) (err error) {
	if err := Config.DB.Where("email = ?", client.Email).Find(client).Error; err != nil {
		return err
	}
	return nil
}

func GetUserInfo(client *Client, id_client string) (err error) {
	if err := Config.DB.Where("id_cliente = ?", id_client).Find(client).Error; err != nil {
		return err
	}
	return nil
}
