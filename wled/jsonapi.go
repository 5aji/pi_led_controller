package wled

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"github.com/gorilla/websocket"
)


func GetWledState(addr string) WledState {
	u := url.URL{Scheme: "http", Host: addr, Path: "/json/state"}

	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result WledState
	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

// Pushes a state payload to the provided address.
func SetWledState(addr string, state interface{}) ([]byte, error) {
	u := url.URL {Scheme: "http", Host: addr, Path: "/json/state"}
	
	encodedPacket, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(encodedPacket))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	return bodyBytes, nil
}

// creates a json payload to enable/disable LEDs
func SetWledLights(state bool) interface{} {
	// construct a json packet
	packet := make(map[string]bool)

	packet["on"] = state

	return packet
}

// creates a wled color setting json payload to be used by a websocket or HTTP POST.
func SetWledColor(colors [3]WledColor) interface{} {
	packet := make(map[string]interface{})
	packet["seg"] = make(map[string]interface{})
	packet["seg"].(map[string]interface{})["col"] = colors

	return packet
}


// creates a websocket connection that can be written to and read from.
// The returned channel passes back WledState structs. 
func WledWebsocket(addr string, state <-chan interface{}, done <-chan bool) <-chan WledState {

	u := url.URL {Scheme: "ws", Host: addr, Path: "/ws"}
	
	resp := make(chan WledState)

	// connect to websocket.

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		panic(err)
	}
	
	go func() {
		for {
			select {
			case msg := <- state:
				err := conn.WriteJSON(msg)
				if err != nil {
					fmt.Printf("error writing message: %v\n", err)
					return
				}
			case <-done:
				return

			}
		}
	}()

	// TODO: rx side doesn't close unless it gets a result.
	go func() {
		var v WledState
		for {
			err := conn.ReadJSON(&v)
			if err != nil {
				fmt.Printf("error reading message: %v\n", err)
				return
			}
			select {
			case resp <- v:

			case <-done:
				conn.Close()
				return
			}

		}
	}()

	return resp
}


