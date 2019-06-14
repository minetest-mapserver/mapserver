package web

import (
	"bytes"
	"encoding/json"
	"mapserver/app"
	"math/rand"
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
		panic(err)
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
		fields := logrus.Fields{
			"err": err,
		}
		logrus.WithFields(fields).Error("ws-upgrade")

	}

	id := rand.Intn(64000)
	ch := make(chan []byte)

	t.mutex.Lock()
	t.channels[id] = ch
	t.clients++
	wsClients.Set(float64(t.clients))
	t.mutex.Unlock()

	for {
		data := <-ch

		if data == nil {
			//how the hell got a nil reference in here..?!
			//related issue: #18
			continue
		}

		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			break
		}
	}

	t.mutex.Lock()
	t.clients--
	wsClients.Set(float64(t.clients))
	delete(t.channels, id)
	close(ch)
	t.mutex.Unlock()
}
