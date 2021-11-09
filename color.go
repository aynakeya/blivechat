package blivechat

import (
	"fmt"
	"strconv"
)

type RGB struct {
	r, g, b int
}

var RGBWhite = RGB{
	r: 255,
	g: 255,
	b: 255,
}

func HexToRGB(hexColor string) (rgb RGB) {
	rgb.r, rgb.g, rgb.b = 255, 255, 255
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
	rgb.r = int(values >> 16)
	rgb.g = int((values >> 8) & 0xFF)
	rgb.b = int(values & 0xFF)
	return
}

func IntToRGB(intColor int) (rgb RGB) {
	rgb.r, rgb.g, rgb.b = 255, 255, 255
	values := uint64(intColor)
	rgb.r = int(values >> 16)
	rgb.g = int((values >> 8) & 0xFF)
	rgb.b = int(values & 0xFF)
	return
}

func SetForegroundColor(rgb RGB, msg string) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", rgb.r, rgb.g, rgb.b, msg)
}

func SetBackgroundColor(rgb RGB, msg string) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", rgb.r, rgb.g, rgb.b, msg)
}
