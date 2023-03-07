package utilities

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func DebugValue(v interface{}) {
	fmt.Printf("%+v\n", v)
}

func GetStandardInputString() string {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		str := string(bytes)
		str = strings.Trim(str, "\n\r ")

		return str
	}

	return ""
}

func ParseInt(s string, d int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		i = d
	}

	return i
}
