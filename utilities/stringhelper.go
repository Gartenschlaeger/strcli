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

func ReplaceString(s string, old string, new string, replaceAll bool, ignoreCase bool) string {
	if ignoreCase {
		oldValueLower := strings.ToLower(old)

		matchIndicees := make([]int, 0, strings.Count(strings.ToLower(s), oldValueLower))
		for i := 0; i < len(s); i++ {
			if strings.HasPrefix(strings.ToLower(s[i:]), oldValueLower) {
				matchIndicees = append(matchIndicees, i)
				i += len(oldValueLower)

				if !replaceAll {
					break
				}
			}
		}

		result := []rune(s)
		for i := len(matchIndicees) - 1; i >= 0; i-- {
			prefix := result[:matchIndicees[i]]
			newRunes := []rune(new)
			suffix := result[matchIndicees[i]+len(old):]
			result = append(prefix, append(newRunes, suffix...)...)
		}

		return string(result)
	} else {
		if replaceAll {
			return strings.ReplaceAll(s, old, new)
		} else {
			return strings.Replace(s, old, new, 1)
		}
	}
}

func StringEquals(s string, o string, ignoreCasing bool) bool {
	if ignoreCasing {
		return strings.EqualFold(s, o)
	}

	return s == o
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
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(prefix), length-len(s)) + s
}

func PadInt(i int, length int) string {
	return PadString(strconv.Itoa(1+i), '0', length)
}
