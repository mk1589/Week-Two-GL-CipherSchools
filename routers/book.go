package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mk1589/Week-Two-GL-CipherSchools/database"
	"github.com/mk1589/Week-Two-GL-CipherSchools/handler"
)

func Router() *gin.Engine {
	router := gin.Default() // get the default engine for further customization
	api := handler.Handler{
		DB: database.GetDB(), // set the handeler DB
	}
	router.GET("/books", api.GetBooks) //set the function for http://locahost:8080/book : Get request
	//while calling handler function, gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	return router

}
