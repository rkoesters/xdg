package ini

import (
	"strings"
	"testing"
)

const testParser = `
# This right here is a test file.
[Header 1]
key=value
	# This line is used to test extraneous white space.
	cat = dog 
[Header 2]
# This line tests for extra equal signs.
man = bear = pig
`

func TestParser(t *testing.T) {
	r := strings.NewReader(testFile)
	m, err := New(r)
	if err != nil {
		t.Error(err)
	}
	t.Log(m)

	if m.Get("Header 1", "key") != "value" {
		t.Error("basic usage")
	}
	if m.Get("Header 1", "cat") != "dog" {
		t.Error("whitespace")
	}
	if m.Get("Header 2", "man") != "bear = pig" {
		t.Error("equal signs")
	}
}

const testInvalid = `
# This example will have an invalid line.
[Header 1]
key=value
hello world!
`

func TestInvalid(t *testing.T) {
	r := strings.NewReader(testInvalid)
	_, err := New(r)
	if err.Error() != "invalid format" {
		t.Fail()
	}
}
