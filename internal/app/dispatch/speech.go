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
	state.TutorialBoard: func(s *state.State, _ int) {
		s.Game.Board.ShowDirections = true
	},
	state.TutorialSticks3: func(s *state.State, _ int) {
		s.Game.Sticks.HasThrown = false
		s.Game.Turn = 0
		s.Game.HasTurn = true
	},
}

var onUnsetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialBoard: func(s *state.State, player int) {
		s.Game.Board.ShowDirections = false
	},
}

func SetSpeechBubbleKind(s *state.State, player int, kind state.SpeechBubbleKind) {
	if s.Game.Players[player].SpeechBubble != nil {
		if onUnSet, ok := onUnsetSpeechBubbleKind[s.Game.Players[player].SpeechBubble.Kind]; ok {
			onUnSet(s, player)
		}
	}
	s.Game.Players[player].SpeechBubble = &state.SpeechBubble{
		Kind: kind,
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

var onToggleSpeechBubbleByKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialPlayers2: func(s *state.State, player int) {
		if !s.Game.Players[player].SpeechBubble.Closed {
			SetSpeechBubbleKind(s, player, state.TutorialGoal)
		}
	},
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

		if onToggle, ok := onToggleSpeechBubbleByKind[newState.Game.Players[player].SpeechBubble.Kind]; ok {
			onToggle(newState, player)
		}

		return newState
	}
}
