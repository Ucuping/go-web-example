package controllers

import (
	"github.com/Ucuping/go-web-example/helpers"
	"github.com/Ucuping/go-web-example/models"
	"github.com/Ucuping/go-web-example/validations"
	"github.com/gofiber/fiber/v2"
)

func FetchPost(c *fiber.Ctx) error {
	_, claims, _, err := helpers.CheckToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	var posts []models.Post
	helpers.DB.Find(&posts, "user_id = ?", claims["sub"])

	return c.JSON(fiber.Map{
		"posts": posts,
	})
	// return c.Render("index", fiber.Map{
	// 	"posts": posts,
	// })
}

func ShowPost(c *fiber.Ctx) error {
	_, claims, _, err := helpers.CheckToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	postId := c.Params("id")
	var post models.Post
	helpers.DB.First(&post, "user_id = ? AND id = ?", claims["sub"], postId)

	// if result.Error != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	// }

	return c.JSON(fiber.Map{
		"post": post,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	_, claims, _, err := helpers.CheckToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	var body validations.PostValidation
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Unprocess Entity")
	}
	errors := helpers.CustomValidator(&body)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	postId := c.Params("id")
	var post models.Post
	helpers.DB.Model(&post).Where("id = ? AND user_id = ?", postId, claims["sub"]).Updates(models.Post{Title: body.Title, Content: body.Content})

	// if result.Error != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	// }

	return c.JSON(fiber.Map{
		"post": post,
	})
}

func CreatePost(c *fiber.Ctx) error {
	_, claims, _, err := helpers.CheckToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	var body validations.PostValidation
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Unprocess Entity")
	}
	errors := helpers.CustomValidator(&body)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	// return c.JSON(fiber.Map{
	// 	"data": uint(claims["sub"].(float64)),
	// })
	post := models.Post{Title: body.Title, Content: body.Content, UserID: uint(claims["sub"].(float64))}
	result := helpers.DB.Create(&post)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(fiber.Map{
		"post": post,
	})
}

func DeletePost(c *fiber.Ctx) error {
	_, claims, _, err := helpers.CheckToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	postId := c.Params("id")
	post := models.Post{}
	result := helpers.DB.Where("user_id = ?", claims["sub"]).Delete(&post, postId)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.JSON(fiber.Map{
		"post": post,
	})
}
