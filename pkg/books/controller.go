package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/book")
	routes.POST("/", h.CreateBook)
	routes.DELETE("/:id", h.DeleteBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
}
