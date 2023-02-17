// Date: 2018-01-01 15:00:00 +0000 UTC
// License: MIT
// Author: caffeines<sadat.talks@gmail.com>
// Description: Blacklisted usernames that you might not want to grant as a username.

/*
Package username_blacklist checks usernames, that you might not want to grant as a username.

Example:

	package main

	import (
		"fmt"
		"github.com/caffeines/username-blacklist"
	)

	func main() {
		blacklisted := username_blacklist.Blacklisted("admin")
		fmt.Println("blacklisted:", blacklisted)
	}

Requests or bugs?
https://github.com/caffeines/username-blacklist/issues
*/
package username_blacklist
