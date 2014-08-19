package main

import (
	"fmt"
	"github.com/bobtfish/sensuplugingo/handler"
	"github.com/bobtfish/sensuplugingo"
)

func FormatMessage(e sensuplugingo.Event) string {
	return("foo")
}

func GetChannels() (c []string) {
	return c
}

func Send(channel string, message string) error {
	// system('/nail/sys/bin/nodebot', channel, message)
	return(nil)
}

func main() {
	event, err := handler.GetEventFromStdin()
	if err != nil {
		panic("Error getting Event")
	}
	message := FormatMessage(event)
	channels := GetChannels()
	for _, channel := range channels {
		Send(channel, message)	
	}
	fmt.Println(event)
}
