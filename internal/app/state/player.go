package state

type Player struct {
	Name         string
	SpeechBubble *SpeechBubble
}

func (p Player) Clone() *Player {
	return &Player{
		Name:         p.Name,
		SpeechBubble: p.SpeechBubble.Clone(),
	}
}
