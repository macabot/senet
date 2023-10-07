package app

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Run(element window.Element) {
	hypp.App(hypp.AppProps[*state.State]{
		Init: &state.State{
			Page: state.StartPage,
		},
		View: component.Senet,
		Node: element,
	})
}
