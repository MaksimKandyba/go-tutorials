package main

import (
	"fmt"

	"github.com/MaksimKandyba/go-tutorials/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
