package slice

func IndexOfStrings(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}

	return -1
}

func IndexOfInts(slice []int, val int) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}

	return -1
}

func ContainsString(slice []string, val string) bool {
	return IndexOfStrings(slice, val) >= 0
}

func ContainsInt(slice []int, val int) bool {
	return IndexOfInts(slice, val) >= 0
}
