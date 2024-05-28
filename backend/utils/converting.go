package utilities

import "strings"

func ConvertToSlice(entries string) []string {
	// Split the entries by the new line character
	entriesSlice := strings.Split(entries, "\n")

	// Return the slice
	return entriesSlice
}

func Split(entries, separator string) []string {
	// Split the entries by the separator
	entriesSlice := strings.Split(entries, separator)

	// Return the slice
	return entriesSlice
}

func FindIndex(buffer, startKey []byte) int {
	// Find the index of the start key
	index := -1
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == startKey[0] {
			if string(buffer[i:i+len(startKey)]) == string(startKey) {
				index = i
				break
			}
		}
	}

	// Return the index
	return index
}
