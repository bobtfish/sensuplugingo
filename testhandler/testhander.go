package main

import (
	"fmt"
	"github.com/bobtfish/sensuplugingo/handler"
)

func main() {
	event := handler.GetEventFromStdin
	//if err != nil {
	//	panic("Error getting Event")
	//}
	fmt.Println(event)
}
