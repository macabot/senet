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
	signalingStep := state.SignalingStepDefault
	var signalingError *hypp.VNode
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
		signalingStep = s.Signaling.Step
		signalingError = molecule.SignalingError(s.Signaling.Error)
	}

	var status string
	onClickNext := dispatch.NoOp
	cta := false
	switch signalingStep {
	case state.SignalingStepDefault:
		status = ""
	case state.SignalingStepConnectingToWebSocket:
		status = "Creating room..."
	case state.SignalingStepIsConnectedToWebSocket:
		status = "Waiting for opponent..."
	case state.SignalingStepOpponentIsConnectedToWebSocket,
		state.SignalingStepNewGameOffer,
		state.SignalingStepNewGameAnswer,
		state.SignalingStepJoinGameOffer,
		state.SignalingStepJoinGameAnswer:
		status = "Connecting to opponent..."
	case state.SignalingStepHasWebRTCConnection:
		status = "Connected"
		onClickNext = dispatch.GoToWhoGoesFirstScreen
		cta = true
	default:
		status = "Unknown signaling step"
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
		signalingError,
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
