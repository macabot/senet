package state

type SpeechBubbleKind int

const (
	TutorialStart SpeechBubbleKind = iota + 1
	TutorialPlayers1
	TutorialPlayers2
	TutorialGoal
	TutorialBoard
	TutorialEnd
)

type SpeechBubble struct {
	Kind           SpeechBubbleKind
	ButtonDisabled bool
}

func (b *SpeechBubble) Clone() *SpeechBubble {
	if b == nil {
		return nil
	}
	return &SpeechBubble{
		Kind:           b.Kind,
		ButtonDisabled: b.ButtonDisabled,
	}
}
