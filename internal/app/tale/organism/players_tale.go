package organism

import (
	"fmt"

	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/organism"
	"github.com/macabot/senet/internal/app/dispatch"
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

var bubbles = mycontrol.LabeledSlice[state.SpeechBubbleKind]{
	{
		Label: "No speech bubble",
	},
	{
		Label: "TutorialStart",
		V:     state.TutorialStart,
	},
	{
		Label: "TutorialPlayers1",
		V:     state.TutorialPlayers1,
	},
	{
		Label: "TutorialPlayers2",
		V:     state.TutorialPlayers2,
	},
	{
		Label: "TutorialBoard",
		V:     state.TutorialBoard1,
	},
	{
		Label: "TutorialEnd",
		V:     state.TutorialEnd,
	},
}

func speechBubbleKind(player int) *control.Select[*state.State, int] {
	return control.NewSelect(
		fmt.Sprintf("Player %d speech bubble", player+1),
		func(s *state.State, option int) hypp.Dispatchable {
			if bubbles[option].V == 0 {
				s.Game.Players[player].SpeechBubble = nil
				return s
			} else {
				kind := bubbles[option].V
				return hypp.ActionAndPayload[*state.State]{
					Action: dispatch.SetSpeechBubbleKind,
					Payload: dispatch.PlayerAndKind{
						Player: player,
						Kind:   kind,
					},
				}
			}
		},
		func(s *state.State) int {
			current := s.Game.Players[player].SpeechBubble
			for i, bubble := range bubbles {
				if current == nil {
					if bubble.V == 0 {
						return i
					}
				} else {
					if bubble.V == current.Kind {
						return i
					}
				}
			}
			return -1
		},
		bubbles.SelectOptions(),
	)
}

func drawAttention(player int) *control.Checkbox[*state.State] {
	return control.NewCheckbox(
		fmt.Sprintf("Player %d draw attention", player+1),
		func(s *state.State, drawAttention bool) hypp.Dispatchable {
			s.Game.Players[player].DrawAttention = drawAttention
			return s
		},
		func(s *state.State) bool {
			return s.Game.Players[player].DrawAttention
		},
	)
}

func TalePlayers() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Players",
		&state.State{
			Game: state.NewGame(),
		},
		func(s *state.State) *hypp.VNode {
			return organism.Players(organism.CreatePlayersProps(s))
		},
	).WithControls(
		mycontrol.PlayerTurn(),
		playerPoints(0),
		playerPoints(1),
		speechBubbleKind(0),
		speechBubbleKind(1),
		drawAttention(0),
		drawAttention(1),
	)
}
