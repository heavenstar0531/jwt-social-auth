package controllers

import (
	"jwt-go/database"
	"jwt-go/helpers"
	"jwt-go/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	if err := db.Debug().Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Email already exists",
			})

			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJson {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	} else {
		if err := c.ShouldBind(&user); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	originalPassword := user.Password
	if err := db.Debug().Where("email = ?", user.Email).First(&user).Take(&user).Error; err != nil {
		panic("User not found")
	}

	if isValid := helpers.CheckPasswordHash([]byte(user.Password), []byte(originalPassword)); !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})

		return
	}

	jwt := helpers.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})

}
