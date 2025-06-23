package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
)

func JoinGame(s *state.State) *hypp.VNode {
	hasConnection := true
	roomName := ""
	var nextButton *hypp.VNode
	cta := hypp.HProps{"class": "cta"}
	if !hasConnection {
		nextButton = atom.Button("Connect", nil, cta)
	} else {
		nextButton = atom.Button("Next", nil, cta)
	}
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
	}

	// TODO Refactor this once Hypp supports fragments.
	children := []*hypp.VNode{
		html.H1(nil, hypp.Text("Online - Join Game")),
	}
	children = append(
		children,
		molecule.RoomNameField(molecule.RoomNameFieldProps{
			RoomName:  roomName,
			AutoFocus: true,
		})...,
	)
	children = append(
		children,
		html.Div(
			nil,
			molecule.CancelToStartPageButton(),
			nextButton,
		),
	)
	return html.Main(
		hypp.HProps{
			"class": "screen",
		},
		children...,
	)
}
