package cmp

func LessUint32(a, b uint32) bool {
	return int32(a-b) < 0
}

func LessUint64(a, b uint32) bool {
	return int64(a-b) < 0
}
