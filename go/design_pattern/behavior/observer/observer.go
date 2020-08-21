package observer

import (
	"fmt"
)

/*
	大到消息队列的订阅发布模式就是用的观察者模式思想
*/

//Observer 观察者
type Observer interface {
	update(message string)
}

//Subject 被观察者应该实现的接口
type Subject interface {
	// 注册一个观察者
	registerObserver(observer Observer)
	// 移除一个观察者
	removeObserver(observer Observer)
	// 通知观察者
	notifyObserver(message string)
}

//ConcreteSubject 被观察则对象
type ConcreteSubject struct {
	observerList []Observer
}

func (c *ConcreteSubject) registerObserver(observer Observer) {
	c.observerList = append(c.observerList, observer)
}

func (c *ConcreteSubject) removeObserver(observer Observer) {
	var index int
	var item Observer
	for index, item = range c.observerList {
		if item == observer {
			break
		}
	}
	c.observerList = append(c.observerList[0:index], c.observerList[index:]...)
}

func (c *ConcreteSubject) notifyObserver(message string) {
	for _, observer := range c.observerList {
		// 如果每个观察者的调用比较耗费时间，可以使用异步非阻塞的方式
		observer.update(message)
	}
}

//ConcreteObserverOne 观察者1
type ConcreteObserverOne struct {
}

func (o *ConcreteObserverOne) update(message string) {
	fmt.Println("ConcreteObserverOne is notified, sub say: ." + message)
}

//ConcreteObserverTwo 观察者2
type ConcreteObserverTwo struct {
}

func (o *ConcreteObserverTwo) update(message string) {
	fmt.Println("ConcreteObserverTwo is notified, sub say: " + message)
}
