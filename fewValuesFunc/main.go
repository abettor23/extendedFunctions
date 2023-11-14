package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* использовал такой подход, исходя из условия задания, что функции обязательно должны
   возвращать несколько значений
*/

// функция-генератор трёх случайных координат:

func pointGeneration() ([]int, []int, []int) {
	rand.Seed(time.Now().UnixNano())
	points := make([][]int, 3)

	for i := 0; i < len(points); i++ {
		points[i] = []int{rand.Intn(100), rand.Intn(100)}
	}

	return points[0], points[1], points[2]
}

// функция преобразующая входящие координаты по формулам и возвращающая новые значения:

func transforming(a, b, c []int) ([]int, []int, []int) {

	pointsTransform := [][]int{a, b, c}

	for i := 0; i < len(pointsTransform); i++ {
		pointsTransform[i][0] = 2*pointsTransform[i][0] + 10
		pointsTransform[i][1] = -3*pointsTransform[i][1] - 5
	}

	return pointsTransform[0], pointsTransform[1], pointsTransform[2]
}

func main() {
	p1, p2, p3 := pointGeneration()
	fmt.Println("Исходные точки:")
	fmt.Printf("Точка 1: %v\nТочка 2: %v\nТочка 3: %v\n", p1, p2, p3)

	p1t, p2t, p3t := transforming(p1, p2, p3)
	fmt.Println("Преобразованные точки:")
	fmt.Printf("Точка 1: %v\nТочка 2: %v\nТочка 3: %v\n", p1t, p2t, p3t)
}
