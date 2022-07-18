package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m0ep/kekw-as-a-service-go/kekw"
	"strconv"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("🚀👍")
	}).Get("/favicon.ico", func(c *fiber.Ctx) error {
		return kekw.KekwStatic(c)
	}).Get("/", func(c *fiber.Ctx) error {
		angleQuery := c.Query("angle", "")

		if "" == angleQuery {
			return kekw.KekwRandom(c)
		} else {
			angle, err := strconv.Atoi(angleQuery)
			if nil != err {
				angle = 0
			}

			if -359 > angle {
				angle = -359
			} else if 359 < angle {
				angle = 359
			}

			return kekw.KekwAngle(c, angle)
		}
	})

	_ = app.Listen(":3000")
}
