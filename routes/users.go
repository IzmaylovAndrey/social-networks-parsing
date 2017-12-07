package routes

import (
	"fmt"
	"net/http"

	"github.com/IzmaylovAndrey/social-networks-parsing/models"
	"github.com/IzmaylovAndrey/social-networks-parsing/utils"

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
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func CreatingUser(c *gin.Context) {
	var json SignUpJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%s", json)
	db := models.OpenConnection()
	defer models.CloseConnection(*db)
	user := models.Users{}
	if err := user.Create(json.Email, json.Name, json.Password, *db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vkChanel := make(chan utils.ChanelResult)
	facebookChanel := make(chan utils.ChanelResult)
	gitHubChanel := make(chan utils.ChanelResult)
	go utils.GithubSearch(user.Name, gitHubChanel)
	go utils.FBSearch(user.Name, facebookChanel)
	go utils.VKSearch(user.Name, vkChanel)

	errors :=make([]error, 3)
	select {
	case data := <-vkChanel:
		if data.Error != nil{
			errors = append(errors, data.Error)
		}
		account := models.Accounts{}
		account.Create(user.ID, "vk", data.Message, *db)
	case data := <-facebookChanel:
		if data.Error != nil{
			errors = append(errors, data.Error)
		}
		account := models.Accounts{}
		account.Create(user.ID, "facebook", data.Message, *db)
	case data := <-gitHubChanel:
		if data.Error != nil{
			errors = append(errors, data.Error)
		}
		account := models.Accounts{}
		account.Create(user.ID, "github", data.Message, *db)
	}

	if err := utils.SendEmail(user.Login, user.Name); err != nil {
		fmt.Printf("Email message for %s was not sent", user.Login)
	}
	if err := utils.SendToTelegram(user.Login, user.Accounts); err != nil {
		fmt.Printf("Telegram message about %s was not sent", user.Login)
	}
	c.JSON(http.StatusCreated, user)
}
