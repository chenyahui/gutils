package slice

func AsStrings(sli []interface{}) []string {
	var result []string
	for _, item := range sli {
		result = append(result, item.(string))
	}
	return result
}
