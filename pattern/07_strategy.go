package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
import "fmt"

type Strategy interface {
	Execute(int, int) int
}

type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

func main() {
	context := &Context{}

	addStrategy := &ConcreteStrategyAdd{}
	context.SetStrategy(addStrategy)
	result := context.ExecuteStrategy(10, 5)
	fmt.Println(result) // Output: 15

	subtractStrategy := &ConcreteStrategySubtract{}
	context.SetStrategy(subtractStrategy)
	result = context.ExecuteStrategy(10, 5)
	fmt.Println(result) // Output: 5
}
