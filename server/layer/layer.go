package layer

type Layer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	From int    `json:"from"`
	To   int    `json:"to"`
}

func FindLayerById(layers []*Layer, id int) *Layer {
	for _, l := range layers {
		if l.Id == id {
			return l
		}
	}
	return nil
}

func FindLayerByY(layers []*Layer, y int) *Layer {
	for _, l := range layers {
		if y >= l.From && y <= l.To {
			return l
		}
	}
	return nil
}
