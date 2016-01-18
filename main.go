package main

import (
	"log"
)

func main() {
	log.Fatal(NewGoApp().ListenAndServe(":8080"))
}
