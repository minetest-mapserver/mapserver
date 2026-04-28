package eventbus

import (
	"sync"
)

const (
	//TODO: move to separate package
	MAPBLOCK_RENDERED = "mapblock-rendered"
	TILE_RENDERED     = "rendered-tile"
)

type Listener interface {
	OnEvent(eventtype string, o interface{})
}

type Eventbus struct {
	mutex     *sync.RWMutex
	listeners []Listener
}

func New() *Eventbus {
	eb := Eventbus{}
	eb.mutex = &sync.RWMutex{}
	eb.listeners = make([]Listener, 0)

	return &eb
}

func (b *Eventbus) Emit(eventtype string, o interface{}) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	for _, l := range b.listeners {
		l.OnEvent(eventtype, o)
	}
}

func (b *Eventbus) AddListener(l Listener) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.listeners = append(b.listeners, l)
}
