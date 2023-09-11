package state

type SpeechBubbleKind int

const (
	DefaultSpeechBubble SpeechBubbleKind = iota
	TutorialStart
	TutorialPlayers1
	TutorialPlayers2
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
	TutorialBlockingPiece1
	TutorialBlockingPiece2
	TutorialReturnToStart1
	TutorialReturnToStart2
	TutorialReturnToStart3
	TutorialMoveBackwards1
	TutorialMoveBackwards2
	TutorialNoMove1
	TutorialNoMove2
	TutorialOffTheBoard1
	TutorialOffTheBoard2
	TutorialOffTheBoard3
	TutorialEnd
)

type SpeechBubble struct {
	Kind        SpeechBubbleKind
	Closed      bool
	DoNotRender bool
}

func (b *SpeechBubble) Clone() *SpeechBubble {
	if b == nil {
		return nil
	}
	return &SpeechBubble{
		Kind:        b.Kind,
		Closed:      b.Closed,
		DoNotRender: b.DoNotRender,
	}
}
