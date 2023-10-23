package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/mazznoer/colorgrad"
	"image/png"
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

	app.Get("/:w/:h/hex/:hex/:hex2", func(c *fiber.Ctx) error {
		w, h, err := getWH(c)
		if err != nil {
			return err
		}

		hex, err := parseHexColor(c.Params("hex"))
		if err != nil {
			return err
		}

		hex2, err := parseHexColor(c.Params("hex2"))
		if err != nil {
			return err
		}

		grad, err := colorgrad.NewGradient().
			Colors(hex, hex2).
			Build()
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createGradientImage(w, h, grad)); err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Disposition", "inline")
		c.Set("Content-Type", "image/png")
		return c.SendStream(buf)
	})

	app.Get("/:w/:h/rgb/:r/:g/:b/:r2/:g2/:b2", func(c *fiber.Ctx) error {
		w, h, err := getWH(c)
		if err != nil {
			return err
		}

		rgb1, err := parseRGB(c, "r", "g", "b")
		if err != nil {
			return err
		}

		rgb2, err := parseRGB(c, "r2", "g2", "b2")
		if err != nil {
			return err
		}

		grad, err := colorgrad.NewGradient().
			Colors(*rgb1, *rgb2).
			Build()
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createGradientImage(w, h, grad)); err != nil {
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

		rgba, err := parseRGB(c, "r", "g", "b")
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createSolidImage(w, h, *rgba)); err != nil {
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

		rgba, err := parseRGBA(c, "r", "g", "b", "a")
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createSolidImage(w, h, *rgba)); err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Disposition", "inline")
		c.Set("Content-Type", "image/png")
		return c.SendStream(buf)
	})

	app.Get("/:w/:h/rgba/:r/:g/:b/:a/:r2/:g2/:b2/:a2", func(c *fiber.Ctx) error {
		w, h, err := getWH(c)
		if err != nil {
			return err
		}

		rgba1, err := parseRGBA(c, "r", "g", "b", "a")
		if err != nil {
			return err
		}

		rgba2, err := parseRGBA(c, "r2", "g2", "b2", "a2")
		if err != nil {
			return err
		}

		grad, err := colorgrad.NewGradient().
			Colors(*rgba1, *rgba2).
			Build()
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer([]byte{})
		if err := png.Encode(buf, createGradientImage(w, h, grad)); err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Disposition", "inline")
		c.Set("Content-Type", "image/png")
		return c.SendStream(buf)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(
			"/<w>/<h>/hex/<hex> - generate a solid color\n" +
				"/<w>/<h>/hex/<hex>/<hex> - generate a color gradient\n" +
				"/<w>/<h>/rgb/<r>/<g>/<b> - generate a solid color\n" +
				"/<w>/<h>/rgb/<r>/<g>/<b>/<r2>/<g2>/<b2> - generate a color gradient\n" +
				"/<w>/<h>/rgba/<r>/<g>/<b>/<a> - generate a solid color\n" +
				"/<w>/<h>/rgba/<r>/<g>/<b>/<a>/<r2>/<g2>/<b2>/<a2> - generate a color gradient\n",
		)
	})

	app.Listen(":3000")
}
