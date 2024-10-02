package configs

import (
	"log"

	"github.com/haikelz/asmaul-husna-api/internal/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dsn := "host=localhost user=postgres password=debian123 dbname=asmaulhusna port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		 PrepareStmt: true,
	})

	if err != nil {
		log.Fatalln("Error while load data from Database!", err)
	}

	db.Model(&entities.AsmaulHusna{})
	db.AutoMigrate(&entities.AsmaulHusna{})

	return db
}
