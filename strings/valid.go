package stringsalgs

func IsValid(s string) bool {
	countOfOpened := map[rune]int{
		'(': 0,
		'{': 0,
		'[': 0,
	}

	countOfClosed := map[rune]int{
		')': 0,
		'}': 0,
		']': 0,
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	lastOpened := make([]rune, 0)
	for _, symbool := range s {
		if _, open := countOfOpened[symbool]; open {
			countOfOpened[symbool] += 1
			lastOpened = append(lastOpened, symbool)
		}

		if _, close := countOfClosed[symbool]; close {
			if len(lastOpened) == 0 ||
				symbool != pairs[rune(lastOpened[len(lastOpened)-1])] {
				return false
			}
			lastOpened = lastOpened[:len(lastOpened)-1]
			countOfClosed[symbool] += 1
		}

		for open, close := range pairs {
			if countOfOpened[open]-countOfClosed[close] < 0 {
				return false
			}
		}
	}

	for open, close := range pairs {
		if countOfOpened[open]-countOfClosed[close] > 0 {
			return false
		}
	}

	return true
}
