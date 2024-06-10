package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sharanrohit7/authService/controller"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
