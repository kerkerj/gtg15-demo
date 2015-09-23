package main

import (
	"github.com/d4l3k/go-pry/pry"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"time"
)

func main() {
	gbot := gobot.NewGobot()
	sphero_name := "/dev/tty.Sphero-RGG-AMP-SPP"

	adaptor := sphero.NewSpheroAdaptor("sphero", sphero_name)
	driver := sphero.NewSpheroDriver(adaptor, "sphero")

	pry.Pry()

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			//driver.Roll(30, uint16(gobot.Rand(360)))
			driver.SetRGB(uint8(gobot.Rand(255)), uint8(gobot.Rand(255)), uint8(gobot.Rand(255)))
		})
	}

	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
