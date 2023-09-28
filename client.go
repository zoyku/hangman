package main

import (
	"github.com/reiver/go-telnet"
)

func main() {
	var caller telnet.Caller = telnet.StandardCaller

	telnet.DialToAndCall("127.0.0.1:5555", caller)
}
