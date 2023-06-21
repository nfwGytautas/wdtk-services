package context

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identifier string
	Password   string
	Role       string
}

type Tokens struct {
	gorm.Model
	User  string
	Token string
}

type AuthData struct {
	dbConn *gorm.DB
}

var Context AuthData

func (ad *AuthData) SetupDatabase(connectionString string) error {
	var err error
	ad.dbConn, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return err
	}

	// Create tables
	ad.dbConn.AutoMigrate(&User{})

	return nil
}

func (ad *AuthData) GetUser(identifier string) (User, error) {
	u := User{}

	err := ad.dbConn.Model(User{}).Where("identifier = ?", identifier).Take(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (ad *AuthData) GetUserByID(identifier uint) (User, error) {
	u := User{}

	err := ad.dbConn.Model(User{}).First(&u, identifier).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (ad *AuthData) CreateUser(user *User) error {
	err := ad.dbConn.Create(user).Error
	return err
}
