package main

import (
	"p2src/database"
	"p2src/routes"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()
	gin.HTMLRender = ginview.Default()
	database.ConnectDB()
	//r.LoadHTMLGlob("templates/*.tmpl")
	routes.Routes(gin)
	gin.Run(":9000")
}
