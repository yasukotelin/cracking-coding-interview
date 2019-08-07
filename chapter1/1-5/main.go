package main

func main() {
}

func canConvert(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var hash map[rune]int
	if s1 > s2 {
		hash = makeHash(s2)
		diffCnt := 0
		for _, r := range s1 {
			_, ok := hash[r]
			if !ok {
				diffCnt++
				if diffCnt > 1 {
					return false
				}
			}
		}
		return true
	} else {
		hash = makeHash(s1)
		diffCnt := 0
		for _, r := range s2 {
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
}

func makeHash(s string) map[rune]int {
	hash := make(map[rune]int)
	for _, r := range s {
		hash[r] = 1
	}
	return hash
}
