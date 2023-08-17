package db

import (
	"fmt"
	"money-treez/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
}

func NewDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connect(creds *model.DBCredentials) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.User, creds.Password, creds.DBName, creds.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
