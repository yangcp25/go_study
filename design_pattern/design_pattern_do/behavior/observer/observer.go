package observer

import "fmt"

type ISubject interface {
	Register(observer IObserve)
	Remove(observer IObserve)
	Notify(msg string)
}

type IObserve interface {
	Update(msg string)
}

type Subject struct {
	observer []IObserve
}

func (s *Subject) Register(observer IObserve) {
	s.observer = append(s.observer, observer)
}

func (s *Subject) Remove(observer IObserve) {
	for i, observe := range s.observer {
		if observe == observer {
			s.observer = append(s.observer[:i], s.observer[:i+1]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, observe := range s.observer {
		observe.Update(msg)
	}
}

type Observe1 struct{}

func (o Observe1) Update(msg string) {
	fmt.Println("Observe1 Update:", msg)
}

type Observe2 struct{}

func (o Observe2) Update(msg string) {
	fmt.Println("Observe2 Update:", msg)
}
