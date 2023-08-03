package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func StartTutorialAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
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
		return newState
	}
}
