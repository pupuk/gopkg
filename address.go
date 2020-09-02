package address

import "fmt"

// Decompose a string to map
func Decompose(address string) string {
	return "Decompose"
}

// Parse a tring to a map
func Parse(address string) string {
	return "Parse"
}

// Smart fucniton, include decompose ,then Parse
func Smart(address string) string {
	return "Smart"
}

// Test then echo a string
func Test(str string) string {
	fmt.Println("This is test string")
}
