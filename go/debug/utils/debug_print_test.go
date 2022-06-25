package utils

import (
	"testing"
)

func Test_DBG_PRINTF(t *testing.T) {
	t.Log("go test -v or go test -v -tags debug")
	DBG_PRINTF("This a debug printf func\n")
}
