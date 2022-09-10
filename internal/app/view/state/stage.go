package state

import (
	"github.com/macabot/senet/internal/app/model"
	"github.com/macabot/senet/internal/pkg/set"
)

type Meta struct {
	Selected    *model.Position
	Highlighted set.Set[model.Position]
}
