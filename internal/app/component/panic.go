package component

import (
	"fmt"
	"runtime/debug"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/state"
)

func RecoverPanic(component func(s *state.State) *hypp.VNode, s *state.State) (vNode *hypp.VNode) {
	defer func() {
		if r := recover(); r != nil {
			panicStackTrace := fmt.Sprintf("%v\n%s", r, string(debug.Stack()))
			window.Console().Error(panicStackTrace)
			s.PanicStackTrace = &panicStackTrace
			vNode = Senet(s)
		}
	}()

	return component(s)
}
