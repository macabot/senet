package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func RoomNameField(roomName string, readOnly bool) []*hypp.VNode {
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
				"readonly":  readOnly,
				"maxlength": 4,
				"class":     "room-name-input",
				"value":     roomName,
			},
		),
	}
}
