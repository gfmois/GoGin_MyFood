package Models

import (
	"fmt"
	"time"

	"gx_myfood/Config"

	_ "github.com/go-sql-driver/mysql"
)

func GetReserves(reserve *[]Reserve) (err error) {
	if err = Config.DB.Find(reserve).Error; err != nil {
		return err
	}

	return nil
}

func GetBannedDays(bannedDays *[]string, comensales string, servicio string) (err error) {
	var reserves = []Reserve{}
	var holidays []time.Time

	if r_err := Config.DB.Raw("SELECT r.fecha FROM reservas r GROUP BY r.fecha, r.tipo HAVING (SUM(r.n_comensales) + ?) > 50 AND r.tipo = ?", comensales, servicio).Scan(&reserves).Error; err != nil {
		return r_err
	}

	if h_err := Config.DB.Raw("SELECT fecha FROM festivos").Find(&holidays).Error; h_err != nil {
		return h_err
	}

	fmt.Println(holidays)

	for _, value := range reserves {
		*bannedDays = append(*bannedDays, value.Fecha)
	}

	return nil
}
