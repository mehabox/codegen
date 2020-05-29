package main

import "fmt"

// consumer is a function that does something with the generated code
type consumer func(string)

func main() {
	alphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	getCombinations(alphabet, 4, func(s string) { fmt.Println(s) })
}

// getCombinations performs permutations to all of the letters generating numChars-sized strings
func getCombinations(letters []byte, numChars int, consumer consumer) {
	var curChar, idx int
	chars := make([]int, numChars+1)
	maxLetter := len(letters)
	for idx = 1; idx <= numChars; idx++ {
		chars[idx] = 1
	}
	chars[numChars] = 0
	curChar = numChars
	for {
		if chars[curChar] == maxLetter {
			curChar--
			if curChar == 0 {
				break
			}
		} else {
			chars[curChar]++
			for {
				if curChar >= numChars {
					break
				}
				curChar++
				chars[curChar] = 1

			}
			// assembling the code to a string
			var code []byte
			for idx = 1; idx <= numChars; idx++ {
				code = append(code, letters[chars[idx]-1])
			}
			// code is ready, call the consumer function
			consumer(string(code))
		}
	}
}
