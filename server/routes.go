package server

import (
	"fruit/server/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {

	//r.POST("/add_fruit", newFruit)
	r.GET("/fruits", handlers.AllFruits)
	// r.PUT("/update", updateFruit)
	// r.DELETE("/delete", deleteFruit)

}
