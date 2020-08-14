package singleton

import "sync"

var singleton *Singleton
var once sync.Once

//Singleton 单例   这种实现方式是，支持延迟加载的单例模式
type Singleton struct{}

//GetInstance 获取单例
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
