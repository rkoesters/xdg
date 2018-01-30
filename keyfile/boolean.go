package keyfile

import (
	"strconv"
)

// Bool returns the value as a bool.
func (kf *KeyFile) Bool(g, k string) (bool, error) {
	return strconv.ParseBool(kf.Value(g, k))
}
