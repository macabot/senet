package dispatch

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func disableSpeechBubbleButton(s *state.State, player int) {
	s.Game.Players[player].SpeechBubble.ButtonDisabled = true
}

var onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialPlayers2: disableSpeechBubbleButton,
}

func setSpeechBubbleKind(s *state.State, player int, kind state.SpeechBubbleKind) {
	s.Game.Players[player].SpeechBubble = &state.SpeechBubble{
		Kind: kind,
	}
	if onSet, ok := onSetSpeechBubbleKind[kind]; ok {
		onSet(s, player)
	}
}

func SetSpeechBubbleKindAction(player int, kind state.SpeechBubbleKind) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		if event, ok := payload.(hypp.Event); ok {
			event.StopPropagation()
		} else {
			// FIXME this happens when changing the speech bubble kind using the tale control.
			fmt.Println("UNEXPECTED payload type")
		}

		newState := s.Clone()
		setSpeechBubbleKind(newState, player, kind)
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
			setSpeechBubbleKind(s, player, state.TutorialGoal)
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
