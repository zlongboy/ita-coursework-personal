package general

func main() {
	SumAll([]int{1, 2}, []int{3, 6})
}

func SumAll(sOne, sTwo []int) []int {
	var finalSum []int
	finalSum = append(finalSum, Sum(sOne), Sum(sTwo))
	return finalSum
}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
