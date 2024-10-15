package common

import (
	"os"
	"path/filepath"
)

func HintResouce(p string) (string, bool) {
	for _, d := range []string{".", "..", "../.."} {
		d = filepath.Join(d, p)
		if _, err := os.Stat(d); err == nil {
			return d, true
		}
	}
	return p, false
}
