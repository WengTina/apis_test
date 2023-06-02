package roles

import (
	"eirc.app/internal/v1/middleware"
	presenter "eirc.app/internal/v1/presenter/roles"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("roles")
	{
		v10.POST("", middleware.Transaction(db), controller.Created)
		v10.GET("", controller.List)
		v10.GET(":rolesID", controller.GetByID)
		v10.DELETE(":rolesID", controller.Delete)
		v10.PATCH(":rolesID", controller.Updated)
	}

	return route
}
