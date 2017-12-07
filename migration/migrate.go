package main

import "github.com/IzmaylovAndrey/social-networks-parsing/models"

func main() {
	db := models.OpenConnection()
	models.Migrate(*db)
	models.CloseConnection(*db)
}
