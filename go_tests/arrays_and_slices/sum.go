package arrays_and_slices

func SumSlice(slice []int) (sum int) {
	for _, num := range slice {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) (result []int) {
	for _, slice := range slices {
		result = append(result, SumSlice(slice))
	}
	return result
}

func SumAllTails(slices ...[]int) (result []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			result = append(result, 0)
			continue
		}
		result = append(result, SumSlice(slice[1:]))
	}
	return result
}
