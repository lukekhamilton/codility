package StacksAndQueues

/*
A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

        S is empty;
        S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
        S has the form "VW" where V and W are properly nested strings.

For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

    func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

        N is an integer within the range [0..200,000];
        string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".
*/
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
