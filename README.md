# username-blacklist 
[![Go Reference](https://pkg.go.dev/badge/github.com/caffeines/username-blacklist.svg)](https://pkg.go.dev/github.com/caffeines/username-blacklist)
[![Tests](https://github.com/caffeines/username-blacklist/actions/workflows/go.yml/badge.svg)](https://github.com/caffeines/username-blacklist/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/caffeines/username-blacklist)](https://goreportcard.com/report/github.com/caffeines/username-blacklist)
[![License](https://img.shields.io/badge/license-MPL2.0-blue.svg)](https://raw.githubusercontent.com/caffeines/username-blacklist/master/LICENSE)

**Username Blacklist** checks usernames, that you might not want to grant as a username.

## Installation
```sh
go get github.com/caffeines/username-blacklist
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/caffeines/username-blacklist"
)

func main() {
	blacklisted := username_blacklist.Blacklisted("admin")
	fmt.Println("blacklisted:", blacklisted)
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[![License](https://img.shields.io/badge/license-MPL2.0-blue.svg)](https://raw.githubusercontent.com/caffeines/username-blacklist/master/LICENSE)

This project is released under the MPL2.0 License. See the bundled LICENSE file for details.

## Credits
- [github.com/marteinn/The-Big-Username-Blacklist](https://github.com/marteinn/The-Big-Username-Blacklist)




