package keyfile

import (
	"strings"
	"testing"
)

const testNumber = `
[Header 1]
zero=0
one=1
ten=10
negative-five=-5
pi=3.1415
`

func TestNumber(t *testing.T) {
	kf, err := New(strings.NewReader(testNumber))
	if err != nil {
		t.Error(err)
	}
	t.Log(kf)

	n, err := kf.Number("Header 1", "zero")
	if n != 0 || err != nil {
		t.Fail()
	}

	n, err = kf.Number("Header 1", "one")
	if n != 1 || err != nil {
		t.Fail()
	}

	n, err = kf.Number("Header 1", "ten")
	if n != 10 || err != nil {
		t.Fail()
	}

	n, err = kf.Number("Header 1", "negative-five")
	if n != -5 || err != nil {
		t.Fail()
	}

	n, err = kf.Number("Header 1", "pi")
	if n != 3.1415 || err != nil {
		t.Fail()
	}
}
