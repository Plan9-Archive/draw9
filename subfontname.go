package draw9

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
 * Default version: convert to file name
 */

func subfontname(cfname, fname string, maxdepth int) string {
	t := cfname
	if cfname == "*default*" {
		return t
	}
	if strings.HasPrefix(t, ".") {
		fdir := filepath.Dir(fname)
		ffile := filepath.Base(fname)

		ffile = strings.Replace(ffile, "unicode.", "", 1)
		ffile = strings.Replace(ffile, ".font", "", 1)
		
		t = filepath.Join(fdir, ffile) + t
	}
	if !strings.HasPrefix(t, "/") {
		dir := filepath.Dir(fname)
		t = filepath.Join(dir, t)
	}
	if maxdepth > 8 {
		maxdepth = 8
	}
	for i := 3; i >= 0; i-- {
		if 1<<uint(i) > maxdepth {
			continue
		}
		// try i-bit grey
		tmp2 := fmt.Sprintf("%s.%d", t, i)
		if _, err := os.Stat(tmp2); err == nil {
			return tmp2
		}
	}

	// try default
	if strings.HasPrefix(t, "/mnt/font/") {
		return t
	}
	if _, err := os.Stat(t); err == nil {
		return t
	}

	return ""
}
