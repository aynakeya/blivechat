package blivechat

import (
	"fmt"
	"github.com/aynakeya/gocui"
	"time"
)

func viewPrintln(v *gocui.View,a interface{}) {
	fmt.Fprintln(v,a)
}

func viewPrint(v *gocui.View,a interface{}) {
	fmt.Fprint(v,a)
}

func viewPrintWithTime(v *gocui.View,a interface{}) {
	fmt.Fprintf(v,"%s >\n%s\n",time.Now().Format("2006/01/02 15:04:05"),a)
}