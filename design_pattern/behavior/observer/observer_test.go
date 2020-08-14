package observer

import "testing"

func TestObserver(t *testing.T) {
	sub := &ConcreteSubject{
		observerList: make([]Observer, 0),
	}

	observer1 := &ConcreteObserverOne{}
	observer2 := &ConcreteObserverTwo{}

	sub.registerObserver(observer1)
	sub.registerObserver(observer2)

	sub.notifyObserver("我准备去上个厕所...")
}
