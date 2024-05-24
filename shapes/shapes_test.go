package shapes

import "testing"

type Rectangle struct {
    Width float64
    Height float64
}

func TestPerimeter(t *testing.T) {

    errLogHelper := func(t testing.TB, got, expected float64) {
        t.Helper()
        if got != expected {
            t.Errorf("got %.2f expected %.2f", got, expected)
        }
    }

    t.Run("calc perimeter with height and width", func(t *testing.T) {
        rectangle := Rectangle{10.0, 10.0}
        got := Perimeter(rectangle)
        expected := 40.0
        errLogHelper(t, got, expected)
    })
}

func TestArea(t *testing.T) {
    rectangle := Rectangle{5.0, 2.0}
    got := Area(rectangle)
    expected := 10.0
    if got != expected {
        t.Errorf("got %.2f expected %.2f", got, expected)
    }
}
