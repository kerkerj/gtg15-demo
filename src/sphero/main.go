package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "/dev/tty.Sphero-RGG-AMP-SPP"
	}

	s := NewSphero("kerkerj", port)
	defer s.Stop()

	go func() {
		api := NewApi(s)
		http.ListenAndServe(":5566", api.Handler())
	}()

	s.Start()
}
