package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-simulation-protocol/pkg/sim"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"time"
)

const maxDatagramSize = 8192

var address = flag.String("address", "localhost:10300", "The address of the simulator control port")

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", *address)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to", *address)

	if err := conn.SetReadBuffer(maxDatagramSize); err != nil {
		log.Printf("Could not set read buffer to %v.", maxDatagramSize)
	}

	bReceive := make([]byte, maxDatagramSize)

	for {
		simCommand := sim.SimulatorCommand{}

		simCommand.Control = &sim.SimulatorControl{}
		simCommand.Control.TeleportBall = &sim.TeleportBall{X: new(float32), Y: new(float32)}
		*simCommand.GetControl().GetTeleportBall().X = 1
		*simCommand.GetControl().GetTeleportBall().Y = -2

		log.Print("Sending: ", proto.MarshalTextString(&simCommand))
		bSend, err := proto.Marshal(&simCommand)
		if _, err := conn.Write(bSend); err != nil {
			log.Print("Failed to write: ", err)
			time.Sleep(1 * time.Second)
			continue
		}

		n, _, err := conn.ReadFrom(bReceive)
		if err != nil {
			log.Print("Could not read: ", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if n >= maxDatagramSize {
			log.Fatal("Buffer size too small")
		}

		simResponse := sim.SimulatorResponse{}
		if err := proto.Unmarshal(bReceive[0:n], &simResponse); err != nil {
			log.Print("Could not unmarshal message: ", err)
			continue
		}

		log.Print("Received: ", proto.MarshalTextString(&simResponse))

		time.Sleep(1 * time.Second)
	}
}
