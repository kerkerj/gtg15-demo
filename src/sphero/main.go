package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "/dev/tty.Sphero-RGG-AMP-SPP"
	}

	s := NewSphero("kerkerj", port)
	defer s.Stop()

	// Beacuse Gobot/Sphero lib will block the main process
	// So, here I created a go routine to deal with the api endpoints.
	go func() {
		api := NewApi(s)
		http.ListenAndServe(":5566", api.Handler())
	}()

	s.Start()
}
