package main

import (
	"fmt"

	greetings "github.com/aderayevans/Go-I-am-comin/packages"
	// greetings "Go-I-am-comin/packages"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
