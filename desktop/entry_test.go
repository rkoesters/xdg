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

	if d.Type != Application {
		t.Error("Type")
	}

	arr := map[string]string{
		d.Version: "1.0",
		d.Name:    "Foo Viewer",
		d.Comment: "The best viewer for Foo objects available!",
		d.TryExec: "fooview",
		d.Exec:    "fooview %F",
		d.Icon:    "fooview",
	}
	for act, exp := range arr {
		if act != exp {
			t.Log("expected: " + exp)
			t.Log("actual:   " + act)
			t.Fail()
		}
	}
}

func TestActions(t *testing.T) {
	d, err := New(strings.NewReader(specExample))
	if err != nil {
		t.Error(err)
	}

	a := d.Actions
	if a[0].Name != "Browse Gallery" {
		t.Fail()
	}
	if a[0].Exec != "fooview --gallery" {
		t.Fail()
	}
	if a[1].Name != "Create a new Foo!" {
		t.Fail()
	}
	if a[1].Exec != "fooview --create-new" {
		t.Fail()
	}
	if a[1].Icon != "fooview-new" {
		t.Fail()
	}
}

const fullExample = `
# This example will use all the keys.
[Desktop Entry]
Type=Application
Version=1.0

Name=fullExample
GenericName=Desktop Entry Test
Comment=This is a test comment.
Icon=test-icon

NoDisplay=true
Hidden=true
OnlyShowIn=Unity;Gnome;
NotShowIn=KDE;xfce;

DBusActivatable=true
TryExec=echo
Exec=echo %F
Path=/
Terminal=true

Actions=NewFile;TacoSalad;
MimeType=text/plain;text/markdown;
Categories=Tests;Golang;
Keywords=full;test;golang;xdg;desktop;

StartupNotify=true
StartupWMClass=test

X-Unity-IconBackgroundColor=#000000
X-Gnome-Something=foo
X-KDE-plasma=workspaces

[Desktop Action NewFile]
Name=New File
Exec=echo
Icon=file

[Desktop Action TacoSalad]
Name=Taco Salad
Exec=echo Taco Salad
Icon=taco
`

func TestPrintIt(t *testing.T) {
	d, err := New(strings.NewReader(fullExample))
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", d)
}

const linkExample = `
[Desktop Entry]
Version=1.0
Type=Link
Name=Go
Name[en_US]=Go
URL=http://www.golang.org/
`

func TestLinkExample(t *testing.T) {
	d, err := New(strings.NewReader(linkExample))
	if err != nil {
		t.Error(err)
	}

	if d.Type != Link {
		t.Error("Type")
	}

	arr := map[string]string{
		d.Version: "1.0",
		d.Name:    "Go",
		d.URL:     "http://www.golang.org/",
	}
	for act, exp := range arr {
		if act != exp {
			t.Log("expected: " + exp)
			t.Log("actual:   " + act)
			t.Fail()
		}
	}
}
