package renderer

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"time"
)

func safeTime(ts int64) time.Time {
	// > 10^12 视为毫秒
	if ts > 1_000_000_000_000 {
		return time.UnixMilli(ts)
	}
	return time.Unix(ts, 0)
}

func int2color(i uint32) lipgloss.Color {
	return lipgloss.Color(fmt.Sprintf("#%06X", i))
}
