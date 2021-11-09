package blivechat

import (
	"fmt"
	"testing"
)

func TestHexToRGB(t *testing.T) {
	fmt.Println(HexToRGB("#ffffff"))
	fmt.Println(HexToRGB("#000000"))
	fmt.Println(HexToRGB("#32a852"))
}

func TestSetUnameColor(t *testing.T) {
	fmt.Println(SetUnameColor("#ffffff", "hello"))
	fmt.Println(SetUnameColor("#000000", "hello"))
	fmt.Println(SetUnameColor("#32a852", "hello"))
}
