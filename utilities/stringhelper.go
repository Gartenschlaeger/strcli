package utilities

import (
	"math/rand"
	"strings"
	"time"
)

func SplitString(s string, separator string, removeEmptyFields bool) []string {
	if removeEmptyFields {
		return strings.FieldsFunc(s, func(r rune) bool {
			return strings.ContainsRune(separator, r)
		})
	} else {
		return strings.Split(s, separator)
	}
}

func SubString(s string, index int, length int) string {
	return s[index : index+length]
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

func ShuffleString(s string) string {
	rand.Seed(time.Now().UnixNano())

	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func replaceStringPortion(input string, index int, oldValue string) string {
	if index+len(oldValue) > len(input) {
		return input
	}

	return input[:index] + oldValue + input[index+len(oldValue):]
}

func ReplaceString(s string, old string, new string, replaceAll bool, ignoreCase bool) string {
	if !ignoreCase {
		if replaceAll {
			return strings.ReplaceAll(s, old, new)
		} else {
			return strings.Replace(s, old, new, 1)
		}
	} else {
		inputLower := strings.ToLower(s)
		oldValueLower := strings.ToLower(old)

		findIndicees := []int{}
		for i := 0; i < len(inputLower); i++ {
			if strings.HasPrefix(inputLower[i:], oldValueLower) {
				findIndicees = append(findIndicees, i)
				i += len(oldValueLower)

				if !replaceAll {
					break
				}
			}
		}

		result := s
		for _, v := range findIndicees {
			result = replaceStringPortion(result, v, new)
		}

		return result
	}
}
