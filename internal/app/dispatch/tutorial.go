package dispatch

import "github.com/macabot/senet/internal/app/state"

func init() {
	onThrowSticks = append(onThrowSticks, func(_, newState *state.State) {
		bubble := newState.Game.Players[1].SpeechBubble
		if bubble != nil && bubble.Kind == state.TutorialSticks3 {
			SetSpeechBubbleKind(newState, 1, state.TutorialMove1)
		}
	})
}
