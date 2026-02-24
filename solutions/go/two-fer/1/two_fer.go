// Package twofer implements twofer logic
package twofer

import "fmt"

// ShareWith returns a string describing who gets the cookies
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
