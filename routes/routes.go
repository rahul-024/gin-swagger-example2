package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahul-024/gin-swagger-example2/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", handlers.GetBooks)
		v1.GET("/books/:isbn", handlers.GetBookByISBN)
		// router.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
		// router.PUT("/books/:isbn", handlers.UpdateBookByISBN)
		v1.POST("/books", handlers.PostBook)
	}

	return router
}
