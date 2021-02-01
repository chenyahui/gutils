package str

func ReplaceSubString(old string, from int, to int, new string) string {
	return old[0:from] + new + old[to:]
}
