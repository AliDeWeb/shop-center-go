package routes

import (
	auth "github.com/alideweb/shop-center-go/modules/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	auth.Routes(engine)
}
