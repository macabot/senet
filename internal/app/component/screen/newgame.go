package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func NewGame(s *state.State) *hypp.VNode {
	hasConnection := true
	roomName := ""
	if s.Signaling != nil {
		roomName = s.Signaling.RoomName
	}
	var status *hypp.VNode
	var onClickNext hypp.Dispatchable
	if hasConnection {
		status = html.P(nil, hypp.Text("Connected"))
		onClickNext = dispatch.GoToWhoGoesFirstPage
	} else {
		status = html.P(nil, hypp.Text("Waiting..."))
	}
	return html.Main(
		hypp.HProps{
			"class": "online",
		},
		html.H1(nil, hypp.Text("Online - New Game")),
		html.Input(hypp.HProps{"readonly": true}, hypp.Text(roomName)),
		status,
		html.Div(
			nil,
			molecule.CancelToStartPageButton(),
			atom.Button("Next", onClickNext, hypp.HProps{"class": "cta"}),
		),
	)
}
