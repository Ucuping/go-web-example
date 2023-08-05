package controllers

import (
	"os"
	"time"

	"github.com/Ucuping/go-web-example/helpers"
	"github.com/Ucuping/go-web-example/models"
	"github.com/Ucuping/go-web-example/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// func AuthIndex(c *fiber.Ctx) error {
// 	return c.Render("login", fiber.Map{
// 		"title": "Login",
// 	}, "login")
// }

func Register(c *fiber.Ctx) error {
	var body validations.RegisterValidation

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Unprocess Entity")
	}
	errors := helpers.CustomValidator(&body)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	var checkUser models.User
	helpers.DB.First(&checkUser, "username = ? OR email = ?", body.Username, body.Email)

	if checkUser.ID != 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status": "fail",
			"errors": "Username or email already exist",
		})
	}

	passwordHash, err := helpers.GeneratePasswordHash(body.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	user := models.User{Name: body.Name, Email: body.Email, Username: body.Username, Password: string(passwordHash)}
	result := helpers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

func Login(c *fiber.Ctx) error {
	var body validations.LoginValidation
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Unprocess Entity")
	}
	errors := helpers.CustomValidator(&body)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": errors,
		})
	}

	var user models.User
	helpers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		var errors []*helpers.ErrorResponse
		var error helpers.ErrorResponse
		error.Field = "Password"
		error.Message = "Invalid username or password"
		errors = append(errors, &error)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": errors,
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		var errors []*helpers.ErrorResponse
		var error helpers.ErrorResponse
		error.Field = "Password"
		error.Message = "Invalid username or password"
		errors = append(errors, &error)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": errors,
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30).Local()
	cookie.HTTPOnly = true
	cookie.SameSite = fiber.CookieSameSiteLaxMode
	cookie.Secure = false

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})

	// return c.Redirect("/posts")
}

func Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour).Local()
	cookie.HTTPOnly = true
	cookie.SameSite = fiber.CookieSameSiteLaxMode
	cookie.Secure = false

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"status": "success",
	})
	// return c.Redirect("/auth/login")
}

func Verify(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}
