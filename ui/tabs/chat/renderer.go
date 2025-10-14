package chat

import (
	"blivechat/model"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"time"
)

type ChatRenderer interface {
	UseStyle(style *Style)
	Styles() *Style // return current styles, modifiable
	RoomTitle(msg *model.RoomInfo) string
	Danmuku(msg *model.Danmaku) string
	LiveStart(msg *model.LiveStart) string
	LiveStop(msg *model.LiveStop) string
	Gift(msg *model.Gift) string
	GuardBuy(msg *model.GuardBuy) string
	SuperChat(msg *model.SuperChat) string
	SystemMsg(msg *model.SystemMsg) string
	InteractWord(msg *model.InteractWord) string
}

type DefaultRenderer struct {
	style *Style
}

var _ ChatRenderer = (*DefaultRenderer)(nil)

func NewDefaultRenderer() *DefaultRenderer {
	style := DefaultStyles()
	return &DefaultRenderer{
		style: &style,
	}
}

func (r *DefaultRenderer) UseStyle(style *Style) {
	r.style = style
}

func (r *DefaultRenderer) Styles() *Style {
	return r.style
}

func (r *DefaultRenderer) RoomTitle(msg *model.RoomInfo) string {
	status := "未开播"
	if msg.Data.LiveStatus == 1 {
		status = r.style.LiveStatus.Start.Render("直播中")
	} else {
		status = r.style.LiveStatus.Stop.Render("未开播")
	}

	liveTime := ""
	if msg.Data.LiveTime != "" {
		liveTime = fmt.Sprintf(" | 开播时间: %s", msg.Data.LiveTime)
	}

	roomID := msg.Data.RoomId
	if msg.Data.ShortId != 0 {
		roomID = msg.Data.ShortId
	}

	return fmt.Sprintf(
		"%s %d %s | %s%s",
		fmt.Sprintf("[%d]", roomID),
		msg.Data.Uid,
		status,
		msg.Data.Title,
		liveTime,
	)
}

func (r *DefaultRenderer) renderTime(t time.Time) string {
	return r.style.Danmaku.Timestamp.Render(t.Format("[15:04:05]"))
}

func (r *DefaultRenderer) renderTimestamp(tsec int64) string {
	return r.renderTime(safeTime(tsec))
}

// Medal 渲染勋章 [名 LvX]（颜色取勋章色）
func (r *DefaultRenderer) Medal(name string, level int, colorDec uint32) string {
	if name == "" || level <= 0 {
		return ""
	}
	c := int2color(colorDec)
	left := r.style.Medal.Bracket.Render("[")
	right := r.style.Medal.Bracket.Render("]")
	body := lipgloss.JoinHorizontal(lipgloss.Top,
		r.style.Medal.Name.Foreground(c).Render(name),
		lipgloss.NewStyle().Render(" "),
		r.style.Medal.Level.Foreground(c).Render(fmt.Sprintf("Lv%d", level)),
	)
	return left + body + right + " "
}

// 渲染大航海徽标；0 不显示，1=舰长 2=提督 3=总督
func (r *DefaultRenderer) guardBadge(level int) string {
	switch level {
	case 1:
		return r.style.GuardBadge.Governor.Render("总")
	case 2:
		return r.style.GuardBadge.Admiral.Render("提")
	case 3:
		return r.style.GuardBadge.Captain.Render("舰")
	default:
		return ""
	}
}

func (r *DefaultRenderer) Danmuku(msg *model.Danmaku) string {
	var (
		medalName  string
		medalLevel int
		medalColor uint32
		dmColor    int
		username   string
		guardLv    int
	)
	username = ""
	if msg.Sender != nil {
		username = msg.Sender.Uname
		guardLv = msg.Sender.GuardLevel // ← 读取大航海等级（0/1/2/3）
		if msg.Sender.Medal != nil {
			medalName = msg.Sender.Medal.Name
			medalLevel = msg.Sender.Medal.Level
			medalColor = (msg.Sender.Medal.V2ColorLevel >> 8) & 0xffffff
		}
	}
	ts := r.renderTimestamp(msg.Timestamp)

	// 新增：渲染大航海徽标
	guard := ""
	if guardLv > 0 {
		guard = r.guardBadge(guardLv) + " "
	}

	if msg.Sender.Admin {
		guard = r.style.User.Admin.Render("管") + " "
	}

	medal := r.Medal(medalName, medalLevel, medalColor)

	usernameStyle := r.style.Danmaku.Username
	if msg.Sender.UserColor != 0 {
		usernameStyle = usernameStyle.Foreground(int2color(msg.Sender.UserColor))
	}
	user := usernameStyle.Render(username)

	if msg.Extra != nil {
		dmColor = msg.Extra.Color
	}
	dmStyle := r.style.Danmaku.Message
	if dmColor != 0 {
		dmStyle = dmStyle.Foreground(int2color(uint32(msg.Extra.Color)))

	}
	dm := dmStyle.Render(msg.Content)
	return fmt.Sprintf("%s %s%s%s: %s", ts, medal, guard, user, dm)
}

func (r *DefaultRenderer) LiveStart(msg *model.LiveStart) string {
	tsv := r.renderTimestamp(int64(msg.LiveTime))
	return fmt.Sprintf("%s %s", tsv, r.style.LiveStatus.Start.Render("[直播开始]"))
}

func (r *DefaultRenderer) LiveStop(msg *model.LiveStop) string {
	tsv := r.renderTimestamp(msg.SendTime)
	return fmt.Sprintf("%s %s", tsv, r.style.LiveStatus.Stop.Render("[直播结束]"))
}

func (r *DefaultRenderer) Gift(msg *model.Gift) string {
	tsv := r.renderTimestamp(msg.Timestamp)
	tag := r.style.Gift.Tag.Render("[礼物]")
	user := r.style.Gift.User.Render(msg.Uname)
	name := r.style.Gift.Name.Render(msg.GiftName)
	cnt := r.style.Gift.Count.Render(fmt.Sprintf("×%d", msg.Num))

	totalYuan := float64(msg.Price*msg.Num) / 1000.0
	price := r.style.Gift.Price.Render(fmt.Sprintf("共 ¥%.2f", totalYuan))
	if !(msg.CoinType == "gold") {
		// 银瓜子礼物不显示金额（或换个提示）
		price = r.style.Gift.Price.Faint(true).Render("(银瓜子)")
	}

	core := lipgloss.JoinHorizontal(lipgloss.Top, tag, lipgloss.NewStyle().Render(" "), user, lipgloss.NewStyle().Render(" 赠送 "), name, lipgloss.NewStyle().Render(" "), cnt, lipgloss.NewStyle().Render(" "), price)
	return lipgloss.JoinHorizontal(lipgloss.Top, tsv, lipgloss.NewStyle().Render(" "), core)
}

func (r *DefaultRenderer) GuardBuy(msg *model.GuardBuy) string {
	tsv := r.renderTime(time.Now())
	tag := r.style.Guard.Tag.Render("[大航海]")
	user := r.style.Guard.User.Render(msg.Username)
	lv := r.style.Guard.Level.Render(fmt.Sprintf("Lv%d", msg.GuardLevel))
	price := r.style.Guard.Price.Render(fmt.Sprintf("¥%.2f", float64(msg.Price)/1000.0))

	core := lipgloss.JoinHorizontal(lipgloss.Top, tag, lipgloss.NewStyle().Render(" "), user, lipgloss.NewStyle().Render(" 开通 "), lv, lipgloss.NewStyle().Render(" "), price)
	return lipgloss.JoinHorizontal(lipgloss.Top, tsv, lipgloss.NewStyle().Render(" "), core)
}

func (r *DefaultRenderer) SuperChat(msg *model.SuperChat) string {
	tsv := r.renderTimestamp(msg.Ts)

	log.Info(msg)

	container := r.style.SuperChat.Container
	// todo: use sc background color
	//if bgColorDec != 0 {
	//	container = container.Background(s.ColorFromInt(uint32(bgColorDec), false))
	//}

	tag := r.style.SuperChat.Tag.Render("[SC]")
	price := r.style.SuperChat.Price.Render(fmt.Sprintf("¥%d", msg.Price))
	user := r.style.SuperChat.User.Render(msg.UserInfo.Uname)
	text := r.style.SuperChat.Msg.Render(msg.Message)

	core := lipgloss.JoinHorizontal(lipgloss.Top, tag, lipgloss.NewStyle().Render(" "), price, lipgloss.NewStyle().Render(" "), user, lipgloss.NewStyle().Render(": "), text)
	return lipgloss.JoinHorizontal(lipgloss.Top, tsv, lipgloss.NewStyle().Render(" "), container.Render(core))
}

func (r *DefaultRenderer) SystemMsg(msg *model.SystemMsg) string {
	ts := r.renderTime(time.Now())
	var body string
	switch msg.Level {
	case model.SystemMsgWarning:
		body = r.style.SystemMsg.Warn.Render(msg.Msg)
	case model.SystemMsgError:
		body = r.style.SystemMsg.Error.Render(msg.Msg)
	default:
		body = r.style.SystemMsg.Info.Render(msg.Msg)
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, ts, lipgloss.NewStyle().Render(" "), body)
}

func (r *DefaultRenderer) InteractWord(msg *model.InteractWord) string {
	if msg.MsgType != 1 {
		return ""
	}
	ts := r.renderTimestamp(int64(msg.Timestamp))
	name := msg.Uname
	userStyle := r.style.Danmaku.Username
	if msg.Uinfo.Base.NameColor != 0 {
		userStyle = userStyle.Foreground(int2color(uint32(msg.Uinfo.Base.NameColor)))
	}
	user := userStyle.Render(name)

	// 新增：渲染大航海徽标
	guard := ""
	if msg.Uinfo.Guard.Level > 0 {
		guard = r.guardBadge(msg.Uinfo.Guard.Level) + " "
	}

	medal := r.Medal(msg.FansMedal.MedalName, msg.FansMedal.MedalLevel, (hexToUint32(msg.Uinfo.Medal.V2MedalColorLevel, 1)>>8)&0xffffff)

	return fmt.Sprintf("%s %s%s%s %s", ts, medal, guard, user, r.style.Danmaku.Message.Render("进入直播间"))
}
