package layer

type Layer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	To   int    `json:"to"`
	From int    `json:"from"`
}
