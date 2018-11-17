package keyfile

import (
	"strings"
	"testing"
)

const testBool = `
[Header 1]
yes=true
no=false
`

func TestBool(t *testing.T) {
	kf, err := New(strings.NewReader(testBool))
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
