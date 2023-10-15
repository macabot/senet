package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func SignalingPage(s *state.State) *hypp.VNode {
	var modal *hypp.VNode
	if s.Signaling != nil {
		switch s.Signaling.Step {
		case state.SignalingStepNewGameOffer:
			modal = signalingModal(s, signalingNewGameOffer)
		case state.SignalingStepNewGameAnswer: // TODO
		case state.SignalingStepJoinGameOffer: // TODO
		case state.SignalingStepJoinGameAnswer: // TODO
		}
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.Button(
			hypp.HProps{
				"class":   "signaling cta",
				"onclick": dispatch.SetSignalingStepNewGameOfferAction(),
			},
			hypp.Text("New game"),
		),
		html.Button(
			hypp.HProps{
				"class": "signaling cta",
				// TODO "onclick"
			},
			hypp.Text("Join game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.ToStartPageAction(),
			},
			hypp.Text("Back"),
		),
		modal,
	)
}

func signalingModal(s *state.State, f func(s *state.State) *hypp.VNode) *hypp.VNode {
	return html.Div(
		hypp.HProps{
			"class": "signaling-modal",
		},
		f(s),
	)
}

func signalingNewGameOffer(s *state.State) *hypp.VNode {
	offer := "[error: Signaling is nil]"
	if s.Signaling != nil {
		if s.Signaling.Loading {
			offer = "[Loading...]"
		} else if s.Signaling.Offer == "" {
			offer = "[error: Signaling.Offer is empty]"
		} else {
			offer = s.Signaling.Offer
		}
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.P(nil, hypp.Text("Copy the text below and send it to your opponent.")),
		html.Textarea(
			hypp.HProps{
				"id":       "offer-textarea",
				"readonly": true,
				"onclick":  dispatch.EffectsAction(dispatch.SelectTextareaEffect("offer-textarea")),
			},
			hypp.Text(offer),
		),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.ToSignalingPageAction(),
			},
			hypp.Text("Back"),
		),
	)
}
