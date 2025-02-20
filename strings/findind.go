package stringsalgs

func StrStr(haystack string, needle string) int {
	haystackIndex := 0
	for haystackIndex < len(haystack) {
		needleIndex := 0
		for needleIndex < len(needle) {
			if len(haystack) <= haystackIndex+needleIndex ||
				haystack[haystackIndex+needleIndex] != needle[needleIndex] {
				break
			}
			needleIndex++
		}
		if needleIndex == len(needle) {
			return haystackIndex
		}
		haystackIndex++
	}
	return -1
}
