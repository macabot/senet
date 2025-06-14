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
	roomName := ""
	opponentIsConnected := false
	isDoneSignaling := false // FIXME
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
		opponentIsConnected = s.Signaling.Step == state.SignalingStepOpponentIsConnectedToWebSocket
	}

	var status string
	onClickNext := dispatch.NoOp
	cta := false
	if roomName == "" {
		status = "Creating room..."
	} else if !opponentIsConnected {
		status = "Waiting for opponent..."
	} else if !isDoneSignaling {
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
			RoomName: roomName,
			ReadOnly: true,
			Disabled: roomName == "",
		})...,
	)
	children = append(
		children,
		html.P(nil, hypp.Text("Share the room name with your opponent.")),
		html.P(hypp.HProps{"class": "status"}, hypp.Text(status)),
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
