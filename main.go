package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m0ep/kekw-as-a-service-go/kekw"
	"strconv"
)

func main() {
	app := fiber.New()

	app.Add("GET", "/", func(c *fiber.Ctx) error {
		angleQuery := c.Query("angle", "")

		if "" == angleQuery {
			return kekw.KekwRandom(c)
		} else {
			angle, err := strconv.Atoi(angleQuery)
			if nil != err {
				angle = 0
			}

			return kekwFromAngle(c, angle)
		}

	})
	app.Add("GET", "/:angle", func(c *fiber.Ctx) error {
		angleParam := c.Params("angle", "0")
		angle, err := strconv.Atoi(angleParam)
		if nil != err {
			angle = 0
		}

		return kekwFromAngle(c, angle)
	})

	_ = app.Listen(":3000")
}

func kekwFromAngle(c *fiber.Ctx, angle int) error {
	switch {
	case -359 > angle:
		angle = -359
	case 359 < angle:
		angle = 359
	}

	return kekw.Kekw(c, angle)
}
