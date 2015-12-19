package main

import (
	"log"
)

var debugMode bool
var debug debugger

type debugger bool

func (d debugger) Printf(format string, args ...interface{}) {
	if debugMode {
		log.Printf(format, args...)
	}
}
