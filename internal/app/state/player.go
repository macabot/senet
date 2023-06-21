package state

type Player struct {
	Name string
}

func (p Player) Clone() *Player {
	return &Player{
		Name: p.Name,
	}
}
