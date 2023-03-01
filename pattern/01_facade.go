package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Facade struct {
	Subsystem1 *Subsystem1
	Subsystem2 *Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		Subsystem1: NewSubsystem1(),
		Subsystem2: NewSubsystem2(),
	}
}

func (f *Facade) Operation() {
	f.Subsystem1.Operation1()
	f.Subsystem2.Operation2()
}

type Subsystem1 struct{}

func NewSubsystem1() *Subsystem1 {
	return &Subsystem1{}
}

func (*Subsystem1) Operation1() {
	fmt.Println("Subsystem1 Operation1")
}

type Subsystem2 struct{}

func NewSubsystem2() *Subsystem2 {
	return &Subsystem2{}
}

func (*Subsystem2) Operation2() {
	fmt.Println("Subsystem2 Operation2")
}
