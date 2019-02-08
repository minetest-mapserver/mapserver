package web

import (
	"bytes"
	"encoding/json"
	"mapserver/app"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	wsClients = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ws_client_count",
		Help: "Websocket client count",
	})
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
	prometheus.MustRegister(wsClients)

	return &ws
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
	conn, _ := upgrader.Upgrade(resp, req, nil)

	id := rand.Intn(64000)
	ch := make(chan []byte)

	t.mutex.Lock()
	t.channels[id] = ch
	t.clients++
	wsClients.Set(float64(t.clients))
	t.mutex.Unlock()

	for {
		data := <-ch
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
