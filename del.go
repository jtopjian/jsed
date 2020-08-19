package main

import (
	"fmt"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/urfave/cli/v2"
)

var cmdDel cli.Command

type delKeyOptions struct {
	json      *gabs.Container
	path      string
	delimiter string
}

func init() {
	cmdDel = cli.Command{
		Name:  "del",
		Usage: "delete data from a json file",
		Subcommands: []*cli.Command{
			{
				Name:   "key",
				Usage:  "Delete a new key/value pair.",
				Action: actionDelKey,
				Flags: []cli.Flag{
					&flagFile,
					&flagPath,
					&flagDelimiter,
					&flagPretty,
				},
			},
		},
	}
}

func actionDelKey(c *cli.Context) error {
	j, err := readInput(c.String("file"))
	if err != nil {
		return err
	}

	options := delKeyOptions{
		json:      j,
		path:      c.String("path"),
		delimiter: getDelimiter(c.String("delimiter")),
	}

	j, err = delKey(options)
	if err != nil {
		return err
	}

	if pretty {
		fmt.Printf(j.StringIndent("", "  "))
	} else {
		fmt.Printf(j.String())
	}
	return nil
}

func delKey(options delKeyOptions) (*gabs.Container, error) {
	j := options.json
	path := strings.Split(options.path, options.delimiter)
	err := j.Delete(path...)
	return j, err
}
