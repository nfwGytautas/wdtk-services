package context

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `gorm:"primarykey;type:uint"`
	Identifier string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Role       string `gorm:"not null"`
}

type Tokens struct {
	UserID uint `gorm:"unique,type:uint"`
	Token  string

	User User `gorm:"foreignKey:UserID"`
}

type AuthData struct {
	mx               sync.Mutex
	dbConn           *gorm.DB
	connectionString string
}

var Context AuthData

func (ad *AuthData) SetupDatabase(connectionString string) error {
	ad.connectionString = connectionString
	ad.dbConn = nil
	return ad.verifyConnection()
}

func (ad *AuthData) establishConnection() error {
	var err error
	ad.dbConn, err = gorm.Open(mysql.Open(ad.connectionString), &gorm.Config{})

	if err != nil {
		ad.dbConn = nil
		return err
	}

	// Create tables
	// ad.dbConn.Migrator().DropTable(&User{}, &Tokens{})
	err = ad.dbConn.AutoMigrate(&User{}, &Tokens{})
	if err != nil {
		log.Panicln("[ERROR] Failed to migrate")
		ad.dbConn = nil
		return err
	}

	return nil
}

func (ad *AuthData) verifyConnection() error {
	ad.mx.Lock()
	defer ad.mx.Unlock()

	if ad.dbConn != nil {
		return nil
	}

	// Try connect
	return ad.establishConnection()
}
