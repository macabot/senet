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
		case state.SignalingStepNewGameAnswer:
			modal = signalingModal(s, signalingNewGameAnswer)
		case state.SignalingStepJoinGameOffer:
			modal = signalingModal(s, signalingJoinGameOffer)
		case state.SignalingStepJoinGameAnswer:
			modal = signalingModal(s, signalingJoinGameAnswer)
		}
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.Button(
			hypp.HProps{
				"class":   "cta",
				"onclick": dispatch.SetSignalingStepNewGameOfferAction(),
			},
			hypp.Text("New game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "cta",
				"onclick": dispatch.SetSignalingStepJoinGameOfferAction(),
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

func connectionStates(s *state.State) *hypp.VNode {
	iceConnectionState := "[unset]"
	connectionState := "[unset]"
	if s.Signaling != nil {
		if s.Signaling.ICEConnectionState != "" {
			iceConnectionState = s.Signaling.ICEConnectionState
		}
		if s.Signaling.ConnectionState != "" {
			connectionState = s.Signaling.ConnectionState
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
	)
}

func signalingNewGameOffer(s *state.State) *hypp.VNode {
	offer := "[error: Signaling is nil]"
	nextDisabled := true
	if s.Signaling != nil {
		if s.Signaling.Loading {
			offer = "[Loading...]"
		} else if s.Signaling.Offer == "" {
			offer = "[error: Signaling.Offer is empty]"
		} else {
			offer = s.Signaling.Offer
			nextDisabled = false
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
				"value":    offer,
			},
		),
		connectionStates(s),
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
					"disabled": nextDisabled,
					"onclick":  dispatch.SetSignalingStepNewGameAnswerAction(),
				},
				hypp.Text("Next"),
			),
		),
	)
}

func signalingNewGameAnswer(s *state.State) *hypp.VNode {
	connectDisabled := true
	connectionState := ""
	if s.Signaling != nil {
		connectDisabled = s.Signaling.Answer == ""
		connectionState = s.Signaling.ConnectionState
	}
	var ctaButton *hypp.VNode
	if connectionState == "connecting" || connectionState == "connected" {
		ctaButton = html.Button(
			hypp.HProps{
				"class":    "cta",
				"disabled": connectionState == "connecting",
				// "onclick"
			},
			hypp.Text("Play"),
		)
	} else {
		ctaButton = html.Button(
			hypp.HProps{
				"class":    "cta",
				"disabled": connectDisabled,
				"onclick":  dispatch.ConnectNewGameAction(),
			},
			hypp.Text("Connect"),
		)
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.P(nil, hypp.Text("Paste the answer of your opponent below.")),
		html.Textarea(
			hypp.HProps{
				"id":      "answer-textarea",
				"oninput": dispatch.SetSignalingAnswerAction(),
			},
		),
		connectionStates(s),
		html.Div(
			nil,
			html.Button(
				hypp.HProps{
					"onclick": dispatch.SetSignalingStepNewGameOfferAction(),
				},
				hypp.Text("Back"),
			),
			ctaButton,
		),
	)
}

func signalingJoinGameOffer(s *state.State) *hypp.VNode {
	nextDisabled := true
	if s.Signaling != nil {
		nextDisabled = s.Signaling.Offer == ""
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.P(nil, hypp.Text("Paste the offer of your opponent below.")),
		html.Textarea(
			hypp.HProps{
				"id":      "offer-textarea",
				"oninput": dispatch.SetSignalingOfferAction(),
			},
		),
		connectionStates(s),
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
					"disabled": nextDisabled,
					"onclick":  dispatch.SetSignalingStepJoinGameAnswerAction(),
				},
				hypp.Text("Next"),
			),
		),
	)
}

func signalingJoinGameAnswer(s *state.State) *hypp.VNode {
	answer := "[error: Signaling is nil]"
	connectDisabled := true
	if s.Signaling != nil {
		if s.Signaling.Loading {
			answer = "[Loading...]"
		} else if s.Signaling.Answer == "" {
			answer = "[error: Signaling.Answer is empty]"
		} else {
			answer = s.Signaling.Answer
			connectDisabled = false
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
				"id":       "answer-textarea",
				"readonly": true,
				"onclick":  dispatch.EffectsAction(dispatch.SelectTextareaEffect("answer-textarea")),
				"value":    answer,
			},
		),
		connectionStates(s),
		html.Div(
			nil,
			html.Button(
				hypp.HProps{
					"onclick": dispatch.SetSignalingStepJoinGameOfferAction(),
				},
				hypp.Text("Back"),
			),
			html.Button(
				hypp.HProps{
					"class":    "cta",
					"disabled": connectDisabled,
					// TODO onclick
				},
				hypp.Text("Connect"),
			),
		),
	)
}
