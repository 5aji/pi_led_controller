package wled

// this code is used for UDP live streaming. uses goroutines.

import (
	"fmt"
	"net"
	"time"
)

// This function takes a slice of Wled Colors and creates a byte
// array for sending.
func mapToBytes(colors *[]WledColor, buf *[]byte) {
	*buf =	(*buf)[:2] // trim off excess bits.

	for _, color := range *colors {
		*buf = append(*buf, color[0], color[1], color[2])
	}
}

// Create a UDP socket connection with a WLED device. Takes a timeout byte
// and a data channel. Returns a stopping channel to close the connection.
// The last sent state will be sent periodically to prevent timeout.
func StreamLights(addr string, data <-chan *[]WledColor) chan<- bool {

	// create udp connection. TODO: use better port config.
	udpaddr, err := net.ResolveUDPAddr("udp", addr+":21324")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to UDP @ %s\n", addr)

	// create internal buffer and setup header

	buf := make([]byte, 2, 1472)
	buf[0] = 2 // using DRGB mode
	buf[1] = 3

	// create resend ticker to keep colors active (no longer required to maintain outside state)
	dur := 2 * time.Second
	ticker := time.NewTicker(dur)


	// we use the signal channel to stop sending.
	signal := make(chan bool) // we pass this back to caller

	// helper func to send udp packet.
	sendPacket := func() {
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Printf("Error sending UDP packet: %v\n", err)
		}
	}
	go func() {
		defer conn.Close() // do this in here since this func uses it.
		for {
			select {
			case led := <-data:
				mapToBytes(led, &buf)
				sendPacket()
				ticker.Reset(dur) // so we don't double send.
			case <-ticker.C:
				sendPacket() // resend last packet, preventing the lights from reverting.

			case s := <-signal:
				if s {
					return
				}
			}
		}

	}()
	return signal

}
