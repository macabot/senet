package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}

var onUnsetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}

func updateSpeechBubbleKind(s *state.State, player int, kind state.SpeechBubbleKind) {
	currentSpeechBubble := s.Game.Players[player].SpeechBubble
	currentClosed := false
	doNotRender := false
	if currentSpeechBubble != nil {
		currentClosed = currentSpeechBubble.Closed
		if onUnset, ok := onUnsetSpeechBubbleKind[currentSpeechBubble.Kind]; ok {
			onUnset(s, player)
		}
		doNotRender = currentSpeechBubble.Kind != kind && currentClosed
	}
	s.Game.Players[player].SpeechBubble = &state.SpeechBubble{
		Kind:        kind,
		Closed:      currentClosed,
		DoNotRender: doNotRender,
	}
	if s.Game.Players[player].SpeechBubble.Closed {
		s.Game.Players[player].DrawAttention = true
	}
	if onSet, ok := onSetSpeechBubbleKind[kind]; ok {
		onSet(s, player)
	}
}

type PlayerAndKind struct {
	Player int
	Kind   state.SpeechBubbleKind
}

func SetSpeechBubbleKind(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	playerAndKind := payload.(PlayerAndKind)
	newState := s.Clone()
	updateSpeechBubbleKind(newState, playerAndKind.Player, playerAndKind.Kind)
	return newState
}

var onToggleSpeechBubbleByKind = map[state.SpeechBubbleKind]func(s *state.State, player int){}

func ToggleSpeechBubble(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	player := payload.(int)
	newState := s.Clone()
	if newState.Game.Players[player].SpeechBubble == nil {
		newState.Game.Players[player].SpeechBubble = &state.SpeechBubble{
			Closed: true,
		}
	}
	newState.Game.Players[player].DrawAttention = false
	currentClosed := newState.Game.Players[player].SpeechBubble.Closed
	newState.Game.Players[player].SpeechBubble.Closed = !currentClosed
	newState.Game.Players[player].SpeechBubble.DoNotRender = false

	if onToggle, ok := onToggleSpeechBubbleByKind[newState.Game.Players[player].SpeechBubble.Kind]; ok {
		onToggle(newState, player)
	}

	return newState
}
