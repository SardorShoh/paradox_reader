package db

var letters = map[rune]string{
	171:  "«",
	184:  "ё",
	187:  "»",
	8470: "№",
}

//decrypt - matnni cp1251 koderovkadan utf8 koderovkasiga o'tkazadi
func decrypt(str string) string {
	ret := ""
	for _, r := range []rune(str) {
		if r >= 161 {
			switch r {
			case 171:
				fallthrough
			case 187:
				fallthrough
			case 184:
				ret += letters[r]
			case 185:
				ret += letters[8470]
			default:
				ret += string([]rune{r + 848})
			}
		} else {
			ret += string(r)
		}
	}
	return ret
}

//Encrypt - matnni utf8 koderovkadan cp1251 koderovkasiga o'tkazadi
func Encrypt(str string) string {
	ret := ""
	for _, r := range []rune(str) {
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
