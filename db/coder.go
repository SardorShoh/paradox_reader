package db

import (
	"golang.org/x/exp/slices"
)

var letters = map[rune]string{
	171: "«",
	184: "ё",
	187: "»",
	185: "№",
}

// decrypt - matnni cp1251 koderovkadan utf8 koderovkasiga o'tkazadi
func decrypt(str string) string {
	var ret string
	for _, r := range str {
		if r >= 1025 {
			ret += string(r)
		} else {
			if slices.Contains([]int32{171, 184, 187, 185}, r) {
				ret += letters[r]
			} else {
				if r >= 161 {
					ret += string(r + 848)
				} else {
					ret += string(r)
				}
			}
		}
	}
	return ret
}

// Encrypt - matnni utf8 koderovkadan cp1251 koderovkasiga o'tkazadi
func Encrypt(str string) string {
	ret := ""
	for _, s := range str {
		r := rune(s)
		if r >= 1009 {
			if r == 8470 {
				ret += letters[8470]
			} else {
				ret += string([]rune{r - 848})
			}
		} else {
			ret += string(r)
		}
	}
	return ret
}
