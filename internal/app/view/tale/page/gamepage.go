package page

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/page"
	"github.com/macabot/senet/internal/app/view/state"
)

func GamePageTale() *fairy.Tale {
	return fairy.NewTale(
		"GamePage",
		&state.State{
			Game: state.NewGame(),
		},
		page.GamePage,
	)
}
