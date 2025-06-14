package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/dispatch"
)

func CancelToStartPageButton() *hypp.VNode {
	return atom.Button(
		"Cancel",
		dispatch.GoToStartScreen,
		nil,
	)
}
