package main

func main() {
}

func canConvert(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var shorter string
	var longer string
	if len(s1) > len(s2) {
		shorter = s2
		longer = s1
	} else {
		shorter = s1
		longer = s2
	}

	hash := makeHash(longer)
	diffCnt := 0
	for _, r := range shorter {
		_, ok := hash[r]
		if !ok {
			diffCnt++
			if diffCnt > 1 {
				return false
			}
		}
	}
	return true
}

func makeHash(s string) map[rune]int {
	hash := make(map[rune]int)
	for _, r := range s {
		hash[r] = 1
	}
	return hash
}
