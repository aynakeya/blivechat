module blivechat

go 1.16

require (
	github.com/aynakeya/blivedm v0.0.0
	github.com/aynakeya/gocui v0.5.4
)

replace (
	github.com/aynakeya/blivedm => ../blivedm
	github.com/aynakeya/gocui => ../gocui
)
