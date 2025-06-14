package molecule

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
)

func TaleMenuButton() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"MenuButton",
		nil,
		func(s *state.State) *hypp.VNode {
			return molecule.MenuButton()
		},
	)
}
