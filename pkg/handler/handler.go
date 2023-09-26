package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhashkevych/todo-app/pkg/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/zhashkevych/todo-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	route := router.Group("/api/v1", h.userIdentity)
	{
		book := route.Group("/book")
		{
			book.POST("/", h.createBook)          // create book
			book.GET("/", h.GetBookList)          // get book
			book.GET("/genre", h.GetBooksByGenre) //get book by genre
			book.GET("/id", h.GetBookById)        // get book by id
			book.PUT("/id", h.UpdateBook)         // update book
			book.DELETE("/id", h.DeleteBook)      // delete book

			like := book.Group("/like")
			{
				like.POST("/", h.creteLike)
				like.GET("/", h.GetLikedList)
			}

			rating := book.Group("/rating")
			{
				rating.POST("/", h.creteRating)
			}
		}
		genre := route.Group("/genre")
		{
			genre.POST("/", h.creteGenre)    // create genre
			genre.GET("/", h.GetGenreList)   // get genre
			genre.GET("/id", h.GetGenreById) // get genre by id
			genre.PUT("/", h.UpdateGenre)    // update genre
			genre.DELETE("/", h.DeleteGenre) // delete genre
		}
		history := route.Group("history")
		{
			history.GET("/", h.GetHistoryList)
		}
	}

	return router
}
