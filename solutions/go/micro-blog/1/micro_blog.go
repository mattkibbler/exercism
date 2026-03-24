package microblog

import "strings"

func Truncate(phrase string) string {
  	var b strings.Builder

    i := 0
	for _, r := range phrase {
		if i >= 5 {
			break
		}
		b.WriteRune(r)
        i++
	}

	return b.String()
}
