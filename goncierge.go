package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

var (
	host      *string = flag.String("host", "http://192.168.50.145:3000/door_events", "goncierge host")
	room_slug *string = flag.String("room_slug", "foo", "goncierge room slug")
)

func main() {
	flag.Parse()

	gbot := gobot.NewGobot()

	r := raspi.NewRaspiAdaptor("raspi")
	statusLed := gpio.NewLedDriver(r, "statusLed", "11")
	eventLed := gpio.NewLedDriver(r, "eventLed", "7")
	button := gpio.NewButtonDriver(r, "button", "13")

	work := func() {
		gobot.Every(1*time.Second, func() {
			statusLed.Toggle()
		})
		gobot.On(button.Event("push"), func(data interface{}) {
			eventLed.On()
			go toggleDoorState("closed")
		})
		gobot.On(button.Event("release"), func(data interface{}) {
			eventLed.Off()
			go toggleDoorState("open")
		})
	}

	robot := gobot.NewRobot("Goncierge",
		[]gobot.Connection{r},
		[]gobot.Device{button, statusLed, eventLed},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}

func toggleDoorState(state string) {
	now := time.Now().Format(time.RFC3339)
	_, err := http.PostForm(
		*host,
		url.Values{
			"room_slug":  {*room_slug},
			"timestamp":  {now},
			"door_state": {state},
		},
	)
	if err != nil {
		fmt.Printf("[%s] Host Error: %s\n", now, err.Error())
	}
}
