package logs

import "fmt"
import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
        switch(fmt.Sprintf("%U", char)) {
        case "U+2757":
            return "recommendation"
        case "U+1F50D":
            return "search"
        case "U+2600":
            return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var builder strings.Builder
    
    for _, char := range log {
        if char == oldRune {
            builder.WriteRune(newRune)
        } else {
            builder.WriteRune(char)
        }
    }
    return builder.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := 0
    for range log {
        count++
    }
    return count <= limit
}
