package utilities

import (
	"math/rand"
	"strconv"
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

func StringEquals(s string, o string, ignoreCasing bool) bool {
	source := s
	other := o

	if ignoreCasing {
		source = strings.ToLower(source)
		other = strings.ToLower(other)
	}

	return source == other
}

func ShiftString(s string, amount int, placeholder rune, repeat bool) string {
	l := len(s)

	if l == 0 && repeat {
		return ""
	}

	if l == 0 {
		l = amount
	}

	if !repeat {
		if amount < 0 {
			amount = MaxI(-l, amount)
		} else {
			amount = MinI(l, amount)
		}
	}

	buffer := make([]rune, l)
	for i := range buffer {
		buffer[i] = placeholder
	}

	for i := 0; i < l; i++ {
		n := ModI(i+amount, l)

		if repeat || (i+amount >= 0 && i+amount < l) {
			buffer[n] = rune(s[i])
		}
	}

	return string(buffer)
}

func PadString(s string, prefix rune, length int) string {
	pl := length - len(s)
	if pl <= 0 {
		return s
	}

	pr := make([]rune, pl)
	for i := 0; i < pl; i++ {
		pr[i] = prefix
	}

	return string(pr) + s
}

func PadInt(i int, length int) string {
	return PadString(strconv.Itoa(1+i), '0', length)
}
