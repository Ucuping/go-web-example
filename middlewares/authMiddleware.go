package middlewares

import (
	"github.com/Ucuping/go-web-example/helpers"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, c.Status(fiber.StatusUnauthorized).SendString("Unathorized")
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return []byte(os.Getenv("JWT_SECRET")), nil
	// })

	// if err != nil {
	// 	return err
	// }

	// if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	return c.Next()
	// }

	token, _, ok, err := helpers.CheckToken(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		// return c.Redirect("/auth/login")
	}

	if ok && token.Valid {
		return c.Next()
	}

	// return c.Redirect("/auth/login")
	return c.Status(fiber.StatusUnauthorized).SendString("Unathorized")
}
