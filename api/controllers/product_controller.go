package controllers

import (
	"REST-GO-CLEAN/lib"
	"REST-GO-CLEAN/models"
	"net/http"

	"github.com/gin-gonic/gin"

	"REST-GO-CLEAN/api/responses"
	"REST-GO-CLEAN/api/services"
)

type ProductController struct {
	service services.ProductService
	logger  lib.Logger
}

func NewProductController(service services.ProductService, logger lib.Logger) ProductController {
	return ProductController{
		service: service,
		logger:  logger,
	}
}

// GetMember gets one member
func (controller ProductController) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		product := models.Product{}

		if err := c.ShouldBindJSON(&product); err != nil {
			controller.logger.Zap.Error("Product Save Error: ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		if _, err := controller.service.Create(ctx, &product); err != nil {
			controller.logger.Zap.Error("Product Save Error: ", err)
			responses.ErrorJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		responses.JSON(c, http.StatusOK, gin.H{
			"msg": "Product created.",
			"id":  product.ID,
		})
	}
}
