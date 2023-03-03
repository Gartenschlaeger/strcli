package utilities

import (
	"io/ioutil"
	"os"
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

func ClampI(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}

	return v
}

func ClampGetTextRange(text string, index int, length int) (int, int) {
	textLength := len(text)

	var startIndex, endIndex int

	if index < 0 {
		startIndex = ClampI(textLength+index, 0, textLength)
		endIndex = ClampI(startIndex+length, startIndex, textLength)
	} else {
		startIndex = ClampI(index, 0, textLength)
		endIndex = ClampI(index+length, startIndex, textLength)
	}

	return startIndex, endIndex
}
