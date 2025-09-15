package main

import (
	"fmt"
)

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlphanumericSymbol(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || isNumeric(b) {
		return true
	}
	return false
}

func isAlphanumericEqualFold(b1, b2 byte) bool {
	if b1 == b2 {
		return true
	}
	if isNumeric(b1) || isNumeric(b2) {
		return false
	}
	if b1 > b2 && b1-b2 == 'a'-'A' {
		return true
	}
	if b1 < b2 && b2-b1 == 'a'-'A' {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphanumericSymbol(s[i]) {
			i++
			continue
		}
		if !isAlphanumericSymbol(s[j]) {
			j--
			continue
		}
		if !isAlphanumericEqualFold(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama") == true)
	fmt.Println(isPalindrome("race a car") == false)
	fmt.Println(isPalindrome(" ") == true)
	fmt.Println(isPalindrome("0P") == false)
}
