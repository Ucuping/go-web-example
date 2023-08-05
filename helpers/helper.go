package helpers

import (
	"errors"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func CustomValidationMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}

type ErrorResponse struct {
	Field   string
	Message string
}

var validate = validator.New()

func CustomValidator(field interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(field)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Message = CustomValidationMessage(err.Tag())
			errors = append(errors, &element)
		}
	}
	return errors
}

func GeneratePasswordHash(pw string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pw), 10)

	if err != nil {
		// fmt.Println(err.Error())
		return "", err
	}

	return string(passwordHash), nil
}

func CheckToken(c *fiber.Ctx) (*jwt.Token, jwt.MapClaims, bool, error) {
	tokenString := c.Cookies("Authorization")

	if tokenString == "" {
		return nil, nil, false, errors.New("Unauthorized")
	}

	// if ts == "" {
	// 	return nil, nil, false, errors.New("Unauthorized")
	// }

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unauthorized")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, nil, false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	return token, claims, ok, nil
}
