package hrest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"libs/github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"sort"
)

var upgrader = websocket.Upgrader{}

type Service struct {
	Method string
	Action func(w http.ResponseWriter, r *http.Request)
}

type HRest struct {
	ListServices          map[string]Service
	ListWebSocketServices map[string]func(message []byte, c *websocket.Conn)
}

func NewRest() *HRest {
	htools := new(HRest)
	htools.ListServices = make(map[string]Service)
	htools.ListWebSocketServices = make(map[string]func(message []byte, c *websocket.Conn))
	return htools
}

func (h *HRest) CorsGolang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	elem, ok := h.ListServices[r.URL.Path]

	if ok {
		if elem.Method == r.Method {
			h.ListServices[r.URL.Path].Action(w, r)
		} else {
			fmt.Fprintln(w, "Method not allow")
		}

	} else {
		fmt.Fprintln(w, "Service not found")
	}

}

func (h *HRest) RegisterService(path, method string, action func(w http.ResponseWriter, r *http.Request)) {
	h.ListServices[path] = Service{method, action}
	http.HandleFunc(path, h.CorsGolang)
}

func (h *HRest) WebsocketRequestManager(message []byte, c *websocket.Conn) {

	if string(message) == "connect" {

		f, o := h.ListWebSocketServices["connect"]
		if o {
			f([]byte(""), c)
		}

		c.WriteMessage(1, []byte("accept"))

	} else {
		var dat map[string]string

		err := json.Unmarshal(message, &dat)

		if err != nil {
			fmt.Println("error con la data en hrest: ", err)
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

func (h *HRest) StartWebSocketServer(path string) {

	p := func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		defer c.Close()

		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				log.Println("read:", err)
				break
			}

			go h.WebsocketRequestManager(message, c)

		}

	}

	h.RegisterService(path, "GET", p)

}

func (h *HRest) RegisterWebSocket(path string, f func(message []byte, c *websocket.Conn)) {
	h.ListWebSocketServices[path] = f
}

func (h *HRest) Encode(bodyrequest []byte) string {

	var v map[string]string

	json.Unmarshal(bodyrequest, &v)

	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := url.QueryEscape(k) + "="

		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(prefix)
		buf.WriteString(url.QueryEscape(vs))

	}
	return buf.String()
}
