package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func NewGame(s *state.State) *hypp.VNode {
	hasConnection := true
	roomName := ""
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - New Game")),
		html.Input(nil, hypp.Text(roomName)),
		atom.ConditionalNextButton(hasConnection, dispatch.GoToWhoGoesFirstPage),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.GoToStartPage,
			},
			hypp.Text("Cancel"),
		),
		// TODO create molecule.CancelButton?
	)
}
