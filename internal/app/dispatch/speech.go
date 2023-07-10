package dispatch

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func SpeechBubbleAction(action state.Action) hypp.Action[*state.State] {
	switch action.Kind {
	case state.SetSpeechBubble:
		return SetSpeechBubbleAction(action.Data.(state.SetSpeechBubbleData))
	case state.SetPage:
		return SetPageAction(action.Data.(state.Page))
	// case state.ShowBoardFlow:
	default:
		panic(fmt.Errorf("ActionKind %d is not implemented", action.Kind))
	}
}

func SetSpeechBubbleAction(data state.SetSpeechBubbleData) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		speechBubbleFunc, ok := state.SpeechBubbleFuncByName[data.SpeechBubbleName]
		if !ok {
			panic(fmt.Errorf("Could not find speech bubble func for name '%s'.", data.SpeechBubbleName))
		}
		speechBubble := speechBubbleFunc(data.Player)
		newState.Game.Players[data.Player].SpeechBubble = speechBubble
		return newState
	}
}

func SetPageAction(page state.Page) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Page = page
		return newState
	}
}
