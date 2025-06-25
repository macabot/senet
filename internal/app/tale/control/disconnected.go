package control

// func Disconnected() *control.Select[*state.State, string] {
// 	return control.NewSelect(
// 		"Disconnected",
// 		func(s *state.State, option string) hypp.Dispatchable {
// 			switch option {
// 			case "disconnected":
// 				s.Signaling = &state.Signaling{
// 					ConnectionState: "disconnected",
// 				}
// 			case "failed":
// 				s.Signaling = &state.Signaling{
// 					ConnectionState: "failed",
// 				}
// 			default: // no
// 				s.Signaling = nil
// 			}
// 			return s
// 		},
// 		func(s *state.State) string {
// 			if s.Signaling == nil {
// 				return "no"
// 			}
// 			return s.Signaling.ConnectionState
// 		},
// 		[]control.SelectOption[string]{
// 			{Label: "No", Value: "no"},
// 			{Label: "Disconnected", Value: "disconnected"},
// 			{Label: "Failed", Value: "failed"},
// 		},
// 	)
// }
