package backend

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(1, 3)
	if got != 4 {
		t.Errorf("wrong value got => %v", got)
	}
}