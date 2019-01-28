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

func (this *Eventbus) Emit(eventtype string, o interface{}) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	for _, l := range this.listeners {
		l.OnEvent(eventtype, o)
	}
}

func (this *Eventbus) AddListener(l Listener) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.listeners = append(this.listeners, l)
}
