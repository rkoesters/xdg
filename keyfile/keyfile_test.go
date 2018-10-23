package keyfile

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
	kf, err := New(strings.NewReader(testParser))
	if err != nil {
		t.Error(err)
	}
	t.Log(kf)

	s, err := kf.String("Header 1", "key")
	if s != "value" || err != nil {
		t.Error("basic usage")
	}
	s, err = kf.String("Header 1", "cat")
	if s != "dog" || err != nil {
		t.Error("whitespace")
	}
	s, err = kf.String("Header 2", "man")
	if s != "bear = pig" || err != nil {
		t.Error("equal signs")
	}

	// Groups() and Keys() should always lead to valid values.
	for _, group := range kf.Groups() {
		for _, key := range kf.Keys(group) {
			if !kf.ValueExists(group, key) {
				t.Errorf("ValueExists == false for group='%v' key='%v'", group, key)
			}
		}
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
list2=man\;bear;pig\r;
`

func TestBool(t *testing.T) {
	kf, err := New(strings.NewReader(testFormat))
	if err != nil {
		t.Error(err)
	}
	t.Log(kf)

	b, err := kf.Bool("Header 1", "yes")
	if b != true || err != nil {
		t.Fail()
	}
	b, err = kf.Bool("Header 1", "no")
	if b != false || err != nil {
		t.Fail()
	}
}

func TestList(t *testing.T) {
	kf, err := New(strings.NewReader(testFormat))
	if err != nil {
		t.Error(err)
	}
	expect := []string{"man", "bear", "pig"}
	actual, err := kf.ValueList("Header 1", "list")
	if err != nil {
		t.Error(err)
	}
	t.Log(expect)
	t.Log(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Fail()
	}
	expect = []string{"man;bear", "pig\\r"}
	actual, err = kf.ValueList("Header 1", "list2")
	if err != nil {
		t.Error(err)
	}
	t.Log(expect)
	t.Log(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Fail()
	}
}
