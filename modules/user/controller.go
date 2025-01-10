package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CRegisterUser(c *gin.Context) {
	var user MUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Validation failed",
			"errors":  err.Error(),
		})
		return
	}

	_, data, tokens, err := SRegisterUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "Email already exists",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Registration failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"data": gin.H{
			"user":   data,
			"tokens": tokens,
		},
	})
}
