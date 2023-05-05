package integers

func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(slices ...[]int) (sum []int) {
	for _, num := range slices {
		sum = append(sum, Sum(num))
	}
	return
}

func SumAllTail(slices ...[]int) (sum []int) {
	for _, num := range slices {
		if len(num) > 0 {
			sum = append(sum, Sum(num[1:]))
		} else {
			sum = append(sum, 0)
		}
	}
	return
}
