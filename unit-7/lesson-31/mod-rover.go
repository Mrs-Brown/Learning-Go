package main

import (
	"image"
	"log"
	"time"
)

type RoverDriver struct {
	commandc chan command
}

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(2)
	stop  = command(3)
)

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

// drive is a goroutine for driving the rover
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	speed := 1
	for {
		select {
		case c := <-r.commandc: //Waits for commands on the command channel
			switch c {
			case right: //Turns right
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left: //Turns left
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case stop:
				speed = 0
			case start:
				speed = 1
			}
			log.Printf("new direction %v; speed %d", direction, speed)
		case <-nextMove:
			pos = pos.Add(direction.Mul(speed))
			log.Printf("moved to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

// Left turns the rover left
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right
func (r *RoverDriver) Right() {
	r.commandc <- left
}

// Stop halts the rover
func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

// Start gets the rover moving
func (r *RoverDriver) Start() {
	r.commandc <- start
}

func main() {
	r := NewRoverDriver()
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
	r.Stop()
	time.Sleep(3 * time.Second)
	r.Start()
	time.Sleep(3 * time.Second)
}
