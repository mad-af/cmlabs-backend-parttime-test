// Link check code online
// https://onlinegdb.com/zHjqCThlk

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	fmt.Print("input your text: ")
	scanner.Scan()
	text = strings.ToLower(strings.Join(strings.Fields(scanner.Text()), ""))

	fmt.Printf(`Result: %s`, checkPalindrome(text))
}

func checkPalindrome(input string) string {
	var i int
	var len int = len(input)
	for ; i < len/2; i++ {
		if input[i] != input[len-i-1] {
			return `FALSE`
		}
	}
	return `TRUE`

}
