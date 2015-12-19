package main

import (
	"testing"

	"github.com/jeffail/gabs"
)

func TestGetKeySimple(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `"bar"`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := get(j, "foo")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestGetKeyNested(t *testing.T) {
	input := []byte(`{"foo":{"bar":"baz"}}`)
	correctResult := `"baz"`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := get(j, "foo.bar")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestGetKeyNestedArray(t *testing.T) {
	input := []byte(`{"foo":{"bar":[8,7,6,{"omg":"wtf"}]}}`)
	correctResult := `7`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := get(j, "foo.bar.1")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}

	correctResult = `"wtf"`

	result, err = get(j, "foo.bar.3.omg")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestGetArrayContains(t *testing.T) {
	input := []byte(`{"foo":{"bar":[8,7,6,{"omg":"wtf"}]}}`)
	correctResult := `7`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := get(j, "foo.bar")
	if err != nil {
		t.Error(err)
	}

	result, err = contains(result, "7")

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}
