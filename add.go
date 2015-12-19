package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/jeffail/gabs"
)

var cmdAdd cli.Command

func init() {
	cmdAdd = cli.Command{
		Name:  "add",
		Usage: "add data to a json file",
		Subcommands: []cli.Command{
			{
				Name:   "key",
				Usage:  "Add a new key/value pair.",
				Action: actionAddKey,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file,f",
						Usage: "the file to edit. stdin if not specified.",
					},
					cli.StringFlag{
						Name:  "path,p",
						Usage: "the path to insert the data.",
					},
					cli.StringFlag{
						Name:  "value,v",
						Usage: "the value to set.",
					},
				},
			},
			{
				Name:   "array",
				Usage:  "Add a new array",
				Action: actionAddArray,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file,f",
						Usage: "the file to edit. stdin if not specified.",
					},
					cli.StringFlag{
						Name:  "path,p",
						Usage: "the path to insert the data.",
					},
					cli.StringSliceFlag{
						Name:  "value,v",
						Usage: "the value to set. can be used multiple times.",
					},
				},
			},
			{
				Name:   "element",
				Usage:  "Add a new array element",
				Action: actionAddArrayElement,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "file,f",
						Usage: "the file to edit. stdin if not specified.",
					},
					cli.StringFlag{
						Name:  "path,p",
						Usage: "the path to insert the data.",
					},
					cli.BoolFlag{
						Name:  "contains,c",
						Usage: "don't insert if the value already exists.",
					},
					cli.StringFlag{
						Name:  "value,v",
						Usage: "the value to set.",
					},
				},
			},
		},
	}
}

func actionAddKey(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	j, err = addKey(j, c.String("path"), c.String("value"))
	if err != nil {
		errAndExit(err)
	}

	if pretty {
		fmt.Printf(j.StringIndent("", "  "))
	} else {
		fmt.Printf(j.String())
	}

}

func actionAddArray(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	j, err = addArray(j, c.String("path"), c.StringSlice("value"))
	if err != nil {
		errAndExit(err)
	}

	if pretty {
		fmt.Printf(j.StringIndent("", "  "))
	} else {
		fmt.Printf(j.String())
	}
}

func actionAddArrayElement(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	j, err = addArrayElement(j, c.String("path"), c.String("value"), c.Bool("contains"))
	if err != nil {
		errAndExit(err)
	}
}

func addKey(j *gabs.Container, path string, value string) (*gabs.Container, error) {
	var err error
	if i, e := strconv.Atoi(value); e == nil {
		_, err = j.SetP(i, path)
	} else {
		_, err = j.SetP(value, path)
	}

	return j, err
}

func addArray(j *gabs.Container, path string, values []string) (*gabs.Container, error) {
	var err error
	_, err = j.ArrayP(path)
	if err != nil {
		return nil, err
	}

	for _, v := range values {
		if i, e := strconv.Atoi(v); e == nil {
			err = j.ArrayAppendP(i, path)
		} else {
			err = j.ArrayAppendP(v, path)
		}
	}

	return j, err
}

func addArrayElement(j *gabs.Container, path string, value string, contains bool) (*gabs.Container, error) {
	var err error

	if contains {
		arrSize, err := j.ArrayCountP(path)
		if err != nil {
			return j, err
		}

		for i := 0; i < arrSize; i++ {
			v, err := j.ArrayElementP(i, path)
			if err != nil {
				return j, err
			}

			if x, err := strconv.ParseFloat(value, 64); err == nil {
				if x == v.Data().(float64) {
					return j, nil
				}
			} else {
				if value == v.Data() {
					return j, nil
				}
			}
		}
	}

	path_pieces := strings.Split(path, ".")
	if index, err := strconv.Atoi(path_pieces[len(path_pieces)-1]); err == nil {
		if index >= 0 {
			arrPath := strings.Join(path_pieces[:len(path_pieces)-1], ".")
			arr := j.Path(arrPath)

			if i, err := strconv.Atoi(value); err == nil {
				_, err = arr.SetIndex(i, index)
			} else {
				_, err = arr.SetIndex(value, index)
			}
		}
	} else {
		if i, err := strconv.Atoi(value); err == nil {
			err = j.ArrayAppendP(i, path)
		} else {
			err = j.ArrayAppendP(value, path)
		}
	}

	return j, err
}
