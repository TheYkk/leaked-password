package leakedpassword

import (
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"io"
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
)

//go:embed db.db
var DB embed.FS

// CustomReader To support user provided bitset DB
var CustomReader io.Reader

// IsLeaked is checking password in leaked passwords bitset DB
// Checking leak status can cause false positives, it depends on your bitset DB
func IsLeaked(password string) (bool, error) {
	fil := bloom.BloomFilter{}
	passDb, err := DB.Open("db.db")
	if err != nil {
		return false, err
	}

	if CustomReader != nil {
		_, err = fil.ReadFrom(CustomReader)
		if err != nil {
			return false, err
		}
	} else {
		_, err = fil.ReadFrom(passDb)
		if err != nil {
			return false, err
		}
	}

	// Generate sha1 sum of password
	h := sha1.New()
	h.Write([]byte(password))
	data := h.Sum(nil)

	// Convert to hex
	shaStr := hex.EncodeToString(data)

	// Make uppercase hex to filter in bitset
	up := strings.ToUpper(shaStr)

	// return password exist status in bitset
	return fil.TestString(up), nil
}
