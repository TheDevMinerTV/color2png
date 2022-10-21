package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var errInvalidDimension = errors.New("invalid dimensions")

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
	raw := c.Params(param)
	if raw == "" {
		raw = "0"
	}

	v, err := strconv.Atoi(raw)
	if err != nil {
		return 0, err
	}

	return clamp(uint8(v), 0, 255), err
}

func parseDimension(raw string) (int, error) {
	v, err := strconv.Atoi(raw)
	if err != nil {
		return 0, errInvalidDimension
	}
	if v <= 0 {
		return 0, errInvalidDimension
	}
	v = clamp(v, 1, 512)

	return v, nil
}

func getWH(c *fiber.Ctx) (int, int, error) {
	w, err := parseDimension(c.Params("w"))
	if err != nil {
		return 0, 0, err
	}

	h, err := parseDimension(c.Params("h"))
	if err != nil {
		return 0, 0, err
	}

	return w, h, nil
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
