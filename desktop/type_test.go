package desktop

import (
	"testing"
)

func TestType(t *testing.T) {
	m := map[Type]string{
		Application: "Application",
		Link: "Link",
		Directory: "Directory",
		Unknown: "Unknown",
	}

	for k, v := range m {
		if k != ParseType(v) {
			t.Fail()
		}
		if k.String() != v {
			t.Fail()
		}
	}
}
