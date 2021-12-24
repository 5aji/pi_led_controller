package main

import (
	"fmt"
	"math"
	"time"

	// "os"

	"math/rand"

	"github.com/kschamplin/pi_led_controller/lib/wled"
	"github.com/lucasb-eyer/go-colorful"
)


const framerate float64 = 60
const delta = 1

// takes a pointer to an array of wled colors and sets 
// all the colors to the provided colorful color
func setLeds(leds *[]wled.WledColor, color colorful.Color) {

	wColor := wled.ToWled(color)

	for i := range *leds {
		(*leds)[i] = wColor
	}

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
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

	leds := make([]wled.WledColor, 300)

	ticker := time.NewTicker(time.Second / time.Duration(framerate))


	fmt.Printf("starting data")

	defer func() {done <- true}()
	var angle float64 = 0
	var color colorful.Color = colorful.Hsv(angle, 1,1)
	for {
		select {
		case <-ticker.C:
			// update color
			angle = math.Mod(angle + delta, 360)
			color = colorful.Hsv(angle, 1,1)
			setLeds(&leds, color)
			data <- &leds
		}
	}
			


	// wled.SetWledState(addr, wled.SetWledLights(false))
	
}
