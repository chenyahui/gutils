package container

// inspired by gin
type H map[string]interface{}

func (h H) GetInt(key string) int {
	return int(h[key].(float64))
}

func (h H) GetInt64(key string) int64 {
	return int64(h[key].(float64))
}

func (h H) GetDefaultInt64(key string, defaultVal int64) int64 {
	if val, ok := h[key]; ok {
		return int64(val.(float64))
	}
	return defaultVal
}

func (h H) GetString(key string) string {
	return h[key].(string)
}

func (h H) Merge(other H) {
	for key, value := range other {
		if _, ok := h[key]; !ok {
			h[key] = value
		}
	}
}
