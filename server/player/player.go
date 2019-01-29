package player

type Player struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Z    int    `json:"z"`
	Name string `json:"name"`
	HP   int    `json:"hp"`
	Breath int	`json:"breath"`
	//TODO: stamina, skin, etc
}
