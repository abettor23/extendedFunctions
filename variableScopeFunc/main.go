package main

import (
	"fmt"
	"math/rand"
	"time"
)

// подготовливаем глобальные переменные - пустышки.

var length int
var numbers []int

// заготовленные пустышки наполянем случайными значениями: количество элементов и сами их значения.
// таким образом получаем набор (срез) случайной длины из случайных чисел для дальнейших действий.

func sliceProperties() {
	rand.Seed(time.Now().UnixNano())
	length = rand.Intn(100) + 1
	numbers = make([]int, length)

	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}
}

//создаем три функции с операцими сложения

func addition1(n int) int {
	randomAddNum := numbers[rand.Intn(length)]
	addResult := n + randomAddNum
	fmt.Printf("Первое сложение %d и %d. Результат: %d\n", n, randomAddNum, addResult)
	return addResult
}

func addition2(n int) int {
	randomAddNum := numbers[rand.Intn(length)]
	addResult := n + randomAddNum
	fmt.Printf("Второе сложение %d и %d. Результат: %d\n", n, randomAddNum, addResult)
	return addResult
}

func addition3(n int) int {
	randomAddNum := numbers[rand.Intn(length)]
	addResult := n + randomAddNum
	fmt.Printf("Третье сложение %d и %d. Результат: %d\n", n, randomAddNum, addResult)
	return addResult
}

//вызываем функию наполнения глобальных переменных, а далее попарно вызываем функции сложения
//первое число для операци сложения - случайное из среза, а далее результат предыдущего сложения

func main() {
	sliceProperties()

	result1 := addition1(numbers[rand.Intn(length)])
	result2 := addition2(result1)
	result3 := addition3(result2)

	fmt.Println("Финальный результат сложений:", result3)
}
