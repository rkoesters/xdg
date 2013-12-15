package ini

import (
	"testing"
	"strings"
)

const testFile = `
# This right here is a test file.
[Header 1]
key=value
	# This line is used to test extraneous white space.
	cat = dog 
[Header 2]
# This line tests for extra equal signs.
man = bear = pig
`

func TestNew(t *testing.T) {
	r := strings.NewReader(testFile)
	file, err := New(r)
	if err != nil {
		t.Error(err)
	}
	t.Log(file)

	if file["Header 1"]["key"] != "value" {
		t.Error("basic usage")
	}
	if file["Header 1"]["cat"] != "dog" {
		t.Error("whitespace")
	}
	if file["Header 2"]["man"] != "bear = pig" {
		t.Error("equal signs")
	}
}

const testFileInvalid = `
# This example will have an invalid line.
[Header 1]
key=value
hello world!
`

func TestInvalid(t *testing.T) {
	r := strings.NewReader(testFileInvalid)
	_, err := New(r)
	if err.Error() != "invalid format" {
		t.Fail()
	}
}
