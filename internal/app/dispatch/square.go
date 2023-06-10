package dispatch

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func MoveToSquareAction(position state.Position) hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		validMoves := newState.Game.ValidMoves
		fromPositionFound := false
		var fromPosition state.Position
		for from, to := range validMoves {
			if to == position {
				fromPosition = from
				fromPositionFound = true
			}
		}
		if !fromPositionFound {
			panic(fmt.Errorf("MoveToSquare failed: could not find 'from' position corresponding to 'to' position '%d'", position))
		}
		newState.Game.Move(newState.Game.Turn, fromPosition)

		return newState
	}
}
