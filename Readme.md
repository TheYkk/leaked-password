# Leaked password check

With this library you can check the password is probably leaked or not.

Pre generated bitset DB includes 6 Million leaked passwords stored in bloom filter bitset.

To generate your own bitset DB head into [Db-generate](https://github.com/TheYkk/db-generate) and follow instructions.
## Bloom filter

The current configuration is a 1 in 1 Million false positive rates. With bloom filter, you can make sure if data is not stored in the bitset. But bloom filter can generate false positives.

To change error rate, you need to generate your own bitset DB. After that you can change bitset DB like this

```go
package main

import (
	"github.com/theykk/leaked-password"
	"os"
	"io"
)

func main() {
	myDB, _ := os.Open("my.db")
	defer myDB.Close()
	leakedpassword.CustomReader = myDB
}
```
# Usage

```go
package main

import (
	"github.com/theykk/leaked-password"
)

func registerUser(username, password string) error {
	// Check password is leaked
	isLeaked, err := leakedpassword.IsLeaked(password)
	if err != nil {
		return err
	}
		
	if isLeaked {
		// Password is leaked do something
	}
	return nil
}
```
