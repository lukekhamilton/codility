package PrefixSums

/*
A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:
    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8

contains the following example slices:

        slice (1, 2), whose average is (2 + 2) / 2 = 2;
        slice (3, 4), whose average is (5 + 1) / 2 = 3;
        slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.

The goal is to find the starting position of a slice whose average is minimal.

Write a function:

    func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

For example, given array A such that:
    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8

the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [2..100,000];
        each element of array A is an integer within the range [−10,000..10,000].
*/
func MinAvgTwoSlice(A []int) int {
	N := len(A)
	minAvg := float64(A[0]+A[1]) / 2
	minPos := 0
	for i := 2; i < N; i++ {
		avg := float64(A[i-1]+A[i-2]) / 2
		if avg < minAvg {
			minAvg = avg
			minPos = i - 2
		}
	}

	// check last two elements
	if N > 2 {
		last2Avg := float64(A[N-2]+A[N-1]) / 2
		if last2Avg < minAvg {
			minAvg = last2Avg
			minPos = N - 2
		}
	}

	return minPos
}
