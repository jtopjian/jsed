package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/jeffail/gabs"
)

var cmdDel cli.Command

func init() {
	cmdDel = cli.Command{
		Name:  "del",
		Usage: "delete data from a json file",
		Subcommands: []cli.Command{
			{
				Name:   "key",
				Usage:  "Delete a new key/value pair.",
				Action: actionDelKey,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file,f",
						Usage: "the file to edit. stdin if not specified.",
					},
					cli.StringFlag{
						Name:  "path,p",
						Usage: "the path to delete the data.",
					},
				},
			},
		},
	}
}

func actionDelKey(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	j, err = delKey(j, c.String("path"))
	if err != nil {
		errAndExit(err)
	}

	if pretty {
		fmt.Printf(j.StringIndent("", "  "))
	} else {
		fmt.Printf(j.String())
	}
}

func delKey(j *gabs.Container, path string) (*gabs.Container, error) {
	err := j.DeleteP(path)
	return j, err
}
