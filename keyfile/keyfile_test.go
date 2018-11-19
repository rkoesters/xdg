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

	s := kf.Value("Header 1", "key")
	if s != "value" {
		t.Error("basic usage")
	}
	s = kf.Value("Header 1", "cat")
	if s != "dog" {
		t.Error("whitespace")
	}
	s = kf.Value("Header 2", "man")
	if s != "bear = pig" {
		t.Error("equal signs")
	}

	// Groups() and Keys() should always lead to valid values.
	for _, group := range kf.Groups() {
		if !kf.GroupExists(group) {
			t.Errorf("GroupExists == false for group='%v'", group)
		}

		for _, key := range kf.Keys(group) {
			if !kf.KeyExists(group, key) {
				t.Errorf("KeyExists == false for group='%v' key='%v'", group, key)
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
	if err != ErrInvalid {
		t.Fail()
	}
}

const testList = `
# This tests that the formatting functions work.
[Header 1]
list=man;bear;pig;
list2=man\;bear;pig\r;
list3=man;bear;pig
`

func TestList(t *testing.T) {
	kf, err := New(strings.NewReader(testList))
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

	// We currently support lists not ending in ';'. If that
	// behavior changes the following will need to be updated
	// accordingly.
	expect = []string{"man", "bear", "pig"}
	actual, err = kf.ValueList("Header 1", "list3")
	if err != nil {
		t.Error(err)
	}
	t.Log(expect)
	t.Log(actual)
	if !reflect.DeepEqual(actual, expect) {
		t.Fail()
	}
}

func TestExists(t *testing.T) {
	kf, err := New(strings.NewReader(""))
	if err != nil {
		t.Error(err)
	}

	if !kf.GroupExists("") {
		t.Fail()
	}

	if kf.GroupExists("group") {
		t.Fail()
	}

	if kf.KeyExists("", "") {
		t.Fail()
	}

	if kf.KeyExists("group", "") {
		t.Fail()
	}

	if kf.KeyExists("", "key") {
		t.Fail()
	}

	if kf.KeyExists("group", "key") {
		t.Fail()
	}
}

const testListUnexpectedEndOfString = `
list=asdf;asdf;\`

func TestListUnexpectedEndOfString(t *testing.T) {
	kf, err := New(strings.NewReader(testListUnexpectedEndOfString))
	if err != nil {
		t.FailNow()
	}

	_, err = kf.ValueList("", "list")
	if err != ErrUnexpectedEndOfString {
		t.Fail()
	}
}
