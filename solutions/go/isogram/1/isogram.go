package isogram

import "strings"

func IsIsogram(word string) bool {
    word = strings.NewReplacer("-", "", " ", "").Replace(strings.ToLower(word))
	tracker := make(map[rune]int)
    for _, char := range word {
        tracker[char]++
        if(tracker[char] > 1) {
            return false
        }
    }
    return true
}
