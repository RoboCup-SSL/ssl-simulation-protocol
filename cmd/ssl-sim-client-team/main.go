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

var address = flag.String("address", "localhost:10301", "The address of the simulator robot control port (10301 for blue, 10302 for yellow)")

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
		robotControl := sim.RobotControl{}

		id := uint32(0)
		kickSpeed := float32(5)
		robotControl.RobotCommands = []*sim.RobotCommand{
			{
				Id:        &id,
				KickSpeed: &kickSpeed,
			},
		}

		log.Print("Sending: ", proto.MarshalTextString(&robotControl))
		bSend, err := proto.Marshal(&robotControl)
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

		robotControlResponse := sim.RobotControlResponse{}
		if err := proto.Unmarshal(bReceive[0:n], &robotControlResponse); err != nil {
			log.Print("Could not unmarshal message: ", err)
			continue
		}

		log.Print("Received: ", proto.MarshalTextString(&robotControlResponse))

		time.Sleep(1 * time.Second)
	}
}
