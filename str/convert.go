package str

import "strconv"

func FromInt64(val int64) string {
	return strconv.FormatInt(val, 10)
}

func ParseInt(str string, defaultVal int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultVal
	}
	return val
}
