package sum

import "testing"

func TestSum(t *testing.T) {
    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}
        got := Sum(numbers)
        expected := 15

        if got != expected {
            t.Errorf("expected %d but got %d given %v", expected, got, numbers)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}
        got := Sum(numbers)
        expected := 6

        if got != expected {
            t.Errorf("expected %d but got %d given %v", expected, got, numbers)
        }
    })

}
