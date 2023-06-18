package state

type Player struct {
	Points int
}

func (p Player) Clone() *Player {
	return &Player{
		Points: p.Points,
	}
}
