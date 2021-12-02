package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/obidovsamandar/go-task-auth/api/handlers/v1"
	"github.com/obidovsamandar/go-task-auth/api/helpers"
	"github.com/obidovsamandar/go-task-auth/config"
)

func main() {

	cfg := config.Load()
	helpers.ConnectionDB()

	server := gin.Default()

	server.POST("/generate-passcode", v1.GenerateOTPCode)
	server.POST("/signup", v1.SignUp)
	server.POST("/signin", v1.SignIn)
	server.PUT("/update-user", v1.UpdateUser)
	server.POST("/image-upload", v1.ImageUpload)
	server.GET("/get-user", v1.GetUser)

	server.Run(cfg.PORT)

}
