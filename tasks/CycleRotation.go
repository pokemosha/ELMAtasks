package tasks

//2. Циклическая ротация
func CycleRotation(A []int, K int) []int {
	K = K % len(A)
	if len(A) == K && len(A) == 1 {
		return A
	}
	firstSlise := A[len(A)-K:]
	secondSlise := A[:len(A)-K]
	return append(firstSlise, secondSlise...)
}
