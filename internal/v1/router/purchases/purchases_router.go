package purchases

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/purchases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("purchases")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":purchasesID", controller.GetByID)
		v10.DELETE(":purchasesID", controller.Delete)
		v10.PATCH(":purchasesID", controller.Updated)
	}

	return route
}
