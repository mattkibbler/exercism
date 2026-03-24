package luhn

import "strings"

func Valid(id string) bool {
	
    id = strings.NewReplacer(" ", "").Replace(id)

    if len(id) <= 1 {
        return false
    }
    
    runes := []rune(id)
    total := 0
    secondDigit := false
    for i := len(runes) - 1; i >= 0; i-- {
        if runes[i] < '0' || runes[i] > '9' {
            return false
        }
        
        num := int(runes[i] - '0')

        if secondDigit {
        	doubled := num * 2
            total += doubled
            if doubled > 9 {
                total -= 9
            }    
        } else {
            total += num
        }

        secondDigit = !secondDigit
        
    }
    return total % 10 == 0
    
}
