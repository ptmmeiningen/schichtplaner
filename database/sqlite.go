package database

import (
	"errors"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/ptmmeiningen/schichtplaner/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func StartDB() error {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		dbPath = "schichtplaner.db"
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return errors.New("Fehler beim Öffnen der SQLite-Datenbank: " + err.Error())
	}

	return nil
}

func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		err = sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}
}

func AutoMigrate() error {
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		return errors.New("Fehler bei der Datenbank-Migration: " + err.Error())
	}
	return nil
}
