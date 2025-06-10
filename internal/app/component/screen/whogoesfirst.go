package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/component/organism"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func WhoGoesFirst(s *state.State) *hypp.VNode {
	hasDecision := s.CommitmentScheme.CanThrow()
	name0 := ""
	name1 := ""
	isPlayer0 := false
	if hasDecision {
		correctCall := s.CommitmentScheme.CallerPredictions[0] == s.CommitmentScheme.FlipperResults[0]
		if s.CommitmentScheme.IsCaller == correctCall {
			name0 = "You"
			name1 = "Opponent"
			isPlayer0 = true
		} else {
			name0 = "Opponent"
			name1 = "You"
		}
	}
	return html.Main(
		hypp.HProps{
			"class": map[string]bool{
				"who-goes-first-page": true,
				"has-decision":        hasDecision,
			},
		},
		html.H1(nil, hypp.Text("Who goes first?")),
		html.P(
			hypp.HProps{
				"class": "negotioating-text",
			},
			hypp.Text("Negotiating commitment scheme..."),
		),
		molecule.Loader(hasDecision),
		html.Div(
			hypp.HProps{
				"class": "players-wrapper",
			},
			molecule.WhoGoesFirstPlayer(0, name0),
			molecule.WhoGoesFirstPlayer(1, name1),
		),
		html.Div(
			nil,
			molecule.CancelToStartPageButton(),
			atom.Button(
				"Play",
				hypp.ActionAndPayload[*state.State]{
					Action:  dispatch.GoToOnlineScreen,
					Payload: isPlayer0,
				},
				hypp.HProps{
					"class":    "cta",
					"disabled": !hasDecision,
				},
			),
		),
		organism.Disconnected(s),
	)
}
