package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("the output is", gcdOfStrings("AABABAB", "ABAB"))
}

// naive approach
func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	if len1 == len2 {
		if str1 != str2 {
			return ""
		} else {
			return str2
		}
	}

	if len1 < len2 {
		str1, str2 = str2, str1
		len1, len2 = len2, len1
	}

	str2orig := str2

	for i := len2; i > 0; i-- {
		//if no common divisor, we skip to next loop
		mod1 := len1 % i
		mod2 := len2 % i

		if mod1 != 0 && mod2 != 0 {
			continue
		}

		str2 = str2[:i]

		//suppose 1 is abc and 2 is abc, then we need to see how many times we can multiply the two
		count1 := len1 / i
		count2 := len2 / i

		if strings.Repeat(str2, count1) == str1 && strings.Repeat(str2, count2) == str2orig {
			return str2
		}
	}

	return ""
}

// from fastest solution in leetcode submissions - using recursion
func gcdOfStrings2(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	if len(str1) == len(str2) {
		return str1
	}

	if len(str1) > len(str2) {
		return gcdOfStrings2(str1[len(str2):], str2)
	}

	return gcdOfStrings2(str1, str2[len(str1):])
}
