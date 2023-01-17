<p align="center"><a href="https://go.dev" target="_blank"><img src="https://imgs.search.brave.com/UcWtu8lFXd-5WgCgtmkojZPNYxW7Emv_zn12TdenVEg/rs:fit:1200:535:1/g:ce/aHR0cHM6Ly93d3cu/dHJpc29mdC5yby9i/dW5kbGVzL2FwcGZy/b250ZW5kL2ltZy9H/b2xhbmctbG9nby5w/bmc" width="400" alt="Laravel Logo"></a></p>

[go]: https://go.dev/learn/
[gorm]: https://gorm.io/docs/
[myfood]: http://github.com/gfmois/Vue_Laravel_Go_Spring_MyFood.git

## ABOUT THIS PROJECT

> This server is intended for use by customers, who can make reservations for lunch, dinner or take out at the restaurant, create an account to order online, or modify their profile and view their order statistics for the year or the reservations made.

## BUILT WITH

- [Go] - Framework
- [Gorm] - ORM

## API

URL
http://localhost:8080/api/

The available endpoints are:

<table>
  <tr>
    <th>Function</th>
    <th>Request</th>
    <th>Route</th>
    <th>Params</th>
    <th>Body</th>
    <th>Response</th>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Clients</th>
  </tr>
  <tr>
    <td>GetProfile()</td>
    <td>GET</td>
    <td>/client/profile</td>
    <td>{ header: Token }</td>
    <td></td>
    <td>[ User Information + Token ]</td>
  </tr>
  <tr>
    <td>UpdateProfile()</td>
    <td>PUT</td>
    <td>/client/profile</td>
    <td>{ Heder: Token }</td>
    <td>New User Information</td>
    <td>[ User Information + Token ]</td>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Categories</th>
  </tr>
  <tr>
    <td>GetCategories()</td>
    <td>GET</td>
    <td>/categorias</td>
    <td></td>
    <td></td>
    <td>[ Categories ]</td>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Allergens</th> 
  </tr>
  <tr>
    <td>GetAllergens()</td>
    <td>GET</td>
    <td>/productos/alergenos</td>
    <td></td>
    <td></td>
    <td>[ Allergens ]</td>
  </tr>
  <tr>
    <td>GetAllergen()</td>
    <td>GET</td>
    <td>/productos/alergenos/:id-alergeno</td>
    <td></td>
    <td></td>
    <td>[ Allergen ]</td>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Productos</th>
  </tr>
  <tr>
    <td>GetProducts()</td>
    <td>GET</td>
    <td>/productos</td>
    <td></td>
    <td></td>
    <td>[ Products ]</td>
  </tr>
  <tr>
    <td>GetFilteredProducts()</td>
    <td>GET</td>
    <td>/productos/filtro</td>
    <td>[ Categories, Order, Range, Pagination ]</td>
    <td></td>
    <td>[ Products ]</td>
  </tr>
  <tr>
    <td>SearchProducts()</td>
    <td>GET</td>
    <td>/productos/search/:producto</td>
    <td></td>
    <td></td>
    <td>[ Products ]</td>
  </tr>
  <tr>
    <td>GetProductDetails()</td>
    <td>GET</td>
    <td>/productos/:slug_producto</td>
    <td></td>
    <td></td>
    <td>[ Producto ]</td>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Auth</th>
  </tr>
  <tr>
    <td>Register()</td>
    <td>POST</td>
    <td>/auth/register</td>
    <td></td>
    <td>{ User Information }</td>
    <td>[ User Information + Token ]</td>
  </tr>
  <tr>
    <td>Login()</td>
    <td>POST</td>
    <td>/auth/login</td>
    <td></td>
    <td>{ User Information }</td>
    <td>[ User Information + Token ]</td>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Orders</th>
  </tr>
  <tr>
    <td>GetOrders()</td>
    <td>GET</td>
    <td>/pedidos</td>
    <td> { Header: Token } </td>
    <td></td>
    <td>[ User Orders ]</td>
  </tr>
  <tr>
    <td>AddOrder()</td>
    <td>POST</td>
    <td>/pedidos</td>
    <td>{ Header: Token }</td>
    <td>{ Order Information }</td>
    <td>[ { msg: "Pedido Creado" } ]</td>
  </tr>
  <tr>
    <th colspan="6" class="title" style="background: #f2f2f2;">Reserves</th>
  </tr>
  <tr>
    <td>GetReserves()</td>
    <td>GET</td>
    <td>/reservas</td>
    <td> { Header: Token } </td>
    <td></td>
    <td>[ User Reserves ]</td>
  </tr>
  <tr>
    <td>CreateReserve()</td>
    <td>POST</td>
    <td>/reservas</td>
    <td>{ Header: Token }</td>
    <td>{ Reserve Information }</td>
    <td>[ { msg: "Reserva Realizada" } ]</td>
  </tr>
  <tr>
    <td>GetReserve()</td>
    <td>GET</td>
    <td>/reservas/:id_reserva</td>
    <td> { Header: Token }</td>
    <td></td>
    <td>[ Reserve ]</td>
  </tr>
  <tr>
    <td>GetImage()</td>
    <td>GET</td>
    <td>/reservas/image</td>
    <td></td>
    <td></td>
    <td>[ Image ]</td>
  </tr>
  <tr>
    <td>GetPDFReserve()</td>
    <td>GET</td>
    <td>/reservas/pdf/_id_reserva</td>
    <td>{ Header: Token }</td>
    <td></td>
    <td>[ Reserve ]</td>
  </tr>
  <tr>
    <td>GetBannedDays()</td>
    <td>GET</td>
    <td>/GetBannedDays</td>
    <td>{ n_comensales: Number, servicio: String }</td>
    <td></td>
    <td>[ Banned Days to Reserve ]</td>
  </tr>
  
</table>

## FEATURES

- JWT Auth
- Guards On Routes
- Served Dockerized
- Working as a submodule of [MyFood]

## USAGE

This application is created to be used in another project ([MyFood]) but if you want to launch it you should keep in mind that you have to create an `Database.go` inside `Config` folder and within this add the configuration towards the database and a `secret` for the JWT.

An functional example of `Database.go` should be like this:
```go
// Database.go
package Config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "db_host",
		Port:     db_port,
		User:     "db_user",
		Password: "db_password",
		DBName:   "db_name",
	}
	return &dbConfig
}

const NBSecretPassword = "secret"

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", DbURL(BuildDBConfig()))
	if err != nil {
		fmt.Println("Status", err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

```
To be launched simply run `go mod tidy` to install all packages on `go.mod` and then `go run main.go`.
## AUTHORS

<div class="authors" style="display: inline-block;
        align-items: center;
        margin: 10px;
        flex-direction: row;
        justify-content: center;
        gap: 50px;">
<div style="text-align:center;">
  <a href="https://github.com/gfmois"><img style="margin-bottom: 10px;" src="https://avatars.githubusercontent.com/u/102977172?s=400&v=4" alt="Creator gfmois" width="150" height="150"></a>
  <h3>gfmois</h3>
</div>
<div style="text-align: center;">
  <a href="https://github.com/JoaquimRS"><img style="margin-bottom: 10px;" src="https://avatars.githubusercontent.com/u/94555035?v=4" alt="Creator JoaquimRS" width="150" height="150"></a>
  <h3>JoaquimRS</h3>
</div>
</div>
