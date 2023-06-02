package company

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"eirc.app/internal/v1/middleware"
	"eirc.app/internal/v1/presenter/company"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := company.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("companies")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":companyID", controller.GetByID)
		v10.DELETE(":companyID", controller.Delete)
		v10.PATCH(":companyID", controller.Updated)
	}

	return route
}
