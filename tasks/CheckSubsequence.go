package tasks

//1. Проверка последовательности
func CheckSubsequence(A []int) int {
	check := make([]int, len(A)+2)
	for i := 0; i < len(A); i++ {
		if check[A[i]] == 1 {
			return 0
		}
		if A[i] > len(A) {
			return 0
		}
		check[A[i]] = 1
	}
	return 1
}
