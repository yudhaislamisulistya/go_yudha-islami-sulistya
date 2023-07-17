package main

import (
	"belajar-go-echo/route"
)

func main() {
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}
