package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jeffail/gabs"
)

func readInput(f string) (*gabs.Container, error) {
	var bytes []byte
	if f != "" {
		debug.Printf("Attempting to read file %s", f)
		b, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}
		bytes = b
	} else {
		stat, _ := os.Stdin.Stat()
		if stat.Mode()&os.ModeNamedPipe != 0 {
			debug.Printf("Attempting to read from stdin")
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return nil, err
			}
			bytes = b
		}
	}

	if len(bytes) == 0 {
		e := fmt.Errorf("No input received.")
		return nil, e
	}

	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		return nil, err
	}

	return jsonParsed, nil
}

func errAndExit(e error) {
	fmt.Printf("Error: %s\n", e)
	os.Exit(1)
}
