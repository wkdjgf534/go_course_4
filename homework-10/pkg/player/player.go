package player

type Player struct {
	Name  string
	Point int
}

func New(name string) *Player {
	return &Player{
		Name:  name,
		Point: 0,
	}
}
