package Models

import (
	"gx_myfood/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetReserves(reserve *[]Reserve, id_client string) (err error) {
	if err = Config.DB.Where("id_cliente = ?", id_client).Find(reserve).Error; err != nil {
		return err
	}

	return nil
}

func GetBannedDays(bannedDays *[]string, comensales string, servicio string) (err error) {
	var reserves = []Reserve{}
	var holidays []struct {
		Fecha string `json:"fecha"`
	}

	if r_err := Config.DB.Raw("SELECT r.fecha FROM reservas r GROUP BY r.fecha, r.tipo HAVING (SUM(r.n_comensales) + ?) > 50 AND r.tipo = ?", comensales, servicio).Scan(&reserves).Error; err != nil {
		return r_err
	}

	if h_err := Config.DB.Raw("SELECT fecha FROM festivos").Find(&holidays).Error; h_err != nil {
		return h_err
	}

	for _, value := range reserves {
		*bannedDays = append(*bannedDays, value.Fecha)
	}

	for _, value := range holidays {
		*bannedDays = append(*bannedDays, value.Fecha)
	}

	return nil
}

func GetReserve(reserve *Reserve, id_reserva string) (err error) {
	if err = Config.DB.Find(reserve, "id_reserva = ?", id_reserva).Error; err != nil {
		return err
	}

	return nil
}

func CreateReserve(newReserve *Reserve) (err error) {
	if err = Config.DB.Create(newReserve).Error; err != nil {
		return err
	}
	return nil
}

func (b *Reserve) SetId_Reserva(newId string) {
	b.Id_reserva = newId
}
