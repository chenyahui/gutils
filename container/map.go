package container

func MergeMap(dst map[string]interface{}, src map[string]interface{}) {
	for key, value := range src {
		if _, ok := dst[key]; !ok {
			dst[key] = value
		}
	}
}
