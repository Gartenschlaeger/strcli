package utilities

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetStandardInputString() string {
	fi, _ := os.Stdin.Stat()
	if fi.Size() > 0 && (fi.Mode()&os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		str := string(bytes)
		str = strings.Trim(str, "\n ")

		return str
	}

	return ""
}
