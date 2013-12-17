package ini

import (
	"reflect"
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
	m, err := New(strings.NewReader(testParser))
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
	_, err := New(strings.NewReader(testInvalid))
	if err.Error() != "invalid format" {
		t.Fail()
	}
}

const testFormat = `
# This tests that the formatting functions work.
[Header 1]
yes=true
no=false

list=man;bear;pig;
`

func TestBool(t *testing.T) {
	m, err := New(strings.NewReader(testFormat))
	if err != nil {
		t.Error(err)
	}
	t.Log(m)

	if m.Bool("Header 1", "yes") != true {
		t.Fail()
	}
	if m.Bool("Header 1", "no") != false {
		t.Fail()
	}
}

func TestList(t *testing.T) {
	m, err := New(strings.NewReader(testFormat))
	if err != nil {
		t.Error(err)
	}
	expect := []string{"man", "bear", "pig"}
	actual := m.List("Header 1", "list")
	t.Log(expect)
	t.Log(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Fail()
	}
}
