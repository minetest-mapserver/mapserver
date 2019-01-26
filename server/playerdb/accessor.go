package playerdb

type Player struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Z    int    `json:"z"`
	Name string `json:"name"`
	HP   int    `json:"hp"`
	//TODO: stamina, skin, etc
}

type DBAccessor interface {
	GetActivePlayers() ([]*Player, error)
}
