package singleton

import (
	"errors"
	"testing"
)

func TestSingleTon(t *testing.T) {
	s := GetInstance()
	if s.AddOne() != 1 {
		t.Error(errors.New("s count is not 1"))
	}
}
