package str

import (
	"strconv"
	"strings"
)

func ReplaceSubString(old string, from int, to int, new string) string {
	return old[0:from] + new + old[to:]
}

// IsBlank check a string is whitespace
func IsBlank(val string) bool {
	return len(strings.TrimSpace(val)) == 0
}

// IsNumber check a string is number
func IsNumber(val string) bool {
	_, err := strconv.ParseInt(val, 10, 64)
	return err == nil
}
