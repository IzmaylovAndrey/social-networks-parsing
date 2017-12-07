package main

import (
	"github.com/IzmaylovAndrey/social-networks-parsing/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users", routes.GettingAllUsers)
	router.POST("/users", routes.CreatingUser)
	router.GET("/users/:id", routes.GettingUserByID)
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8080")
}
