package encoder

import (
	"encoding/base64"
	"strconv"
)

// URLCodeGenerator takes the number and hashes it and gives back a hashed string
func URLCodeGenerator(urlNumber int) string {
	numberString := strconv.Itoa(urlNumber)
	code := base64.URLEncoding.EncodeToString([]byte(numberString))
	if len(code) > 6 {
		code = code[:6]
	}
	return code
}
