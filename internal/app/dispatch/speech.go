package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func disableSpeechBubbleButton(s *state.State, player int) {
	s.Game.Players[player].SpeechBubble.ButtonDisabled = true
}

var OnSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialPlayers2: disableSpeechBubbleButton,
}

func SetSpeechBubbleKindAction(player int, kind state.SpeechBubbleKind) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		newState.Game.Players[player].SpeechBubble = &state.SpeechBubble{
			Kind: kind,
		}
		if onSet, ok := OnSetSpeechBubbleKind[kind]; ok {
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
