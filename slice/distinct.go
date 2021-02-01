package slice

func DistinctStrings(stringSlice []string) []string {
	var result []string
	tempMap := map[string]bool{}
	for _, e := range stringSlice {
		l := len(tempMap)
		tempMap[e] = true
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
