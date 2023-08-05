package state

type SpeechBubbleKind int

const (
	DefaultSpeechBubble SpeechBubbleKind = iota
	TutorialStart
	TutorialPlayers1
	TutorialPlayers2
	TutorialGoal
	TutorialBoard1
	TutorialBoard2
	TutorialBoard3
	TutorialSticks1
	TutorialSticks2
	TutorialSticks3
	TutorialMove
	TutorialMultiplemoves
	TutorialTradingPlaces1
	TutorialTradingPlaces2
	TutorialTradingPlaces3
	TutorialTradingPlaces4
	TutorialBlockingPiece
	TutorialEnd
)

type SpeechBubble struct {
	Kind   SpeechBubbleKind
	Closed bool
}

func (b *SpeechBubble) Clone() *SpeechBubble {
	if b == nil {
		return nil
	}
	return &SpeechBubble{
		Kind:   b.Kind,
		Closed: b.Closed,
	}
}
