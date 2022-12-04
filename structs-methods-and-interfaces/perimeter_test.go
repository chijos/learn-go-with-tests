package Perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		rectangle    Rectangle
		hasPerimeter float64
	}{
		{name: "square", rectangle: Rectangle{Width: 10.0, Height: 10.0}, hasPerimeter: 40.0},
		{name: "rectangle", rectangle: Rectangle{Width: 20.0, Height: 5.0}, hasPerimeter: 50.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rectangle.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("got %.2f want %.2f", got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
    name string
		shape Shape
		hasArea  float64
	}{
    {name: "rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
    {name: "circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
    {name: "triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      if got != tt.hasArea {
        t.Errorf("got %.2f want %.2f", got, tt.hasArea)
      }
    })
	}

}
