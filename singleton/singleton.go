package singleton

import (
	"fmt"
	"sync"
)

var s *singleton
var once sync.Once

type singleton struct {
}

func newSingleton() *singleton {
	return &singleton{}
}

func (s *singleton) Work() {
	fmt.Println("working")
}

type Instance interface {
	Work()
}

func GetSingleton() Instance {
	once.Do(func() {
		s = newSingleton()
	})

	return s
}
