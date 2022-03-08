package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DBURL"),
		Port:     5432,
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBName:   "gomicro",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
