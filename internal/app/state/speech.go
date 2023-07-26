package state

type SpeechBubbleKind int

const (
	DefaultSpeechBubble SpeechBubbleKind = iota
	TutorialStart
	TutorialPlayers1
	TutorialPlayers2
	TutorialGoal
	TutorialBoard
	TutorialPieces
	TutorialSticks1
	TutorialSticks2
	TutorialEnd
)

type SpeechBubble struct {
	Kind           SpeechBubbleKind
	Closed         bool
	ButtonDisabled bool
}

func (b *SpeechBubble) Clone() *SpeechBubble {
	if b == nil {
		return nil
	}
	return &SpeechBubble{
		Kind:           b.Kind,
		Closed:         b.Closed,
		ButtonDisabled: b.ButtonDisabled,
	}
}
