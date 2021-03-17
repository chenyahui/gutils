package slice

func DumpStrings(strSlice []string) []string {
	result := []string{}
	for _, item := range strSlice {
		result = append(result, item)
	}

	return result
}
