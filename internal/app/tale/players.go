package tale

import (
	"fmt"

	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func playerPoints(player int) *control.Select[*state.State, int] {
	return control.NewSelect(
		fmt.Sprintf("Player %d points", player+1),
		func(s *state.State, points int) hypp.Dispatchable {
			pieces := make([]*state.Piece, points)
			for i := 0; i < points; i++ {
				pieces[i] = &state.Piece{Position: state.Position(30 + i)}
			}
			s.Game.Board.PlayerPieces[player] = state.NewPiecesByPosition(pieces...)
			return s
		},
		func(s *state.State) int {
			return s.Game.Board.Points(player)
		},
		[]control.SelectOption[int]{
			{Label: "0", Value: 0},
			{Label: "1", Value: 1},
			{Label: "2", Value: 2},
			{Label: "3", Value: 3},
			{Label: "4", Value: 4},
			{Label: "5", Value: 5},
		},
	)
}

type SpeechBubbles []*state.SpeechBubble

func (b SpeechBubbles) SelectOptions() []control.SelectOption[int] {
	options := make([]control.SelectOption[int], len(b))
	for i, bubble := range b {
		var label string
		if bubble == nil {
			label = "No speech bubble"
		} else {
			label = bubble.Name
		}
		options[i] = control.SelectOption[int]{
			Label: label,
			Value: i,
		}
	}
	return options
}

var speechBubbles = SpeechBubbles{
	nil,
	state.TutorialStart,
	state.TutorialPlayers1,
	state.TutorialPlayers2,
	state.TutorialGoal,
	state.TutorialBoard,
	state.TutorialEnd,
}

func speechBubble(player int) *control.Select[*state.State, int] {
	return control.NewSelect(
		fmt.Sprintf("Player %d speech bubble", player+1),
		func(s *state.State, option int) hypp.Dispatchable {
			s.Game.Players[player].SpeechBubble = speechBubbles[option]
			return s
		},
		func(s *state.State) int {
			current := s.Game.Players[player].SpeechBubble
			if current == nil {
				return 0
			}
			for i, bubble := range speechBubbles {
				if bubble == nil && current == nil {
					return i
				}
				if bubble != nil && current != nil && bubble.Name == current.Name {
					return i
				}
			}
			return -1
		},
		speechBubbles.SelectOptions(),
	)
}

func Players() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Players",
		&state.State{
			Game: state.NewGame(),
		},
		func(s *state.State) *hypp.VNode {
			return component.Players(component.CreatePlayersProps(s))
		},
	).WithControls(
		mycontrol.PlayerTurn(),
		playerPoints(0),
		playerPoints(1),
		speechBubble(0),
		speechBubble(1),
	)
}
