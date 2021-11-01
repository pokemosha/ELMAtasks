package tasks

import "sort"

//3. Чудные вхождения в массив
func WonderfulArrayEntries(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if A[i] != A[i+1] {
			return A[i]
		}
		i++
	}
	return 0
}
