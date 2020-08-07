package encoder

import (
	"encoding/base32"
	"strconv"
	"strings"
)

// URLCodeGenerator takes the number and hashes it and gives back a hashed string
func URLCodeGenerator(urlNumber int) string {
	numberString := strconv.Itoa(urlNumber)
	code := base32.StdEncoding.EncodeToString([]byte(numberString))
	if len(code) > 6 {
		code = code[:6]
	}
	return strings.ToLower(code)
}
