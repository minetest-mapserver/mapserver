package web

import (
	"bytes"
	"encoding/json"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapblockparser"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WS struct {
	ctx      *app.App
	channels map[int]chan []byte
	mutex    *sync.RWMutex
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
}

func (t *WS) OnParsedMapBlock(block *mapblockparser.MapBlock) {
	t.SendJSON("parsed-mapblock", block)
}

func (t *WS) OnRenderedTile(tc *coords.TileCoords) {
	t.SendJSON("rendered-tile", tc)
}

func (t *WS) SendJSON(eventtype string, o interface{}) {
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
	t.mutex.Unlock()

	defer func() {
		t.mutex.Lock()
		delete(t.channels, id)
		close(ch)
		t.mutex.Unlock()
	}()

	for {

		data := <-ch
		conn.WriteMessage(websocket.TextMessage, data)
	}

}
