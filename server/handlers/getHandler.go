package handlers

import (
	"fruit/pkg"
	"fruit/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllFruits(c *gin.Context) {
	Data := []pkg.Fruits{}

	SQL := "SELECT id,fruit_name from fruit"
	rows, err := utils.DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	defer rows.Close()
	fruits := pkg.Fruits{}

	for rows.Next() {
		rows.Scan(&fruits.Id, &fruits.FruitName)

		Data = append(Data, fruits)

	}
	res := gin.H{
		"fruits": Data,
	}
	c.JSON(http.StatusOK, res)
}
