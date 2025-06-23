package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func JoinGame(s *state.State) *hypp.VNode {
	roomName := ""
	signalingStep := state.SignalingStepDefault
	var signalingError *hypp.VNode
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
		signalingStep = s.Signaling.Step
		signalingError = molecule.SignalingError(s.Signaling.Error)
	}

	var status string
	var ctaButton *hypp.VNode
	inputDisabled := true
	switch signalingStep {
	case state.SignalingStepDefault:
		status = "Waiting for room name"
		ctaButton = html.Button(
			hypp.HProps{
				"class": "cta",
				"type":  "submit",
			},
			hypp.Text("Connect"),
		)
		inputDisabled = false
	case state.SignalingStepConnectingToWebSocket:
		status = "Joining room..."
		ctaButton = html.Button(
			nil,
			hypp.Text("Next"),
		)
	case state.SignalingStepIsConnectedToWebSocket:
	case state.SignalingStepOpponentIsConnectedToWebSocket,
		state.SignalingStepNewGameOffer,
		state.SignalingStepNewGameAnswer,
		state.SignalingStepJoinGameOffer,
		state.SignalingStepJoinGameAnswer:
		status = "Connecting to opponent..."
		ctaButton = html.Button(
			nil,
			hypp.Text("Next"),
		)
	case state.SignalingStepHasWebRTCConnection:
		status = "Connected"
		ctaButton = atom.Button(
			"Next",
			dispatch.JoinGame,
			hypp.HProps{"class": "cta"},
		)
	default:
		status = "Unknown signaling step"
	}

	return html.Main(
		hypp.HProps{
			"class":    "screen",
			"onsubmit": dispatch.JoinGame,
		},
		html.H1(nil, hypp.Text("Online - Join Game")),
		html.Form(
			hypp.HProps{
				"class": "flex-column-center",
			},
			append(
				molecule.RoomNameField(molecule.RoomNameFieldProps{
					RoomName:  roomName,
					AutoFocus: true,
					Disabled:  inputDisabled,
				}),
				html.P(hypp.HProps{"class": "status"}, hypp.Text(status)),
				signalingError,
				html.Div(
					nil,
					molecule.CancelToStartPageButton(),
					ctaButton,
				),
			)...,
		),
	)
}
