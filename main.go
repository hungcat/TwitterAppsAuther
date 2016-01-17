package main

import (
	"log"
)

func main() {
	goapp := NewGoApp()
	goapp.RegisterHandlers()
	log.Fatal(goapp.ListenAndServe(":8080"))
}
