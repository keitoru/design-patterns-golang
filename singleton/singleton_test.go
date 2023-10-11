package singleton

import "testing"

func Test_singleton(t *testing.T) {
	s := GetSingleton()
	s.Work()
}
