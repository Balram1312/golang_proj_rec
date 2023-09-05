package main 

import (
	"github.com/gin-gonic/gin"
	"golang_test/utils/database"
	"golang_test/handlers"
)

func main(){
	router := gin.Default()
	db,err := database.Connect()
	if err != nil{
		panic("not connected")
	}
	app := &handlers.MyApp{
		DB : db,
	}
	
	app.Routes(router)
	router.Run(":8080")
}