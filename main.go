package main

import (
	"ELMAcourses/task1"
	"ELMAcourses/task2"
	"ELMAcourses/task3"
	"ELMAcourses/task4"
	"fmt"
)

func main() {
	input1 := []int{1, 2, 4, 5}
	input2 := []int{3, 8, 9, 7, 6}
	K := 3
	input3 := []int{9, 3, 9, 3, 7, 9, 9}
	input4 := []int{1, 2, 3, 5}
	fmt.Println(task1.Solution_one(input1))    // 0 - не послед, 1 - послед.
	fmt.Println(task2.Solution_two(input2, K)) // сдвиг
	fmt.Println(task3.Solution_three(input3))  // элемент без пары
	fmt.Println(task4.Solution_four(input4))   // как 1, но сам пропущенный элемент
}
