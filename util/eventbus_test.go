package util

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testEvent struct {
	eventtype string
	payload   interface{}
}

type testListener struct {
	mu     sync.Mutex
	events []testEvent
}

func (l *testListener) OnEvent(eventtype string, o interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.events = append(l.events, testEvent{eventtype, o})
}

func TestEventbusEmitWithoutListeners(t *testing.T) {
	assert.NotPanics(t, func() { NewEventbus().Emit("x", nil) })
}

func TestEventbusDeliversToAllListeners(t *testing.T) {
	eb := NewEventbus()
	l1, l2 := &testListener{}, &testListener{}
	eb.AddListener(l1)
	eb.AddListener(l2)

	eb.Emit(MAPBLOCK_RENDERED, "a")
	eb.Emit(TILE_RENDERED, 2)

	expected := []testEvent{{MAPBLOCK_RENDERED, "a"}, {TILE_RENDERED, 2}}
	assert.Equal(t, expected, l1.events)
	assert.Equal(t, expected, l2.events)
}

func TestEventbusConcurrent(t *testing.T) {
	eb := NewEventbus()
	l := &testListener{}
	eb.AddListener(l)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(2)
		go func() { defer wg.Done(); eb.Emit("e", nil) }()
		go func() { defer wg.Done(); eb.AddListener(&testListener{}) }()
	}
	wg.Wait()

	assert.Len(t, l.events, 20)
}
