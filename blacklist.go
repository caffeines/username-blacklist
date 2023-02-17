package username_blacklist

// Blacklisted returns true if the username is blacklisted
func Blacklisted(username string) bool {
	_, ok := blacklist[username]
	return ok
}
