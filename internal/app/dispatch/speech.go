package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var onSetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialPlayers2: func(s *state.State, player int) {
		s.Game.Players[player].DrawAttention = true
	},
	state.TutorialBoard3: func(s *state.State, _ int) {
		s.Game.Board.ShowDirections = true
	},
	state.TutorialSticks3: func(s *state.State, _ int) {
		s.Game.Sticks.HasThrown = false
		s.Game.Turn = 0
	},
	state.TutorialTradingPlaces2: func(s *state.State, _ int) {
		// TODO continue
		s.Game.SetBoard(&state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 7},
					&state.Piece{ID: 3, Position: 5},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 6},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		})
	},
}

var onUnsetSpeechBubbleKind = map[state.SpeechBubbleKind]func(s *state.State, player int){
	state.TutorialBoard3: func(s *state.State, player int) {
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
			s.Game.Players[player].DrawAttention = false
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
