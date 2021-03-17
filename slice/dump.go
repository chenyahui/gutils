package slice

func DumpStrings(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}
