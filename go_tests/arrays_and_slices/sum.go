package arrays_and_slices

func SumArray(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return sum
}
