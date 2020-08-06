package encoder

import (
	"testing"
)

func TestURLEncoder(t *testing.T) {
	code := URLCodeGenerator(1)
	if len(code) > 6 {
		t.Errorf("Number of charater is exceeding, %s", code)
	}
}
