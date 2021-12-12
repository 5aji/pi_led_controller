package main

import (
	"fmt"
	"time"
	// "os"

	"github.com/kschamplin/wledkeeb/wled"
)

func main() {
	addr := "172.16.21.155"

	// state := os.Args[1]
	// var statebool bool

	// if state == "true" {
	// 	statebool = true
	// } else {
	// 	statebool = false
	// }

	// wled.SetWledLights("172.16.21.155", statebool)
	

	// websocket test

// 	out := make(chan interface{})
// 	done := make(chan bool)
// 	in := wled.WledWebsocket(addr, out, done)

// 	packet := make(map[string]bool)

// 	packet["on"] = true

// 	out <- packet
// 	time.Sleep(time.Second)
// 	for {
// 		select {
// 		case newstate := <- in:
// 			fmt.Printf("got new state: %v\n", newstate)
// 		}
// 	}

	// udp test

	// wled.SetWledState(addr, wled.SetWledLights(true))

	data := make(chan *[]wled.WledColor)
	done := wled.StreamLights(addr, data)

	testColor := make([]wled.WledColor, 300)

	fmt.Println("sending data")
	for i := range testColor {
		var color wled.WledColor
		if i < 150 {
			color = wled.WledColor{0,0,0}
		} else {
			color = wled.WledColor{255,255,255}
		}
		testColor[i] = color
	}

	data <- &testColor
	fmt.Println("all data sent, waiting")
	time.Sleep(5 * time.Second)

	// wled.SetWledState(addr, wled.SetWledLights(false))
	done <- true
	
}
