package tasks

// 4. Поиск отсутствующего элемента
func FindingMissingItem(A []int) int {
	A = append(A, 0)
	sch1, sch2 := 0, 0
	for i := 0; i < len(A); i++ {
		sch1 += A[i]
		sch2 += i + 1
	}
	return sch2 - sch1
}

/* 4. Поиск отсутствующего элемента
func Solution_four(A []int) int {
	check := make([]int, len(A)+2)
	for i := 0; i < len(A); i++ {
		check[A[i]] = 1
	}
	for i := 1; i < len(A); i++ {
		if check[i] == 0 {
			return i
		}
	}
	return 0
}*/
