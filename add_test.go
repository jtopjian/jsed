package main

import (
	"testing"

	"github.com/Jeffail/gabs/v2"
)

func TestAddObjectSimple(t *testing.T) {
	input := []byte(`{}`)
	correctResult := `{"foo":"bar"}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addObjectOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		values:    []string{"bar"},
	}

	result, err := addObject(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddObjectWithObject(t *testing.T) {
	input := []byte(`{"foo":{}}`)
	correctResult := `{"foo":{"bar":"baz"}}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addObjectOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		keys:      []string{"bar"},
		values:    []string{"baz"},
	}

	result, err := addObject(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddObjectExisting(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{"baz":"qux","foo":"bar"}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addObjectOptions{
		json:      j,
		path:      "baz",
		delimiter: ".",
		values:    []string{"qux"},
	}

	result, err := addObject(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddObjectNested(t *testing.T) {
	input := []byte(`{"foo":{"bar":"qux"}}`)
	correctResult := `{"foo":{"bar":"qux","omg":"wtf"}}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addObjectOptions{
		json:      j,
		path:      "foo.omg",
		delimiter: ".",
		values:    []string{"wtf"},
	}

	result, err := addObject(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddObjectNumber(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{"foo":1}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addObjectOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		values:    []string{"1"},
	}

	result, err := addObject(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddObjectsAndValues(t *testing.T) {
	input := []byte(`{}`)
	correctResult := `{"baz":"qux","foo":"bar"}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addObjectOptions{
		json:      j,
		delimiter: ".",
		keys:      []string{"foo", "baz"},
		values:    []string{"bar", "qux"},
	}

	result, err := addObject(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArraySimple(t *testing.T) {
	input := []byte(`{}`)
	correctResult := `{"foo":["a","b","c"]}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		values:    []string{"a", "b", "c"},
	}

	result, err := addArray(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArrayExisting(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{"baz":["a","b","c"],"foo":"bar"}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayOptions{
		json:      j,
		path:      "baz",
		delimiter: ".",
		values:    []string{"a", "b", "c"},
	}

	result, err := addArray(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArrayNested(t *testing.T) {
	input := []byte(`{"foo":{"bar":"qux"}}`)
	correctResult := `{"foo":{"bar":"qux","omg":["a","b","c"]}}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayOptions{
		json:      j,
		path:      "foo.omg",
		delimiter: ".",
		values:    []string{"a", "b", "c"},
	}

	result, err := addArray(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArrayNumber(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{"foo":[1,2,3]}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		values:    []string{"1", "2", "3"},
	}

	result, err := addArray(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArrayElementSimple(t *testing.T) {
	input := []byte(`{"foo":[1,2,3]}`)
	correctResult := `{"foo":[1,2,3,4]}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayElementOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		value:     "4",
		exists:    false,
	}

	result, err := addArrayElement(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArrayElementPosition(t *testing.T) {
	input := []byte(`{"foo":[1,2,3]}`)
	correctResult := `{"foo":[1,4,3]}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayElementOptions{
		json:      j,
		path:      "foo.1",
		delimiter: ".",
		value:     "4",
		exists:    false,
	}

	result, err := addArrayElement(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}

	correctResult = `{"foo":[1,"a",3]}`
	options = addArrayElementOptions{
		json:      j,
		path:      "foo.1",
		delimiter: ".",
		value:     "a",
		exists:    false,
	}

	result, err = addArrayElement(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddArrayElementContains(t *testing.T) {
	input := []byte(`{"foo":[1,2,3]}`)
	correctResult := `{"foo":[1,2,3]}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	options := addArrayElementOptions{
		json:      j,
		path:      "foo",
		delimiter: ".",
		value:     "3",
		exists:    true,
	}

	result, err := addArrayElement(options)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}
