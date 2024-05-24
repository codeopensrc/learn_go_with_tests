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
    errLogHelper := func(t testing.TB, got, expected float64) {
        t.Helper()
        if got != expected {
            t.Errorf("got %g expected %g", got, expected)
        }
    }
    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{5.0, 2.0}
        got := rectangle.Area()
        expected := 10.0
        errLogHelper(t, got, expected)
    })
    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        got := circle.Area()
        expected := 314.1592653589793
        errLogHelper(t, got, expected)
    })
}
