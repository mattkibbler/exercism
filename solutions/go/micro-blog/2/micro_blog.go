package microblog


func Truncate(phrase string) string {
  	s := []rune(phrase)
    return string(s[0:min(len(s), 5)])
}
