package purchases_products

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/purchases_products"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("purchases_products")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":purchases_productsID", controller.GetByID)
		v10.DELETE(":purchases_productsID", controller.Delete)
		v10.PATCH(":purchases_productsID", controller.Updated)
	}

	return route
}
