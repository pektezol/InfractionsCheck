package routes

import (
	"p2src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/infractions", controllers.Home)
	r.GET("/infractions/add", controllers.Add)
}
