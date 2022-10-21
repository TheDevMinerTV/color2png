package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/gofiber/fiber/v2"
)

func parseHexColor(s string) (c color.RGBA, err error) {
	switch len(s) {
	case 8:
		_, err = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
		return c, err

	case 6:
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
		c.A = 255
		return c, err

	case 4:
		if _, err = fmt.Sscanf(s, "%1x%1x%1x%1x", &c.R, &c.G, &c.B, &c.A); err != nil {
			return color.RGBA{}, err
		}

		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A *= 17
		return c, nil

	case 3:
		if _, err = fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B); err != nil {
			return color.RGBA{}, err
		}

		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A = 255
		return c, nil

	default:
		return color.RGBA{}, fmt.Errorf("invalid length, must be 8, 6, 4 or 3")
	}
}

func getUint8Param(c *fiber.Ctx, param string) (uint8, error) {
	v, err := c.ParamsInt(param, 0)
	if err != nil {
		return 0, err
	}

	return uint8(v), err
}

func getWH(c *fiber.Ctx) (int, int, error) {
	w, err := c.ParamsInt("w", 32)
	if err != nil {
		return 0, 0, err
	}

	h, err := c.ParamsInt("h", 32)
	if err != nil {
		return 0, 0, err
	}

	return w, h, err
}

func createSolidImage(w, h int, c color.RGBA) image.Image {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i.Set(x, y, c)
		}
	}

	return i
}

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

		b := bytes.NewBuffer([]byte{})
		if err := png.Encode(b, createSolidImage(w, h, rgba)); err != nil {
			return c.SendStatus(500)
		}

		return c.SendStream(b)
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

		return c.SendStream(buf)
	})

	app.Listen(":3000")
}
