package username_blacklist

import "fmt"

func Example() {
	exists := Blacklisted("admin")
	fmt.Println(exists)
	exists = Blacklisted("caffeines")
	fmt.Println(exists)
	// Output: true
	// false
}
