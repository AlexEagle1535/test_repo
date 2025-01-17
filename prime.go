package main

import "fmt"

// PrintPrimes печатает простые числа меньше или равные n.
// Функция возвращает количество найденных простых чисел.
func PrintPrimes(n int) int {
	// сразу создадим слайс, в котором будем отмечать true составные числа
	// если в процессе перебора s[i] равно false, то i — простое число
	s := make([]bool, n+1)

	count := 0 // считаем количество найденных чисел
	for i := 2; i <= n; i++ {
		if !s[i] { // у этого числа нет делителей - нашли простое число
			count++
			// %3d будет дополнять число слева пробелами до 3 символов
			fmt.Printf("%3d ", i)
			for k := i + i; k <= n; k += i {
				// если k делится на i без остатка - отмечаем
				// что делитель для k - найден
				s[k] = true
			}
			// вывод по 10 чисел в строке, а затем перевод строки
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
	return count
}

func main() {
	count := PrintPrimes(300)
	fmt.Printf("\nКоличество простых чисел: %d\n", count)
}
