package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("strings must be the same length")
    }
    result := 0
    for i, char := range a {
        if char != rune(b[i]) {
            result++
        }
    }
    return result, nil
}
