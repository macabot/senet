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
				"onclick": dispatch.SetSignalingStepNewGameOffer,
			},
			hypp.Text("New game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "cta",
				"onclick": dispatch.SetSignalingStepJoinGameOffer,
			},
			hypp.Text("Join game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.GoToStartPage,
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
	readyState := "[unset]"
	if s.Signaling != nil {
		if s.Signaling.ICEConnectionState != "" {
			iceConnectionState = s.Signaling.ICEConnectionState
		}
		if s.Signaling.ConnectionState != "" {
			connectionState = s.Signaling.ConnectionState
		}
		if s.Signaling.ReadyState != "" {
			readyState = s.Signaling.ReadyState
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
		html.Div(
			nil,
			html.Span(nil, hypp.Text("Ready state:")),
			html.B(nil, hypp.Text(readyState)),
		),
	)
}

func signalingError(signaling *state.Signaling, isOffer bool) *hypp.VNode {
	if signaling == nil || signaling.Error == nil {
		return nil
	}

	summary := "Invalid answer"
	if isOffer {
		summary = "Invalid offer"
	}
	return html.Details(
		hypp.HProps{"class": "error"},
		html.Summary(nil, hypp.Text(summary)),
		html.P(nil, hypp.Text(signaling.Error.Error())),
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
		html.P(nil, hypp.Text("Copy the offer below and send it to your opponent.")),
		html.P(
			hypp.HProps{"class": "warning"},
			hypp.Text("The offer contains your public IP. Only send the offer to someone you trust."),
		),
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
					"onclick": dispatch.GoToStartPage,
				},
				hypp.Text("Cancel"),
			),
			html.Button(
				hypp.HProps{
					"class":   "cta",
					"onclick": dispatch.SetSignalingStepNewGameAnswer,
				},
				hypp.Text("Next"),
			),
		),
	)
}

func signalingNewGameAnswer(s *state.State) *hypp.VNode {
	connectionState := ""
	readyState := ""
	if s.Signaling != nil {
		connectionState = s.Signaling.ConnectionState
		readyState = s.Signaling.ReadyState
	}
	var ctaButton *hypp.VNode
	if connectionState == "connecting" || connectionState == "connected" {
		buttonText := "Next"
		if connectionState != "connected" || readyState != "open" {
			buttonText = "Connecting..."
		}
		ctaButton = html.Button(
			hypp.HProps{
				"class": "cta",
				"onclick": hypp.ActionAndPayload[*state.State]{
					Action:  dispatch.GoToWhoGoesFirstPage,
					Payload: true,
				},
			},
			hypp.Text(buttonText),
		)
	} else {
		ctaButton = html.Button(
			hypp.HProps{
				"class":   "cta",
				"onclick": dispatch.ConnectNewGame,
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
		signalingError(s.Signaling, false),
		html.Textarea(
			hypp.HProps{
				"id":      "answer-textarea",
				"oninput": dispatch.SetSignalingAnswer,
			},
		),
		connectionStates(s),
		html.Div(
			nil,
			html.Button(
				hypp.HProps{
					"onclick": dispatch.GoToStartPage,
				},
				hypp.Text("Cancel"),
			),
			ctaButton,
		),
	)
}

func signalingJoinGameOffer(s *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.P(nil, hypp.Text("Paste the offer of your opponent below.")),
		signalingError(s.Signaling, true),
		html.Textarea(
			hypp.HProps{
				"id":      "offer-textarea",
				"oninput": dispatch.SetSignalingOffer,
			},
		),
		connectionStates(s),
		html.Div(
			nil,
			html.Button(
				hypp.HProps{
					"onclick": dispatch.GoToStartPage,
				},
				hypp.Text("Cancel"),
			),
			html.Button(
				hypp.HProps{
					"class":   "cta",
					"onclick": dispatch.SetSignalingStepJoinGameAnswer,
				},
				hypp.Text("Next"),
			),
		),
	)
}

func signalingJoinGameAnswer(s *state.State) *hypp.VNode {
	answer := "[error: Signaling is nil]"
	buttonText := "Next"
	if s.Signaling != nil {
		if s.Signaling.Loading {
			answer = "[Loading...]"
		} else if s.Signaling.Answer == "" {
			answer = "[error: Signaling.Answer is empty]"
		} else {
			answer = s.Signaling.Answer
		}
		if s.Signaling.ConnectionState != "connected" || s.Signaling.ReadyState != "open" {
			buttonText = "Connecting..."
		}
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.P(nil, hypp.Text("Copy the answer below and send it to your opponent.")),
		html.P(
			hypp.HProps{"class": "warning"},
			hypp.Text("The answer contains your public IP. Only send the answer to someone you trust."),
		),
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
					"onclick": dispatch.GoToStartPage,
				},
				hypp.Text("Cancel"),
			),
			html.Button(
				hypp.HProps{
					"class": "cta",
					"onclick": hypp.ActionAndPayload[*state.State]{
						Action:  dispatch.GoToWhoGoesFirstPage,
						Payload: false,
					},
				},
				hypp.Text(buttonText),
			),
		),
	)
}
