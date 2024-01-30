package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func Disconnected(s *state.State) *hypp.VNode {
	if s.Signaling == nil || s.Signaling.ConnectionState == "connected" {
		return nil
	}
	return html.Div(
		hypp.HProps{
			"class": "disconnected",
		},
		html.H3(
			nil,
			hypp.Text("There is a problem with the connection"),
		),
		html.P(
			hypp.HProps{
				"class": "warning",
			},
			hypp.Text("If you refresh the page, all progress will be lost."),
		),
		connectionStates(s),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.ToStartPageAction(),
			},
			hypp.Text("Start page"),
		),
	)
}
