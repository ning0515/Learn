package integers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func TestSum(t *testing.T) {
	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	// 	sum := Sum(numbers)
	// 	want := 15

	// 	if sum != want {
	// 		t.Errorf("want %d but got %d given %d", want, sum, numbers)
	// 	}
	// })

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		want := 6

		if sum != want {
			t.Errorf("want %d but got %d given %d", want, sum, numbers)
		}

	})

	t.Run("collection of 1 slice", func(t *testing.T) {
		sum := SumAll([]int{1, 2}, []int{1, 2, 3})
		want := []int{3, 6}

		if !reflect.DeepEqual(sum, want) {
			t.Errorf("want %v but sum %v", want, sum)
		}
	})

}

func TestSumAllTail(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	}
	t.Run("常规测试", func(t *testing.T) {
		sum := SumAllTail([]int{1, 2}, []int{1, 2, 3})
		want := []int{2, 5}

		checkSums(t, sum, want)
	})

	t.Run("测试传入空切片", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	//Output: 6
}
