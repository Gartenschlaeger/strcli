package utilities

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetStandardInputString() string {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		str := string(bytes)
		str = strings.Trim(str, "\n ")

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

func ClampI(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}

	return v
}

func ClampStringPartion(input string, index int, length int) (startIndex int, endIndex int) {
	tl := len(input)

	if index < 0 {
		// x characters from end
		startIndex = ClampI(tl+index, 0, tl)
		endIndex = ClampI(startIndex+length, startIndex, tl)
	} else {
		// x characters from start
		startIndex = ClampI(index, 0, tl)
		endIndex = ClampI(startIndex+length, startIndex, tl)
	}

	return startIndex, endIndex
}
