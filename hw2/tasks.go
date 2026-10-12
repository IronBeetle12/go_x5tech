package lesson2

// FizzBuzz возвращает FizzBuzz, Fizz, Buzz или пустую строку для обычного числа.
// Для n <= 0 возвращает invalid.
func FizzBuzz(n int) string {
	// TODO: Реализовать через switch без выражения.
	switch {
	case n <= 0:
		return "invalid"
	case (n%15 == 0):
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	}
	return ""
}

// Sum возвращает сумму. Для пустого среза и nil возвращает 0.
func Sum(values []int64) int64 {
	if len(values) == 0 || values == nil {
		return 0
	}
	var sum int64
	for _, curVal := range values {
		sum += curVal
	}
	return sum
	// TODO: Пройти по значениям циклом.
}

// Min возвращает минимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Min(values []int64) int64 {
	if len(values) == 0 || values == nil {
		return 0
	}
	minElement := values[0]
	for _, tmp := range values {
		if tmp < minElement {
			minElement = tmp
		}
	}
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	return minElement
}

// Max возвращает максимум. Для пустого среза и nil возвращает 0 по условию задачи.
func Max(values []int64) int64 {
	if len(values) == 0 || values == nil {
		return 0
	}
	maxElement := values[0]
	for _, tmp := range values {
		if tmp > maxElement {
			maxElement = tmp
		}
	}
	// TODO: Учесть пустой ввод, затем сравнить элементы.
	return maxElement
}

// Average возвращает среднее как float64. Для пустого среза и nil возвращает 0.
func Average(values []int64) float64 {
	if len(values) == 0 || values == nil {
		return 0
	}
	var mid float64
	for _, tmp := range values {
		mid += float64(tmp)
	}
	// TODO: Выполнить преобразования до деления.
	return mid / float64(len(values))
}

// Category возвращает invalid, small, regular или large по сумме в копейках.
func Category(amount int64) string {
	// TODO: Проверить границы через switch без выражения.
	switch {
	case amount < 0:
		return "invalid"
	case amount >= 0 && amount < 10000:
		return "small"
	case amount >= 10000 && amount < 50000:
		return "regular"
	case amount >= 50000:
		return "large"
	}
	return ""
}

// RuneCount считает кодовые точки Unicode в строке через range.
func RuneCount(text string) int {
	// TODO: Не использовать len(text) как количество кодовых точек.
	var len int
	for range text {
		len++
	}
	return len
}
