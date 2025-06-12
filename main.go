package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizzsv/go-jwt/controllers"
	"github.com/rizzsv/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}


func main() {
  router := gin.Default()

  router.POST("/signup", controllers.Signup)
  router.Run() // listen and serve on 0.0.0.0:8080
}