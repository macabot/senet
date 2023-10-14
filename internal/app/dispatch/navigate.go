package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func resetListeners() {
	onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}
	onUnsetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}
	onToggleSpeechBubbleByKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}
	onMoveToSquare = []func(s, newState *state.State){}
	onNoMove = []func(s, newState *state.State){}
	onThrowSticks = []func(s, newState *state.State){}
}

func resetForNavigation() {
	resetListeners()
	resetSignaling()
}

func ToTutorialAction() hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Page = state.GamePage
		newState.Game = state.NewGame()
		newState.Game.Players[0].Name = "You"
		newState.Game.Players[1].Name = "Tutor"
		newState.Game.Turn = 1
		newState.Game.TurnMode = state.IsPlayer1
		newState.Game.Players[1].SpeechBubble = &state.SpeechBubble{
			Kind: state.TutorialStart,
		}
		newState.Game.Sticks.GeneratorKind = state.TutorialSticksGeneratorKind
		resetForNavigation()
		registerTutorial()
		return newState
	}
}

func ToLocalPlayerVsPlayerAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Page = state.GamePage
		newState.Game = state.NewGame()
		newState.Game.TurnMode = state.IsBothPlayers
		resetForNavigation()
		return newState
	}
}

func toPageAction(page state.Page) hypp.Action[*state.State] {
	return func(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
		resetForNavigation()
		return &state.State{
			Page: page,
		}
	}
}

func ToStartPageAction() hypp.Action[*state.State] {
	return toPageAction(state.StartPage)
}

func ToSignalingPageAction() hypp.Action[*state.State] {
	return func(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
		resetForNavigation()
		initSignaling()
		return &state.State{
			Page: state.SignalingPage,
		}
	}
}

func ToSignalingNewGamePageAction() hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Page = state.SignalingNewGamePage
		if newState.Signaling == nil {
			newState.Signaling = &state.Signaling{}
		}
		newState.Signaling.Loading = true
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: []hypp.Effect{CreatePeerConnectionOfferEffect()},
		}
	}
}
