package kekw

import (
	"bytes"
	_ "embed"
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"image/color"
	"math/rand"
)

//go:embed resources/kekw.png
var kekwImageData []byte

func Kekw(c *fiber.Ctx, angle int) error {
	srcImage, _ := imaging.Decode(bytes.NewReader(kekwImageData))
	rotatedImage := imaging.Rotate(srcImage, float64(angle), color.Transparent)
	c.Response().Header.Add("Content-Type", "image/png")
	return imaging.Encode(c, rotatedImage, imaging.PNG)
}

func KekwRandom(c *fiber.Ctx) error {
	return Kekw(c, rand.Intn(360))
}
