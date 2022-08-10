package r2h

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	alphaRegexString  = "^[a-zA-Z]+$"
	numberRegexString = "^[0-9０-９]+$"
)

var (
	alphaRegex  = regexp.MustCompile(alphaRegexString)
	numberRegex = regexp.MustCompile(numberRegexString)
)

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

func isNumber(s string) bool {
	return numberRegex.MatchString(s)
}

func isHiragana(s string) bool {
	for _, r := range s {
		if b := isHiraganaRune(r); !b {
			return false
		}
	}
	return true
}

func isHiraganaRune(r rune) bool {
	return unicode.In(r, unicode.Hiragana)
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

func convertWords(s string, strict bool) (result string, isCompleted bool, err error) {
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
				if kana != " " && !isNumber(kana) && !isHiragana(kana) {
					isCompleted = false
					if strict {
						return "", false, fmt.Errorf("%s is not romaji", kana)
					}
				}
			}
		}
		result += kana
		s = substr(s, l)
	}
	return
}

// Convert romaji to hiragana
func Convert(s string) (result string, isCompleted bool) {
	result, isCompleted, _ = convertWords(s, false)
	return
}

// ConvertStrict converts romaji to hiragana. If non-romaji letter are mixed, an error will occur.
func ConvertStrict(s string) (result string, err error) {
	result, _, err = convertWords(s, true)
	return
}
