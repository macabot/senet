package screen

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/screen"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func TaleWhoGoesFirst() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"WhoGoesFirst",
		&state.State{},
		screen.WhoGoesFirst,
	).WithControls(
		control.NewSelect(
			"Decision",
			func(s *state.State, option int) hypp.Dispatchable {
				switch option {
				case 0:
					s.CommitmentScheme = state.CommitmentScheme{}
				case 1:
					s.CommitmentScheme = state.CommitmentScheme{
						IsCaller:             true,
						HasCallerPredictions: true,
						CallerPredictions:    [4]bool{true},
						HasFlipperResults:    true,
						FlipperResults:       [4]bool{true},
					}
				case 2:
					s.CommitmentScheme = state.CommitmentScheme{
						IsCaller:             false,
						HasCallerPredictions: true,
						CallerPredictions:    [4]bool{true},
						HasFlipperResults:    true,
						FlipperResults:       [4]bool{true},
					}
				}
				return s
			},
			func(s *state.State) int {
				if !s.CommitmentScheme.CanThrow() {
					return 0
				}
				correctCall := s.CommitmentScheme.CallerPredictions[0] == s.CommitmentScheme.FlipperResults[0]
				if s.CommitmentScheme.IsCaller == correctCall {
					return 1
				}
				return 2
			},
			[]control.SelectOption[int]{
				{Label: "No decision", Value: 0},
				{Label: "You are player 1", Value: 1},
				{Label: "You are player 2", Value: 2},
			},
		),
		mycontrol.Disconnected(),
	)
}
