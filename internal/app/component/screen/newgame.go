package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func NewGame(s *state.State) *hypp.VNode {
	var status string
	var onClickNext hypp.Dispatchable
	cta := false

	if s.RoomName == "" || !s.WebSocketConnected {
		status = "Creating room..."
	} else if !s.OpponentWebsocketConnected {
		status = "Waiting for opponent..."
	} else if !s.WebRTCConnected {
		status = "Connecting to opponent..."
	} else {
		status = "Connected"
		onClickNext = dispatch.GoToWhoGoesFirstScreen
		cta = true
	}

	// TODO Refactor this once Hypp supports fragments.
	children := []*hypp.VNode{
		html.H1(nil, hypp.Text("Online - New Game")),
	}
	children = append(
		children,
		molecule.RoomNameField(molecule.RoomNameFieldProps{
			RoomName: s.RoomName,
			ReadOnly: true,
			Disabled: s.RoomName == "",
		})...,
	)
	children = append(
		children,
		html.P(nil, hypp.Text("Share the room name with your opponent.")),
		html.P(hypp.HProps{"class": "status"}, hypp.Text(status)),
		molecule.SignalingErrors(s.SignalingErrors),
		html.Div(
			nil,
			molecule.CancelToStartPageButton(),
			atom.Button(
				"Next",
				onClickNext,
				hypp.HProps{
					"class": map[string]bool{
						"cta": cta,
					},
				},
			),
		),
	)
	return html.Main(
		hypp.HProps{
			"class": "screen",
		},
		children...,
	)
}
