package jsonutil

import (
	"encoding/json"
	"gounchpad/pkg/common/chk"
)

// Marshal marshal
func Marshal(v interface{}) string {
	b, err := json.Marshal(v)
	chk.SE(err)
	return string(b)
}
