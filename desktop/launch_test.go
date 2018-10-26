package desktop

import (
	"bytes"
	"testing"
)

func TestExpandExec(t *testing.T) {
	de := &Entry{
		Name: "Test",
		Exec: "echo hello,   world  %i %U",
		Icon: "icon",
	}

	expansion, err := de.expandExec("file.txt", "https://example.com")
	if err != nil {
		t.Errorf("expandExec returned err=%v", err)
	}

	t.Log(argvsToString(expansion))

	de = &Entry{
		Name: "Test",
		Exec: "echo hello,   world  %i %u",
		Icon: "icon",
	}

	expansion, err = de.expandExec("file.txt", "https://example.com")
	if err != nil {
		t.Errorf("expandExec returned err=%v", err)
	}

	t.Log(argvsToString(expansion))

	de = &Entry{
		Name:     "Test",
		Exec:     "echo hello,   world  %i %u",
		Icon:     "icon",
		Terminal: true,
	}

	expansion, err = de.expandExec("file.txt", "https://example.com")
	if err != nil {
		t.Errorf("expandExec returned err=%v", err)
	}

	t.Log(argvsToString(expansion))
}

func argvsToString(argvs [][]string) string {
	var buf1 bytes.Buffer
	for _, argv := range argvs {
		if buf1.Len() != 0 {
			buf1.WriteString(", ")
		} else {
			buf1.WriteRune('[')
		}

		var buf2 bytes.Buffer
		for _, arg := range argv {
			if buf2.Len() != 0 {
				buf2.WriteString(", ")
			} else {
				buf2.WriteRune('[')
			}
			buf2.WriteRune('\'')
			buf2.WriteString(arg)
			buf2.WriteRune('\'')
		}
		buf2.WriteRune(']')

		buf1.WriteString(buf2.String())
	}
	buf1.WriteRune(']')

	return buf1.String()
}
