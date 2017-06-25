package main

import (
	"fmt"
)

func main() {
	printer := messagePrinter{"foobar"}

	printer.printMessage()

}

type messagePrinter struct {
	message string
}

func (printer *messagePrinter) printMessage() {
	fmt.Println(printer.message)
}
