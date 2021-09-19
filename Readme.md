# Leaked password check

With this library you can check password is leaked or not.

Pre generated db includes 6 Million leaked passwords stored in bloom filter bitset.

To generate your own db head into [Db-generate](https://github.com/TheYkk/db-generate) and follow instructions.

# Usage

```go
package main

import (
	"github.com/theykk/leaked-password"
)

func registerUser(username, password string) error {
	// Check password is leaked
	isLeaked, err := leakedpassword.Leaked(password)
	if err != nil {
		return err
	}
		
	if isLeaked {
		// Password is leaked do something
	}
	return nil
}
```