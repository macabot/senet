package app

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Run(driver hypp.Driver, node hypp.Node) {
	hypp.App(hypp.AppProps[*state.State]{
		Driver: driver,
		Init: &state.State{
			Page: state.StartPage,
		},
		View: component.Senet,
		Node: node,
	})
}
