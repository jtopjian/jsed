package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var pretty bool
var flagFile, flagPath, flagDelimiter cli.StringFlag
var flagMultiKey, flagMultiValue cli.StringSliceFlag
var flagDebug, flagPretty, flagExists cli.BoolFlag

func init() {
	flagFile = cli.StringFlag{
		Name:  "file,f",
		Usage: "the file to edit. stdin if not specified.",
	}

	flagPath = cli.StringFlag{
		Name:  "path,p",
		Usage: "the path to the data being acted upon.",
	}

	flagDelimiter = cli.StringFlag{
		Name:  "delimiter,d",
		Usage: "the path delimiter",
	}

	flagMultiKey = cli.StringSliceFlag{
		Name:  "key,k",
		Usage: "the key to set. can be used multiple times.",
	}

	flagMultiValue = cli.StringSliceFlag{
		Name:  "value,v",
		Usage: "the value to set. can be used multiple times.",
	}

	flagDebug = cli.BoolFlag{
		Name:        "debug,d",
		Usage:       "debug mode",
		Destination: &debugMode,
	}

	flagPretty = cli.BoolFlag{
		Name:        "pretty,r",
		Usage:       "pretty print the result.",
		Destination: &pretty,
	}

	flagExists = cli.BoolFlag{
		Name:  "exists,e",
		Usage: "don't insert if value already exists",
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "jsed"
	app.Usage = "a simple json utility"
	app.Version = "0.2.2"

	app.Flags = []cli.Flag{
		&flagDebug,
	}

	app.Commands = []cli.Command{
		cmdGet,
		cmdContains,
		cmdAdd,
		cmdDel,
	}

	app.Run(os.Args)
}
