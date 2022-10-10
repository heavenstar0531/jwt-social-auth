package controllers

import (
	"jwt-go/database"
	"jwt-go/helpers"
	"jwt-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	AppJson = "application/json"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == AppJson {
		if err := c.ShouldBindJSON(&product); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&product); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	product.UserID = userID
	if err := db.Debug().Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	product := models.Product{}

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid product id",
		})
		return
	}

	userID := uint(userData["id"].(float64))

	if contentType == AppJson {
		if err := c.ShouldBindJSON(&product); err != nil {
			c.AbortWithError(400, err)
			return
		}
	} else {
		if err := c.ShouldBind(&product); err != nil {
			c.AbortWithError(400, err)
			return
		}
	}

	product.UserID = userID
	product.ID = uint(productId)

	err = db.Debug().Where("id = ?", productId).Updates(models.Product{Title: product.Title, Description: product.Description}).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func FindProductByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid product id",
		})
		return
	}

	userID := uint(userData["id"].(float64))

	product := models.Product{}
	err = db.Debug().Where("id = ? AND user_id = ?", productId, userID).First(&product).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func FindAllProducts(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	products := []models.Product{}
	err := db.Debug().Where("user_id = ?", userID).Find(&products).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	} else if len(products) == 0 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "You don't have any products",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid product id",
		})
		return
	}

	userID := uint(userData["id"].(float64))

	product := models.Product{}
	err = db.Debug().Where("id = ? AND user_id = ?", productId, userID).First(&product).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Delete(&product).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
