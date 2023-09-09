package dispatch

import "github.com/macabot/senet/internal/app/state"

func replaceCurrentBubbleWithNext(current, next state.SpeechBubbleKind) func(s, newState *state.State) {
	return func(_, newState *state.State) {
		bubble := newState.Game.Players[1].SpeechBubble
		if bubble != nil && bubble.Kind == current {
			SetSpeechBubbleKind(newState, 1, next)
		}
	}
}

func init() {
	// TutorialPlayers2
	onSetSpeechBubbleKind[state.TutorialPlayers2] = func(s *state.State, player int) {
		s.Game.Players[player].DrawAttention = true
	}
	onToggleSpeechBubbleByKind[state.TutorialPlayers2] = func(s *state.State, player int) {
		if !s.Game.Players[player].SpeechBubble.Closed {
			s.Game.Players[player].DrawAttention = false
			SetSpeechBubbleKind(s, player, state.TutorialGoal)
		}
	}
	// TutorialBoard3
	onSetSpeechBubbleKind[state.TutorialBoard3] = func(s *state.State, _ int) {
		s.Game.Board.ShowDirections = true
	}
	onUnsetSpeechBubbleKind[state.TutorialBoard3] = func(s *state.State, player int) {
		s.Game.Board.ShowDirections = false
	}
	// TutorialSticks3
	onSetSpeechBubbleKind[state.TutorialSticks3] = func(s *state.State, _ int) {
		s.Game.Sticks.HasThrown = false
		s.Game.Turn = 0
	}
	onThrowSticks = append(
		onThrowSticks,
		replaceCurrentBubbleWithNext(state.TutorialSticks3, state.TutorialMove),
	)
	// TutorialMove
	onMoveToSquare = append(onMoveToSquare, replaceCurrentBubbleWithNext(state.TutorialMove, state.TutorialMultiplemoves))
	// TutorialTradingPlaces2
	onSetSpeechBubbleKind[state.TutorialTradingPlaces2] = func(s *state.State, _ int) {
		s.Game.SetBoard(&state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 13},
					&state.Piece{ID: 2, Position: 33},
					&state.Piece{ID: 3, Position: 32},
					&state.Piece{ID: 4, Position: 6},
					&state.Piece{ID: 5, Position: 31},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 14},
					&state.Piece{ID: 7, Position: 7},
					&state.Piece{ID: 8, Position: 28},
					&state.Piece{ID: 9, Position: 5},
					&state.Piece{ID: 10, Position: 30},
				),
			},
		})
	}
	// TutorialTradingPlaces4
	onSetSpeechBubbleKind[state.TutorialTradingPlaces4] = func(s *state.State, _ int) {
		s.Game.Sticks.HasThrown = false
		s.Game.Turn = 0
	}
	onMoveToSquare = append(onMoveToSquare, replaceCurrentBubbleWithNext(state.TutorialTradingPlaces4, state.TutorialBlockingPiece1))
	// TutorialBlockingPiece1
	onSetSpeechBubbleKind[state.TutorialBlockingPiece1] = func(s *state.State, _ int) {
		b := false
		s.Game.OverwriteHasTurn = &b
	}
	onUnsetSpeechBubbleKind[state.TutorialBlockingPiece1] = func(s *state.State, _ int) {
		s.Game.OverwriteHasTurn = nil
	}
	// TutorialBlockingPiece2
	onMoveToSquare = append(
		onMoveToSquare,
		replaceCurrentBubbleWithNext(state.TutorialBlockingPiece2, state.TutorialReturnToStart1),
	)
	// TutorialReturnToStart2
	onSetSpeechBubbleKind[state.TutorialReturnToStart2] = func(s *state.State, _ int) {
		// _ p p x p _ _ B _ _
		// _ _ B B R R _ _ _ _
		// _ _ _ _ R B _ R B R
		s.Game.SetBoard(&state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 22},
					&state.Piece{ID: 2, Position: 12},
					&state.Piece{ID: 3, Position: 13},
					&state.Piece{ID: 4, Position: 4},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 15},
					&state.Piece{ID: 7, Position: 14},
					&state.Piece{ID: 8, Position: 5},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		})
	}
	// TutorialReturnToStart3
	onSetSpeechBubbleKind[state.TutorialReturnToStart3] = func(s *state.State, player int) {
		s.Game.Sticks.HasThrown = false
		s.Game.Turn = 0
		s.Game.Players[player].DrawAttention = true
	}
	onToggleSpeechBubbleByKind[state.TutorialReturnToStart3] = func(s *state.State, player int) {
		s.Game.Players[player].DrawAttention = false
	}
	onMoveToSquare = append(
		onMoveToSquare,
		replaceCurrentBubbleWithNext(state.TutorialReturnToStart3, state.TutorialMoveBackwards1),
	)
	// TutorialMoveBackwards1
	onSetSpeechBubbleKind[state.TutorialMoveBackwards1] = func(s *state.State, player int) {
		b := false
		s.Game.OverwriteHasTurn = &b
	}
	onUnsetSpeechBubbleKind[state.TutorialMoveBackwards1] = func(s *state.State, _ int) {
		s.Game.OverwriteHasTurn = nil
	}
	// TutorialMoveBackwards2
	onSetSpeechBubbleKind[state.TutorialMoveBackwards2] = func(s *state.State, player int) {
		s.Game.Players[player].DrawAttention = true
	}
	onToggleSpeechBubbleByKind[state.TutorialMoveBackwards2] = func(s *state.State, player int) {
		s.Game.Players[player].DrawAttention = false
	}
}
