package chat

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strconv"
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

func hexToUint32(s string, prefixLen int) uint32 {
	if len(s) <= prefixLen {
		return 0
	}
	parseInt, err := strconv.ParseInt(s[prefixLen:], 16, 64)
	if err != nil {
		return 0
	}
	return uint32(parseInt)
}
