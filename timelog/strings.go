package timelog

// CloneStrings duplicates []string in memory.
func CloneStrings(data []string) []string {
	r := make([]string, len(data))
	copy(r, data)
	return r
}
