package main

import "fmt"

func compareString(a string, b string) bool {
	for n := 0; n < len(a); n++ {
		if a[n] == b[n] {
			return true
		}
	}
	return false
}

func commonSubstring(a []string, b []string) {
	// Write your code here
	for i := 0; i < len(a); i++ {
		if compareString(a[i], b[i]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	a := []string{"ab", "ae", "cf"}
	b := []string{"cb", "dd", "cg"}
	commonSubstring(a, b)
}
