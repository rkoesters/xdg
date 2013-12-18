package trash

import (
	"strings"
	"testing"
)

const trashinfo = `[Trash Info]
Path=/home/user/file.txt
DeletionDate=2006-01-02T15:04:05
`

func TestInfo(t *testing.T) {
	r := strings.NewReader(trashinfo)
	info, err := NewInfo(r)
	if err != nil {
		t.Error(err)
	}
	t.Log(trashinfo)
	t.Log(info)

	if info.String() != trashinfo {
		t.Fail()
	}
}
