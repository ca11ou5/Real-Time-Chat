package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"realtimeChat/internal/domain"
)

func NewDatabase() *gorm.DB {
	db, err := sql.Open("postgres", "user=postgres password=root host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("no response from database")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to db gorm")
	}

	err = gormDB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("migration error")
	}

	return gormDB

}
