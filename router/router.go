package router

import (
	"github.com/gin-gonic/gin"
	"myProject/handler"
)

func Router(r *gin.Engine, db *handler.StoreStruct) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/", db.GetProducts)
		productGroup.GET("/:id", db.GetProductById)
		productGroup.DELETE("/:id", db.DelProductById)
		productGroup.PUT("/:id", db.UpdateProduct)
		productGroup.POST("/", db.CreateProduct)
	}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/", nil)
	}

	commentGroup := r.Group("/comment")
	{
		commentGroup.GET("/", nil)
	}
}
