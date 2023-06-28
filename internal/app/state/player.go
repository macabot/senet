package state

type Player struct {
	Name   string
	Speech string
}

func (p Player) Clone() *Player {
	return &Player{
		Name:   p.Name,
		Speech: p.Speech,
	}
}
