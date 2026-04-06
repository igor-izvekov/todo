package migrations

import (
	"log"

	"gorm.io/gorm"

	"github.com/igor-izvekov/todo/pkg/models"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Запускаем миграции")

	err := db.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	if err != nil {
		log.Printf("Ошибка миграции: %v", err)
		return err
	}

	log.Println("Завершение миграции")
	return nil
}
