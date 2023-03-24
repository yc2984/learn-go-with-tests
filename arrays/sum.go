package main // TODO: runtime.main_main·f: function main is undeclared in the main package. This message appears if main() is removed
import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for index, number := range numbers {
		sum += number
		fmt.Println(index)
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum{
		sums[i] = Sum(numbers)
	}
	return sums
}

func main() {
	array := []int{1, 2, 3, 4, 5} // TODO: how to give it directly in the function? 
	fmt.Println(Sum(array))
}