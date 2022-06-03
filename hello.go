package main

import (
	"fmt"

	greetings "Go-I-am-comin/packages"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
