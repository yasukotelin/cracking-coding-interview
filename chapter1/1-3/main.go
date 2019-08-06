package main

func main() {
}

const whiteSpace = rune(' ')

// whiteSpacesToPer20 は文字列s中の空白をすべて%20に置換して新しい文字列を返却する
func whiteSpacesToPer20(s string) string {
	whiteSpacesNum := countWhiteSpaces(s)

	runes := make([]rune, 0, len(s)+whiteSpacesNum*2)
	for _, r := range s {
		if r == whiteSpace {
			runes = append(runes, []rune{'%', '2', '0'}...)
		} else {
			runes = append(runes, r)
		}
	}
	return string(runes)
}

func countWhiteSpaces(s string) int {
	var count int
	for _, r := range s {
		if r == whiteSpace {
			count++
		}
	}
	return count
}
