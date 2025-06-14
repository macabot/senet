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
	isConnected := state.Scaledrone.IsConnected()
	opponentIsConnected := false
	roomName := ""
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
	}
	var status string
	onClickNext := dispatch.NoOp
	roomNameDisabled := false
	cta := false
	if !isConnected {
		status = "Connecting..."
		roomNameDisabled = true
	} else if !opponentIsConnected {
		status = "Waiting for opponent..."
	} else {
		status = "Connected"
		onClickNext = dispatch.GoToWhoGoesFirstScreen
		cta = true
	}

	children := []*hypp.VNode{
		html.H1(nil, hypp.Text("Online - New Game")),
	}
	children = append(
		children,
		molecule.RoomNameField(molecule.RoomNameFieldProps{
			RoomName: roomName,
			Disabled: roomNameDisabled,
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
