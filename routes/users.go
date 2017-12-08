package routes

import (
	"fmt"
	"net/http"
	"sync"

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
	case "createdAt":
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

	wg := sync.WaitGroup{}
	wg.Add(3)
	result := utils.APIHandlersResult{}
	go utils.GithubSearch(user.Name, &result, &wg)
	go utils.FBSearch(user.Name, &result, &wg)
	go utils.VKSearch(user.Name, &result, &wg)
	wg.Wait()
	fmt.Printf("Result: %v", result)
	if len(result.Facebook) != 0 {
		account := models.Accounts{}
		account.Create(user.ID, "facebook", result.Facebook, *db)
	}
	if len(result.Github) != 0 {
		account := models.Accounts{}
		account.Create(user.ID, "github", result.Github, *db)
	}
	if len(result.VK) != 0 {
		account := models.Accounts{}
		account.Create(user.ID, "vk", result.VK, *db)
	}

	if err := utils.SendEmail(user.Login, user.Name); err != nil {
		fmt.Printf("Email message for %s was not sent", user.Login)
	}

	accounts, err := models.GetAccountsByUserID(user.ID, *db)
	if err != nil {
		fmt.Printf("Error while getting user accounts from db - %s", err)
	}

	if err := utils.SendToTelegram(user.Login, accounts); err != nil {
		fmt.Printf("Telegram message about %s was not sent", user.Login)
	}
	c.JSON(http.StatusCreated, user)
}
