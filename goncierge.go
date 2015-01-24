package main

import (
	"fmt"
	//"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()

	api.NewAPI(gbot).Start()

	r := raspi.NewRaspiAdaptor("raspi")
	//pin := gpio.NewDirectPinDriver(r, "pin", "13")
	led := gpio.NewLedDriver(r, "led", "7")
	button := gpio.NewButtonDriver(r, "button", "13")

	work := func() {
		gobot.On(button.Event("push"), func(data interface{}) {
			fmt.Println("Button Pushed")
			led.On()
		})
		gobot.On(button.Event("release"), func(data interface{}) {
			fmt.Println("Button Released")
			led.Off()
		})
	}

	/*
		work := func() {
			gobot.Every(500*time.Millisecond, func() {
				v, err := pin.DigitalRead()
				if err != nil {
					fmt.Printf("Digital Read Error: %s\n", err.Error())
				}
				if v == 1 {
					fmt.Println("Motion Detected")
					led.On()
				} else {
					fmt.Println("No Motion Detected")
					led.Off()
				}
			})
		}
	*/

	robot := gobot.NewRobot("Goncierge",
		[]gobot.Connection{r},
		//[]gobot.Device{pin, led},
		[]gobot.Device{button, led},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
