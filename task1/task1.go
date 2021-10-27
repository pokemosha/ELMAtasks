package task1

//1. Проверка последовательности
func Solution_one(A []int) int {
	check := make([]int, len(A)+2)
	for i := 0; i < len(A); i++ {
		check[A[i]] = 1
	}
	for i := 1; i < len(A); i++ {
		if check[i] == 0 {
			return 0
		}
	}
	return 1
}
