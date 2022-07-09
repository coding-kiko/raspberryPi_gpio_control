package main

import (
	// std lib
	"log"
	"net/http"

	// third party
	"github.com/stianeikeland/go-rpio"
)

const pinNumber = 1

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatalln("Unable to open gpio connection")
	}
	defer rpio.Close()

	pin := rpio.Pin(pinNumber)

	mux := http.NewServeMux()
	mux.Handle("/toggle", http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		log.Println("toggled")
		pin.Toggle()
	}))
	log.Fatalln(http.ListenAndServe("0.0.0.0:8888", mux))
}

/* func TogglePin(pin rpio.Pin) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pin.Toggle()
	})
} */
