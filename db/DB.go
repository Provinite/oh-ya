package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB, DBError = gorm.Open(postgres.Open("postgres://postgres:password@host.docker.internal:5432/postgres"), &gorm.Config{})
