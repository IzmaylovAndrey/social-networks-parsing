package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Accounts struct {
	ID       		string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID 			string `gorm:"index;type:uuid"`
	SocialNetwork 	string
	Data 			string
}

func Add (account Accounts, db gorm.DB) error {
	if err := db.Create(&account).Error; err != nil {
		fmt.Printf("Account adding error: %s", err)
		return err
	}
	return nil
}

func GetAccountsByUserID (userID string, db gorm.DB) ([]Accounts, error) {
	var accounts []Accounts
	if err := db.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		fmt.Printf("User list getting error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", accounts[0].UserID)
	return accounts, nil
}