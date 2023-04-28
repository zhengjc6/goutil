package strtool

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomString(16))
	}
}
