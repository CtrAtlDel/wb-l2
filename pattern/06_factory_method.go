package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
import "fmt"

// Shape is an interface for shapes that can be drawn
type Shape interface {
	Draw()
}

// Circle is a struct representing a circle
type Circle struct{}

// Draw implements the Draw method of the Shape interface for Circle
func (c *Circle) Draw() {
	fmt.Println("Drawing a circle")
}

// Rectangle is a struct representing a rectangle
type Rectangle struct{}

// Draw implements the Draw method of the Shape interface for Rectangle
func (r *Rectangle) Draw() {
	fmt.Println("Drawing a rectangle")
}

// ShapeFactory is a function that returns a new Shape based on the input string
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
