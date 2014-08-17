package main

import (
	"fmt"
	"github.com/bobtfish/sensuplugingo/handler"
)

func main() {
	fmt.Println("Hello")
	e := handler.GetEventFromStdin
	//if err != nil {
	//	panic("Error getting Event")
	//}
	fmt.Println(e)
}
