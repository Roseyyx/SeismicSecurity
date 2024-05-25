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
