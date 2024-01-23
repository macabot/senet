package dispatch

import (
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

var onMoveToSquare = []func(s, newState *state.State, from, to state.Position) []hypp.Effect{}

func MoveToSquareAction(fromPosition, toPosition state.Position) hypp.Action[*state.State] {
	return func(s *state.State, _ hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		nextMove, err := newState.Game.Move(newState.Game.Turn, fromPosition, toPosition)
		if err != nil {
			panic(err)
		}
		var effects []hypp.Effect
		if nextMove != nil {
			effects = append(effects, DelayedAction(
				MoveToSquareAction(nextMove.From, nextMove.To),
				time.Second,
			))
		}
		for _, f := range onMoveToSquare {
			fEffects := f(s, newState, fromPosition, toPosition)
			effects = append(effects, fEffects...)
		}
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: effects,
		}
	}
}

var onNoMove = []func(s, newState *state.State) []hypp.Effect{}

func NoMoveAction() hypp.Action[*state.State] {
	return func(s *state.State, payload hypp.Payload) hypp.Dispatchable {
		newState := s.Clone()
		if err := newState.Game.NoMove(newState.Game.Turn); err != nil {
			panic(err)
		}
		var effects []hypp.Effect
		for _, f := range onNoMove {
			fEffects := f(s, newState)
			effects = append(effects, fEffects...)
		}
		return hypp.StateAndEffects[*state.State]{
			State:   newState,
			Effects: effects,
		}
	}
}
