package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("number is less than 1")
	}
	num := n
	steps := 0
	for num > 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = (num * 3) + 1
		}
		steps++
	}
	return steps, nil
}
