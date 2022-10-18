package migrations

import (
	"golang-web-service/assignment-2/models"

	"gorm.io/gorm"
)

var ModelMigrations = []interface{}{
	&models.Orders{},
	&models.Items{},
}

func AutoMigrate(tx *gorm.DB) {
	tx.AutoMigrate(ModelMigrations...)
}
