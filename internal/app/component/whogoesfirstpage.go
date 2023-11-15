package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func WhoGoesFirstPage(s *state.State) *hypp.VNode {
	hasDecision := s.CommitmentScheme.CanThrow()
	name0 := ""
	name1 := ""
	if hasDecision {
		correctCall := s.CommitmentScheme.CallerPredictions[0] == s.CommitmentScheme.FlipperResults[0]
		if s.CommitmentScheme.IsCaller == correctCall {
			name0 = "You"
			name1 = "Opponent"
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
		html.Div(
			hypp.HProps{
				"class": "players-wrapper",
			},
			player(0, Player{Player: &state.Player{Name: name0}}, false),
			player(0, Player{Player: &state.Player{Name: name1}}, false),
		),
	)
}
