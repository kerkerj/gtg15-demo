package main

import (
	"net/http"
	"os"
	"path/filepath"
	"log"
	"fmt"
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

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
          log.Fatal(err)
  }
  fmt.Println(dir)

	s.Start()
}
