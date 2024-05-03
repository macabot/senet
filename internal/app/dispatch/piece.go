package dispatch

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func SelectPiece(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	id := payload.(int)
	newState := s.Clone()
	piece := newState.Game.Board.FindPieceByID(id)
	if piece == nil {
		panic(fmt.Errorf("could not find piece by id '%d'", id))
	}
	if newState.Game.Selected != nil && newState.Game.Selected.ID == piece.ID {
		newState.Game.SetSelected(nil)
	} else {
		newState.Game.SetSelected(piece)
	}
	return newState
}
