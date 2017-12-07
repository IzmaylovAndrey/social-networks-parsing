package models

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Accounts struct {
	ID            string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	User      	  Users  `gorm:"ForeignKey:UserID"`
	UserID        string `gorm:"index;type:uuid" json:"-"`
	SocialNetwork string
	Data          string
}

func (acc *Accounts) Create(userID string, socialNetwork string, data []string, db gorm.DB) (error) {
	acc.UserID = userID
	acc.SocialNetwork = socialNetwork
	preparedData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error in marshalling account data in JSON. Error: %s", err)
	}
	acc.Data = string(preparedData)
	fmt.Printf("Data: %s", acc.Data)
	return acc.insert(db)
}

func (acc *Accounts) insert(db gorm.DB) error {
	if err := db.Create(acc).Error; err != nil {
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