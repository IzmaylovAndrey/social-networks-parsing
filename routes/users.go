package routes

import (
	"net/http"

	"github.com/IzmaylovAndrey/social-networks-parsing/models"

	"github.com/gin-gonic/gin"
)

type SignUpJSON struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
}

func GettingAllUsers(c *gin.Context) {
	orderBy := c.Query("orderBy")
	desc := c.DefaultQuery("desc", "")
	db := models.OpenConnection()
	defer models.CloseConnection(*db)
	var err error
	var users []models.Users
	switch orderBy {
	case "login":
		if desc == "" || desc == "0" {
			users, err = models.GetAllUsersOrderByLogin(*db)
		} else if desc == "1" {
			users, err = models.GetAllUsersOrderByLoginDecs(*db)
		}
	case "id":
		if desc == "" || desc == "0" {
			users, err = models.GetAllUsersOrderByID(*db)
		} else if desc == "1" {
			users, err = models.GetAllUsersOrderByIDDesc(*db)
		}
	case "created_at":
		if desc == "" || desc == "0" {
			users, err = models.GetAllUsersOrderByCreation(*db)
		} else if desc == "1" {
			users, err = models.GetAllUsersOrderByCreationDesc(*db)
		}
	default:
		users, err = models.GetAllUsers(*db)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	}
	c.JSON(http.StatusOK, users)
}

func CreatingUser(c *gin.Context) {
	var json SignUpJSON
	if err := c.ShouldBindJSON(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	db := models.OpenConnection()
	defer models.CloseConnection(*db)
	user := models.Users{}
	if err := user.Create(json.Email, json.Name, json.Password, *db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//TODO: goroutines with API handlers
	//TODO: sending to telegram
	//TODO: ??? mail for user
	c.JSON(http.StatusCreated, user)
}
