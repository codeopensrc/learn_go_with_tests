package sum

func Sum(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func SumAll(numbersToSum ...[]int) []int {
    arr_len := len(numbersToSum)
    sums := make([]int, arr_len)

    for i, numbers := range numbersToSum {
        sums[i] = Sum(numbers)
    }
    return sums
}
