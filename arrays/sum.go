package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlices(slicesToSum ...[]int) []int {
	var slicesSum []int
	for _, nums := range slicesToSum {
		slicesSum = append(slicesSum, Sum(nums))
	}
	return slicesSum
}

func SumTails(slicesTails ...[]int) []int {
	var tailsSum []int
	for _, nums := range slicesTails {
		if len(nums) == 0 {
			tailsSum = append(tailsSum, 0)
		} else {
			tailsSum = append(tailsSum, Sum(nums[1:]))
		}
	}
	return tailsSum
}

func SumSliceInner(slices ...[]int) []int {
	var result []int
	for _, nums := range slices {
		result = append(result, Sum(nums[1:len(nums)-1]))
	}
	return result
}
