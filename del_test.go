package main

import (
	"testing"

	"github.com/jeffail/gabs"
)

func TestDelKeySimple(t *testing.T) {
	input := []byte(`{"foo":"bar"}`)
	correctResult := `{}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := delKey(j, "foo")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestDelKeyNested(t *testing.T) {
	input := []byte(`{"foo":{"bar":"baz"}}`)
	correctResult := `{"foo":{}}`

	j, err := gabs.ParseJSON(input)
	if err != nil {
		t.Error(err)
	}

	result, err := delKey(j, "foo.bar")
	if err != nil {
		t.Error(err)
	}

	if result.String() != correctResult {
		t.Errorf("Wanted %s\nGot %s", correctResult, result)
	}
}
