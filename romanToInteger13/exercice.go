package romanToInteger13

import (
	"fmt"
)

func Run() {
	fmt.Println(romanToInt("III") == 3)
	fmt.Println(romanToInt("LVIII") == 58)
	fmt.Println(romanToInt("MCMXCIV") == 1994)
}

func romanToInt(s string) int {
	total := 0
	length := len(s)

	for i := 0; i < length; i++ {
		if i < length-1 {
			if s[i] == 'I' && s[i+1] == 'V' {
				total += 4
				i++
				continue
			} else if s[i] == 'I' && s[i+1] == 'X' {
				total += 9
				i++
				continue
			} else if s[i] == 'X' && s[i+1] == 'L' {
				total += 40
				i++
				continue
			} else if s[i] == 'X' && s[i+1] == 'C' {
				total += 90
				i++
				continue
			} else if s[i] == 'C' && s[i+1] == 'D' {
				total += 400
				i++
				continue
			} else if s[i] == 'C' && s[i+1] == 'M' {
				total += 900
				i++
				continue
			}
		}

		if s[i] == 'I' {
			total += 1
		} else if s[i] == 'V' {
			total += 5
		} else if s[i] == 'X' {
			total += 10
		} else if s[i] == 'L' {
			total += 50
		} else if s[i] == 'C' {
			total += 100
		} else if s[i] == 'D' {
			total += 500
		} else if s[i] == 'M' {
			total += 1000
		}
	}

	return total
}
