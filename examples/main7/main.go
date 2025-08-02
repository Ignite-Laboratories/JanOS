package main

import (
	"fmt"
	"github.com/gvalkov/golang-evdev"
	"log"
)

func main() {
	fmt.Println("evdev Mouse Position Example")

	dev, err := evdev.Open("/dev/input/event1") // Replace "event0" with your device
	if err != nil {
		log.Fatalf("Failed to open input device: %v", err)
	}

	fmt.Printf("Reading raw input events from: %v\n", dev.Name)
	for {
		events, err := dev.Read()
		if err != nil {
			log.Fatal(err)
		}

		for _, ev := range events {
			if ev.Type == evdev.EV_REL {
				if ev.Code == evdev.REL_X || ev.Code == evdev.REL_Y {
					fmt.Printf("Mouse event: %v, Value: %v\n", ev.Code, ev.Value)
				}
			}
		}
	}
}
