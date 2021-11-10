package blivechat

import (
	"fmt"
	"strconv"
)

type RGB struct {
	R, G, B int
}

func (c RGB) ToHex() string {
	return fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
}

var RGBWhite = RGB{
	R: 255,
	G: 255,
	B: 255,
}

func HexToRGB(hexColor string) (rgb RGB) {
	rgb.R, rgb.G, rgb.B = 255, 255, 255
	if len(hexColor) < 6 {
		return
	}
	if len(hexColor) > 6 {
		hexColor = hexColor[len(hexColor)-6:]
	}
	values, err := strconv.ParseUint(hexColor, 16, 32)
	if err != nil {
		return
	}
	rgb.R = int(values >> 16)
	rgb.G = int((values >> 8) & 0xFF)
	rgb.B = int(values & 0xFF)
	return
}

func IntToRGB(intColor int) (rgb RGB) {
	rgb.R, rgb.G, rgb.B = 255, 255, 255
	values := uint64(intColor)
	rgb.R = int(values >> 16)
	rgb.G = int((values >> 8) & 0xFF)
	rgb.B = int(values & 0xFF)
	return
}

func SetForegroundColor(rgb RGB, msg string) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", rgb.R, rgb.G, rgb.B, msg)
}

func SetBackgroundColor(rgb RGB, msg string) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", rgb.R, rgb.G, rgb.B, msg)
}
