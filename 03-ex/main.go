package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("shit happens: %v", ce.info)
}

func main() {

	ce := customErr{
		info: "Nya",
	}

	foo(ce)
}

func foo(e error) {
	fmt.Println(e)
}
