package route

import (
	"go-chat/controllers"
	"go-chat/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", controllers.GetIndex)
	r.GET("/users", controllers.Users)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.PATCH("/user/:id", controllers.UpdateUser)

	return r
}
