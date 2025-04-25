package palindromeNumber9

import (
	"fmt"
	"strconv"
)

func Run() {
	fmt.Println(isPalindrome(121) == true)
	fmt.Println(isPalindrome(-121) == false)
	fmt.Println(isPalindrome(10) == false)
}

func isPalindrome(x int) bool {
	cadena := []rune(strconv.Itoa(x))
	mida := len(cadena)

	for index, caracter := range cadena {
		if caracter != cadena[mida-1-index] {
			return false
		}
	}

	return true
}
