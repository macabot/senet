package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}

var onUnsetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}

func SetSpeechBubbleKind(s *state.State, player int, kind state.SpeechBubbleKind) {
	currentSpeechBubble := s.Game.Players[player].SpeechBubble
	currentClosed := false
	if currentSpeechBubble != nil {
		currentClosed = currentSpeechBubble.Closed
		if onUnSet, ok := onUnsetSpeechBubbleKind[currentSpeechBubble.Kind]; ok {
			onUnSet(s, player)
		}
	}
	s.Game.Players[player].SpeechBubble = &state.SpeechBubble{
		Kind:   kind,
		Closed: currentClosed,
	}
	if s.Game.Players[player].SpeechBubble.Closed {
		s.Game.Players[player].DrawAttention = true
	}
	if onSet, ok := onSetSpeechBubbleKind[kind]; ok {
		onSet(s, player)
	}
}

func SetSpeechBubbleKindAction(player int, kind state.SpeechBubbleKind) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		SetSpeechBubbleKind(newState, player, kind)
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

var onToggleSpeechBubbleByKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}

func ToggleSpeechBubble(player int) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if newState.Game.Players[player].SpeechBubble == nil {
			newState.Game.Players[player].SpeechBubble = &state.SpeechBubble{
				Closed: true,
			}
		}
		currentClosed := newState.Game.Players[player].SpeechBubble.Closed
		if currentClosed {
			newState.Game.Players[player].DrawAttention = false
		}
		newState.Game.Players[player].SpeechBubble.Closed = !currentClosed

		if onToggle, ok := onToggleSpeechBubbleByKind[newState.Game.Players[player].SpeechBubble.Kind]; ok {
			onToggle(newState, player)
		}

		return newState
	}
}
