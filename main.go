package main

import (
	"github.com/alifhaider/BCS-Journey-Server/controllers"
	"github.com/alifhaider/BCS-Journey-Server/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadENV()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	 r := gin.Default()
	  r.POST("/signup", controllers.SignUp)
		r.POST("/login", controllers.Login)
	 r.Run()

}