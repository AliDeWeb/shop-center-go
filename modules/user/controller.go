package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CRegisterUser(c *gin.Context) {
	var user MUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid body request: %s", err.Error())})
		return
	}

	_, data, err := SRegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("duplicate email: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User data received",
		"user":    data,
	})
}
