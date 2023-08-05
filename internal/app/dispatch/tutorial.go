package dispatch

import "github.com/macabot/senet/internal/app/state"

func setBubbleOnMove(current, next state.SpeechBubbleKind) func(s, newState *state.State) {
	return func(_, newState *state.State) {
		bubble := newState.Game.Players[1].SpeechBubble
		if bubble != nil && bubble.Kind == current {
			SetSpeechBubbleKind(newState, 1, next)
		}
	}
}

func init() {
	onThrowSticks = append(onThrowSticks, func(_, newState *state.State) {
		bubble := newState.Game.Players[1].SpeechBubble
		if bubble != nil && bubble.Kind == state.TutorialSticks3 {
			SetSpeechBubbleKind(newState, 1, state.TutorialMove)
		}
	})
	onMoveToSquare = append(
		onMoveToSquare,
		setBubbleOnMove(state.TutorialMove, state.TutorialMultiplemoves),
		setBubbleOnMove(state.TutorialTradingPlaces4, state.TutorialBlockingPiece1),
		setBubbleOnMove(state.TutorialBlockingPiece2, state.TutorialReturnToStart),
	)
}
