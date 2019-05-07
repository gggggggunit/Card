package handlers



type Player struct {
	ID   string // 123123123dd-23423sfdfsf-23423e-2323
	Name string // finalist
	Cards [] string
}

type Room struct {
	ID      string // 123-ewdsd-1123123
	Game    string // pocker, black_jack...
	Players []*Player
}

type Roomss struct {
	Roooms []*Room
}