package database

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var SqlServerDb *gorm.DB

func InitSqlServer() error {
	var err error

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	SqlServerDb, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}
