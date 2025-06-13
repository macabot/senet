package molecule

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
)

func TaleRoomNameField() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"RoomNameField",
		&state.State{Signaling: &state.Signaling{RoomName: "R2D2"}},
		func(s *state.State) *hypp.VNode {
			return html.Div(nil, molecule.RoomNameField(s.Signaling.RoomName, false)...)
		},
	).WithControls(
		control.NewTextInput(
			"Room name",
			func(s *state.State, roomName string) hypp.Dispatchable {
				newState := s.Clone()
				newState.Signaling.RoomName = roomName
				return newState
			},
			func(s *state.State) string {
				return s.Signaling.RoomName
			},
		),
	)
}
