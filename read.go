package leakedpassword

import (
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"strings"

	"github.com/bits-and-blooms/bloom/v3"
)

//go:embed db.db
var DB embed.FS

//IsLeaked is checking password in leaked passwords bitset db
// To find password leaked status first we need to generate sum of file
// and make uppercase because db is stored with uppercase sums
func IsLeaked(password string) (bool, error) {
	fil := bloom.NewWithEstimates(6000000, 0.0001)
	passDb, err := DB.Open("db.db")
	if err != nil {
		return false, err
	}
	_, err = fil.ReadFrom(passDb)
	if err != nil {
		return false, err
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
