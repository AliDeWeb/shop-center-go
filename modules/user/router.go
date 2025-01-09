package auth

import "github.com/gin-gonic/gin"

func Routes(c *gin.Engine) {
	routesGroup := c.Group("/user")

	routesGroup.POST("/auth/register", CRegisterUser)
}
