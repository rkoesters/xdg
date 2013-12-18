package userdirs

import (
	"os/exec"
	"strings"
	"testing"
)

// XdgUserDir runs the xdg-user-dir command with the given argument.
func XdgUserDir(s string) string {
	out, err := exec.Command("xdg-user-dir", s).Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func doTest(t *testing.T, s1, s2 string) {
	t.Log(s1)
	t.Log(s2)
	if s1 != s2 {
		t.Fail()
	}
}

func TestDesktop(t *testing.T) {
	doTest(t, Desktop, XdgUserDir("DESKTOP"))
}

func TestDocuments(t *testing.T) {
	doTest(t, Documents, XdgUserDir("DOCUMENTS"))
}

func TestDownload(t *testing.T) {
	doTest(t, Download, XdgUserDir("DOWNLOAD"))
}

func TestMusic(t *testing.T) {
	doTest(t, Music, XdgUserDir("MUSIC"))
}

func TestPictures(t *testing.T) {
	doTest(t, Pictures, XdgUserDir("PICTURES"))
}

func TestPublicShare(t *testing.T) {
	doTest(t, PublicShare, XdgUserDir("PUBLICSHARE"))
}

func TestTemplates(t *testing.T) {
	doTest(t, Templates, XdgUserDir("TEMPLATES"))
}

func TestVideos(t *testing.T) {
	doTest(t, Videos, XdgUserDir("VIDEOS"))
}
