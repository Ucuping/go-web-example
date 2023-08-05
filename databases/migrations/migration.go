package migrations

import (
	"github.com/Ucuping/go-web-example/helpers"
	"github.com/Ucuping/go-web-example/models"
)

func MigrateTable() {
	helpers.DB.AutoMigrate(&models.User{}, &models.Post{})
}

func DropTable() {
	helpers.DB.Migrator().DropTable(&models.User{}, &models.Post{})
}
