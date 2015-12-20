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

type getOptions struct {
	json      *gabs.Container
	path      string
	delimiter string
}

func init() {
	cmdGet = cli.Command{
		Name:   "get",
		Usage:  "extract an path from a json file",
		Action: actionGet,
		Flags: []cli.Flag{
			&flagFile,
			&flagPath,
			&flagDelimiter,
			&flagPretty,
		},
	}
}

func actionGet(c *cli.Context) {
	j, err := readInput(c.String("file"))
	if err != nil {
		errAndExit(err)
	}

	options := getOptions{
		json:      j,
		path:      c.String("path"),
		delimiter: getDelimiter(c.String("delimiter")),
	}

	j, err = get(options)
	if err != nil {
		errAndExit(err)
	}

	switch j.Data().(type) {
	case string:
		fmt.Printf("%s", j.Data())
	default:
		if pretty {
			fmt.Println(j.StringIndent("", "  "))
		} else {
			fmt.Println(j.String())
		}
	}
}

// get retrieves a path from a JSON structure.
// The path is specified in dotted notation:
// {"foo":{"bar":{"baz":"xyz"}}} = foo.bar.baz
// {"foo":{"bar":["a","b","c"]}} = foo.bar.2
// {"foo":{"bar":{"baz":"xyz"}}} = foo.bar.baz=xyz
// {"foo":{"bar":{"baz":"xyz"}}} = foo.*
// {"foo":{"bar":[{"a":"b"},{"c":"d"}]}} = foo.bar.*.a
func get(options getOptions) (*gabs.Container, error) {
	var err error
	var value string

	j := options.json
	pathPieces := strings.Split(options.path, options.delimiter)

	for i := 0; i < len(pathPieces); i++ {
		p := pathPieces[i]

		// Check if a value was specified
		kv := strings.Split(p, "=")
		if len(kv) > 1 {
			p = kv[0]
			if len(kv) > 2 {
				value = strings.Join(kv[1:], "=")
			} else {
				value = kv[1]
			}
		}

		debug.Printf("Path piece: %+v", p)
		debug.Printf("Path value: %+v", value)

		if _, ok := j.Data().([]interface{}); ok {
			debug.Printf("%+v is an array", j)
			if p == "*" {
				debug.Printf("glob used")
				children, err := j.Children()
				if err != nil {
					return nil, err
				}

				for _, c := range children {
					debug.Printf("Child: %+v", c)
					newPath := strings.Join(pathPieces[i+1:], ".")
					debug.Printf("New path: %+v", newPath)
					newOptions := getOptions{
						json:      c,
						path:      newPath,
						delimiter: options.delimiter,
					}
					if j, err := get(newOptions); err != nil {
						continue
					} else {
						return j, nil
					}
				}
			} else {
				j, err = checkArray(j, p)
				if err != nil {
					return nil, err
				}
			}
		} else {
			j = j.Path(p)
		}

		// if a value was given, see if the returned value matches
		// if the returned value is an array, check and see if the value exists in the array
		if value != "" {
			j, err = compareValues(j, value)
			if err != nil {
				return j, err
			}
		}
	}

	if j.Data() == nil {
		return nil, fmt.Errorf("No match found.")
	}

	return j, nil
}

// if the path piece is a number:
// check and see if it can be used as an array index
// if the current path is not an array, use it as a key
func checkArray(j *gabs.Container, pathPiece string) (*gabs.Container, error) {
	i, err := strconv.Atoi(pathPiece)
	if err != nil {
		return nil, fmt.Errorf("Non-numerical index.")
	}

	if i < 0 {
		return nil, fmt.Errorf("Array index out of bounds.")
	}

	if i > len(j.Data().([]interface{})) {
		return nil, fmt.Errorf("Array index out of bounds.")
	}

	return j.Index(i), nil
}

func compareValues(j *gabs.Container, value string) (*gabs.Container, error) {
	debug.Printf("[compareValues] j = %+v, v = %+v\n", j, value)
	var err error
	switch j.Data().(type) {
	case []interface{}:
		if value == "[]" {
			return j, nil
		}

		array := j.Data().([]interface{})
		for i, _ := range array {
			_, err = compareValues(j.Index(i), value)
			if err == nil {
				return j.Index(i), nil
			}
		}
	case map[string]interface{}:
		if value == "{}" {
			return j, nil
		}
	case string:
		if value == j.Data().(string) {
			return j, nil
		}
	default:
		if valueFloat64, err := strconv.ParseFloat(value, 64); err == nil {
			if jFloat64, ok := j.Data().(float64); ok {
				if valueFloat64 == jFloat64 {
					return j, nil
				}
			}
		}

		if value == j.String() {
			return j, nil
		}
	}

	return nil, fmt.Errorf("No match found.")
}
