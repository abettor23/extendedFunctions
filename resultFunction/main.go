package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evenOrNo(n int) bool {
	if n%2 == 0 && n != 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Введите число:")

	reader := bufio.NewReader(os.Stdin)

	var number int
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var err error
		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите корректное целое число.")
			continue
		}
		break
	}

	result := evenOrNo(number)
	fmt.Println("Проверка на чётность:", result)
}
