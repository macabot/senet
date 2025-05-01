package state

type Player struct {
	Name          string
	SpeechBubble  *SpeechBubble
	DrawAttention bool
}

func (p Player) Clone() *Player {
	return &Player{
		Name:          p.Name,
		SpeechBubble:  p.SpeechBubble.Clone(),
		DrawAttention: p.DrawAttention,
	}
}
