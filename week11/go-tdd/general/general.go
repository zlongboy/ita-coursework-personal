package general

func main() {
	SumAll([]int{1, 2}, []int{3, 6})
}

func SumAll(toSum ...[]int) []int { // variadic arguments
	var finalSum []int

	for _, nums := range toSum {
		finalSum = append(finalSum, Sum(nums))
	}

	return finalSum
}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
