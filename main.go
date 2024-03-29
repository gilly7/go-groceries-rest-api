package main

import (
	"log"

	"github.com/gilly7/go-groceries-rest-api/grocery"
	"github.com/gilly7/go-groceries-rest-api/model"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	router.GET("/groceries", grocery.GetGroceries)
	router.GET("/grocery/:id", grocery.GetGrocery)
	router.POST("/grocery", grocery.PostGrocery)
	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	router.DELETE("/grocery/:id", grocery.DeleteGrocery)

	log.Fatal(router.Run(":10000"))
}
