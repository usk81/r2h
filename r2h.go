package r2h

import (
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

const alphaRegexString = "^[a-zA-Z]+$"

var alphaRegex = regexp.MustCompile(alphaRegexString)

func substr(s string, start int, end ...int) string {
	if len(end) > 0 {
		return string([]rune(s)[start:end[0]])
	}
	return string([]rune(s)[start:])
}

func charAt(s string, index int) string {
	return string([]rune(s)[index])
}

func isAlpha(s string) bool {
	return alphaRegex.MatchString(s)
}

func convertLetter(us string) (kana string, length int) {
	min := int(math.Min(3, float64(utf8.RuneCountInString(us))))
	for min > 0 {
		l := substr(us, 0, min)
		if kana, ok := dict[l]; ok {
			return kana, utf8.RuneCountInString(l)
		}
		min--
	}
	return
}

func Convert(s string) (result string, isCompleted bool) {
	isCompleted = true
	for utf8.RuneCountInString(s) > 0 {
		us := strings.ToUpper(s)
		kana, l := convertLetter(us)
		if kana == "" {
			if utf8.RuneCountInString(us) >= 3 {
				head := charAt(us, 0)
				next := charAt(us, 1)
				if isAlpha(head) && head == next {
					kana, l = convertLetter(substr(us, 0))
					if kana == "" {
						kana = dict["LTU"]
						l = 1
					}
				}
			}
			if kana == "" {
				kana = charAt(s, 0)
				l = 1
				if kana != " " {
					isCompleted = false
				}
			}
		}
		result += kana
		s = substr(s, l)
	}
	return
}
