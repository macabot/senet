package organism

import (
	"encoding/json"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func PanicModal(s *state.State) *hypp.VNode {
	panicStackTrace := "[Unknown error]"
	if s.PanicStackTrace != "" {
		panicStackTrace = "[Error]\n" + s.PanicStackTrace
	}

	newState := s.Clone()
	seeAbove := "[See above]"
	newState.PanicStackTrace = seeAbove
	var stateJSON string
	if b, err := json.MarshalIndent(newState, "", "  "); err == nil {
		stateJSON = "[State]\n" + string(b)
	} else {
		stateJSON = "[Could not JSON encode state]\n" + err.Error()
	}

	details := panicStackTrace + "\n\n" + stateJSON

	return html.Div(
		hypp.HProps{
			"class": "panic-modal",
		},
		html.H1(
			nil,
			hypp.Text("Sorry, something went wrong!"),
		),
		html.Details(
			nil,
			html.Summary(nil, hypp.Text("Details")),
			html.Pre(nil, hypp.Text(details)),
		),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.ReloadPage,
			},
			hypp.Text("Reload"),
		),
	)
}
