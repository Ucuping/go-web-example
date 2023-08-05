package databases

import (
	"os"

	"github.com/Ucuping/go-web-example/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	helpers.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database")
	}
}
