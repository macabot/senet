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
		resetListeners()
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
		resetListeners()
		return newState
	}
}

func ToStartPageAction() hypp.Action[*state.State] {
	return func(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
		resetListeners()
		return &state.State{
			Page: state.StartPage,
		}
	}
}

func ToSignalingPageAction() hypp.Action[*state.State] {
	return func(_ *state.State, payload hypp.Payload) hypp.Dispatchable {
		resetListeners()
		return &state.State{
			Page: state.SignalingPage,
		}
	}
}
