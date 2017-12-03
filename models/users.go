package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Users struct {
	ID 				string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Login 			string
	PasswordHash 	string
	Salt 			string
	CreatedAt 		time.Time
}

func Create (user Users, db gorm.DB) error {
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("User creation error: %s", err)
		return err
	}
	return nil
}

func GetByID (id string, db gorm.DB) (*Users, error) {
	user := Users {ID: id}
	if err := db.First(&user).Error; err != nil {
		fmt.Printf("User getting by id error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", user.Login)
	return &user, nil
}

func GetByLogin (login string, db gorm.DB) (*Users, error) {
	var user Users
	if err := db.Where("login = ?", login).First(&user).Error; err != nil {
		fmt.Printf("User getting by login error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", user.Login)
	return &user, nil
}

func GetAll (db gorm.DB) (*[]Users, error) {
	var users []Users
	if err := db.Find(&users).Error; err != nil {
		fmt.Printf("User list getting error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", users[0].Login)
	return &users, nil
}

func GetAllOrderbyID (db gorm.DB) (*[]Users, error) {
	var users []Users
	if err := db.Order("id").Find(&users).Error; err != nil {
		fmt.Printf("User list ordered by id getting error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", users[0].Login)
	return &users, nil
}

func GetAllOrderbyLogin (db gorm.DB) (*[]Users, error) {
	var users []Users
	if err := db.Order("login").Find(&users).Error; err != nil {
		fmt.Printf("User list ordered by login getting error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", users[0].Login)
	return &users, nil
}

func GetAllOrderbyCreation (db gorm.DB) (*[]Users, error) {
	var users []Users
	if err := db.Order("created_at").Find(&users).Error; err != nil {
		fmt.Printf("User list ordered by creation date getting error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", users[0].Login)
	return &users, nil
}

func GetAllOrderbyCreationDesc (db gorm.DB) (*[]Users, error) {
	var users []Users
	if err := db.Order("created_at desc").Find(&users).Error; err != nil {
		fmt.Printf("User list ordered by creation date desc getting error: %s", err)
		return nil, err
	}
	fmt.Printf("%s", users[0].Login)
	return &users, nil
}