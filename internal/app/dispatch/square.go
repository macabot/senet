package dispatch

import (
	"fmt"
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func MoveToSquareAction(toPosition state.Position) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		validMoves := newState.Game.ValidMoves
		fromPositionFound := false
		var fromPosition state.Position
		for from, to := range validMoves {
			if to == toPosition {
				fromPosition = from
				fromPositionFound = true
			}
		}
		if !fromPositionFound {
			panic(fmt.Errorf("MoveToSquare failed: could not find 'from' position corresponding to 'to' position '%d'", toPosition))
		}
		nextMove, err := newState.Game.Move(newState.Game.Turn, fromPosition, toPosition)
		if err != nil {
			panic(err)
		}
		var effects []hypp.Effect
		if nextMove != nil {
			effects = append(effects, DelayedAction(
				MoveToSquareAction(nextMove.To),
				time.Second,
			))
		}
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: effects,
		}
	}
}
