package state

type Status int

const (
	Created Status = iota
	Ready
	InProgress
	Finished
)

type Game struct {
	Board  Board
	Sticks Sticks
	You    int
	Turn   int
	Status Status
}
