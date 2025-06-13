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
	var status string
	var onClickNext hypp.Dispatchable
	if hasConnection {
		status = "Connected"
		onClickNext = dispatch.GoToWhoGoesFirstScreen
	} else {
		status = "Waiting..."
	}

	statusNode := html.P(hypp.HProps{"class": "status"}, hypp.Text(status))

	children := []*hypp.VNode{
		html.H1(nil, hypp.Text("Online - New Game")),
	}
	children = append(children, molecule.RoomNameField(roomName, false)...)
	children = append(
		children,
		html.P(nil, hypp.Text("Share the room name with your opponent.")),
		statusNode,
		html.Div(
			nil,
			molecule.CancelToStartPageButton(),
			atom.Button("Next", onClickNext, hypp.HProps{"class": "cta"}),
		),
	)
	return html.Main(
		hypp.HProps{
			"class": "screen",
		},
		children...,
	)
}
