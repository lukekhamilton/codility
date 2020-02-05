package Arrays

func CyclicRotation(A []int, K int) []int {
	N := len(A)
	if N == 0 {
		return A
	}
	K = K % N
	rotated := make([]int, N)

	// shift start of array
	for i := 0; i < K; i++ {
		rotated[i] = A[N-K+i]
	}
	// shift end of array
	for i := K; i < N; i++ {
		rotated[i] = A[i-K]
	}

	return rotated
}
