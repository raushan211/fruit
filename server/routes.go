package server

import (
	"fruit/server/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {

	r.POST("/add_fruit", handlers.NewFruit)
	r.GET("/fruits", handlers.AllFruits)
	r.PUT("/update", handlers.UpdateFruit)
	r.DELETE("/delete", handlers.DeleteFruit)

}
