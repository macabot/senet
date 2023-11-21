package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func WhoGoesFirstPage(s *state.State) *hypp.VNode {
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
		whoGoesFirstLoader(),
		html.Div(
			hypp.HProps{
				"class": "players-wrapper",
			},
			whoGoesFirstPlayer(0, name0),
			whoGoesFirstPlayer(1, name1),
		),
		html.Div(
			nil,
			html.Button(
				hypp.HProps{
					"onclick": dispatch.ToSignalingPageAction(),
				},
				hypp.Text("Back"),
			),
			html.Button(
				hypp.HProps{
					"class":    "cta",
					"disabled": !hasDecision,
					"onclick":  dispatch.ToOnlinePlayerVsPlayerAction(isPlayer0),
				},
				hypp.Text("Play"),
			),
		),
	)
}

func whoGoesFirstLoader() *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": "who-loader-wrapper",
		},
		html.Div(hypp.HProps{"class": "who-loader"}),
	)
}

func whoGoesFirstPlayer(playerIndex int, name string) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"player-wrapper":                      true,
				fmt.Sprintf("player-%d", playerIndex): true,
			},
		},
		html.Div(
			hypp.HProps{
				"class": "player",
			},
			hypp.Text(name),
		),
	)
}
