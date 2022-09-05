package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/model"
	"github.com/macabot/senet/internal/app/view/render/component"
)

func BoardTale() *fairy.Tale {
	return fairy.NewTale(
		"Board",
		model.NewBoard(),
		component.Board,
	)
}
