package dispatch

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
)

func SelectTextareaEffect(id string) hypp.Effect {
	return hypp.Effect{
		Effecter: func(_ hypp.Dispatch, _ hypp.Payload) {
			window.RequestAnimationFrame(func() {
				window.Document().GetElementById(id).Value.Call("select")
			})
		},
	}
}
