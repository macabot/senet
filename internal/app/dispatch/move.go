package dispatch

import (
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

type Move struct {
	From state.Position
	To   state.Position
}

var onMoveToSquare = []func(s, newState *state.State, from, to state.Position) []hypp.Effect{}

func MoveToSquare(s *state.State, payload hypp.Payload) hypp.Dispatchable {
	move := payload.(Move)
	newState := s.Clone()
	nextMove, err := newState.Game.Move(newState.Game.Turn, move.From, move.To)
	if err != nil {
		panic(err)
	}
	var effects []hypp.Effect
	if nextMove != nil {
		effects = append(effects, Delayed(
			hypp.ActionAndPayload[*state.State]{
				Action:  MoveToSquare,
				Payload: Move{From: nextMove.From, To: nextMove.To},
			},
			time.Second,
		))
	}
	for _, f := range onMoveToSquare {
		fEffects := f(s, newState, move.From, move.To)
		effects = append(effects, fEffects...)
	}
	return hypp.StateAndEffects[*state.State]{
		State:   newState,
		Effects: effects,
	}
}

var onNoMove = []func(s, newState *state.State) []hypp.Effect{}

func NoMove(s *state.State, _ hypp.Payload) hypp.Dispatchable {
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
