package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var pretty bool

func main() {
	app := cli.NewApp()
	app.Name = "jsed"
	app.Usage = "a simple json utility"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug,d",
			Usage:       "debug mode",
			Destination: &debugMode,
		},
		cli.BoolFlag{
			Name:        "pretty,p",
			Usage:       "pretty print",
			Destination: &pretty,
		},
	}

	app.Commands = []cli.Command{
		cmdGet,
		cmdContains,
		cmdAdd,
		cmdDel,
	}

	app.Run(os.Args)
}
