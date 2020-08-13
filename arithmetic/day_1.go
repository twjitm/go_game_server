package arithmetic

import "fmt"

//duplicates sorted array
func DuplicatesArray(data []int) (int, []int) {

	length := len(data)

	var index = 0
	maps := make(map[int]int)
	for i := 0; i < length; i++ {
		maps[data[i]] = data[i]
	}
	keys, values := MapKv(maps)
	fmt.Println(len(maps), keys)
	return index, values
}

func MapKv(data map[int]int) ([]int, []int) {

	keys := make([]int, 0, len(data))
	values := make([]int, 0, len(data))
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func HalfSearch(data []int, target int) (int, int) {
	if len(data) == 0 {
		return -1, -1
	}
	min := 0
	max := len(data) - 1
	mid := max / 2
	for max > min {
		if target > data[mid] {
			min = mid
			mid = (max + min) / 2
		} else if target < data[mid] {
			max = mid
			mid = (max + min) / 2
		} else {
			return target, mid
		}
	}
	return -1, -1
}




