package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func NoOp(s *state.State, _ hypp.Payload) hypp.Dispatchable {
	return s
}
