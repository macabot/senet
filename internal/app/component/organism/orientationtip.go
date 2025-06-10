package organism

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func OrientationTip(s *state.State) *hypp.VNode {
	if s.HideOrientationTip {
		return nil
	}
	return html.Div(
		hypp.HProps{
			"class": "orientation-tip",
		},
		html.P(
			nil,
			hypp.Text("Rotate or enlarge your screen for a better experience."),
		),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.ToggleOrientationTip,
			},
			hypp.Text("OK"),
		),
	)
}
