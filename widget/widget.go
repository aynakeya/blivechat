package widget

import "github.com/awesome-gocui/gocui"

type DynamicSizeFunc func(x, y int) (x0, y0, x1, y1 int)

type BaseWidget struct {
	ViewName       string // name for gocui.View
	ParentViewName string
	DisplayName    string // display name, title
	GetSize        DynamicSizeFunc
}

type LinkedWidget struct {
	NextViewName string
}

type DoubleLinkedWidget struct {
	PrevViewName string
	NextViewName string
}

type KeyCombo struct {
	Key      gocui.Key
	Modifier gocui.Modifier
}
