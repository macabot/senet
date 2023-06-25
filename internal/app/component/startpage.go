package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func StartPage(s *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "start-page",
		},
	)
}
