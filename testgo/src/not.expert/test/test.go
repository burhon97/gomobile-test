package test

import (
	"fmt"
)

func Greeting(name string) string {
	message := fmt.Sprintf("Hello from Gomobile, %s", name)
	return message
}
