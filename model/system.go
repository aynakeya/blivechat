package model

const (
	SystemMsgInfo = iota
	SystemMsgWarning
	SystemMsgError
)

type SystemMsg struct {
	Level int
	Msg   string
}

type DebugLineMsg struct {
	Line string
}
