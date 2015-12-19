package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/jeffail/gabs"
)

var cmdGet cli.Command
var cmdContains cli.Command

func init() {
	cmdGet = cli.Command{
		Name:   "get",
		Usage:  "extract an path from a json file",
		Action: actionGet,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file,f",
				Usage: "the file to search.",
			},
			cli.StringFlag{
				Name:  "path,p",
				Usage: "the search path.",
			},
		},
	}

	cmdContains = cli.Command{
		Name:   "contains",
		Usage:  "determine if a value is contained in an array",
		Action: actionContains,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file,f",
				Usage: "the file to search.",
			},
			cli.StringFlag{
				Name:  "path,p",
				Usage: "the search path.",
			},
			cli.StringFlag{
				Name:  "value,v",
				Usage: "the value contained in the array.",
			},
		},
	}
}

func actionGet(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	j, err = get(j, c.String("path"))
	if err != nil {
		errAndExit(err)
	}

	switch j.Data().(type) {
	case string:
		fmt.Printf("%s", j.Data())
	default:
		if pretty {
			fmt.Printf(j.StringIndent("", "  "))
		} else {
			fmt.Printf(j.String())
		}
	}
}

func actionContains(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	j, err = get(j, c.String("path"))
	if err != nil {
		errAndExit(err)
	}

	j, err = contains(j, c.String("value"))
	if err != nil {
		errAndExit(err)
	}

	if j != nil {
		switch j.Data().(type) {
		case string:
			fmt.Printf("%s", j.Data())
		default:
			if pretty {
				fmt.Printf(j.StringIndent("", "  "))
			} else {
				fmt.Printf(j.String())
			}
		}
	}
}

func get(j *gabs.Container, path string) (*gabs.Container, error) {
	for _, p := range strings.Split(path, ".") {
		debug.Printf("Path piece: %+v", p)
		if i, err := strconv.Atoi(p); err == nil {
			if array, ok := j.Data().([]interface{}); ok {
				debug.Printf("%+v is an array", j)
				if i < 0 {
					return nil, gabs.ErrNotArray
				}

				if i < len(array) {
					j = j.Index(i)
				} else {
					return nil, gabs.ErrOutOfBounds
				}
			} else {
				j = j.Path(p)
			}
		} else {
			j = j.Path(p)
		}
	}

	return j, nil
}

func contains(j *gabs.Container, value string) (*gabs.Container, error) {
	if array, ok := j.Data().([]interface{}); ok {
		for i, a := range array {
			if f, err := strconv.ParseFloat(value, 64); err == nil {
				if f == a {
					return j.Index(i), nil
				}
			} else {
				if value == a {
					return j.Index(i), nil
				}
			}
		}
	}

	return nil, nil
}
