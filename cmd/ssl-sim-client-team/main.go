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
var command = flag.String("command", "local", "Command to send")

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

	go receive(conn)

	for {
		robotControl := sim.RobotControl{}

		id := uint32(0)
		kickSpeed := float32(4)
		kickAngle := float32(0)
		dribbleSpeed := float32(0)
		var robotCommand sim.RobotMoveCommand
		if *command == "local" {
			forward := float32(0)
			left := float32(0)
			angular := float32(0)
			robotCommand = sim.RobotMoveCommand{
				Command: &sim.RobotMoveCommand_LocalVelocity{
					LocalVelocity: &sim.MoveLocalVelocity{
						Forward: &forward,
						Left:    &left,
						Angular: &angular,
					},
				},
			}
		} else {
			globalVelX := float32(0.0)
			globalVelY := float32(0.0)
			globalVelAngular := float32(0.0)
			robotCommand = sim.RobotMoveCommand{
				Command: &sim.RobotMoveCommand_GlobalVelocity{
					GlobalVelocity: &sim.MoveGlobalVelocity{
						X:       &globalVelX,
						Y:       &globalVelY,
						Angular: &globalVelAngular,
					},
				},
			}
		}
		robotControl.RobotCommands = []*sim.RobotCommand{
			{
				Id:            &id,
				MoveCommand:   &robotCommand,
				KickSpeed:     &kickSpeed,
				KickAngle:     &kickAngle,
				DribblerSpeed: &dribbleSpeed,
			},
		}

		log.Print("Sending: ", proto.MarshalTextString(&robotControl))
		bSend, _ := proto.Marshal(&robotControl)
		if _, err := conn.Write(bSend); err != nil {
			log.Print("Failed to write: ", err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Printf("%d bytes sent", len(bSend))

		time.Sleep(1 * time.Second)
		return
	}
}

func receive(conn *net.UDPConn) {
	bReceive := make([]byte, maxDatagramSize)
	for {
		log.Printf("Receiving...")
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

		log.Printf("Received %d bytes: %s", n, proto.MarshalTextString(&robotControlResponse))
	}
}
