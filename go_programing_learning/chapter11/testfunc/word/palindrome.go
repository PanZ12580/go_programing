/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letter = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			continue
		}
		if unicode.IsLetter(r) {
			letter = append(letter, unicode.ToLower(r))
		} else {
			letter = append(letter, r)
		}
	}
	for i := 0; i < len(letter) / 2; i++ {
		if letter[i] != letter[len(letter) - i - 1] {
			return false
		}
	}
	return true
}
