package web

import (
	"bytes"
	"encoding/json"
	"mapserver/app"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WS struct {
	ctx      *app.App
	channels map[int]chan []byte
	mutex    *sync.RWMutex
	clients  int
	nextid   int
}

func NewWS(ctx *app.App) *WS {
	ws := WS{}
	ws.mutex = &sync.RWMutex{}
	ws.channels = make(map[int]chan []byte)

	return &ws
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (t *WS) OnEvent(eventtype string, o interface{}) {
	data, err := json.Marshal(o)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "type": eventtype}).Error("ws-marshal")
		return
	}

	buf := new(bytes.Buffer)
	buf.Write([]byte("{\"type\":\""))
	buf.Write([]byte(eventtype))
	buf.Write([]byte("\",\"data\":"))
	buf.Write(data)
	buf.Write([]byte("}"))

	t.mutex.RLock()
	defer t.mutex.RUnlock()

	for _, c := range t.channels {
		select {
		case c <- buf.Bytes():
		default:
		}
	}
}

func (t *WS) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("ws-upgrade")
		return
	}
	defer conn.Close()

	ch := make(chan []byte, 32)

	t.mutex.Lock()
	t.nextid++
	id := t.nextid
	t.channels[id] = ch
	t.clients++
	wsClients.Set(float64(t.clients))
	t.mutex.Unlock()

	defer func() {
		t.mutex.Lock()
		t.clients--
		wsClients.Set(float64(t.clients))
		delete(t.channels, id)
		t.mutex.Unlock()
	}()

	// detect closed connections: the read fails as soon as the client disconnects
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case data := <-ch:
			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}
		}
	}
}
