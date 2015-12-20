package main

import (
	"testing"
)

func runGetTest(t *testing.T, options getOptions, correctResult string) {
	result, err := get(options)
	if err != nil {
		t.Fatal(err)
	}

	if result.String() != correctResult {
		t.Fatalf("Wanted %s\nGot %s", correctResult, result)
	}
}

func TestGetKeySimple(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.name",
		delimiter: ".",
	}

	runGetTest(t, options, `"redis_mysql"`)
}

func TestGetKeyNestedArray(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.tags.0",
		delimiter: ".",
	}

	runGetTest(t, options, `"master"`)

	options = getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.checks.0.script",
		delimiter: ".",
	}

	runGetTest(t, options, `"/usr/local/bin/check_redis.py"`)
}

func TestGetKeyValueString(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.name=redis_mysql",
		delimiter: ".",
	}

	runGetTest(t, options, `"redis_mysql"`)
}

func TestGetKeyValueNumber(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.port=8000",
		delimiter: ".",
	}

	runGetTest(t, options, "8000")
}

func TestGetHash(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.foo={}",
		delimiter: ".",
	}

	runGetTest(t, options, `{"bar":"baz"}`)
}

func TestGetGlobArray(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service.checks.*.interval",
		delimiter: ".",
	}

	runGetTest(t, options, `"10s"`)
}

func TestGetGlobArrayValue(t *testing.T) {
	options := getOptions{
		json:      readTestFile(t, "test.json"),
		path:      "service..checks..*..script=/usr/local/bin/check_mysql.py",
		delimiter: "..",
	}

	runGetTest(t, options, `"/usr/local/bin/check_mysql.py"`)
}
