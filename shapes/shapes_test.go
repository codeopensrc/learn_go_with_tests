package shapes

import "testing"
import "math"

type Rectangle struct {
    Width float64
    Height float64
}
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
    Base float64
    Height float64
}
func (t Triangle) Area() float64 {
    return (t.Base * t.Height) * 0.5
}

type Shape interface {
    Area() float64
}


func TestPerimeter(t *testing.T) {
    errLogHelper := func(t testing.TB, got, expected float64) {
        t.Helper()
        if got != expected {
            t.Errorf("got %.2f expected %.2f", got, expected)
        }
    }
    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{10.0, 10.0}
        got := Perimeter(rectangle)
        expected := 40.0
        errLogHelper(t, got, expected)
    })
}

func TestArea(t *testing.T) {
    areaTests := []struct {
        name     string
        shape    Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 5.0, Height: 2.0}, hasArea: 10.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
    }

    for _, tt := range areaTests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
            }
        })
    }

}
