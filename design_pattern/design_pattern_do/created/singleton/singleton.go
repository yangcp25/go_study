package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

var lazeSingleton *Singleton
var once = &sync.Once{}

func GetLazyInstance() *Singleton {
	if lazeSingleton == nil {
		once.Do(func() {
			lazeSingleton = &Singleton{}
		})
	}
	return lazeSingleton
}
