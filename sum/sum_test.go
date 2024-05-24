package sum

import "testing"
import "reflect"

func TestSum(t *testing.T) {
    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}
        got := Sum(numbers)
        expected := 6

        if got != expected {
            t.Errorf("expected %d but got %d given %v", expected, got, numbers)
        }
    })

}

func TestSumAllTails(t *testing.T) {

    checkSums := func(t testing.TB, got, expected []int) {
        t.Helper()
        if !reflect.DeepEqual(got, expected) {
            t.Errorf("expected %v but got %v", expected, got)
        }
    }

    t.Run("make the sum of some slices", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        expected := []int{2, 9}
        checkSums(t, got, expected)
    })
    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        expected := []int{0, 9}
        checkSums(t, got, expected)
    })
}
