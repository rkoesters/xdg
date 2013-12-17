package desktop

import (
	"strings"
	"testing"
)

// specExample is the example desktop file given in the spec.
const specExample = `
[Desktop Entry]
Version=1.0
Type=Application
Name=Foo Viewer
Comment=The best viewer for Foo objects available!
TryExec=fooview
Exec=fooview %F
Icon=fooview
MimeType=image/x-foo;
Actions=Gallery;Create;

[Desktop Action Gallery]
Exec=fooview --gallery
Name=Browse Gallery

[Desktop Action Create]
Exec=fooview --create-new
Name=Create a new Foo!
Icon=fooview-new
`

func TestSpecExample(t *testing.T) {
	d, err := New(strings.NewReader(specExample))
	if err != nil {
		t.Error(err)
	}

	if d.Type() != Application {
		t.Error("Type")
	}

	arr := map[Key]string{
		Version: "1.0",
		Name:    "Foo Viewer",
		Comment: "The best viewer for Foo objects available!",
		TryExec: "fooview",
		Exec:    "fooview %F",
		Icon:    "fooview",
	}
	for k, exp := range arr {
		if d.String(k) != exp {
			t.Log("expected: " + exp)
			t.Log("actual:   " + d.String(k))
			t.Fail()
		}
	}
}

func TestActions(t *testing.T) {
	d, err := New(strings.NewReader(specExample))
	if err != nil {
		t.Error(err)
	}

	acts := d.Actions()
	if acts[0].Name != "Browse Gallery" {
		t.Fail()
	}
	if acts[0].Exec != "fooview --gallery" {
		t.Fail()
	}
	if acts[1].Name != "Create a new Foo!" {
		t.Fail()
	}
	if acts[1].Exec != "fooview --create-new" {
		t.Fail()
	}
	if acts[1].Icon != "fooview-new" {
		t.Fail()
	}
}
