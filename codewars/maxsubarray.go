package codewars

// Caddan's Classic Algorythm
func MaxSequence(arr []int) int {
	maxSoFar, maxEndingHere := 0, 0
	for _, v := range arr {
		maxEndingHere = max(0, maxEndingHere+v)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 func MaximumSubarraySum(numbers []int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
} */
