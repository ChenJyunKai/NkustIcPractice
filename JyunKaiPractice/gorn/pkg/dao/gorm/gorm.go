package gorm

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func New() (*gorm.DB, error) {
	const config string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

	sources := fmt.Sprintf(config,
		os.Getenv("SOURCES_HOST"),
		os.Getenv("SOURCES_PORT"),
		os.Getenv("SOURCES_USER"),
		os.Getenv("SOURCES_PASSWORD"),
		os.Getenv("SOURCES_DATABASE"),
		os.Getenv("SOURCES_SSLMODE"),
	)

	replicas := fmt.Sprintf(config,
		os.Getenv("REPLICAS_HOST"),
		os.Getenv("REPLICAS_PORT"),
		os.Getenv("REPLICAS_USER"),
		os.Getenv("REPLICAS_PASSWORD"),
		os.Getenv("REPLICAS_DATABASE"),
		os.Getenv("REPLICAS_SSLMODE"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  sources,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.New(postgres.Config{
			DSN:                  replicas,
			PreferSimpleProtocol: true,
		})},
	}))

	return db, nil
}
