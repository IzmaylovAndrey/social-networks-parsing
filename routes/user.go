package routes

import (
	"net/http"

	"github.com/IzmaylovAndrey/social-networks-parsing/models"

	"github.com/gin-gonic/gin"
)


func GettingUserByID(c *gin.Context){
	id := c.Param("id")
	db := models.OpenConnection()
	defer models.CloseConnection(*db)
	user, err := models.GetByID(id, *db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if user.Login == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	}
	c.JSON(http.StatusOK, user)
}