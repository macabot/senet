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
		newState.Game.TurnMode = state.IsPlayer0
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
		return newState
	}
}

func ToStartPageAction() hypp.Action[*state.State] {
	return func(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := &state.State{
			Page: state.StartPage,
		}
		resetListeners()
		resetSignaling(newState)
		return newState
	}
}

func ToSignalingPageAction() hypp.Action[*state.State] {
	return func(_ *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := &state.State{
			Page: state.SignalingPage,
		}
		resetSignaling(newState)
		initSignaling(newState)
		return newState
	}
}

func ToOnlinePlayerVsPlayerAction(isPlayer0 bool) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Page = state.GamePage
		newState.Game = state.NewGame()
		if isPlayer0 {
			newState.Game.TurnMode = state.IsPlayer0
			newState.Game.Players[0].Name = "You"
			newState.Game.Players[1].Name = "Opponent"
		} else {
			newState.Game.TurnMode = state.IsPlayer1
			newState.Game.Players[0].Name = "Opponent"
			newState.Game.Players[1].Name = "You"
		}
		newState.Game.Sticks.GeneratorKind = state.CommitmentSchemeGeneratorKind
		newState.CommitmentScheme = state.CommitmentScheme{
			IsCaller: isPlayer0,
		}
		if !isPlayer0 {
			return sendFlipperSecret(newState)
		}
		return newState
	}
}

func ToWhoGoesFirstPageAction(isCaller bool) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Page = state.WhoGoesFirstPage
		newState.CommitmentScheme.IsCaller = isCaller
		if !isCaller {
			return sendFlipperSecret(newState)
		}
		return newState
	}
}
