package main

import (
	"fmt"
	"log"

	greetings "github.com/aderayevans/Go-I-am-comin/hello/packages"
	// greetings "Go-I-am-comin/packages"
	"golang.org/x/example/stringutil"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samatha", "Darius"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Println(stringutil.Reverse("Hello"))
}
