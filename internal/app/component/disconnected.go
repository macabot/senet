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
	var message *hypp.VNode
	if s.Signaling.ConnectionState == "disconnected" {
		message = html.P(
			hypp.HProps{
				"class": "warning",
			},
			hypp.Text("If you refresh the page, all progress will be lost. Wait to see if the problem will resolve itself."),
		)
	} else {
		message = html.P(
			hypp.HProps{
				"class": "error",
			},
			hypp.Text("The connection is lost. Go back to the start page."),
		)
	}

	return html.Div(
		hypp.HProps{
			"class": "disconnected",
		},
		html.H3(
			nil,
			hypp.Text("There is a problem with the connection"),
		),
		message,
		connectionStates(s),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.GoToStartPage,
			},
			hypp.Text("Start page"),
		),
	)
}
