package requisitions

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/requisitions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("requisitions")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":requisitionsID", controller.GetByID)
		v10.DELETE(":requisitionsID", controller.Delete)
		v10.PATCH(":requisitionsID", controller.Updated)
	}

	return route
}
