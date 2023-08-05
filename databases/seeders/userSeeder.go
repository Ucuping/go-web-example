package seeders

import (
	"github.com/Ucuping/go-web-example/helpers"
	"github.com/Ucuping/go-web-example/models"
)

func UserSeeder() {
	passwordHash, err := helpers.GeneratePasswordHash("root")
	if err != nil {
		panic("Data seeder error")
	}

	var users = []models.User{
		{
			Name:     "Developer",
			Email:    "dev@example.com",
			Username: "root",
			Password: passwordHash,
		},
		{
			Name:     "Test",
			Email:    "test@example.com",
			Username: "test",
			Password: passwordHash,
		},
	}

	result := helpers.DB.Create(&users)

	if result.Error != nil {
		panic(result.Error)
	}
}
