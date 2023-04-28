package csvparse

import (
	"testing"
)

func TestExcel2Go(t *testing.T) {
	Excel2Go("./excel", "godata", "godata", true)
}

func TestExcel2Lua(t *testing.T) {
	Excel2Lua("./excel", "luadata", true)
} 
