package main

// import required modules
import (
	"fmt"
)

// main function
func main() {
	var str string = "Lorem ipsum dolor sit amet"
	fmt.Println(str[6:11])

	// Creating a slice of string
	s := []string{"abc", "def"}

	// Appending elements to existing slice
	s = append(s, "ghi")
	fmt.Println(s)
}
