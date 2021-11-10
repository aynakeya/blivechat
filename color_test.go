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
