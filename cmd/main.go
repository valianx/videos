package main

import (
	"github.com/valianx/videos/internal"
	"os"
	_ "os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	internal.Connect()
	r := internal.Routes(port)
	r.Run()
}
