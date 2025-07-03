package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
)

type RoomNameFieldProps struct {
	RoomName  string
	ReadOnly  bool
	Disabled  bool
	AutoFocus bool
}

func RoomNameField(props RoomNameFieldProps) []*hypp.VNode {
	title := ""
	var explanation *hypp.VNode
	if !props.ReadOnly {
		title = "2 letters alternated with 2 numbers. For example, \"R2D2\""
		explanation = html.P(
			hypp.HProps{"class": "explanation"},
			hypp.Text("The room name consists of 2 letters alternated with 2 numbers. For example, \"R2D2\"."),
		)
	}
	id := "room-name"
	return []*hypp.VNode{
		html.Label(
			hypp.HProps{
				"for":   id,
				"class": "room-name-label",
			},
			hypp.Text("Room name"),
		),
		html.Input(
			hypp.HProps{
				"id":        id,
				"readonly":  props.ReadOnly,
				"disabled":  props.Disabled,
				"autofocus": props.AutoFocus,
				"maxlength": 4,
				"pattern":   "[A-Z][0-9][A-Z][0-9]",
				"required":  true,
				"title":     title,
				"class":     "room-name-input",
				"value":     props.RoomName,
				"oninput":   dispatch.SetRoomNameByEvent,
			},
		),
		explanation,
	}
}
