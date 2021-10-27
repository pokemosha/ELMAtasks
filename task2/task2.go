package task2

//2. Циклическая ротация
func Solution_two(A []int, K int) []int {
	if len(A) == K {
		return A
	}
	firstSlise := A[len(A)-K:]
	secondSlise := A[:len(A)-K]
	return append(firstSlise, secondSlise...)
}
