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
	} else {
		result := SRegisterUser(&user)

		c.JSON(http.StatusOK, gin.H{
			"message": "User data received",
			"user":    result,
		})
	}

}
