package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"mapserver/app"
	"mapserver/coords"
	"mapserver/mapblockparser"
	"math/rand"
	"net/http"
	"sync"
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

type ParsedMapBlockEvent struct {
	Eventtype string                   `json:"type"`
	Block     *mapblockparser.MapBlock `json:"block"`
}

type RenderedTileEvent struct {
	Eventtype string             `json:"type"`
	Tc        *coords.TileCoords `json:"tilepos"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (t *WS) OnParsedMapBlock(block *mapblockparser.MapBlock) {
	e := &ParsedMapBlockEvent{"parsed-mapblock", block}
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	t.mutex.RLock()
	defer t.mutex.RUnlock()

	for _, c := range t.channels {
		select {
		case c <- data:
		default:
		}
	}
}

func (t *WS) OnRenderedTile(tc *coords.TileCoords) {

	e := &RenderedTileEvent{"rendered-tile", tc}
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	t.mutex.RLock()
	defer t.mutex.RUnlock()

	for _, c := range t.channels {
		select {
		case c <- data:
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

	fmt.Print("Socket opened: ")
	fmt.Println(id)

	for {

		data := <-ch
		conn.WriteMessage(websocket.TextMessage, data)
	}

}
