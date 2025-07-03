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
	var status string
	ctaButton := html.Button(
		nil,
		hypp.Text("Next"),
	)
	inputDisabled := true

	if !s.ConnectingToWebSocket {
		status = "Waiting for room name"
		ctaButton = html.Button(
			hypp.HProps{
				"class": "cta",
				"type":  "submit",
			},
			hypp.Text("Connect"),
		)
		inputDisabled = false
	} else if !s.WebSocketConnected {
		status = "Joining room..."
	} else if !s.OpponentWebsocketConnected {
		status = "Waiting for opponent..."
	} else if !s.WebRTCConnected {
		status = "Connecting to opponent..."
	} else {
		status = "Connected"
		ctaButton = atom.Button(
			"Next",
			dispatch.GoToWhoGoesFirstScreen,
			hypp.HProps{"class": "cta"},
		)
	}

	return html.Main(
		hypp.HProps{
			"class":    "screen",
			"onsubmit": dispatch.JoinGame,
		},
		html.H1(nil, hypp.Text("Online - Join Game")),
		html.Form(
			hypp.HProps{
				"class":    "flex-column-center",
				"onsubmit": dispatch.JoinGame,
			},
			append(
				molecule.RoomNameField(molecule.RoomNameFieldProps{
					RoomName:  s.RoomName,
					AutoFocus: true,
					Disabled:  inputDisabled,
				}),
				html.P(hypp.HProps{"class": "status"}, hypp.Text(status)),
				molecule.SignalingErrors(s.SignalingErrors),
				html.Div(
					nil,
					molecule.CancelToStartPageButton(),
					ctaButton,
				),
			)...,
		),
	)
}
