package main

// LEEDCODE: 125

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(line string) bool {
	var reg = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	fmt.Println(line)
	s := reg.ReplaceAllString(line, "")
	fmt.Println(s)
	s = strings.Replace(s, " ", "", -1)
	fmt.Println(s)
	s = strings.ToLower(s)
	fmt.Println(s)

	var result string
	for _, v := range s {

		result = string(v) + result
	}
	fmt.Println("flipped:", result)
	if s == result {
		return true
	}

	return false
}

/*

Answer in Python

def isPalindrome(line):
    output = "".join(filter(str.isalnum, line)).lower()
    print(output)
    test = output
    test = test[::-1]
    if test == output:
        return True

    return False

def main():
    line = "A man, a plan, a canal: Panama"
    print(isPalindrome(line))

main()
*/
