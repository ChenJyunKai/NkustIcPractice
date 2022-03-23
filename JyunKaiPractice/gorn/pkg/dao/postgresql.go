package dao

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	DB = Open()
}

func Open() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable password=usa960092"
	// host := os.Getenv("HS_GORM_HOST")
	// port := os.Getenv("HS_GORM_PORT")
	// user := os.Getenv("HS_GORM_USER")
	// pass := os.Getenv("HS_GORM_PASSWORD")
	// dbname := os.Getenv("HS_GORM_DATABASE")
	// sslmode := os.Getenv("HS_GORM_SSLMODE")

	// dsn := "host=" + host + " user=" + user + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode + " password=" + pass

	mode := os.Getenv("GIN_MODE")
	config := &gorm.Config{}
	if mode != "release" {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
