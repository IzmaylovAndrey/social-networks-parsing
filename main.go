package main

import (
	"./models"
	//"time"
	//"fmt"
	//"./utils"
)

func main() {
	db := models.OpenConnection()
	//models.Migrate(*db)
	//user := models.Users {ID: "0b454430-5292-449e-9c84-2ba3e6e6578e", Login: "maggy93@mail.ru", PasswordHash: "1234", Salt: "3456", CreatedAt: time.Now()}
	//models.Create(user, *db)
	//models.GetByID("0d29fcb9-1129-4b23-9084-763012be284a", *db)
	//models.GetByLogin("helpik94@ysndex.ru", *db)
	//users, _ := models.GetAll(*db)
	//users, _ := models.GetAllOrderbyCreationDesc(*db)
	//fmt.Printf("%s", (*users)[0].Login)
	//account := models.Accounts{ID: "86f18aaf-3340-40ea-9b04-769c65808d11", UserID: "0b454430-5292-449e-9c84-2ba3e6e6578e", SocialNetwork: "vk", Data: "{}"}
	//models.Add(account, *db)
	//accounts, _ := models.GetAccountsByUserID("0b454430-5292-449e-9c84-2ba3e6e6578e", *db)
	//utils.SendToTelegram("maggy93@mail.ru", *accounts)
	models.CloseConnection(*db)
}
