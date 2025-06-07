package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func ConnectionStates(s *state.State) *hypp.VNode {
	iceConnectionState := "[unset]"
	connectionState := "[unset]"
	readyState := "[unset]"
	if s.Signaling != nil {
		if s.Signaling.ICEConnectionState != "" {
			iceConnectionState = s.Signaling.ICEConnectionState
		}
		if s.Signaling.ConnectionState != "" {
			connectionState = s.Signaling.ConnectionState
		}
		if s.Signaling.ReadyState != "" {
			readyState = s.Signaling.ReadyState
		}
	}
	return html.Div(
		hypp.HProps{"class": "connection-state"},
		html.Div(
			nil,
			html.Span(nil, hypp.Text("ICE connection state:")),
			html.B(nil, hypp.Text(iceConnectionState)),
		),
		html.Div(
			nil,
			html.Span(nil, hypp.Text("Connection state:")),
			html.B(nil, hypp.Text(connectionState)),
		),
		html.Div(
			nil,
			html.Span(nil, hypp.Text("Ready state:")),
			html.B(nil, hypp.Text(readyState)),
		),
	)
}
