package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-simulation-protocol/pkg/sim"
	"github.com/golang/protobuf/proto"
	"log"
	"math"
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
		simCommand.Control.TeleportBall = &sim.TeleportBall{
			X:  new(float32),
			Y:  new(float32),
			Z:  new(float32),
			Vx: new(float32),
			Vy: new(float32),
			Vz: new(float32),
		}
		*simCommand.GetControl().GetTeleportBall().X = 1
		*simCommand.GetControl().GetTeleportBall().Y = -2
		*simCommand.GetControl().GetTeleportBall().Z = 1
		*simCommand.GetControl().GetTeleportBall().Vx = 2
		*simCommand.GetControl().GetTeleportBall().Vy = 0.0
		*simCommand.GetControl().GetTeleportBall().Vz = 1

		teleBot := sim.TeleportRobot{
			Id: &sim.RobotId{
				Id:   new(uint32),
				Team: new(sim.Team),
			},
			X:           new(float32),
			Y:           new(float32),
			Orientation: new(float32),
			VX:          new(float32),
			VY:          new(float32),
			VAngular:    new(float32),
		}
		*teleBot.Id.Id = 0
		*teleBot.Id.Team = sim.Team_YELLOW
		*teleBot.X = 2
		*teleBot.Y = 3
		*teleBot.Orientation = math.Pi / 2
		*teleBot.VX = 0.0
		*teleBot.VY = 0.0
		*teleBot.VAngular = 0

		simCommand.Control.TeleportRobot = append(simCommand.Control.TeleportRobot, &teleBot)

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

		break
	}
}
