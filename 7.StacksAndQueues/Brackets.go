package __StacksAndQueues

func Brackets(S string) int {
	N := len(S)

	// Odd number S cant be nested
	if N%2 == 1 {
		return 0
	}

	opening := map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}

	closing := map[string]bool{
		"}": true,
		"]": true,
		")": true,
	}

	match := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	var stack []string
	for i := 0; i < N; i++ {

		if opening[string(S[i])] {
			stack = append(stack, string(S[i]))
			continue
		}

		if closing[string(S[i])] {
			if len(stack) != 0 {
				// pop
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				// match
				if match[pop] != string(S[i]) {
					return 0
				}
			} else {
				return 0
			}
		}
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
