package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()

	//Use at your own risk! This is untested on pi
	r := raspi.NewRaspiAdaptor("raspi")
	pin := gpio.NewDirectPinDriver(r, "pin", "13")
	work := func() {
		gobot.Every(1*time.Second, func() {
			if v, err := pin.DigitalRead(); err != nil {
				fmt.Printf("Signal Detected: %v\n", v)
			}
		})
	}

	robot := gobot.NewRobot("pinBot",
		[]gobot.Connection{r},
		[]gobot.Device{pin},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
