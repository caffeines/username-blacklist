package username_blacklist_test

import (
	"fmt"
	username_blacklist "github.com/caffeines/username-blacklist"
)

func Example() {
	exists := username_blacklist.Blacklisted("admin")
	fmt.Println(exists)
	exists = username_blacklist.Blacklisted("caffeines")
	fmt.Println(exists)
	// Output: true
	// false
}
