package blivechat

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestHexToRGB(t *testing.T) {
	fmt.Println(HexToRGB("#ffffff"))
	fmt.Println(HexToRGB("#000000"))
	fmt.Println(HexToRGB("#32a852"))
}

func TestIntToHex(t *testing.T) {
	fmt.Println(IntToRGB(cast.ToInt("5816798")).ToHex())
	fmt.Println(HexToRGB("#32a852"), HexToRGB("#32a852").ToHex())
}
