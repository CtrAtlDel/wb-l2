package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	В отличие от других шаблонов создания, Builder не требует, чтобы продукты имели общий интерфейс. Это позволяет производить различные изделия, используя один и тот же процесс изготовления.
*/

type IBuilder interface {
	CallSoulGoodman(str string)
	CallWalterWhite(str string)
}

type Phone struct {
	builder IBuilder
}

func (c *Phone) Phone() {
	c.builder.CallSoulGoodman("Soul")
	c.builder.CallWalterWhite("Walter")
}

type Builder struct {
	line *Line
}

func (b *Builder) CallSoulGoodman(str string) {
	b.line.Message = str + "scamer"
}

func (b *Builder) CallWalterWhite(str string) {
	b.line.Message = str + "not scamer"
}

type Line struct {
	Message string
}

func (l *Line) Show() {
	fmt.Println(l.Message)
}
