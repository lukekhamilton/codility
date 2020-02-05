package TimeComplexity

import (
	"math"
	"math/bits"
)

// 53% https://app.codility.com/demo/results/trainingKJ6BZC-QXD/
func TapeEquilibrium0(A []int) int {
	minimalDifference := math.MaxFloat64

	leftSum := 0
	for i := 0; i < len(A); i++ {
		leftSum += A[i]

		rightSum := 0
		for ii := i + 1; ii < len(A); ii++ {
			rightSum += A[ii]
		}

		diff := math.Abs(float64(leftSum - rightSum))
		if minimalDifference > diff {
			minimalDifference = diff
		}
	}

	return int(minimalDifference)
}

// 100% https://app.codility.com/demo/results/trainingFWTBSC-C7X/
func TapeEquilibrium(A []int) int {
	sumAll := sum(A)
	leftSum := 0
	rightSum := 0
	minDiff := 1<<bits.UintSize/2 - 1
	currentDiff := 1<<bits.UintSize/2 - 1

	for i := 0; i < len(A)-1; i++ {
		leftSum += A[i]
		rightSum = sumAll - leftSum
		currentDiff = int(math.Abs(float64(leftSum) - float64(rightSum)))
		minDiff = int(math.Min(float64(currentDiff), float64(minDiff)))
	}
	return minDiff
}

func sum(A []int) int {
	sum := 0
	for _, e := range A {
		sum += e
	}
	return sum
}
