package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func multiplication(n *int) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	fmt.Printf("Умножаю %d на %d\n", *n, random)
	*n *= random
}

func addition(n int) (result int) {
	random := rand.Intn(100)
	fmt.Printf("Складываю %d и %d\n", n, random)
	result = n + random
	return
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

	multiplication(&number)
	result := addition(number)
	fmt.Println("Ваше преобразованное число:", result)
}
