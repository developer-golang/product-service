package migration

import (
	"command/models"
	"log"

	"gorm.io/gorm"
)

func PostgresMigration(Db *gorm.DB) (err error) {
	err = Db.AutoMigrate(&models.ProductDetails{})
	if err != nil {
		log.Println(err.Error() + "001")
		return
	}
	return
}
