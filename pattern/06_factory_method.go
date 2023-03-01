package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a circle")
}

type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle")
}

func ShapeFactory(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return &Circle{}
	case "rectangle":
		return &Rectangle{}
	default:
		return nil
	}
}

func main() {
	circle := ShapeFactory("circle")
	rectangle := ShapeFactory("rectangle")

	circle.Draw()    // Output: Drawing a circle
	rectangle.Draw() // Output: Drawing a rectangle
}
