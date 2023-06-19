package state

type Player struct {
	Name   string
	Points int
}

func (p Player) Clone() *Player {
	return &Player{
		Name:   p.Name,
		Points: p.Points,
	}
}
