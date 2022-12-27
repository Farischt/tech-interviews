package arrays_hash

// Time Complexity O(n)
// Space Complexity O(n)
func LongestConsecutive(nums []int) int {
	// Bucket sort
	set := make(map[int]bool)
	longestSequenceLen := 0

	for _, el := range nums {
		if _, exist := set[el]; !exist {
			set[el] = true
		}
	}

	currentLen := 1
	for _, el := range nums {
		// Does the current number is a sequence start ?
		if _, exist := set[el-1]; !exist {
			// In this case this is a sequence start
			currentLen = 1

			_, nextNumberExist := set[el+currentLen]
			for nextNumberExist {
				currentLen++
				nextNumberExist = set[el+currentLen]
			}

			if currentLen > longestSequenceLen {
				longestSequenceLen = currentLen
			}

		}

	}

	return longestSequenceLen
}
