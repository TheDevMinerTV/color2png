package main

import (
	"bytes"
	"image/color"
	"image/png"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:w/:h/hex/:hex", func(c *fiber.Ctx) error {
		w, h, err := getWH(c)
		if err != nil {
			return err
		}

		rgba, err := parseHexColor(c.Params("hex"))
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createSolidImage(w, h, rgba)); err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Disposition", "inline")
		c.Set("Content-Type", "image/png")
		return c.SendStream(buf)
	})

	app.Get("/:w/:h/rgb/:r/:g/:b", func(c *fiber.Ctx) error {
		w, h, err := getWH(c)
		if err != nil {
			return err
		}

		r, err := getUint8Param(c, "r")
		if err != nil {
			return err
		}
		g, err := getUint8Param(c, "g")
		if err != nil {
			return err
		}
		b, err := getUint8Param(c, "b")
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createSolidImage(w, h, color.RGBA{
			R: r,
			G: g,
			B: b,
			A: 255,
		})); err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Disposition", "inline")
		c.Set("Content-Type", "image/png")
		return c.SendStream(buf)
	})

	app.Get("/:w/:h/rgba/:r/:g/:b/:a", func(c *fiber.Ctx) error {
		w, h, err := getWH(c)
		if err != nil {
			return err
		}

		r, err := getUint8Param(c, "r")
		if err != nil {
			return err
		}
		g, err := getUint8Param(c, "g")
		if err != nil {
			return err
		}
		b, err := getUint8Param(c, "b")
		if err != nil {
			return err
		}
		a, err := getUint8Param(c, "a")
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createSolidImage(w, h, color.RGBA{
			R: r,
			G: g,
			B: b,
			A: a,
		})); err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Disposition", "inline")
		c.Set("Content-Type", "image/png")
		return c.SendStream(buf)
	})

	app.Listen(":3000")
}
