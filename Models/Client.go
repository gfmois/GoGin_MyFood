package Models

import "gx_myfood/Config"

func UpdateProfile(client *Client) (err error) {
	if err := Config.DB.Save(&client).Error; err != nil {
		return err
	}
	return nil
}
