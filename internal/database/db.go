package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"test-Grpc/internal/module/user/entities"
)

type DatabaseInfo struct {
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSLMode  string `json:"ssl_mode"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func Migrate(database *gorm.DB) (err error) {
	err = database.AutoMigrate(
		&entities.User{},
	)

	return err
}

var once sync.Once
var Connections *gorm.DB

// var Connections map[string]*gorm.DB
//var Connection *gorm.DB

func ConnectToDB(database *DatabaseInfo) (err error) {
	once.Do(func() {
		dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			database.Username, database.Password, database.Host, database.Port, database.Name, database.SSLMode)

		Connections, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
		postgreDB, err := Connections.DB()
		//Connection, err = gorm.Open(postgres.Open(dbSource), &gorm.Config{})
		//postgresDB, err := Connection.DB()
		postgreDB.SetMaxIdleConns(30)
		postgreDB.SetMaxOpenConns(300)

		if err != nil {
			log.Printf("Connect database fail with error: %v", err.Error())
		}

		err = Migrate(Connections)

		if err != nil {
			log.Printf("Migrate database fail with err: %v", err.Error())
		}

	})

	if err != nil {
		return err
	}
	return nil
}
