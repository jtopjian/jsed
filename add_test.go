package main

import (
	"testing"

	"github.com/jeffail/gabs"
)

func TestAddKeySimple(t *testing.T) {
	input := []byte(`{}`)
	correctResult := `{"foo":"bar"}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := addKey(j, "foo", "bar")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddKeyExisting(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{"baz":"qux","foo":"bar"}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := addKey(j, "baz", "qux")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddKeyNested(t *testing.T) {
	input := []byte(`{"foo":{"bar":"qux"}}`)
	correctResult := `{"foo":{"bar":"qux","omg":"wtf"}}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := addKey(j, "foo.omg", "wtf")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestAddKeyNumber(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{"foo":1}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := addKey(j, "foo", "1")
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

	result, err := addArray(j, "foo", []string{"a", "b", "c"})
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

	result, err := addArray(j, "baz", []string{"a", "b", "c"})
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

	result, err := addArray(j, "foo.omg", []string{"a", "b", "c"})
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

	result, err := addArray(j, "foo", []string{"1", "2", "3"})
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

	result, err := addArrayElement(j, "foo", "4", false)
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

	result, err := addArrayElement(j, "foo.1", "4", false)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}

	correctResult = `{"foo":[1,"a",3]}`
	result, err = addArrayElement(j, "foo.1", "a", false)
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

	result, err := addArrayElement(j, "foo", "3", true)
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}
