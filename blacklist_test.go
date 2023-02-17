package username_blacklist

import (
	"fmt"
	"testing"
)

func TestBlacklisted(t *testing.T) {
	t.Run("should return true if the username is blacklisted", func(t *testing.T) {
		got := Blacklisted("admin")
		if got != true {
			t.Errorf("got %v want %v", got, true)
		}
	})

	t.Run("should return false if the username is not blacklisted", func(t *testing.T) {
		got := Blacklisted("caffeines")
		if got != false {
			t.Errorf("got %v want %v", got, false)
		}
	})
}

func ExampleBlacklisted() {
	exists := Blacklisted("admin")
	fmt.Println(exists)
	exists = Blacklisted("caffeines")
	fmt.Println(exists)
	// Output: true
	// false
}
