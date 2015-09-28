package main

import (
  "github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
)

func NewSphero(name, port string) Sphero {
	return &s_robot{Name: name, Port: port}
}

type Sphero interface {
	Start()
	Stop()
	SetRGB(r, g, b uint8)
}

type s_robot struct {
	Name       string
	Port       string
  device     *sphero.SpheroDriver
	connection *sphero.SpheroAdaptor
}

func (s *s_robot) Start() {
  // Connect Sphero with Gobot
	gbot := gobot.NewGobot()
  adaptor := sphero.NewSpheroAdaptor(s.Name, s.Port)
	driver := sphero.NewSpheroDriver(adaptor, "sphero")

  s.connection = adaptor
	s.device = driver

  robot := gobot.NewRobot(s.Name,
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
	)

  gbot.AddRobot(robot)
	gbot.Start()
}

func (s *s_robot) Stop() {
	s.connection.Disconnect()
}

func (s *s_robot) SetRGB(r, g, b uint8) {
	s.device.SetRGB(r, g, b)
}
