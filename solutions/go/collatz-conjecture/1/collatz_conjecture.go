package collatzconjecture

func CollatzConjecture(n int) (int, error) {
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
