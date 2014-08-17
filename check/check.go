package check

import (
	"fmt"
	"os"
)

func OK(m string) {
	fmt.Println(m)
	os.Exit(0)
}

func WARNING(m string) {
	fmt.Println(m)
	os.Exit(1)
}

func CRITICAL(m string) {
	fmt.Println(m)
	os.Exit(2)
}
