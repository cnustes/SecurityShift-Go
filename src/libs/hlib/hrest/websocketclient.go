package hrest

import (
	"encoding/json"
	"fmt"
	"libs/github.com/gorilla/websocket"
	"time"
)

type HWSClient struct {
	urlws                 string
	ListWebSocketServices map[string]func(message []byte, c *websocket.Conn)
}

func NWSClient(urlws string) *HWSClient {
	htools := new(HWSClient)
	htools.ListWebSocketServices = make(map[string]func(message []byte, c *websocket.Conn))
	htools.urlws = urlws
	htools.StarClient()
	return htools
}

func (h *HWSClient) connect() *websocket.Conn {
	var dialer = websocket.Dialer{}
	c, _, err := dialer.Dial(h.urlws, nil)
	if err != nil {
		time.Sleep(20000 * time.Millisecond)
		c = h.connect()
	} else {
		c.WriteMessage(websocket.TextMessage, []byte("connect"))
	}

	return c
}

func (h *HWSClient) RegisterWebSocket(path string, f func(message []byte, c *websocket.Conn)) {
	h.ListWebSocketServices[path] = f
}

func (h *HWSClient) processCommand(message []byte, c *websocket.Conn) {
	if string(message) == "accept" {
		f, o := h.ListWebSocketServices["connect"]
		if o {
			f([]byte(""), c)
		}
	} else {
		var dat map[string]string
		fmt.Println("message client >>> ", string(message))
		err := json.Unmarshal(message, &dat)

		if err != nil {
			fmt.Println("error con la data en websockets cliente: ", err)
		} else {

			method, ok := dat["command"]
			if ok {
				f, o := h.ListWebSocketServices[method]
				if o {
					f(message, c)
				}
			}
		}
	}
}

func (h *HWSClient) StarClient() {

	connection := h.connect()

	//defer connection.Close()

	go func() {
		defer connection.Close()
		for {
			_, message, err := connection.ReadMessage()

			if err != nil {
				fmt.Println("start.....", err)
				connection = h.connect()
			} else {
				h.processCommand(message, connection)
			}

		}
	}()

}
