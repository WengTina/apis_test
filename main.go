package main

import (
	"fmt"
	"net/http"

	_ "eirc.app/api"
	"eirc.app/internal/pkg/dao/gorm"
	"eirc.app/internal/pkg/log"

	"eirc.app/internal/v1/router"
	"eirc.app/internal/v1/router/account"
	"eirc.app/internal/v1/router/company"
	"eirc.app/internal/v1/router/departments"
	"eirc.app/internal/v1/router/employees"
	"eirc.app/internal/v1/router/products"
	"eirc.app/internal/v1/router/purchases"
	"eirc.app/internal/v1/router/purchases_products"
	"eirc.app/internal/v1/router/requisitions"
	"eirc.app/internal/v1/router/roles"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title hsmaster SYSTEM API
// @version 0.1
// @description 企業系統整合管理平台
// @termsOfService https://eirc.app/

// @contact.name API System Support
// @contact.url https://eirc.app/
// @contact.email mingzong.lyu@gmail.com

// @license.name AGPL 3.0
// @license.url https://www.gnu.org/licenses/agpl-3.0.en.html

// @host api.eirc.app
// @BasePath /
// @schemes https
func main() {

	db, err := gorm.New()

	if err != nil {

		log.Error(err)
		return
	}

	route := router.Default()
	route = account.GetRoute(route, db)
	route = company.GetRoute(route, db)
	route = requisitions.GetRoute(route, db)
	route = purchases.GetRoute(route, db)
	route = purchases_products.GetRoute(route, db)
	route = products.GetRoute(route, db)
	route = roles.GetRoute(route, db)
	route = employees.GetRoute(route, db)
	route = departments.GetRoute(route, db)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Fatal(http.ListenAndServe(":8080", route))
}
