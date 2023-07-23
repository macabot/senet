package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func disableSpeechBubbleButton(s *state.State, player int) {
	s.Game.Players[player].SpeechBubble.ButtonDisabled = true
}

var onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialPlayers2: disableSpeechBubbleButton,
}

func SetSpeechBubbleKindAction(player int, kind state.SpeechBubbleKind) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Game.Players[player].SpeechBubble = &state.SpeechBubble{
			Kind: kind,
		}
		if onSet, ok := onSetSpeechBubbleKind[kind]; ok {
			onSet(newState, player)
		}
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

func ToggleSpeechBubble(player int) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Game.Players[player].SpeechBubble == nil {
			newState.Game.Players[player].SpeechBubble = &state.SpeechBubble{
				Closed: true,
			}
		}
		newState.Game.Players[player].SpeechBubble.Closed = !newState.Game.Players[player].SpeechBubble.Closed
		return newState
	}
}
