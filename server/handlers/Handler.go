package handlers

import (
	"fmt"
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
func DeleteFruit(c *gin.Context) {
	Data := pkg.Fruits{}
	c.Bind(&Data)
	SQL := "DELETE from fruit where id=$1"
	_, err := utils.DB.Query(SQL, Data.Id)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return
	}
	//defer rows.Close()
	//nurse := Nurses{}

	res := gin.H{
		"fruits": Data,
	}
	c.JSON(http.StatusOK, res)
}
func NewFruit(c *gin.Context) {
	reqBody := pkg.Fruits{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"Error": err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result, _ := NewFruits(reqBody)
	if !result {
		res := gin.H{
			"Error": "Something is wromg",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := gin.H{
		"message": "successfully inserted",
		"status":  200,
	}
	c.JSON(http.StatusCreated, res)
}
func NewFruits(reqbody pkg.Fruits) (bool, string) {
	var result = true
	var err_responce = ""

	sqlStatement := `
INSERT INTO fruit(fruit_name)
VALUES ($1)`
	_, err2 := utils.DB.Exec(sqlStatement, reqbody.FruitName)
	fmt.Println(err2)
	if err2 != nil {

		err_responce = "Something went wrong"
		result = false
		return result, err_responce
	}
	return result, err_responce
}

func UpdateFruit(c *gin.Context) {

	reqBody := pkg.Fruits{}

	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusBadGateway, res)
	}
	SQL := `UPDATE fruit SET  fruit_name=$1 WHERE id=$2`

	_, err2 := utils.DB.Exec(SQL, reqBody.FruitName, reqBody.Id)

	result := true
	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		fmt.Println(err2)
		result = false
	}

	//return true

	if result == false {
		res := gin.H{
			"error": "Something went wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"message": "Updated Successfully",
	}
	c.JSON(http.StatusOK, res)
	return

}
