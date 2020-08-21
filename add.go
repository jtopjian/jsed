package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/urfave/cli/v2"
)

var cmdAdd cli.Command

type addObjectOptions struct {
	json      *gabs.Container
	path      string
	delimiter string
	keys      []string
	values    []string
}

type addArrayOptions struct {
	json      *gabs.Container
	path      string
	delimiter string
	values    []string
}

type addArrayElementOptions struct {
	json      *gabs.Container
	path      string
	delimiter string
	key       string
	value     string
	exists    bool
}

func init() {
	cmdAdd = cli.Command{
		Name:  "add",
		Usage: "add data to a json file",
		Subcommands: []*cli.Command{
			{
				Name:   "object",
				Usage:  "Add a new key/value pair.",
				Action: actionAddObject,
				Flags: []cli.Flag{
					&flagFile,
					&flagPath,
					&flagDelimiter,
					&flagMultiKey,
					&flagMultiValue,
					&flagPretty,
				},
			},
			{
				Name:   "array",
				Usage:  "Add a new array",
				Action: actionAddArray,
				Flags: []cli.Flag{
					&flagFile,
					&flagPath,
					&flagDelimiter,
					&flagMultiValue,
					&flagPretty,
				},
			},
			{
				Name:   "element",
				Usage:  "Add a new array element",
				Action: actionAddArrayElement,
				Flags: []cli.Flag{
					&flagFile,
					&flagPath,
					&flagDelimiter,
					&flagMultiKey,
					&flagMultiValue,
					&flagExists,
					&flagPretty,
				},
			},
		},
	}
}

func actionAddObject(c *cli.Context) error {
	j, err := readInput(c.String("file"))
	if err != nil {
		return err
	}

	options := addObjectOptions{
		json:      j,
		path:      c.String("path"),
		delimiter: getDelimiter(c.String("delimiter")),
		keys:      c.StringSlice("key"),
		values:    c.StringSlice("value"),
	}

	j, err = addObject(options)
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

func actionAddArray(c *cli.Context) error {
	j, err := readInput(c.String("file"))
	if err != nil {
		return err
	}

	options := addArrayOptions{
		json:      j,
		path:      c.String("path"),
		delimiter: getDelimiter(c.String("delimiter")),
		values:    c.StringSlice("value"),
	}

	j, err = addArray(options)
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

func actionAddArrayElement(c *cli.Context) error {
	j, err := readInput(c.String("file"))
	if err != nil {
		return err
	}

	options := addArrayElementOptions{
		json:      j,
		path:      c.String("path"),
		delimiter: getDelimiter(c.String("delimiter")),
		key:       c.String("key"),
		value:     c.String("value"),
		exists:    c.Bool("exists"),
	}

	j, err = addArrayElement(options)
	if err != nil {
		return err
	}
	return nil
}

func addObject(options addObjectOptions) (*gabs.Container, error) {
	var err error
	var isArray bool
	kv := make(map[string]interface{})
	j := options.json

	// at least one value is required
	if len(options.values) < 1 {
		return nil, fmt.Errorf("A value is required")
	}

	splitPath := []string{}
	if options.path != "" {
		splitPath = strings.Split(options.path, options.delimiter)
	}

	// Check and see if the destination path is an array
	getOpt := getOptions{
		json:      options.json,
		path:      options.path,
		delimiter: options.delimiter,
	}
	p, err := get(getOpt)
	if err == nil {
		if _, ok := p.Data().([]interface{}); ok {
			isArray = true
		}
	}

	// if no keys were specified, only allow a single value
	for i, v := range options.values {
		if isArray {
			if len(options.keys) > i {
				if x, e := strconv.Atoi(v); e == nil {
					kv[options.keys[i]] = x
				} else if b, e := strconv.ParseBool(v); e == nil {
					kv[options.keys[i]] = b
				} else if v == "{}" {
					kv[options.keys[i]] = map[string]interface{}{}
				} else if v == "[]" {
					kv[options.keys[i]] = []interface{}{}
				} else {
					kv[options.keys[i]] = v
				}
			}
		} else {
			sp := []string{}
			if len(options.keys) > i {
				sp = append(splitPath, options.keys[i])
			} else {
				if i > 1 {
					break
				}
				sp = splitPath
			}

			if i, e := strconv.Atoi(v); e == nil {
				_, err = j.Set(i, sp...)
			} else if b, e := strconv.ParseBool(v); e == nil {
				_, err = j.Set(b, sp...)
			} else if v == "{}" {
				_, err = j.Object(sp...)
			} else if v == "[]" {
				_, err = j.Array(sp...)
			} else {
				_, err = j.Set(v, sp...)
			}
		}
	}

	if isArray {
		err := j.ArrayAppend(kv, splitPath...)
		if err != nil {
			return nil, err
		}
	}

	return j, err
}

func addArray(options addArrayOptions) (*gabs.Container, error) {
	var err error
	j := options.json
	splitPath := strings.Split(options.path, options.delimiter)
	_, err = j.Array(splitPath...)
	if err != nil {
		return nil, err
	}

	for _, v := range options.values {
		if i, e := strconv.Atoi(v); e == nil {
			err = j.ArrayAppend(i, splitPath...)
		} else {
			err = j.ArrayAppend(v, splitPath...)
		}
	}

	return j, err
}

func addArrayElement(options addArrayElementOptions) (*gabs.Container, error) {
	var err error

	j := options.json
	splitPath := strings.Split(options.path, options.delimiter)

	if options.exists {
		var searchPath string
		if options.key != "" {
			searchPath = fmt.Sprintf("%s%s*%s%s=%s", options.path, options.delimiter, options.delimiter, options.key, options.value)
		} else {
			searchPath = fmt.Sprintf("%s=%s", options.path, options.value)
		}

		getOptions := getOptions{
			json:      j,
			path:      searchPath,
			delimiter: options.delimiter,
		}

		if _, err = get(getOptions); err == nil {
			return j, nil
		}
	}

	pathPieces := strings.Split(options.path, options.delimiter)
	if index, err := strconv.Atoi(pathPieces[len(pathPieces)-1]); err == nil {
		if index >= 0 {
			arrPath := strings.Join(pathPieces[:len(pathPieces)-1], options.delimiter)
			arr := j.Path(arrPath)

			if i, err := strconv.Atoi(options.value); err == nil {
				_, err = arr.SetIndex(i, index)
			} else {
				_, err = arr.SetIndex(options.value, index)
			}
		}
	} else {
		if i, err := strconv.Atoi(options.value); err == nil {
			err = j.ArrayAppend(i, splitPath...)
		} else {
			err = j.ArrayAppend(options.value, splitPath...)
		}
	}

	return j, err
}
