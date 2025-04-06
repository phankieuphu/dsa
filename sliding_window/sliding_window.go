package sliding_window

func SlidingWindow(s string) int {
	// Create a map to store the last index of each character
	lastIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	for end, char := range s {
		// If the character is already in the map and its index is greater than or equal to start
		if lastIdx, exists := lastIndex[char]; exists && lastIdx >= start {
			// Move the start pointer to the right of the last occurrence
			start = lastIdx + 1
		}
		// Update the last index of the character
		lastIndex[char] = end
		// Calculate the maximum length of the substring
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}
