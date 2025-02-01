package tool

func GetKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m)) // Preallocate slice for efficiency
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
