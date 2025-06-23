package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

type RoomNameFieldProps struct {
	RoomName  string
	ReadOnly  bool
	Disabled  bool
	AutoFocus bool
}

func RoomNameField(props RoomNameFieldProps) []*hypp.VNode {
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
				"class":     "room-name-input",
				"value":     props.RoomName,
			},
		),
	}
}
